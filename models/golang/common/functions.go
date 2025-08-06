package common

import "encoding/json"


func IsValidFeedName(name string) bool {
	switch FeedName(name) {
	case Feed_FreeBikeStatus, Feed_GbfsVersions, Feed_GeofencingZones, Feed_Gbfs, Feed_StationInformation, Feed_StationStatus, Feed_SystemAlerts, Feed_SystemCalendar, Feed_SystemHours, Feed_SystemInformation, Feed_SystemPricingPlans, Feed_SystemRegions, Feed_VehicleStatus, Feed_VehicleTypes, Feed_Manifest:
		return true
	}
	return false
}

func UnmarshalVersion(data []byte) (VersionStruct, error) {
	var r VersionStruct
	err := json.Unmarshal(data, &r)
	return r, err
}