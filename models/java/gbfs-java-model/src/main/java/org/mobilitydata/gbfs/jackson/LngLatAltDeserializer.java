package org.mobilitydata.gbfs.jackson;

import org.mobilitydata.gbfs.geojson.LngLatAlt;
import tools.jackson.core.JsonParser;
import tools.jackson.databind.DeserializationContext;
import tools.jackson.databind.JsonNode;
import tools.jackson.databind.ValueDeserializer;

/**
 * Reads a GeoJSON position array, {@code [lng, lat]} or {@code [lng, lat, alt]}, into a {@link LngLatAlt}.
 * <p>
 * Elements after the third are ignored, since the GBFS schema allows any number of elements from two up.
 * Non-arrays, non-numeric elements, and arrays with fewer than two elements are rejected.
 */
public class LngLatAltDeserializer extends ValueDeserializer<LngLatAlt> {

    @Override
    public LngLatAlt deserialize(JsonParser p, DeserializationContext ctxt) {
        JsonNode node = ctxt.readTree(p);
        if (!node.isArray() || node.size() < 2) {
            return (LngLatAlt) ctxt.handleUnexpectedToken(LngLatAlt.class, p);
        }
        for (JsonNode element : node) {
            if (!element.isNumber()) {
                return (LngLatAlt) ctxt.handleUnexpectedToken(LngLatAlt.class, p);
            }
        }
        Double altitude = node.size() >= 3 ? node.get(2).doubleValue() : null;
        return new LngLatAlt(node.get(0).doubleValue(), node.get(1).doubleValue(), altitude);
    }
}
