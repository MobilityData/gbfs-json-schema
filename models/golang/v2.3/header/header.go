// Copyright 2024 MobilityData
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package header

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
)

type HeaderStruct struct {
	// Last time the data in the feed was updated in POSIX time (epoch seconds).
	LastUpdated common.PosixTimestamp `json:"last_updated"`
	// Number of seconds before the data in the feed will be updated again (0 if the data should always be refreshed).
	TTL uint64 `json:"ttl"`
	// GBFS version number to which the feed conforms, according to the versioning framework
	// (added in v1.1).
	Version common.Version `json:"version"`
}
