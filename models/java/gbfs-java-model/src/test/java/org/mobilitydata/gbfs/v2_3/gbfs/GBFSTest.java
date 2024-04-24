package org.mobilitydata.gbfs.v2_3.gbfs;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class GBFSTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/gbfs.json", GBFS.class);
    }
}
