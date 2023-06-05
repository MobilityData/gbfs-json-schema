package org.entur.gbfs.v2_3.system_calendar;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class SystemCalendarTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/system_calendar.json", GBFSSystemCalendar.class);
    }
}
