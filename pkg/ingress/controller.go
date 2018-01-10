package ingress

import (
	"fmt"
	"log"
	"reflect"
	"time"

	// Actor
	"github.com/AsynkronIT/protoactor-go/actor"

	// Ingress controller
	"github.com/asridharan/edgelb-k8s/pkg/state"

	// RxGo
	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"

	// K8s
	V1api "k8s.io/api/core/v1"
	V1Beta1api "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/listers/extensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

const (
	ADD    = iota
	UPDATE = iota
	DEL    = iota
)

type Operation struct {
	Op int
}

type k8sIngressMsg struct {
	Op      Operation
	Ingress V1Beta1api.Ingress
}

type k8sServiceMsg struct {
	Op      Operation
	Service V1api.Service
}

type k8sEndpointsMsg struct {
	Op        Operation
	Endpoints V1api.Endpoints
}

type ingressRuleMsg struct {
	Op          Operation
	Namespace   string
	IngressRule V1Beta1api.IngressRule
}

// Used to add/del/update a `Host` on this controller.
type hostMsg struct {
	Op   Operation
	Host string
}

// Used to add/del/update a `state.Service` on this controller.
type serviceMsg struct {
	Op      Operation
	Service state.ServiceName
}

// Used to add/del/update a service endpoint on this controller.
type endpointMsg struct {
	Op       Operation
	Service  state.ServiceName
	Endpoint string
}

type Controller interface {
	Start()
}

type controller struct {
	pid              *actor.PID // PID of the controller.
	si               informers.SharedInformerFactory
	endpoints        v1.EndpointsLister        // All the endpoints that are availabe in a k8s cluster.
	ingressResources v1beta1.IngressLister     // Ingress resource that define the config for the controller.
	services         map[string]*state.Service // Services for which the controller is asked to setup ingress.
	vhosts           map[string]*state.VHost
	//Observable channels
	ingressMsgs   chan k8sIngressMsg
	serviceMsgs   chan k8sServiceMsg
	endpointsMsgs chan k8sEndpointsMsg
}

func NewController(clientset *kubernetes.Clientset) (ctrl Controller, err error) {
	resyncPeriod := 30 * time.Minute
	si := informers.NewSharedInformerFactory(clientset, resyncPeriod)

	ingressCtrl := controller{
		si:               si,
		endpoints:        si.Core().V1().Endpoints().Lister(),
		ingressResources: si.Extensions().V1beta1().Ingresses().Lister(),
		services:         make(map[string]*state.Service),
		vhosts:           make(map[string]*state.VHost),
		ingressMsgs:      make(chan k8sIngressMsg),
		serviceMsgs:      make(chan k8sServiceMsg),
		endpointsMsgs:    make(chan k8sEndpointsMsg),
	}

	ctrl = &ingressCtrl

	// Add watchers for endpoints.
	si.Core().V1().Endpoints().Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				ingressCtrl.endpointsMsgs <- k8sEndpointsMsg{
					Op:        Operation{Op: ADD},
					Endpoints: *(obj.(*V1api.Endpoints)),
				}
			},
			UpdateFunc: func(old interface{}, new interface{}) {
				ingressCtrl.endpointsMsgs <- k8sEndpointsMsg{
					Op:        Operation{Op: UPDATE},
					Endpoints: *(new.(*V1api.Endpoints)),
				}
			},
			DeleteFunc: func(obj interface{}) {
				ingressCtrl.endpointsMsgs <- k8sEndpointsMsg{
					Op:        Operation{Op: DEL},
					Endpoints: *(obj.(*V1api.Endpoints)),
				}
			},
		},
	)

	// Add watchers for services.
	si.Core().V1().Services().Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				ingressCtrl.serviceMsgs <- k8sServiceMsg{
					Op:      Operation{Op: ADD},
					Service: *(obj.(*V1api.Service)),
				}
			},
			UpdateFunc: func(old interface{}, new interface{}) {
				ingressCtrl.serviceMsgs <- k8sServiceMsg{
					Op:      Operation{Op: UPDATE},
					Service: *(new.(*V1api.Service)),
				}
			},
			DeleteFunc: func(obj interface{}) {
				ingressCtrl.serviceMsgs <- k8sServiceMsg{
					Op:      Operation{Op: DEL},
					Service: *(obj.(*V1api.Service)),
				}
			},
		},
	)

	// Add watchers for ingress
	si.Extensions().V1beta1().Ingresses().Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				ingressCtrl.ingressMsgs <- k8sIngressMsg{
					Op:      Operation{Op: ADD},
					Ingress: *(obj.(*V1Beta1api.Ingress)),
				}
			},
			UpdateFunc: func(old interface{}, new interface{}) {
				ingressCtrl.ingressMsgs <- k8sIngressMsg{
					Op:      Operation{Op: UPDATE},
					Ingress: *(new.(*V1Beta1api.Ingress)),
				}
			},
			DeleteFunc: func(obj interface{}) {
				ingressCtrl.ingressMsgs <- k8sIngressMsg{
					Op:      Operation{Op: DEL},
					Ingress: *(obj.(*V1Beta1api.Ingress)),
				}
			},
		},
	)

	// Setup observers so that we can process the different k8s messages we are
	// interested in.
	sink := observer.Observer{
		NextHandler: func(item interface{}) {
			ingressCtrl.pid.Tell(&item)
		},
	}

	k8sIngressSource, _ := iterable.New(ingressCtrl.ingressMsgs)
	k8sServiceSource, _ := iterable.New(ingressCtrl.serviceMsgs)
	k8sEndpointsSource, _ := iterable.New(ingressCtrl.endpointsMsgs)

	observable.From(k8sIngressSource).Subscribe(sink)
	observable.From(k8sServiceSource).Subscribe(sink)
	observable.From(k8sEndpointsSource).Subscribe(sink)

	return
}

func (ctrl *controller) Start() {
	// Store the PID and spawn the actor.
	ctrl.pid = actor.Spawn(actor.FromInstance(ctrl))

	// Start the watchers.
	ctrl.si.Start(wait.NeverStop)
	<-wait.NeverStop
}

func (ctrl *controller) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *k8sIngressMsg:
		ingressMsg, _ := ctx.Message().(*k8sIngressMsg)
		switch operation := ingressMsg.Op.Op; operation {
		case ADD, UPDATE:
			ctrl.ingressCreateAndUpdate(ingressMsg.Ingress)
		case DEL:
			ctrl.ingressDeleted(ingressMsg.Ingress)
		default:
			log.Printf("Undefined operation %d requested on `k8sIngressMsg`", operation)
		}
	case *ingressRuleMsg:
		ingressRuleMsg, _ := ctx.Message().(*ingressRuleMsg)
		switch operation := ingressRuleMsg.Op.Op; operation {
		case ADD, UPDATE:
			host := ctrl.ingressRuleCreateAndUpdate(ingressRuleMsg.Namespace, ingressRuleMsg.IngressRule)
			// Tell the controller to process this host.
			ctrl.pid.Tell(&hostMsg{Op: Operation{Op: ADD}, Host: host})
		case DEL:
			host := ctrl.ingressRuleDeleted(ingressRuleMsg.Namespace, ingressRuleMsg.IngressRule)
			// Tell the controller to delete this host.
			ctrl.pid.Tell(&hostMsg{Op: Operation{Op: DEL}, Host: host})
		default:
			log.Printf("Undefined operation %d requested on `IngressRuleMsg`", operation)
		}
	case *k8sServiceMsg, *k8sEndpointsMsg:
		var id state.ServiceName
		var op Operation
		switch ctx.Message().(type) {
		case *k8sServiceMsg:
			service := (ctx.Message().(*k8sServiceMsg)).Service
			id = state.ServiceName{Namespace: service.GetNamespace(), Name: service.GetName()}
			op = (ctx.Message().(*k8sServiceMsg)).Op
		case *k8sEndpointsMsg:
			endpoints := (ctx.Message().(*k8sEndpointsMsg)).Endpoints
			id = state.ServiceName{Namespace: endpoints.GetNamespace(), Name: endpoints.GetName()}
			op = (ctx.Message().(*k8sEndpointsMsg)).Op
		}
		switch op.Op {
		case ADD, UPDATE:
			hosts := ctrl.updateServiceEndpoints(id)
			for _, host := range hosts {
				// Tell the controller to process this host.
				ctrl.pid.Tell(&hostMsg{Op: Operation{Op: ADD}, Host: host})
			}
		case DEL:
			hosts := ctrl.deleteService(id)
			for _, host := range hosts {
				// Tell the controller to delete this host.
				ctrl.pid.Tell(&hostMsg{Op: Operation{Op: DEL}, Host: host})
			}
		default:
			log.Printf("Undefined operation %d requested on `k8ServiceMsg/k8sEndpointsMsg`", op)
		}
	case *hostMsg:
		hostMsg, _ := ctx.Message().(*hostMsg)
		switch operation := hostMsg.Op.Op; operation {
		case ADD, UPDATE:
			log.Printf("Will send update for host:%s to the load-balancer", hostMsg.Host)
		case DEL:
			log.Printf("Will delete host:%s from the load-balancer", hostMsg.Host)
		default:
			log.Printf("Undefined operation %d requested on `k8ServiceMsg/k8sEndpointsMsg`", operation)
		}
	default:
		log.Printf("Unsopported message received by %s", ctrl.pid)
	}
}

func (ctrl *controller) endpointCreateUpdateAndDelete(obj interface{}) {
	endpoint, ok := obj.(*V1api.Endpoints)
	if !ok {
		log.Printf("Expected an object of type `*v1.Endpoint`, but got object of type: %s", reflect.TypeOf(obj))
		return
	}

	service := state.ServiceName{
		Name:      endpoint.GetName(),
		Namespace: endpoint.GetNamespace(),
	}

	serviceMsg := &serviceMsg{Op: Operation{Op: UPDATE}, Service: service}

	// We won't do anything specific to this endpoint. We will just ask the Service to recreate all the endpoints
	// belonging to this service if the service is actually being exposed.
	ctrl.pid.Tell(serviceMsg)
}

func (ctrl *controller) serviceCreateUpdateAndDelete(obj interface{}, Op Operation) {
	service, ok := obj.(*V1api.Service)
	if !ok {
		log.Printf("Expected an object of type `*v1.Service`, but got object of type: %s", reflect.TypeOf(obj))
		return
	}

	id := state.ServiceName{
		Name:      service.GetName(),
		Namespace: service.GetNamespace(),
	}

	serviceMsg := &serviceMsg{Op: Op, Service: id}

	// We won't do anything specific to this endpoint. We will just ask the Service to recreate all the endpoints
	// belonging to this service if the service is actually being exposed.
	ctrl.pid.Tell(serviceMsg)

}

func (ctrl *controller) ingressCreateAndUpdate(ingress V1Beta1api.Ingress) {
	namespace := ingress.GetNamespace()

	it, _ := iterable.New(ingress.Spec.Rules)

	// Process all the rules.
	observable.From(it).Subscribe(observer.Observer{
		// For every VHost that we get, register it with the load-balancer.
		NextHandler: func(item interface{}) {
			ingressRule := item.(*V1Beta1api.IngressRule)
			// Ask the controller to process this rule.
			ctrl.pid.Tell(&ingressRuleMsg{
				Op:          Operation{Op: ADD},
				Namespace:   namespace,
				IngressRule: *ingressRule})
		},
	})
}

func (ctrl *controller) ingressDeleted(ingress V1Beta1api.Ingress) {
	it, _ := iterable.New(ingress.Spec.Rules)

	// Process all the rules.
	observable.From(it).Subscribe(observer.Observer{
		// For every VHost that we get, register it with the load-balancer.
		NextHandler: func(item interface{}) {
			ingressRule := item.(*V1Beta1api.IngressRule)
			// Ask the controller to process this rule.
			ctrl.pid.Tell(&ingressRuleMsg{
				Op:          Operation{Op: DEL},
				Namespace:   ingress.GetNamespace(),
				IngressRule: *ingressRule})
		},
	})
}

// Create a `VHost` based on an `IngressRule`.
// Returns the host added/deleted in this updated.
func (ctrl *controller) ingressRuleCreateAndUpdate(namespace string, rule V1Beta1api.IngressRule) (host string) {
	// If a VHost already exists delete it since we will be re-creating it here.
	if _, ok := ctrl.vhosts[rule.Host]; ok {
		delete(ctrl.vhosts, rule.Host)
	}

	vhost := &state.VHost{Host: rule.Host}

	for _, path := range rule.HTTP.Paths {
		url := state.URL{Host: rule.Host, Path: path.Path}
		id := state.ServiceName{Namespace: namespace, Name: path.Backend.ServiceName}

		// Store the association of the service with the URI
		if _, ok := ctrl.services[id.String()]; !ok {
			ctrl.services[id.String()] = &state.Service{ServiceName: id}
		}

		ctrl.services[id.String()].URLs[url.String()] = url

		ctrl.updateServiceEndpoints(id)

		// Append the route to the VHost.
		route := state.Route{Path: url.Path, ServiceName: id}
		vhost.Routes[route.String()] = route
	}

	ctrl.vhosts[vhost.Host] = vhost

	host = vhost.Host

	return
}

// Delete a `Vhost` based on an IngressRule.
// Returns the `host` that got deleted.
func (ctrl *controller) ingressRuleDeleted(namespace string, rule V1Beta1api.IngressRule) (host string) {
	host = rule.Host

	// If a VHost already exists delete it since we will be re-creating it here.
	if _, ok := ctrl.vhosts[rule.Host]; ok {
		delete(ctrl.vhosts, rule.Host)
	}

	for _, path := range rule.HTTP.Paths {
		url := state.URL{Host: rule.Host, Path: path.Path}
		id := state.ServiceName{Namespace: namespace, Name: path.Backend.ServiceName}
		service, ok := ctrl.services[id.String()]
		// Delete any association of the service with this host
		if ok {
			delete(service.URLs, url.String())
		}
	}

	return
}

// Takes a `Service` and updates the endpoints of the service.
// Returns the affected vhosts.
func (ctrl *controller) updateServiceEndpoints(id state.ServiceName) (host []string) {
	service, ok := ctrl.services[id.String()]

	if !ok {
		// We don't have a VHost corresponding to this servcie so we don't need to do anything.
		log.Printf("Found a service(%s) that is not part of any ingress reousrce, hence skipping", id)
		return
	}

	// Since we might actually be adding/removing existing endpoints to the service,
	// remove the existing endpoints from the service before adding new ones.
	service.Endpoints = nil

	// Look at the service name, and get the corresponding endpoints for this service name.
	endpoints, err := ctrl.endpoints.Endpoints(service.Namespace).Get(service.Name)
	if err != nil {
		for _, endpoint := range endpoints.Subsets {
			for _, address := range endpoint.Addresses {
				for _, port := range endpoint.Ports {
					service.Endpoints = append(
						service.Endpoints,
						state.Endpoint{
							ServiceName: id,
							Address:     fmt.Sprintf("%s:%d", address.IP, port.Port),
						})
				}
			}

			for _, address := range endpoint.NotReadyAddresses {
				for _, port := range endpoint.Ports {
					service.Endpoints = append(
						service.Endpoints,
						state.Endpoint{
							ServiceName: id,
							Address:     fmt.Sprintf("%s:%d", address.IP, port.Port),
						})
				}
			}

		}
	} else {
		log.Printf("Unable to retrieve the endpoints for service:%s, error:%s", *service, err)
	}

	for vhost, _ := range service.URLs {
		host = append(host, vhost)
	}

	return
}

func (ctrl *controller) deleteService(id state.ServiceName) (host []string) {
	service, ok := ctrl.services[id.String()]

	if !ok {
		// We don't have a VHost corresponding to this servcie so we don't need to do anything.
		return
	}

	for vhost, _ := range service.URLs {
		host = append(host, vhost)
	}

	delete(ctrl.services, id.String())

	return
}
