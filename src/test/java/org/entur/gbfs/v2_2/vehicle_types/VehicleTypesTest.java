package org.entur.gbfs.v2_2.vehicle_types;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class VehicleTypesTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/vehicle_types.json", GBFSVehicleTypes.class);
    }
}
