package org.entur.gbfs.v2_3.free_bike_status;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class FreeBikeStatusTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/free_bike_status.json", GBFSFreeBikeStatus.class);
    }
}
