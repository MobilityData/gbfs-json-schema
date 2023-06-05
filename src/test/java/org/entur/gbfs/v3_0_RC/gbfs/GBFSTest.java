package org.entur.gbfs.v3_0_RC.gbfs;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class GBFSTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v3_0_RC/gbfs.json", GBFSGbfs.class);
    }
}
