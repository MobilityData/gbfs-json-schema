package org.mobilitydata.gbfs.v2_2.gbfs;

import org.mobilitydata.gbfs.v2_2.station_information.GBFSStationInformation;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class GBFSFeedNameTest {

    @Test
    void value() {
        assertEquals("station_information", GBFSFeedName.StationInformation.value());
    }

    @Test
    void implementingClass() {
        assertEquals(GBFSStationInformation.class, GBFSFeedName.StationInformation.implementingClass());
    }

    @Test
    void fromValue() {
        assertEquals(GBFSFeedName.StationInformation, GBFSFeedName.fromValue("station_information"));
    }

    @Test
    void fromClass() {
        assertEquals(GBFSFeedName.StationInformation, GBFSFeedName.fromClass(GBFSStationInformation.class));
    }
}