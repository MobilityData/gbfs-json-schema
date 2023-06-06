package org.entur.gbfs.v2_3.vehicle_types;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class VehicleTypesTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/vehicle_types.json", GBFSVehicleTypes.class);
    }
}
