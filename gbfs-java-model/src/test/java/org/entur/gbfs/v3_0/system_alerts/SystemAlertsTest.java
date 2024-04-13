package org.entur.gbfs.v3_0.system_alerts;

import org.entur.gbfs.TestBase;
import org.entur.gbfs.v3_0.system_alerts.GBFSSystemAlerts;
import org.junit.jupiter.api.Test;

class SystemAlertsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/system_alerts.json", GBFSSystemAlerts.class);
    }
}
