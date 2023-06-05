package org.entur.gbfs.v2_3.gbfs;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class GBFSTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/gbfs.json", GBFS.class);
    }
}
