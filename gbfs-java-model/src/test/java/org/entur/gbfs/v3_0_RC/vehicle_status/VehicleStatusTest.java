package org.entur.gbfs.v3_0_RC.vehicle_status;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class VehicleStatusTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0_RC/vehicle_status.json", GBFSVehicleStatus.class);
    }
}
