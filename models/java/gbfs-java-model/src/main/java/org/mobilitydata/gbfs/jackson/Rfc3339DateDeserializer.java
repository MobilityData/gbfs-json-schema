package org.mobilitydata.gbfs.jackson;

import tools.jackson.core.JsonParser;
import tools.jackson.core.JsonToken;
import tools.jackson.databind.DeserializationContext;
import tools.jackson.databind.ValueDeserializer;

import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;
import java.util.Date;

/**
 * Reads an RFC 3339 date-time, e.g. {@code 2023-07-17T13:34:13+02:00}, into a {@link Date}.
 * <p>
 * Anything else is rejected, including numbers (which would otherwise be read as epoch milliseconds),
 * values without an offset and date-only values.
 */
public class Rfc3339DateDeserializer extends ValueDeserializer<Date> {

    @Override
    public Date deserialize(JsonParser p, DeserializationContext ctxt) {
        if (p.currentToken() != JsonToken.VALUE_STRING) {
            return (Date) ctxt.handleUnexpectedToken(Date.class, p);
        }
        String text = p.getString();
        try {
            return Date.from(OffsetDateTime.parse(text, DateTimeFormatter.ISO_OFFSET_DATE_TIME).toInstant());
        } catch (DateTimeParseException e) {
            throw ctxt.weirdStringException(text, Date.class, "not an RFC 3339 date-time: " + e.getMessage());
        }
    }
}
