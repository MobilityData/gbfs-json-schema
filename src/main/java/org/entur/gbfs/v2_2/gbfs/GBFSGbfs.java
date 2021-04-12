
package org.entur.gbfs.v2_2.gbfs;

import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyDescription;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonValue;

import java.io.Serializable;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


/**
 * Auto-discovery file that links to all of the other files published by the system.
 * 
 */
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder({
    "last_updated",
    "ttl",
    "version",
    "data"
})
public class GBFSGbfs implements Serializable
{

    /**
     * Last time the data in the feed was updated in POSIX time.
     * (Required)
     * 
     */
    @JsonProperty("last_updated")
    @JsonPropertyDescription("Last time the data in the feed was updated in POSIX time.")
    private Integer lastUpdated;
    /**
     * Number of seconds before the data in the feed will be updated again (0 if the data should always be refreshed).
     * (Required)
     * 
     */
    @JsonProperty("ttl")
    @JsonPropertyDescription("Number of seconds before the data in the feed will be updated again (0 if the data should always be refreshed).")
    private Integer ttl;
    /**
     * GBFS version number to which the feed conforms, according to the versioning framework (added in v1.1).
     * (Required)
     * 
     */
    @JsonProperty("version")
    @JsonPropertyDescription("GBFS version number to which the feed conforms, according to the versioning framework (added in v1.1).")
    private Version version;
    /**
     * Response data in the form of name:value pairs.
     * (Required)
     * 
     */
    @JsonProperty("data")
    @JsonPropertyDescription("Response data in the form of name:value pairs.")
    private Map<String, GBFSFeeds> data;

    private final static long serialVersionUID = -6130628496027595734L;

    /**
     * Last time the data in the feed was updated in POSIX time.
     * (Required)
     * 
     */
    @JsonProperty("last_updated")
    public Integer getLastUpdated() {
        return lastUpdated;
    }

    /**
     * Last time the data in the feed was updated in POSIX time.
     * (Required)
     * 
     */
    @JsonProperty("last_updated")
    public void setLastUpdated(Integer lastUpdated) {
        this.lastUpdated = lastUpdated;
    }

    /**
     * Number of seconds before the data in the feed will be updated again (0 if the data should always be refreshed).
     * (Required)
     * 
     */
    @JsonProperty("ttl")
    public Integer getTtl() {
        return ttl;
    }

    /**
     * Number of seconds before the data in the feed will be updated again (0 if the data should always be refreshed).
     * (Required)
     * 
     */
    @JsonProperty("ttl")
    public void setTtl(Integer ttl) {
        this.ttl = ttl;
    }

    /**
     * GBFS version number to which the feed conforms, according to the versioning framework (added in v1.1).
     * (Required)
     * 
     */
    @JsonProperty("version")
    public Version getVersion() {
        return version;
    }

    /**
     * GBFS version number to which the feed conforms, according to the versioning framework (added in v1.1).
     * (Required)
     * 
     */
    @JsonProperty("version")
    public void setVersion(Version version) {
        this.version = version;
    }

    /**
     * Response data in the form of name:value pairs.
     * (Required)
     * 
     */
    @JsonProperty("data")
    public Map<String, GBFSFeeds> getData() {
        return data;
    }

    /**
     * Response data in the form of name:value pairs.
     * (Required)
     * 
     */
    @JsonProperty("data")
    public void setData(Map<String, GBFSFeeds> data) {
        this.data = data;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(GBFSGbfs.class.getName()).append('@').append(Integer.toHexString(System.identityHashCode(this))).append('[');
        sb.append("lastUpdated");
        sb.append('=');
        sb.append(((this.lastUpdated == null)?"<null>":this.lastUpdated));
        sb.append(',');
        sb.append("ttl");
        sb.append('=');
        sb.append(((this.ttl == null)?"<null>":this.ttl));
        sb.append(',');
        sb.append("version");
        sb.append('=');
        sb.append(((this.version == null)?"<null>":this.version));
        sb.append(',');
        sb.append("data");
        sb.append('=');
        sb.append(((this.data == null)?"<null>":this.data));
        sb.append(',');
        if (sb.charAt((sb.length()- 1)) == ',') {
            sb.setCharAt((sb.length()- 1), ']');
        } else {
            sb.append(']');
        }
        return sb.toString();
    }

    @Override
    public int hashCode() {
        int result = 1;
        result = ((result* 31)+((this.lastUpdated == null)? 0 :this.lastUpdated.hashCode()));
        result = ((result* 31)+((this.data == null)? 0 :this.data.hashCode()));
        result = ((result* 31)+((this.ttl == null)? 0 :this.ttl.hashCode()));
        result = ((result* 31)+((this.version == null)? 0 :this.version.hashCode()));
        return result;
    }

    @Override
    public boolean equals(Object other) {
        if (other == this) {
            return true;
        }
        if ((other instanceof GBFSGbfs) == false) {
            return false;
        }
        GBFSGbfs rhs = ((GBFSGbfs) other);
        return ((((((this.lastUpdated == rhs.lastUpdated)||((this.lastUpdated!= null)&&this.lastUpdated.equals(rhs.lastUpdated)))&&((this.data == rhs.data)||((this.data!= null)&&this.data.equals(rhs.data))))&&((this.ttl == rhs.ttl)||((this.ttl!= null)&&this.ttl.equals(rhs.ttl))))&&((this.version == rhs.version)||((this.version!= null)&&this.version.equals(rhs.version)))));
    }


    /**
     * GBFS version number to which the feed conforms, according to the versioning framework (added in v1.1).
     * 
     */
    public enum Version {

        _2_1_RC("2.1-RC"),
        _2_1_RC_2("2.1-RC2"),
        _2_1("2.1"),
        _2_2("2.2"),
        _3_0_RC("3.0-RC"),
        _3_0("3.0");
        private final String value;
        private final static Map<String, Version> CONSTANTS = new HashMap<String, Version>();

        static {
            for (Version c: values()) {
                CONSTANTS.put(c.value, c);
            }
        }

        private Version(String value) {
            this.value = value;
        }

        @Override
        public String toString() {
            return this.value;
        }

        @JsonValue
        public String value() {
            return this.value;
        }

        @JsonCreator
        public static Version fromValue(String value) {
            Version constant = CONSTANTS.get(value);
            if (constant == null) {
                throw new IllegalArgumentException(value);
            } else {
                return constant;
            }
        }

    }

}
