package org.entur.gbfs.v3_0.gbfs;

import org.entur.gbfs.TestBase;
import org.entur.gbfs.v3_0.gbfs.GBFSGbfs;
import org.junit.jupiter.api.Test;

class GBFSTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/gbfs.json", GBFSGbfs.class);
    }
}
