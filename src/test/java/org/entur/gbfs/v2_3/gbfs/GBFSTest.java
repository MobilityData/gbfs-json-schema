package org.entur.gbfs.v2_3.gbfs;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class GBFSTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/gbfs.json", GBFS.class);
    }
}
