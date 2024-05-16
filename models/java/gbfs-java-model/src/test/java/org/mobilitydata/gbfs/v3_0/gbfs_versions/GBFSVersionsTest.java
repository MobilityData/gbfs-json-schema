package org.mobilitydata.gbfs.v3_0.gbfs_versions;

import org.mobilitydata.gbfs.TestBase;

import org.junit.jupiter.api.Test;

class GBFSVersionsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/gbfs_versions.json", GBFSGbfsVersions.class);
    }
}
