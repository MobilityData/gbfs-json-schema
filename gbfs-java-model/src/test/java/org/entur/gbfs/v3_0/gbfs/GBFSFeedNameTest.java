package org.entur.gbfs.v3_0.gbfs;

import org.entur.gbfs.v3_0.gbfs.GBFSFeed;
import org.entur.gbfs.v3_0.station_information.GBFSStationInformation;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class GBFSFeedNameTest {

    @Test
    void value() {
        Assertions.assertEquals("station_information", org.entur.gbfs.v3_0.gbfs.GBFSFeed.Name.STATION_INFORMATION.value());
    }

    @Test
    void implementingClass() {
        assertEquals(GBFSStationInformation.class, org.entur.gbfs.v3_0.gbfs.GBFSFeedName.implementingClass(org.entur.gbfs.v3_0.gbfs.GBFSFeed.Name.STATION_INFORMATION));
    }

    @Test
    void fromValue() {
        assertEquals(org.entur.gbfs.v3_0.gbfs.GBFSFeed.Name.STATION_INFORMATION, org.entur.gbfs.v3_0.gbfs.GBFSFeed.Name.fromValue("station_information"));
    }

    @Test
    void fromClass() {
        assertEquals(GBFSFeed.Name.STATION_INFORMATION, GBFSFeedName.fromClass(GBFSStationInformation.class));
    }
}