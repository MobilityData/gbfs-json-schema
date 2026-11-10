module org.mobilitydata.gbfs {
    requires transitive tools.jackson.databind; // also brings in tools.jackson.core and com.fasterxml.jackson.annotation
    requires transitive geojson.jackson; // org.geojson.MultiPolygon is part of the public API

    // Serializers referenced from @JsonSerialize/@JsonDeserialize on the generated classes
    exports org.mobilitydata.gbfs.jackson to tools.jackson.databind;

    exports org.mobilitydata.gbfs.v2_2.free_bike_status;
    exports org.mobilitydata.gbfs.v2_2.gbfs;
    exports org.mobilitydata.gbfs.v2_2.gbfs_versions;
    exports org.mobilitydata.gbfs.v2_2.geofencing_zones;
    exports org.mobilitydata.gbfs.v2_2.station_information;
    exports org.mobilitydata.gbfs.v2_2.station_status;
    exports org.mobilitydata.gbfs.v2_2.system_alerts;
    exports org.mobilitydata.gbfs.v2_2.system_calendar;
    exports org.mobilitydata.gbfs.v2_2.system_hours;
    exports org.mobilitydata.gbfs.v2_2.system_information;
    exports org.mobilitydata.gbfs.v2_2.system_pricing_plans;
    exports org.mobilitydata.gbfs.v2_2.system_regions;
    exports org.mobilitydata.gbfs.v2_2.vehicle_types;

    exports org.mobilitydata.gbfs.v2_3.free_bike_status;
    exports org.mobilitydata.gbfs.v2_3.gbfs;
    exports org.mobilitydata.gbfs.v2_3.gbfs_versions;
    exports org.mobilitydata.gbfs.v2_3.geofencing_zones;
    exports org.mobilitydata.gbfs.v2_3.station_information;
    exports org.mobilitydata.gbfs.v2_3.station_status;
    exports org.mobilitydata.gbfs.v2_3.system_alerts;
    exports org.mobilitydata.gbfs.v2_3.system_calendar;
    exports org.mobilitydata.gbfs.v2_3.system_hours;
    exports org.mobilitydata.gbfs.v2_3.system_information;
    exports org.mobilitydata.gbfs.v2_3.system_pricing_plans;
    exports org.mobilitydata.gbfs.v2_3.system_regions;
    exports org.mobilitydata.gbfs.v2_3.vehicle_types;

    exports org.mobilitydata.gbfs.v3_0.gbfs;
    exports org.mobilitydata.gbfs.v3_0.gbfs_versions;
    exports org.mobilitydata.gbfs.v3_0.geofencing_zones;
    exports org.mobilitydata.gbfs.v3_0.manifest;
    exports org.mobilitydata.gbfs.v3_0.station_information;
    exports org.mobilitydata.gbfs.v3_0.station_status;
    exports org.mobilitydata.gbfs.v3_0.system_alerts;
    exports org.mobilitydata.gbfs.v3_0.system_information;
    exports org.mobilitydata.gbfs.v3_0.system_pricing_plans;
    exports org.mobilitydata.gbfs.v3_0.system_regions;
    exports org.mobilitydata.gbfs.v3_0.vehicle_status;
    exports org.mobilitydata.gbfs.v3_0.vehicle_types;

    exports org.mobilitydata.gbfs.v3_1_RC3.gbfs;
    exports org.mobilitydata.gbfs.v3_1_RC3.gbfs_versions;
    exports org.mobilitydata.gbfs.v3_1_RC3.geofencing_zones;
    exports org.mobilitydata.gbfs.v3_1_RC3.manifest;
    exports org.mobilitydata.gbfs.v3_1_RC3.station_information;
    exports org.mobilitydata.gbfs.v3_1_RC3.station_status;
    exports org.mobilitydata.gbfs.v3_1_RC3.system_alerts;
    exports org.mobilitydata.gbfs.v3_1_RC3.system_information;
    exports org.mobilitydata.gbfs.v3_1_RC3.system_pricing_plans;
    exports org.mobilitydata.gbfs.v3_1_RC3.system_regions;
    exports org.mobilitydata.gbfs.v3_1_RC3.vehicle_availability;
    exports org.mobilitydata.gbfs.v3_1_RC3.vehicle_status;
    exports org.mobilitydata.gbfs.v3_1_RC3.vehicle_types;
}
