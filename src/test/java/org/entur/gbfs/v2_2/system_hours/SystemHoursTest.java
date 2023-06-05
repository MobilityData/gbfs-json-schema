package org.entur.gbfs.v2_2.system_hours;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemHoursTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/system_hours.json", GBFSSystemHours.class);
    }
}
