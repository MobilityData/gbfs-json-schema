package org.entur.gbfs.v3_0_RC.gbfs;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class GBFSTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0_RC/gbfs.json", GBFSGbfs.class);
    }
}
