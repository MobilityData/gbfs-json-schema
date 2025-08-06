package vehicleavailability


import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/header"
)

// Describes the capacity and rental availability of the station
type VehicleAvailability struct {
	header.HeaderStruct

	// Array that contains one object per station as defined below.
	Data Data `json:"data"`
}

// Array that contains one object per station as defined below.
type Data struct {
	Vehicles []Vehicle `json:"vehicles"`
}

type Vehicle struct {
	// Identifier of a vehicle. The vehicle_id identifier MUST be rotated to a random string after each trip to protect user privacy (as of v2.0).
	// Use of persistent vehicle IDs poses a threat to user privacy. The vehicle_id identifier SHOULD only be rotated once per trip.
	// The vehicle_id SHOULD be the same as in vehicle_status.json if the file has been defined and the vehicle is currently available.
	VehicleID common.ID `json:"vehicle_id"`
	// Unique identifier of a vehicle type as defined in vehicle_types.json.
	VehicleTypeID common.ID `json:"vehicle_type_id"`
	// The station_id of the station where this vehicle is located when available as defined in station_information.json.
	StationID common.ID `json:"station_id"`
	// The plan_id of the pricing plan this vehicle is eligible for as described in system_pricing_plans.json. 
	// If this field is defined it supersedes default_pricing_plan_id in vehicle_types.json. 
	// This field SHOULD be used to override default_pricing_plan_id in vehicle_types.json to define pricing plans for individual vehicles when necessary.
	PricingPlanID *common.ID `json:"pricing_plan_id,omitempty"`
	// List of vehicle equipment provided by the operator in addition to the accessories already provided in the vehicle 
	// (field vehicle_accessories of vehicle_types.json) but subject to more frequent updates.
	VehicleEquipment []*common.VehicleEquipment `json:"vehicle_equipment,omitempty"`
	// Array of time slots during which the specified vehicle is available.
	Availabilities []Availability `json:"availabilities"`
}

type Availability struct {
	// Start date and time of available time slot.
	From common.DateTime `json:"from"`
	// End date and time of available time slot. If this field is empty, it means that the vehicle is available all the time from the date in the from field.
	Until *common.DateTime `json:"until,omitempty"`
}
