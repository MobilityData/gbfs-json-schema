package org.entur.gbfs.v2_2.system_alerts;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemAlertsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/system_alerts.json", GBFSSystemAlerts.class);
    }
}
