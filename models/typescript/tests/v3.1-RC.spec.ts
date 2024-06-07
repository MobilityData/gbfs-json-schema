import { createCheckers } from "ts-interface-checker";

// checker model
import GbfsVersionsTI from '../v3.1-RC/test-type-checkers/gbfs_versions-ti';
import GbfsTI from '../v3.1-RC/test-type-checkers/gbfs-ti';
import GeofencingZonesTI from '../v3.1-RC/test-type-checkers/geofencing_zones-ti';
import ManifestTI from '../v3.1-RC/test-type-checkers/manifest-ti';
import StationsInformationTI from '../v3.1-RC/test-type-checkers/station_information-ti';
import StationStatusTI from '../v3.1-RC/test-type-checkers/station_status-ti';
import SystemAlertsTI from "../v3.1-RC/test-type-checkers/system_alerts-ti";
import SystemInformationTI from '../v3.1-RC/test-type-checkers/system_information-ti';
import SystemPricingPlansTI from '../v3.1-RC/test-type-checkers/system_pricing_plans-ti'
import SystemRegionsTI  from "../v3.1-RC/test-type-checkers/system_regions-ti";
import VehicleStatusTI from "../v3.1-RC/test-type-checkers/vehicle_status-ti";
import VehicleTypesTI from '../v3.1-RC/test-type-checkers/vehicle_types-ti';

// checkers
const { GbfsVersions } = createCheckers(GbfsVersionsTI);
const { Gbfs } = createCheckers(GbfsTI);
const { GeofencingZones } = createCheckers(GeofencingZonesTI);
const { Manifest } = createCheckers(ManifestTI);
const { StationInformation } = createCheckers(StationsInformationTI);
const { StationStatus } = createCheckers(StationStatusTI);
const { SystemAlerts } = createCheckers(SystemAlertsTI);
const { SystemInformation } = createCheckers(SystemInformationTI);
const { SystemPricingPlans } = createCheckers(SystemPricingPlansTI);
const { SystemRegions } = createCheckers(SystemRegionsTI);
const { VehicleStatus } = createCheckers(VehicleStatusTI);
const { VehicleTypes } = createCheckers(VehicleTypesTI);

// json test data: these are gbfs with no errors for v3.1-RC
import gbfsVersionsJson from '../../../testFixtures/v3.1-RC/gbfs_versions.json';
import gbfsJson from '../../../testFixtures/v3.1-RC/gbfs.json';
import geofencingZonesJson from '../../../testFixtures/v3.1-RC/geofencing_zones.json';
import manifestJson from '../../../testFixtures/v3.1-RC/manifest.json';
import stationInformationJson from '../../../testFixtures/v3.1-RC/station_information.json';
import stationStatusJson from '../../../testFixtures/v3.1-RC/station_status.json';
import systemAlertsJson from '../../../testFixtures/v3.1-RC/system_alerts.json';
import systemInformationJson from '../../../testFixtures/v3.1-RC/system_information.json';
import systemPricingPlansJson from '../../../testFixtures/v3.1-RC/system_pricing_plans.json';
import systemRegionsJson from '../../../testFixtures/v3.1-RC/system_regions.json';
import vehicleStatusJson from '../../../testFixtures/v3.1-RC/vehicle_status.json';
import vehicleTypesJson from '../../../testFixtures/v3.1-RC/vehicle_types.json';

// Date objects cannot be represented in JSON
// Manual checks for dates are required
describe('GBFS Validator v3.1-RC', () => {
    it('should check if gbfs_versions is valid', () => {
        expect(() => {
            GbfsVersions.check(gbfsVersionsJson);
        }).not.toThrow();
    });

    it('should check if gbfs is valid', () => {
        expect(() => {
            Gbfs.check(gbfsJson);
        }).not.toThrow();
    });

    it('should check if geofencing_zones is valid', () => {
        expect(() => {
            GeofencingZones.check(geofencingZonesJson);
        }).not.toThrow();
    });

    it('should check if manifest is valid', () => {
        expect(() => {
            Manifest.check(manifestJson);
        }).not.toThrow();
    });

    it('should check if station_information is valid', () => {
        expect(() => {
            StationInformation.check(stationInformationJson);
        }).not.toThrow();
    });

    it('should check if station_status is valid', () => {
        expect(() => {
            StationStatus.check(stationStatusJson);
        }).not.toThrow();
    });

    it('should check if system_alerts is valid', () => {
        expect(() => {
            SystemAlerts.check(systemAlertsJson);
        }).not.toThrow();
     });

    it('should check if system_information is valid', () => {
        expect(() => {
            SystemInformation.check(systemInformationJson);
        }).not.toThrow();
    });

    it('should check if system_pricing_plans is valid', () => { 
        expect(() => {
            SystemPricingPlans.check(systemPricingPlansJson);
        }).not.toThrow();
    });

    it('should check if system_regions is valid', () => { 
        expect(() => {
            SystemRegions.check(systemRegionsJson);
        }).not.toThrow();
    });

    it('should check if vehicle_status is valida', () => {
        expect(() => {
            VehicleStatus.check(vehicleStatusJson);
        }).not.toThrow();
    });

    it('should check if vehicle_types is valid', () => {
        expect(() => {
            VehicleTypes.check(vehicleTypesJson);
        }).not.toThrow();
    });
});