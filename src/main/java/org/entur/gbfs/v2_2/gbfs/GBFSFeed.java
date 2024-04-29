package org.entur.gbfs.v2_2.gbfs;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyDescription;

import java.io.Serializable;
import java.net.URI;
import java.util.Objects;

public class GBFSFeed implements Serializable {

    /**
     * Key identifying the type of feed this is. The key must be the base file name defined in the spec for the corresponding feed type.
     * (Required)
     *
     */
    @JsonProperty("name")
    @JsonPropertyDescription("Key identifying the type of feed this is. The key must be the base file name defined in the spec for the corresponding feed type.")
    private GBFSFeedName name;

    /**
     * URL for the feed.
     * (Required)
     *
     */
    @JsonProperty("url")
    @JsonPropertyDescription("URL for the feed.")
    private URI url;

    /**
     * Key identifying the type of feed this is. The key must be the base file name defined in the spec for the corresponding feed type.
     * (Required)
     *
     */
    public GBFSFeedName getName() {
        return name;
    }

    /**
     * Key identifying the type of feed this is. The key must be the base file name defined in the spec for the corresponding feed type.
     * (Required)
     *
     */
    public void setName(GBFSFeedName name) {
        this.name = name;
    }

    /**
     * URL for the feed.
     * (Required)
     *
     */
    public URI getUrl() {
        return url;
    }

    /**
     * URL for the feed.
     * (Required)
     *
     */
    public void setUrl(URI url) {
        this.url = url;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(GBFSFeed.class.getName()).append('@').append(Integer.toHexString(System.identityHashCode(this))).append('[');
        sb.append("name");
        sb.append('=');
        sb.append(((this.name == null)?"<null>":this.name));
        sb.append(',');
        sb.append("url");
        sb.append('=');
        sb.append(((this.url == null)?"<null>":this.url.toString()));
        sb.append(',');
        if (sb.charAt((sb.length()- 1)) == ',') {
            sb.setCharAt((sb.length()- 1), ']');
        } else {
            sb.append(']');
        }
        return sb.toString();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        GBFSFeed gbfsFeed = (GBFSFeed) o;
        return name == gbfsFeed.name && url.equals(gbfsFeed.url);
    }

    @Override
    public int hashCode() {
        return Objects.hash(name, url);
    }
}
