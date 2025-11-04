package org.mobilitydata.gbfs.v3_1_RC2.system_alerts;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemAlertsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_1_RC2/system_alerts.json", GBFSSystemAlerts.class);
    }
}
