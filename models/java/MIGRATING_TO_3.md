# Migrating `gbfs-java-model` from 2.x to 3.x

Version 3.0.0 replaces the `org.geojson` types in the model's public API with types the model owns.
The `geojson-jackson` dependency is removed.

## Code changes

Geofencing zone geometries (`geometry`) and station areas (`station_area`) are now
`org.mobilitydata.gbfs.geojson.MultiPolygon`, not `org.geojson.MultiPolygon`. The coordinates are
`List<List<List<LngLatAlt>>>`, as before, with `LngLatAlt` now in `org.mobilitydata.gbfs.geojson`.

- Change imports from `org.geojson.MultiPolygon` and `org.geojson.LngLatAlt` to the
  `org.mobilitydata.gbfs.geojson` package.
- `LngLatAlt` has `getLongitude()`, `getLatitude()` and `getAltitude()` (`null` when the position has no
  altitude). The altitude is a `Double` instead of a primitive `double` that was `NaN` when absent, so
  `hasAltitude()` becomes a `null` check. `getAdditionalElements()` is not carried over.
- `MultiPolygon` has `getType()`, which is always `"MultiPolygon"`, and `getCoordinates()`. Other `org.geojson`
  members (`getBbox()`, `getCrs()` and the `GeoJsonObject` base class) are not carried over.
  GBFS does not define them for these properties.
- Reading a geometry whose `type` is not `MultiPolygon`, or a position that is not an array of at least two
  numbers, now fails. Position elements after the third are ignored. Unknown properties on a geometry, such as
  `bbox`, are ignored whatever the mapper's `FAIL_ON_UNKNOWN_PROPERTIES` setting.

## Dependencies

- `de.grundid.opendatalab:geojson-jackson` is no longer needed by the model. If your code uses `org.geojson`
  itself, add that dependency to your own build, and convert between the two types where the model is involved.
