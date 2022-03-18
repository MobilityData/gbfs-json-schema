package org.entur.gbfs.v2_3_RC2.gbfs;

import org.entur.gbfs.v2_3_RC2.station_information.GBFSStationInformation;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

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