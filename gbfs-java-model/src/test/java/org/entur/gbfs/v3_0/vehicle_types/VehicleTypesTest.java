package org.entur.gbfs.v3_0.vehicle_types;

import org.entur.gbfs.TestBase;
import org.entur.gbfs.v3_0.vehicle_types.GBFSVehicleTypes;
import org.junit.jupiter.api.Test;

class VehicleTypesTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/vehicle_types.json", GBFSVehicleTypes.class);
    }
}
