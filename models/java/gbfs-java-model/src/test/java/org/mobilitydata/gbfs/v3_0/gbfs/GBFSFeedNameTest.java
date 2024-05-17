package org.mobilitydata.gbfs.v3_0.gbfs;

import org.mobilitydata.gbfs.v3_0.station_information.GBFSStationInformation;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class GBFSFeedNameTest {

    @Test
    void value() {
        Assertions.assertEquals("station_information", org.mobilitydata.gbfs.v3_0.gbfs.GBFSFeed.Name.STATION_INFORMATION.value());
    }

    @Test
    void implementingClass() {
        assertEquals(GBFSStationInformation.class, org.mobilitydata.gbfs.v3_0.gbfs.GBFSFeedName.implementingClass(org.mobilitydata.gbfs.v3_0.gbfs.GBFSFeed.Name.STATION_INFORMATION));
    }

    @Test
    void fromValue() {
        assertEquals(org.mobilitydata.gbfs.v3_0.gbfs.GBFSFeed.Name.STATION_INFORMATION, org.mobilitydata.gbfs.v3_0.gbfs.GBFSFeed.Name.fromValue("station_information"));
    }

    @Test
    void fromClass() {
        assertEquals(GBFSFeed.Name.STATION_INFORMATION, GBFSFeedName.fromClass(GBFSStationInformation.class));
    }
}