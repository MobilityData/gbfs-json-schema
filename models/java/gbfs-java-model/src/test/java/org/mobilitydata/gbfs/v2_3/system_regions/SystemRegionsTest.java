package org.mobilitydata.gbfs.v2_3.system_regions;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemRegionsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/system_regions.json", GBFSSystemRegions.class);
    }
}
