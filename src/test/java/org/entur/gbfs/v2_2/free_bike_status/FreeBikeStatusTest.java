package org.entur.gbfs.v2_2.free_bike_status;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class FreeBikeStatusTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/free_bike_status.json", GBFSFreeBikeStatus.class);
    }
}
