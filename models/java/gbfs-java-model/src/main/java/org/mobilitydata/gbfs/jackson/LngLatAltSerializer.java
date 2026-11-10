package org.mobilitydata.gbfs.jackson;

import org.mobilitydata.gbfs.geojson.LngLatAlt;
import tools.jackson.core.JsonGenerator;
import tools.jackson.databind.SerializationContext;
import tools.jackson.databind.ValueSerializer;

/**
 * Writes a {@link LngLatAlt} as a GeoJSON position array: {@code [lng, lat]} or {@code [lng, lat, alt]}.
 */
public class LngLatAltSerializer extends ValueSerializer<LngLatAlt> {

    @Override
    public void serialize(LngLatAlt value, JsonGenerator gen, SerializationContext ctxt) {
        gen.writeStartArray();
        gen.writeNumber(value.getLongitude());
        gen.writeNumber(value.getLatitude());
        if (value.getAltitude() != null) {
            gen.writeNumber(value.getAltitude());
        }
        gen.writeEndArray();
    }
}
