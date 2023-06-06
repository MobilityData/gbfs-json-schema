package org.entur.gbfs.v3_0_RC.system_alerts;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemAlertsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0_RC/system_alerts.json", GBFSSystemAlerts.class);
    }
}
