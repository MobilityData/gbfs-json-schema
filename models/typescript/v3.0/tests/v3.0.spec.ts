import { createCheckers } from "ts-interface-checker";

// checker model
import GbfsVersionsTI from './ti/gbfs_versions-ti';
import GbfsTI from './ti/gbfs-ti';
import GeofencingZonesTI from './ti/geofencing_zones-ti';
import ManifestTI from './ti/manifest-ti';
import StationsInformationTI from './ti/station_information-ti';
import StationStatusTI from './ti/station_status-ti';
import SystemAlertsTI from "./ti/system_alerts-ti";
import SystemInformationTI from './ti/system_information-ti';
import SystemPricingPlansTI from './ti/system_pricing_plans-ti'
import SystemRegionsTI  from "./ti/system_regions-ti";
import VehicleStatusTI from "./ti/vehicle_status-ti";
import VehicleTypesTI from './ti/vehicle_types-ti';

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

// json test data: these are gbfs with no errors for v3.0
import gbfsVersionsJson from './data/gbfs_versions.json';
import gbfsJson from './data/gbfs.json';
import geofencingZonesJson from './data/geofencing_zones.json';
import manifestJson from './data/manifest.json';
import stationInformationJson from './data/station_information.json';
import stationStatusJson from './data/station_status.json';
import systemAlertsJson from './data/system_alerts.json';
import systemInformationJson from './data/system_information.json';
import systemPricingPlansJson from './data/system_pricing_plans.json';
import systemRegionsJson from './data/system_regions.json';
import vehicleStatusJson from './data/vehicle_status.json';
import vehicleTypesJson from './data/vehicle_types.json';

// Date objects cannot be represented in JSON
// Manual checks for dates are required
describe('GBFS Validator v3.0', () => {
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