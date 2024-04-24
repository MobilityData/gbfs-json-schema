package org.mobilitydata.gbfs.v2_3.gbfs;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyDescription;

import java.io.Serializable;
import java.util.List;

public class GBFSFeeds implements Serializable {
    @JsonProperty("feeds")
    @JsonPropertyDescription("An array of all of the feeds that are published by the auto-discovery file. Each element in the array is an object with the keys below.")
    private List<GBFSFeed> feeds;

    public List<GBFSFeed> getFeeds() {
        return feeds;
    }

    public void setFeeds(List<GBFSFeed> feeds) {
        this.feeds = feeds;
    }

    @Override
    public String toString() {
        return "GBFSFeeds{" +
                "feeds=" + feeds +
                '}';
    }
}
