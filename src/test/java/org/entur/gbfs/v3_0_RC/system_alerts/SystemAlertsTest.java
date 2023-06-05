package org.entur.gbfs.v3_0_RC.system_alerts;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class SystemAlertsTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v3_0_RC/system_alerts.json", GBFSSystemAlerts.class);
    }
}
