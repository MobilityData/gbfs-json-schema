package org.entur.gbfs.v3_0_RC.gbfs_versions;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class GBFSVersionsTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0_RC/gbfs_versions.json", GBFSGbfsVersions.class);
    }
}
