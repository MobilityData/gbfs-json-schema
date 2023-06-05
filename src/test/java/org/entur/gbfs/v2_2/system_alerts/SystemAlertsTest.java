package org.entur.gbfs.v2_2.system_alerts;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class SystemAlertsTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/system_alerts.json", GBFSSystemAlerts.class);
    }
}
