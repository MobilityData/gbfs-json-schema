package org.mobilitydata.gbfs.v3_1_RC2.vehicle_availability;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class VehicleAvailabilityTest extends TestBase {
	@Test
	void testUnmarshal() {
		assertUnmarshalDoesNotThrow("v3_1_RC2/vehicle_availability.json", GBFSVehicleAvailability.class);
	}
}
