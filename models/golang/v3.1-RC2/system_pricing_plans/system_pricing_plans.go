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
	system_pricing_plans_v31RC "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/system_pricing_plans"
)

// Describes the pricing schemes of the system.
type SystemPricingPlans struct {
	system_pricing_plans_v31RC.SystemPricingPlans
	// Array that contains one object per plan as defined below.
	Data Data `json:"data"`
}

// Array that contains one object per plan as defined below.
type Data struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	system_pricing_plans_v31RC.Plan
	// Object defining a capped fare once a price threshold has been spent within a timeframe.
	// The same fare cap applies to each subsequent timeframe. For example, a fare capped at 15.00 CAD per 12-hour period.
	Fare_capping *FareCapping `json:"fare_capping,omitempty"`
}

type FareCapping struct {
	// Amount of time in minutes during which the fare is capped.
	Duration uint64 `json:"duration"`
	// The maximum fare threshold for the current timeframe, in the unit specified by currency.
	Price float64 `json:"price"`
}
