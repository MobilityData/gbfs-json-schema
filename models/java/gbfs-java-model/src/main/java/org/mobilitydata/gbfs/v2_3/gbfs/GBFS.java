package org.mobilitydata.gbfs.v2_3.gbfs;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyDescription;

import java.io.Serializable;
import java.util.Map;

public class GBFS extends GBFSGbfs implements Serializable {

    @JsonProperty("data")
    @JsonPropertyDescription("Response data in the form of name:value pairs.")
    private Map<String, GBFSFeeds> data;


    @JsonProperty("data")
    public Map<String, GBFSFeeds> getFeedsData() {
        return data;
    }

    @JsonProperty("data")
    public void setFeedsData(Map<String, GBFSFeeds> data) {
        this.data = data;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(GBFS.class.getName()).append('@').append(Integer.toHexString(System.identityHashCode(this))).append('[');
        sb.append("lastUpdated");
        sb.append('=');
        sb.append(((this.getLastUpdated() == null)?"<null>":this.getLastUpdated()));
        sb.append(',');
        sb.append("ttl");
        sb.append('=');
        sb.append(((this.getTtl() == null)?"<null>":this.getTtl()));
        sb.append(',');
        sb.append("version");
        sb.append('=');
        sb.append(((this.getVersion() == null)?"<null>":this.getVersion()));
        sb.append(',');
        sb.append("data");
        sb.append('=');
        sb.append(((this.data == null)?"<null>":this.data));
        sb.append(',');
        sb.append("additionalProperties");
        sb.append('=');
        sb.append(((this.getAdditionalProperties() == null)?"<null>":this.getAdditionalProperties()));
        sb.append(',');
        if (sb.charAt((sb.length()- 1)) == ',') {
            sb.setCharAt((sb.length()- 1), ']');
        } else {
            sb.append(']');
        }
        return sb.toString();
    }
}
