package org.entur.gbfs.v3_0_RC.manifest;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class ManifestTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0_RC/manifest.json", GBFSManifest.class);
    }
}
