package org.entur.gbfs.v3_0.manifest;

import org.entur.gbfs.TestBase;
import org.entur.gbfs.v3_0.manifest.GBFSManifest;
import org.junit.jupiter.api.Test;

class ManifestTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/manifest.json", GBFSManifest.class);
    }
}
