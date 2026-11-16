package org.mobilitydata.gbfs.geojson;

import org.mobilitydata.gbfs.jackson.LngLatAltDeserializer;
import org.mobilitydata.gbfs.jackson.LngLatAltSerializer;
import tools.jackson.databind.annotation.JsonDeserialize;
import tools.jackson.databind.annotation.JsonSerialize;

import java.io.Serializable;
import java.util.Objects;

/**
 * A GeoJSON position: {@code [longitude, latitude]} with an optional altitude as a third element.
 */
@JsonSerialize(using = LngLatAltSerializer.class)
@JsonDeserialize(using = LngLatAltDeserializer.class)
public class LngLatAlt implements Serializable {

    private static final long serialVersionUID = 1L;

    private double longitude;
    private double latitude;
    private Double altitude;

    public LngLatAlt() {}

    public LngLatAlt(double longitude, double latitude) {
        this.longitude = longitude;
        this.latitude = latitude;
    }

    public LngLatAlt(double longitude, double latitude, Double altitude) {
        this(longitude, latitude);
        this.altitude = altitude;
    }

    public double getLongitude() {
        return longitude;
    }

    public void setLongitude(double longitude) {
        this.longitude = longitude;
    }

    public double getLatitude() {
        return latitude;
    }

    public void setLatitude(double latitude) {
        this.latitude = latitude;
    }

    /** The altitude, or {@code null} if the position has none. */
    public Double getAltitude() {
        return altitude;
    }

    public void setAltitude(Double altitude) {
        this.altitude = altitude;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof LngLatAlt other)) return false;
        return Double.compare(longitude, other.longitude) == 0
                && Double.compare(latitude, other.latitude) == 0
                && Objects.equals(altitude, other.altitude);
    }

    @Override
    public int hashCode() {
        return Objects.hash(longitude, latitude, altitude);
    }

    @Override
    public String toString() {
        return "LngLatAlt{longitude=" + longitude + ", latitude=" + latitude + ", altitude=" + altitude + '}';
    }
}
