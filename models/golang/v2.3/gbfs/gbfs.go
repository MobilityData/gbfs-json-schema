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

package gbfs

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Auto-discovery file that links to all of the other files published by the system.
type Gbfs struct {
	header.HeaderStruct
	// Response data in the form of name:value pairs.
	Data Data `json:"data"`
}

// The language that will be used throughout the rest of the files. It MUST match the value in the system_information.json file.
type Data map[common.Language]LanguageFeeds

type LanguageFeeds struct {
	// An array of all of the feeds that are published by this auto-discovery file. Each element in the array is an object with the keys below.
	Feeds []common.Feed `json:"feeds"`
}
