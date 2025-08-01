# Test Fixtures for GBFS

This directory contains JSON files used to test GBFS models in various programming languages for versions v2.3, v3.0, and v3.1-RC. The test fixtures help ensure that the models are correctly implemented and can handle the expected data formats.

These test fixtures were directly copied from the GitHub specification examples for v2.3, v3.0, and v3.1-RC.

## Deviations

In some situations, it is necessary to deviate from the original specification examples. These deviations are documented below.

### v3.0:
- In geofencing_zones.json, the ride_through_allowed field was added to data.geofencing_zones.features[0].properties.rules[0] to comply with the schema. This change was made in a later version (https://github.com/MobilityData/gbfs/pull/651).