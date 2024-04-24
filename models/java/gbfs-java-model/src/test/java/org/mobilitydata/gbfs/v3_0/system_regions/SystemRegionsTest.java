package org.mobilitydata.gbfs.v3_0.system_regions;

import org.mobilitydata.gbfs.TestBase;
import org.mobilitydata.v3_0.GBFSSystemRegions;
import org.junit.jupiter.api.Test;

class SystemRegionsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/system_regions.json", GBFSSystemRegions.class);
    }
}
