// Copyright 2019-present Open Networking Foundation.
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

package store

import (
	"github.com/onosproject/onos-config/pkg/store/change"
	"sort"
	"strings"
	"time"
)

// ConfigName is an alias for string - is used to qualify identifier for Configuration
type ConfigName string

// Configuration is the connection between a device and Change objects
// The set of ChangeIds define it's content
type Configuration struct {
	Name        ConfigName
	Device      string
	Created     time.Time
	Updated     time.Time
	User        string
	Description string
	Changes     []change.ID
}

// ExtractFullConfig retrieves the full consolidated config for a Configuration
// This gets the change up to and including the latest
// Use "nBack" to specify a number of changes back to go
// If there are not as many changes in the history as nBack nothing is returned
func (b Configuration) ExtractFullConfig(changeStore map[string]*change.Change, nBack int) []change.ConfigValue {

	// Have to use a slice to have a consistent output order
	consolidatedConfig := make([]change.ConfigValue, 0)

	for _, changeID := range b.Changes[0 : len(b.Changes)-nBack] {
		change := changeStore[B64(changeID)]
		for _, changeValue := range change.Config {
			if changeValue.Remove {
				// Delete everything at that path and all below it
				// Have to search through consolidated config
				// Make a list of indices to remove
				indices := make([]int, 0)
				for idx, cce := range consolidatedConfig {
					if strings.Contains(cce.Path, changeValue.Path) {
						indices = append(indices, idx)
					}
				}
				// Remove in reverse
				for i := len(indices) - 1; i >= 0; i-- {
					consolidatedConfig = append(consolidatedConfig[:indices[i]], consolidatedConfig[indices[i]+1:]...)
				}

			} else {
				var alreadyExists bool
				for idx, cv := range consolidatedConfig {
					if changeValue.Path == cv.Path {
						consolidatedConfig[idx].Value = changeValue.Value
						alreadyExists = true
						break
					}
				}
				if !alreadyExists {
					consolidatedConfig = append(consolidatedConfig, changeValue.ConfigValue)
				}
			}
		}
	}

	sort.Slice(consolidatedConfig, func(i, j int) bool {
		return consolidatedConfig[i].Path < consolidatedConfig[j].Path
	})

	return consolidatedConfig
}
