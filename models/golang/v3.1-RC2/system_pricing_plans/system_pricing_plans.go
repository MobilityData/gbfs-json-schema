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
//    systemPricingPlans, err := UnmarshalSystemPricingPlans(bytes)
//    bytes, err = systemPricingPlans.Marshal()

package system_pricing_plans



import "encoding/json"

func UnmarshalSystemPricingPlans(data []byte) (SystemPricingPlans, error) {
	var r SystemPricingPlans
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SystemPricingPlans) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Describes the pricing schemes of the system.
type SystemPricingPlans struct {
	// Array that contains one object per plan as defined below.                                          
	Data                                                                                        Data      `json:"data"`
	// Last time the data in the feed was updated in RFC3339 format.                                      
	LastUpdated                                                                                 string `json:"last_updated"`
	// Number of seconds before the data in the feed will be updated again (0 if the data should          
	// always be refreshed).                                                                              
	TTL                                                                                         int64     `json:"ttl"`
	// GBFS version number to which the feed conforms, according to the versioning framework              
	// (added in v1.1).                                                                                   
	Version                                                                                     Version   `json:"version"`
}

// Array that contains one object per plan as defined below.
type Data struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	// Currency used to pay the fare in ISO 4217 code.                                                        
	Currency                                                                                  string          `json:"currency"`
	// Customer-readable description of the pricing plan.                                                     
	Description                                                                               []Description   `json:"description"`
	// Object defining a capped fare once a price threshold has been spent within a timeframe.                
	// (added in v3.1-RC2)                                                                                    
	FareCapping                                                                               *FareCapping    `json:"fare_capping,omitempty"`
	// Will additional tax be added to the base price?                                                        
	IsTaxable                                                                                 bool            `json:"is_taxable"`
	// Name of this pricing plan.                                                                             
	Name                                                                                      []Name          `json:"name"`
	// Array of segments when the price is a function of distance travelled, displayed in                     
	// kilometers (added in v2.1-RC2).                                                                        
	PerKMPricing                                                                              []PerKMPricing  `json:"per_km_pricing,omitempty"`
	// Array of segments when the price is a function of time travelled, displayed in minutes                 
	// (added in v2.1-RC2).                                                                                   
	PerMinPricing                                                                             []PerMinPricing `json:"per_min_pricing,omitempty"`
	// Identifier of a pricing plan in the system.                                                            
	PlanID                                                                                    string          `json:"plan_id"`
	// Fare price.                                                                                            
	Price                                                                                     float64         `json:"price"`
	// The cost, described as a flat rate, to reserve the vehicle prior to beginning a rental.                
	ReservationPriceFlatRate                                                                  *float64        `json:"reservation_price_flat_rate,omitempty"`
	// The cost, described as per minute rate, to reserve the vehicle prior to beginning a                    
	// rental.                                                                                                
	ReservationPricePerMin                                                                    *float64        `json:"reservation_price_per_min,omitempty"`
	// Is there currently an increase in price in response to increased demand in this pricing                
	// plan? (added in v2.1-RC2)                                                                              
	SurgePricing                                                                              *bool           `json:"surge_pricing,omitempty"`
	// URL where the customer can learn more about this pricing plan.                                         
	URL                                                                                       *string         `json:"url,omitempty"`
}

type Description struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

// Object defining a capped fare once a price threshold has been spent within a timeframe.
// (added in v3.1-RC2)
type FareCapping struct {
	// Amount of time in minutes during which the fare is capped.                                     
	Duration                                                                                  int64   `json:"duration"`
	// The maximum fare threshold for the current timeframe, in the unit specified by currency        
	Price                                                                                     float64 `json:"price"`
}

type Name struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type PerKMPricing struct {
	// The kilometer at which the rate will no longer apply (added in v2.1-RC2).                     
	End                                                                                      *int64  `json:"end,omitempty"`
	// Interval in kilometers at which the rate of this segment is either reapplied                  
	// indefinitely, or if defined, up until (but not including) end kilometer (added in             
	// v2.1-RC2).                                                                                    
	Interval                                                                                 int64   `json:"interval"`
	// Rate that is charged for each kilometer interval after the start (added in v2.1-RC2).         
	Rate                                                                                     float64 `json:"rate"`
	// Number of kilometers that have to elapse before this segment starts applying (added in        
	// v2.1-RC2).                                                                                    
	Start                                                                                    int64   `json:"start"`
}

type PerMinPricing struct {
	// The minute at which the rate will no longer apply (added in v2.1-RC2).                     
	End                                                                                   *int64  `json:"end,omitempty"`
	// Interval in minutes at which the rate of this segment is either reapplied (added in        
	// v2.1-RC2).                                                                                 
	Interval                                                                              int64   `json:"interval"`
	// Rate that is charged for each minute interval after the start (added in v2.1-RC2).         
	Rate                                                                                  float64 `json:"rate"`
	// Number of minutes that have to elapse before this segment starts applying (added in        
	// v2.1-RC2).                                                                                 
	Start                                                                                 int64   `json:"start"`
}

type Version string

const (
	The31Rc2 Version = "3.1-RC2"
)
