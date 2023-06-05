package org.entur.gbfs.v2_3.gbfs_versions;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class GBFSVersionsTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/gbfs_versions.json", GBFSGbfsVersions.class);
    }
}
