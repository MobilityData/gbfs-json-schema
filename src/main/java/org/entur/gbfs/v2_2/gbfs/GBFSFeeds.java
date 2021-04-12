package org.entur.gbfs.v2_2.gbfs;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyDescription;

import java.io.Serializable;
import java.util.List;
import java.util.Objects;

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

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        GBFSFeeds gbfsFeeds = (GBFSFeeds) o;
        return feeds.equals(gbfsFeeds.feeds);
    }

    @Override
    public int hashCode() {
        return Objects.hash(feeds);
    }
}
