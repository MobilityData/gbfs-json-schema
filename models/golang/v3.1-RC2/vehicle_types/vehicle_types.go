// Copyright 2025 MobilityData
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    vehicleTypes, err := UnmarshalVehicleTypes(bytes)
//    bytes, err = vehicleTypes.Marshal()

package vehicle_types



import "encoding/json"

func UnmarshalVehicleTypes(data []byte) (VehicleTypes, error) {
	var r VehicleTypes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VehicleTypes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Describes the types of vehicles that System operator has available for rent (added in
// v2.1-RC).
type VehicleTypes struct {
	// Response data in the form of name:value pairs.                                                     
	Data                                                                                        Data      `json:"data"`
	// Last time the data in the feed was updated in RFC3339 format.                                      
	LastUpdated                                                                                 string `json:"last_updated"`
	// Number of seconds before the data in the feed will be updated again (0 if the data should          
	// always be refreshed).                                                                              
	TTL                                                                                         int64     `json:"ttl"`
	// GBFS version number to which the feed conforms, according to the versioning framework.             
	Version                                                                                     Version   `json:"version"`
}

// Response data in the form of name:value pairs.
type Data struct {
	// Array that contains one object per vehicle type in the system as defined below.              
	VehicleTypes                                                                      []VehicleType `json:"vehicle_types"`
}

type VehicleType struct {
	// The capacity of the vehicle cargo space (excluding passengers), expressed in kilograms.                     
	CargoLoadCapacity                                                                           *int64             `json:"cargo_load_capacity,omitempty"`
	// Cargo volume available in the vehicle, expressed in liters.                                                 
	CargoVolumeCapacity                                                                         *int64             `json:"cargo_volume_capacity,omitempty"`
	// The color of the vehicle. Added in v2.3                                                                     
	Color                                                                                       *string            `json:"color,omitempty"`
	// A plan_id as defined in system_pricing_plans.json added in v2.3-RC.                                         
	DefaultPricingPlanID                                                                        *string            `json:"default_pricing_plan_id,omitempty"`
	// Maximum time in minutes that a vehicle can be reserved before a rental begins added in                      
	// v2.3-RC.                                                                                                    
	DefaultReserveTime                                                                          *int64             `json:"default_reserve_time,omitempty"`
	// Customer-readable description of the vehicle type outlining special features or how-tos.                    
	// An array with one object per supported language with the following keys:                                    
	Description                                                                                 []Description      `json:"description,omitempty"`
	// Vehicle air quality certificate. added in v2.3.                                                             
	EcoLabels                                                                                   []EcoLabel         `json:"eco_labels,omitempty"`
	// The vehicle's general form factor.                                                                          
	FormFactor                                                                                  FormFactor         `json:"form_factor"`
	// Maximum quantity of CO2, in grams, emitted per kilometer, according to the WLTP. Added in                   
	// v2.3                                                                                                        
	GCO2KM                                                                                      *int64             `json:"g_CO2_km,omitempty"`
	// The name of the vehicle manufacturer. An array with one object per supported language                       
	// with the following keys:                                                                                    
	Make                                                                                        []Make             `json:"make,omitempty"`
	// The maximum speed in kilometers per hour this vehicle is permitted to reach in accordance                   
	// with local permit and regulations. Added in v2.3                                                            
	MaxPermittedSpeed                                                                           *int64             `json:"max_permitted_speed,omitempty"`
	// The furthest distance in meters that the vehicle can travel without recharging or                           
	// refueling when it has the maximum amount of energy potential.                                               
	MaxRangeMeters                                                                              *float64           `json:"max_range_meters,omitempty"`
	// The name of the vehicle model. An array with one object per supported language with the                     
	// following keys:                                                                                             
	Model                                                                                       []Model            `json:"model,omitempty"`
	// The public name of this vehicle type. An array with one object per supported language                       
	// with the following keys:                                                                                    
	Name                                                                                        []Name             `json:"name,omitempty"`
	// Array of all pricing plan IDs as defined in system_pricing_plans.json added in v2.3-RC.                     
	PricingPlanIDS                                                                              []string           `json:"pricing_plan_ids,omitempty"`
	// The primary propulsion type of the vehicle. Updated in v2.3 to represent car-sharing                        
	PropulsionType                                                                              PropulsionType     `json:"propulsion_type"`
	// The rated power of the motor for this vehicle type in watts. Added in v2.3                                  
	RatedPower                                                                                  *int64             `json:"rated_power,omitempty"`
	// The conditions for returning the vehicle at the end of the trip. Added in v2.3-RC as                        
	// return_type, and updated to return_constraint in v2.3.                                                      
	ReturnConstraint                                                                            *ReturnConstraint  `json:"return_constraint,omitempty"`
	// The number of riders (driver included) the vehicle can legally accommodate                                  
	RiderCapacity                                                                               *int64             `json:"rider_capacity,omitempty"`
	// Description of accessories available in the vehicle.                                                        
	VehicleAccessories                                                                          []VehicleAccessory `json:"vehicle_accessories,omitempty"`
	// An object where each key defines one of the items listed below added in v2.3-RC.                            
	VehicleAssets                                                                               *VehicleAssets     `json:"vehicle_assets,omitempty"`
	// URL to an image that would assist the user in identifying the vehicle. JPEG or PNG. Added                   
	// in v2.3                                                                                                     
	VehicleImage                                                                                *string            `json:"vehicle_image,omitempty"`
	// Unique identifier of a vehicle type.                                                                        
	VehicleTypeID                                                                               string             `json:"vehicle_type_id"`
	// Number of wheels this vehicle type has. Added in v2.3                                                       
	WheelCount                                                                                  *int64             `json:"wheel_count,omitempty"`
}

type Description struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type EcoLabel struct {
	// Country code following the ISO 3166-1 alpha-2 notation. Added in v2.3.       
	CountryCode                                                              string `json:"country_code"`
	// Name of the eco label. Added in v2.3.                                        
	EcoSticker                                                               string `json:"eco_sticker"`
}

type Make struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type Model struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type Name struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

// An object where each key defines one of the items listed below added in v2.3-RC.
type VehicleAssets struct {
	// Date that indicates the last time any included vehicle icon images were modified or              
	// updated added in v2.3-RC.                                                                        
	IconLastModified                                                                            string  `json:"icon_last_modified"`
	// A fully qualified URL pointing to the location of a graphic icon file that MAY be used to        
	// represent this vehicle type on maps and in other applications added in v2.3-RC.                  
	IconURL                                                                                     string  `json:"icon_url"`
	// A fully qualified URL pointing to the location of a graphic icon file to be used to              
	// represent this vehicle type when in dark mode added in v2.3-RC.                                  
	IconURLDark                                                                                 *string `json:"icon_url_dark,omitempty"`
}

// The vehicle's general form factor.
type FormFactor string

const (
	Bicycle         FormFactor = "bicycle"
	Car             FormFactor = "car"
	CargoBicycle    FormFactor = "cargo_bicycle"
	Moped           FormFactor = "moped"
	Other           FormFactor = "other"
	Scooter         FormFactor = "scooter"
	ScooterSeated   FormFactor = "scooter_seated"
	ScooterStanding FormFactor = "scooter_standing"
)

// The primary propulsion type of the vehicle. Updated in v2.3 to represent car-sharing
type PropulsionType string

const (
	Combustion           PropulsionType = "combustion"
	CombustionDiesel     PropulsionType = "combustion_diesel"
	Electric             PropulsionType = "electric"
	ElectricAssist       PropulsionType = "electric_assist"
	Human                PropulsionType = "human"
	HydrogenFuelCell     PropulsionType = "hydrogen_fuel_cell"
	PlugInHybrid         PropulsionType = "plug_in_hybrid"
	PropulsionTypeHybrid PropulsionType = "hybrid"
)

// The conditions for returning the vehicle at the end of the trip. Added in v2.3-RC as
// return_type, and updated to return_constraint in v2.3.
type ReturnConstraint string

const (
	AnyStation             ReturnConstraint = "any_station"
	FreeFloating           ReturnConstraint = "free_floating"
	ReturnConstraintHybrid ReturnConstraint = "hybrid"
	RoundtripStation       ReturnConstraint = "roundtrip_station"
)

type VehicleAccessory string

const (
	AirConditioning VehicleAccessory = "air_conditioning"
	Automatic       VehicleAccessory = "automatic"
	Convertible     VehicleAccessory = "convertible"
	CruiseControl   VehicleAccessory = "cruise_control"
	Doors2          VehicleAccessory = "doors_2"
	Doors3          VehicleAccessory = "doors_3"
	Doors4          VehicleAccessory = "doors_4"
	Doors5          VehicleAccessory = "doors_5"
	Manual          VehicleAccessory = "manual"
	Navigation      VehicleAccessory = "navigation"
)

type Version string

const (
	The31Rc2 Version = "3.1-RC2"
)
