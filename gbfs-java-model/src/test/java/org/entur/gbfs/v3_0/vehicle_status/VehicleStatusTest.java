package org.entur.gbfs.v3_0.vehicle_status;

import org.entur.gbfs.TestBase;
import org.entur.gbfs.v3_0.vehicle_status.GBFSVehicleStatus;
import org.junit.jupiter.api.Test;

class VehicleStatusTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/vehicle_status.json", GBFSVehicleStatus.class);
    }
}
