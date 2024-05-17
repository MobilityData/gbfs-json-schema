package org.mobilitydata.gbfs.v2_3.system_calendar;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemCalendarTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/system_calendar.json", GBFSSystemCalendar.class);
    }
}
