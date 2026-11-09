# Migrating `gbfs-java-model` from 1.x to 2.x

Version 2.0.0 moves the model from Jackson 2 to [Jackson 3](https://github.com/FasterXML/jackson),
and makes the handling of timestamps in GBFS v3.x feeds follow the specification.

The generated model itself is unchanged: the same classes exist in the same packages, and every
`public`/`protected` declaration has the same signature as in 1.0.14. What you
need to change is your Jackson setup, and (for v3.x feeds) anything that depends on how dates used
to be read or written.

## Dependencies

- The model requires Jackson 3 (`tools.jackson.core:jackson-databind`, Java 17 or newer).
  Jackson 3 keeps using the 2.x `jackson-annotations` (`com.fasterxml.jackson.annotation`).
- GeoJSON moves to `de.grundid.opendatalab:geojson-jackson` **3.0** (1.14 in 1.x). Both versions contain
  the same `org.geojson` classes, so **only one of them can be on the classpath**. If another library pulls
  in 1.x, exclude it or upgrade that library.
- Jackson 2 and 3 can otherwise live side by side, since they use different packages. But **every code path
  that reads or writes the model classes, or `org.geojson` objects, must use a Jackson 3 `ObjectMapper`**. The
  geofencing geometry types (and `geojson-jackson` 3.0) carry Jackson 3 (de)serializers, and the v3.x date
  handling relies on Jackson 3 annotations that Jackson 2 silently ignores.
- This includes code that only writes, such as an HTTP response converter in your web framework that serializes
  model objects returned from a controller. Frameworks that default to Jackson 3 (e.g. Spring Boot 4) can use
  their mapper directly. In frameworks that default to Jackson 2 (e.g. Spring Boot 3), a separate Jackson 3
  mapper for reading is not enough; the converter that writes the responses has to be Jackson 3 too.

## Code changes

The Jackson 2 to 3 changes (new `tools.jackson` package and group, `JsonMapper.builder()` instead of
configuring a mapper, unchecked `JacksonException` instead of `IOException`/`JsonProcessingException`, and
changed defaults) are documented in
[Jackson's migration guide](https://github.com/FasterXML/jackson/blob/main/jackson3/MIGRATING_TO_JACKSON_3.md).
A plain `new ObjectMapper()` still works if you do not need any configuration. Changes that are likely to
affect code using this model:

- `DeserializationFeature.READ_UNKNOWN_ENUM_VALUES_AS_NULL` is now
  `tools.jackson.databind.cfg.EnumFeature.READ_UNKNOWN_ENUM_VALUES_AS_NULL`.
- `objectMapper.readValue(URL, Class)` is removed. Use `objectMapper.readValue(url.openStream(), Class)`.
- Two new Jackson 3 defaults can change how a feed is parsed. `FAIL_ON_TRAILING_TOKENS` is enabled, so extra
  content after the JSON document is an error. `FAIL_ON_UNKNOWN_PROPERTIES` is disabled: most model classes keep
  unknown properties in `additionalProperties` either way, but on the few classes without one (the schema
  disallows extra properties, e.g. v3.1-RC3 `GBFSVehicle`), an unknown field used to fail with a default
  Jackson 2 mapper and is now silently ignored.

Because `JacksonException` is unchecked, you will not get a compile error for a missing `catch`, so look
through the call sites that read feeds.

## Dates in GBFS v3.x

The specification defines timestamps differently per version: integer POSIX timestamps in v2.x, and RFC 3339
`date-time` strings, e.g. `2023-07-17T13:34:13Z`, in v3.x (v3.0 and v3.1-RC3). The v2.x model exposes them
as `Integer` and the v3.x model as `java.util.Date`; neither changes in 2.x.

In 1.x the JSON representation of the v3.x `Date` fields was whatever your `ObjectMapper` was configured to do.
In 2.x these fields (the ones with `format: date-time` in the JSON schema) are annotated with dedicated
serializers, which makes the behaviour independent of the mapper:

| | 1.x | 2.x |
|---|---|---|
| Writing | Depends on the mapper (default Jackson 2 mapper: epoch milliseconds as a number; default Jackson 3 mapper: an ISO string) | Always an RFC 3339 string in UTC with `Z`, e.g. `2023-07-17T13:34:13Z`. Fractional seconds are only written when non-zero, e.g. `2023-07-17T13:34:13.250Z` |
| Reading | Lenient: whatever Jackson's default `Date` handling accepted | Strict: only RFC 3339 date-time strings, with `Z` or a numeric offset such as `+02:00` |

- **Remove your workarounds.** Custom date formats, `WRITE_DATES_AS_TIMESTAMPS` settings and custom `Date`
  serializers registered on the mapper are no longer needed for these fields. They are also not applied to them,
  because the annotations on the model's fields take precedence over mapper-level configuration. (Your
  configuration still applies to any other `Date` you serialize with the same mapper.)
- **The output can change.** This applies to anyone with their own date setup, including those who already read
  1.x with a Jackson 3 mapper. For instance, a mapper configured to write `2020-11-30T09:15:10.000+00:00` now
  produces `2020-11-30T09:15:10Z`. Both are valid RFC 3339 and denote the same instant, but anything comparing
  the strings, such as snapshot tests, needs updating.
- **Non-compliant v3.x input is now rejected.** The following used to be accepted and now fail with a
  `JacksonException` (an `InvalidFormatException` or `MismatchedInputException`):
  - numbers, e.g. `1600000000`. Jackson would have read it as epoch *milliseconds*, giving a date in
    January 1970 for epoch-seconds input
  - numeric strings, e.g. `"1600000000"`
  - strings without an offset, e.g. `2023-07-17T13:34:13`
  - date-only strings, e.g. `2023-07-17`
  - offsets without a colon, e.g. `2023-07-17T13:34:13+0200`
  - strings with leading whitespace

  A single invalid timestamp fails deserialization of the whole feed document. If you read feeds from third
  parties, handle the exception per feed so that one broken feed does not affect the others.
- **v2.x is unchanged.** No v2.x field is a `Date`, and serializing the v2.x test fixtures gives the same output
  as in 1.x. Date-only fields (`format: date`) are still plain `String`s.

## Other differences

- The generated `equals` and `hashCode` methods evaluate fields in a different order, because the generator
  (jsonschema2pojo) was upgraded from 1.2.1 to 1.3.3. `equals` gives the same results, and `hashCode` is still
  consistent with it, but `hashCode` *values* differ from 1.x. Do not rely on hash codes persisted by 1.x, and
  note that the iteration order of `HashMap`/`HashSet` containing model objects can differ (snapshot tests that
  depend on it may need updating).
- A new package, `org.mobilitydata.gbfs.jackson`, contains `Rfc3339DateSerializer` and `Rfc3339DateDeserializer`.
  They are referenced from the annotations of the generated date fields; you do not need to use them directly.
- For contributors: the version of the `rule-factory` module is bumped to 1.1.0 together with this version. It
  still uses the Jackson 2 `JsonNode` API internally, because the jsonschema2pojo plugin API is based on
  Jackson 2. This only affects building the model, not users of the published artifact.
