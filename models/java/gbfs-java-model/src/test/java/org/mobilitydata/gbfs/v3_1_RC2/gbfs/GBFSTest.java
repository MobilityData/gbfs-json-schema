package org.mobilitydata.gbfs.v3_1_RC2.gbfs;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class GBFSTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_1_RC2/gbfs.json", GBFSGbfs.class);
    }
}
