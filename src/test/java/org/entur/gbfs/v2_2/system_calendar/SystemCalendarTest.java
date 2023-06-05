package org.entur.gbfs.v2_2.system_calendar;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemCalendarTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/system_calendar.json", GBFSSystemCalendar.class);
    }
}
