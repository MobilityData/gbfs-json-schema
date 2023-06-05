package org.entur.gbfs.v3_0_RC.manifest;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class ManifestTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v3_0_RC/manifest.json", GBFSManifest.class);
    }
}
