package org.entur.gbfs.v2_2.gbfs_versions;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class GBFSVersionsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/gbfs_versions.json", GBFSGbfsVersions.class);
    }
}
