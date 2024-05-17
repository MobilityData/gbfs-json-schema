package org.mobilitydata.gbfs.v2_3.gbfs;

import org.mobilitydata.gbfs.v2_3.station_information.GBFSStationInformation;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class GBFSFeedNameTest {

    @Test
    void value() {
        Assertions.assertEquals("station_information", GBFSFeedName.StationInformation.value());
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