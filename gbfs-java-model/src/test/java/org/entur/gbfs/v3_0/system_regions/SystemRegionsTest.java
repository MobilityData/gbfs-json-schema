package org.entur.gbfs.v3_0.system_regions;

import org.entur.gbfs.TestBase;
import org.entur.gbfs.v3_0.system_regions.GBFSSystemRegions;
import org.junit.jupiter.api.Test;

class SystemRegionsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/system_regions.json", GBFSSystemRegions.class);
    }
}
