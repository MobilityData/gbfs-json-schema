package org.entur.gbfs.v3_0_RC.gbfs_versions;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class GBFSVersionsTest extends TestBase {

    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v3_0_RC/gbfs_versions.json", GBFSGbfsVersions.class);
    }
}
