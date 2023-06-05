package org.entur.gbfs.v2_2.system_regions;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class SystemRegionsTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/system_regions.json", GBFSSystemRegions.class);
    }
}
