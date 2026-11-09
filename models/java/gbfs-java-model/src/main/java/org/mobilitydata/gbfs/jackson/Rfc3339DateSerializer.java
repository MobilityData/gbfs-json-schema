package org.mobilitydata.gbfs.jackson;

import tools.jackson.core.JsonGenerator;
import tools.jackson.databind.SerializationContext;
import tools.jackson.databind.ValueSerializer;

import java.util.Date;

/**
 * Writes a {@link Date} as an RFC 3339 date-time in UTC, e.g. {@code 2023-07-17T13:34:13Z}.
 */
public class Rfc3339DateSerializer extends ValueSerializer<Date> {

    @Override
    public void serialize(Date value, JsonGenerator gen, SerializationContext ctxt) {
        gen.writeString(value.toInstant().toString());
    }
}
