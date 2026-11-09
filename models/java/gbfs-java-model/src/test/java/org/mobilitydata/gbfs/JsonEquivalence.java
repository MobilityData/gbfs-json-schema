package org.mobilitydata.gbfs;

import tools.jackson.databind.JsonNode;

import java.time.OffsetDateTime;
import java.time.format.DateTimeParseException;
import java.util.Comparator;

/**
 * Compares JSON trees ignoring differences that do not change the meaning of the data:
 * numbers are compared by value (12345 equals 12345.0), and RFC 3339 date-times by instant
 * (13:34:13+02:00 equals 11:34:13Z).
 */
public final class JsonEquivalence {
    private static final Comparator<JsonNode> BY_MEANING = (a, b) -> {
        if (a.isNumber() && b.isNumber()) {
            return a.decimalValue().compareTo(b.decimalValue());
        }
        if (a.isString() && b.isString()) {
            try {
                return OffsetDateTime.parse(a.stringValue()).toInstant()
                        .compareTo(OffsetDateTime.parse(b.stringValue()).toInstant());
            } catch (DateTimeParseException e) {
                // not date-times, compare as plain values below
            }
        }
        return a.equals(b) ? 0 : 1;
    };

    private JsonEquivalence() {
    }

    public static boolean equivalent(JsonNode expected, JsonNode actual) {
        return expected.equals(BY_MEANING, actual);
    }
}
