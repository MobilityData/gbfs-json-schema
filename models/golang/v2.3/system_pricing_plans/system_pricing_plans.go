// Copyright 2024 MobilityData
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

package system_pricing_plans

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Describes the pricing schemes of the system.
type SystemPricingPlans struct {
	header.HeaderStruct
	// Array that contains one object per plan as defined below.
	Data Data `json:"data"`
}

// Array that contains one object per plan as defined below.
type Data struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	// Currency used to pay the fare in ISO 4217 code.
	Currency string `json:"currency"`
	// Customer-readable description of the pricing plan.
	Description string `json:"description"`
	// Will additional tax be added to the base price?
	IsTaxable bool `json:"is_taxable"`
	// Name of this pricing plan.
	Name string `json:"name"`
	// Array of segments when the price is a function of distance travelled, displayed in
	// kilometers (added in v2.2)
	PerKMPricing []common.PerKMPricing `json:"per_km_pricing,omitempty"`
	// Array of segments when the price is a function of time travelled, displayed in minutes
	// (added in v2.2)
	PerMinPricing []common.PerMinPricing `json:"per_min_pricing,omitempty"`
	// Identifier of a pricing plan in the system.
	PlanID common.ID `json:"plan_id"`
	// Fare price. The specification states that it can either be a float OR a string. As this is not
	// compatible with the Go language, we use a float64 type here as thios is the only allowed type
	// starting from v3.0.
	Price float64 `json:"price"`
	// Is there currently an increase in price in response to increased demand in this pricing plan?
	// If this field is empty, it means there is no surge pricing in effect. (added in v2.2)
	// true - Surge pricing is in effect.
	// false - Surge pricing is not in effect.
	SurgePricing *bool `json:"surge_pricing,omitempty"`
	// URL where the customer can learn more about this pricing plan.
	URL *common.URL `json:"url,omitempty"`
}
