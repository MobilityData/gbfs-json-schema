package org.mobilitydata.gbfs.geojson;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/**
 * A GeoJSON MultiPolygon: a list of polygons, each a list of linear rings, each a list of positions.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
@JsonPropertyOrder({"type", "coordinates"})
public class MultiPolygon implements Serializable {

    private static final long serialVersionUID = 1L;

    public static final String TYPE = "MultiPolygon";

    private List<List<List<LngLatAlt>>> coordinates = new ArrayList<>();

    public MultiPolygon() {}

    public MultiPolygon(List<List<List<LngLatAlt>>> coordinates) {
        this.coordinates = coordinates;
    }

    @JsonProperty("type")
    public String getType() {
        return TYPE;
    }

    @JsonProperty("type")
    public void setType(String type) {
        if (!TYPE.equals(type)) {
            throw new IllegalArgumentException("Expected type " + TYPE + " but was " + type);
        }
    }

    @JsonProperty("coordinates")
    public List<List<List<LngLatAlt>>> getCoordinates() {
        return coordinates;
    }

    @JsonProperty("coordinates")
    public void setCoordinates(List<List<List<LngLatAlt>>> coordinates) {
        this.coordinates = coordinates;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof MultiPolygon other)) return false;
        return Objects.equals(coordinates, other.coordinates);
    }

    @Override
    public int hashCode() {
        return Objects.hash(coordinates);
    }

    @Override
    public String toString() {
        return "MultiPolygon{coordinates=" + coordinates + '}';
    }
}
