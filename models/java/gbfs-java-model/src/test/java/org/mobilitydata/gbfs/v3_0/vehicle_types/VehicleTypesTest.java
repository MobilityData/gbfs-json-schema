package org.mobilitydata.gbfs.v3_0.vehicle_types;

import org.mobilitydata.gbfs.TestBase;
import org.mobilitydata.gbfs.v3_0.GBFSVehicleTypes;
import org.junit.jupiter.api.Test;

class VehicleTypesTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/vehicle_types.json", GBFSVehicleTypes.class);
    }
}
