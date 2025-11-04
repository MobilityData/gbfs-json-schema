package org.mobilitydata.gbfs.v3_1_RC2.manifest;

import org.mobilitydata.gbfs.TestBase;

import org.junit.jupiter.api.Test;

class ManifestTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_1_RC2/manifest.json", GBFSManifest.class);
    }
}
