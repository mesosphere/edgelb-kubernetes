package com.mesosphere.sdk.config.validate;


import com.mesosphere.sdk.scheduler.SchedulerFlags;
import com.mesosphere.sdk.specification.*;
import com.mesosphere.sdk.testutils.TestConstants;
import org.junit.Before;
import org.junit.Test;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;

import java.util.Arrays;
import java.util.Collection;
import java.util.Collections;
import java.util.Optional;

import static org.hamcrest.Matchers.*;
import static org.junit.Assert.assertThat;
import static org.mockito.Mockito.when;

public class TLSRequiresServiceAccountTest {

    @Mock
    private PodSpec podWithTLS;
    @Mock
    private TaskSpec taskWithTLS;

    @Mock
    private PodSpec podWithoutTLS;
    @Mock
    private TaskSpec taskWithoutTLS;

    @Mock
    private SchedulerFlags flags;

    private Optional<ServiceSpec> original = Optional.empty();

    @Before
    public void init() {
        MockitoAnnotations.initMocks(this);

        when(taskWithTLS.getTransportEncryption()).thenReturn(
                Arrays.asList(
                        new DefaultTransportEncryptionSpec.Builder()
                                .name("server")
                                .type(TransportEncryptionSpec.Type.TLS).build())
        );
        when(podWithTLS.getTasks()).thenReturn(Arrays.asList(taskWithTLS));
        when(podWithTLS.getType()).thenReturn(TestConstants.POD_TYPE);

        when(taskWithoutTLS.getTransportEncryption()).thenReturn(Collections.emptyList());
        when(podWithoutTLS.getTasks()).thenReturn(Arrays.asList(taskWithoutTLS));
        when(podWithoutTLS.getType()).thenReturn(TestConstants.POD_TYPE);
    }

    private ServiceSpec createServiceSpec(PodSpec podSpec) {
        return DefaultServiceSpec.newBuilder()
                .addPod(podSpec)
                .name(TestConstants.SERVICE_NAME)
                .principal(TestConstants.PRINCIPAL)
                .build();
    }

    @Test
    public void testNoTLSNoServiceAccount() {
        Collection<ConfigValidationError> errors = new TLSRequiresServiceAccount(flags)
                .validate(original, createServiceSpec(podWithoutTLS));
        assertThat(errors, is(empty()));
    }

    @Test
    public void testNoTLSWithServiceAccount() {
        when(flags.getServiceAccountUid()).thenReturn(TestConstants.SERVICE_USER);
        Collection<ConfigValidationError> errors = new TLSRequiresServiceAccount(flags)
                .validate(original, createServiceSpec(podWithoutTLS));
        assertThat(errors, is(empty()));
    }

    @Test
    public void testWithTLSNoServiceAccount() {
        Collection<ConfigValidationError> errors = new TLSRequiresServiceAccount(flags)
                .validate(original, createServiceSpec(podWithTLS));
        assertThat(errors, hasSize(1));
    }

    @Test
    public void testWithTLSWithServiceAccount() {
        when(flags.getServiceAccountUid()).thenReturn(TestConstants.SERVICE_USER);
        Collection<ConfigValidationError> errors = new TLSRequiresServiceAccount(flags)
                .validate(original, createServiceSpec(podWithTLS));
        assertThat(errors, is(empty()));
    }

    @Test
    public void testEmptyServiceAccountUidIsNotValid() {
        when(flags.getServiceAccountUid()).thenReturn("");
        Collection<ConfigValidationError> errors = new TLSRequiresServiceAccount(flags)
                .validate(original, createServiceSpec(podWithTLS));
        assertThat(errors, hasSize(1));
    }

    @Test
    public void testWhitespaceServiceAccountUidIsNotValid() {
        when(flags.getServiceAccountUid()).thenReturn("    ");
        Collection<ConfigValidationError> errors = new TLSRequiresServiceAccount(flags)
                .validate(original, createServiceSpec(podWithTLS));
        assertThat(errors, hasSize(1));
    }

    @Test
    public void testNullServiceAccountUidIsNotValid() {
        when(flags.getServiceAccountUid()).thenReturn(null);
        Collection<ConfigValidationError> errors = new TLSRequiresServiceAccount(flags)
                .validate(original, createServiceSpec(podWithTLS));
        assertThat(errors, hasSize(1));
    }

    @Test
    public void testNullFlagsAreNotValid() {
        // TODO(elezar): How do we guarantee that the constructor is never called with a null value?
        //               Is a @NonNull annotation sufficient?
        Collection<ConfigValidationError> errors = new TLSRequiresServiceAccount(null)
                .validate(original, createServiceSpec(podWithTLS));
        assertThat(errors, hasSize(1));
    }

}
