import { createCheckers } from "ts-interface-checker";

// checker model
import FreeBikeStatusTI from "../v2.3/test-type-checkers/free_bike_status-ti";
import GbfsVersionsTI from "../v2.3/test-type-checkers/gbfs_versions-ti";
import GbfsTI from "../v2.3/test-type-checkers/gbfs-ti";
import GeofencingZonesTI from "../v2.3/test-type-checkers/geofencing_zones-ti";
import StationsInformationTI from "../v2.3/test-type-checkers/station_information-ti";
import StationStatusTI from "../v2.3/test-type-checkers/station_status-ti";
import SystemAlertsTI from "../v2.3/test-type-checkers/system_alerts-ti";
import SystemCalendarTI from "../v2.3/test-type-checkers/system_calendar-ti";
import SystemHoursTI from "../v2.3/test-type-checkers/system_hours-ti";
import SystemInformationTI from "../v2.3/test-type-checkers/system_information-ti";
import SystemPricingPlansTI from "../v2.3/test-type-checkers/system_pricing_plans-ti";
import SystemRegionsTI from "../v2.3/test-type-checkers/system_regions-ti";
import VehicleTypesTI from "../v2.3/test-type-checkers/vehicle_types-ti";

// checkers
const { FreeBikeStatus } = createCheckers(FreeBikeStatusTI);
const { GbfsVersions } = createCheckers(GbfsVersionsTI);
const { Gbfs } = createCheckers(GbfsTI);
const { GeofencingZones } = createCheckers(GeofencingZonesTI);
const { StationInformation } = createCheckers(StationsInformationTI);
const { StationStatus } = createCheckers(StationStatusTI);
const { SystemAlerts } = createCheckers(SystemAlertsTI);
const { SystemCalendar } = createCheckers(SystemCalendarTI);
const { SystemHours } = createCheckers(SystemHoursTI);
const { SystemInformation } = createCheckers(SystemInformationTI);
const { SystemPricingPlans } = createCheckers(SystemPricingPlansTI);
const { SystemRegions } = createCheckers(SystemRegionsTI);
const { VehicleTypes } = createCheckers(VehicleTypesTI);

// json test data: these are gbfs with no errors for v2.3
import freeBikeStatusJson from "../../../testFixtures/v2.3/free_bike_status.json";
import gbfsVersionsJson from "../../../testFixtures/v2.3/gbfs_versions.json";
import gbfsJson from "../../../testFixtures/v2.3/gbfs.json";
import geofencingZonesJson from "../../../testFixtures/v2.3/geofencing_zones.json";
import stationInformationJson from "../../../testFixtures/v2.3/station_information.json";
import stationStatusJson from "../../../testFixtures/v2.3/station_status.json";
import systemAlertsJson from "../../../testFixtures/v2.3/system_alerts.json";
import systemCalendarJson from "../../../testFixtures/v2.3/system_calendar.json";
import systemHoursJson from "../../../testFixtures/v2.3/system_hours.json";
import systemInformationJson from "../../../testFixtures/v2.3/system_information.json";
import systemPricingPlansJson from "../../../testFixtures/v2.3/system_pricing_plans.json";
import systemRegionsJson from "../../../testFixtures/v2.3/system_regions.json";
import vehicleTypesJson from "../../../testFixtures/v2.3/vehicle_types.json";

// Date objects cannot be represented in JSON
// Manual checks for dates are required
describe("GBFS Validator v2.3", () => {
  it("should check if gbfs_versions is valid", () => {
    expect(() => {
      GbfsVersions.check(gbfsVersionsJson);
    }).not.toThrow();
  });

  it("should check if gbfs is valid", () => {
    expect(() => {
      Gbfs.check(gbfsJson);
    }).not.toThrow();
  });

  it("should check if geofencing_zones is valid", () => {
    expect(() => {
      GeofencingZones.check(geofencingZonesJson);
    }).not.toThrow();
  });

  it("should check if station_information is valid", () => {
    expect(() => {
      StationInformation.check(stationInformationJson);
    }).not.toThrow();
  });

  it("should check if station_status is valid", () => {
    expect(() => {
      StationStatus.check(stationStatusJson);
    }).not.toThrow();
  });

  it("should check if system_alerts is valid", () => {
    expect(() => {
      SystemAlerts.check(systemAlertsJson);
    }).not.toThrow();
  });

  it("should check if system_information is valid", () => {
    expect(() => {
      SystemCalendar.check(systemCalendarJson);
    }).not.toThrow();
  });

  it("should check if system_information is valid", () => {
    expect(() => {
      SystemHours.check(systemHoursJson);
    }).not.toThrow();
  });

  it("should check if system_information is valid", () => {
    expect(() => {
      SystemInformation.check(systemInformationJson);
    }).not.toThrow();
  });

  it("should check if system_pricing_plans is valid", () => {
    expect(() => {
      SystemPricingPlans.check(systemPricingPlansJson);
    }).not.toThrow();
  });

  it("should check if system_regions is valid", () => {
    expect(() => {
      SystemRegions.check(systemRegionsJson);
    }).not.toThrow();
  });

  it("should check if vehicle_status is valida", () => {
    expect(() => {
      FreeBikeStatus.check(freeBikeStatusJson);
    }).not.toThrow();
  });

  it("should check if vehicle_types is valid", () => {
    expect(() => {
      VehicleTypes.check(vehicleTypesJson);
    }).not.toThrow();
  });
});
