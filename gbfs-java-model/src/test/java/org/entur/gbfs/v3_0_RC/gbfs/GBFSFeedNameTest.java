package org.entur.gbfs.v3_0_RC.gbfs;

import org.entur.gbfs.v3_0_RC.station_information.GBFSStationInformation;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class GBFSFeedNameTest {

    @Test
    void value() {
        Assertions.assertEquals("station_information", GBFSFeed.Name.STATION_INFORMATION.value());
    }

    @Test
    void implementingClass() {
        assertEquals(GBFSStationInformation.class, org.entur.gbfs.v3_0_RC.gbfs.GBFSFeedName.implementingClass(GBFSFeed.Name.STATION_INFORMATION));
    }

    @Test
    void fromValue() {
        assertEquals(GBFSFeed.Name.STATION_INFORMATION, GBFSFeed.Name.fromValue("station_information"));
    }

    @Test
    void fromClass() {
        assertEquals(GBFSFeed.Name.STATION_INFORMATION, GBFSFeedName.fromClass(GBFSStationInformation.class));
    }
}