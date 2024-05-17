package org.mobilitydata.gbfs.v3_0.gbfs;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class GBFSTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/gbfs.json", GBFSGbfs.class);
    }
}
