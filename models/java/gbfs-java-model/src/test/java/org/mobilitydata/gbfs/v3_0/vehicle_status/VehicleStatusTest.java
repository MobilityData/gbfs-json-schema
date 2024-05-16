package org.mobilitydata.gbfs.v3_0.vehicle_status;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class VehicleStatusTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/vehicle_status.json", GBFSVehicleStatus.class);
    }
}
