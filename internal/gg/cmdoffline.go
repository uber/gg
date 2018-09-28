// Copyright (c) 2018 Uber Technologies, Inc.
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
// Copyright (c) 2018 Uber Technologies, Inc.
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

package gg

import "context"

const offlineUsage UsageError = `Usage: gg offline/off online/on
Example: gg off up write

Enables or disables offline mode.  When gg is offline, it will not fetch
upgrades for dependency git repositories.  Some commands will not work if they
cannot find a suitable version of a dependency in the cache.
`

func offlineCommand() Command {
	return Command{
		Names: []string{
			"offline",
			"off",
		},
		Usage: offlineUsage,
		Niladic: func(ctx context.Context, driver *Driver) error {
			driver.memo.Offline = true
			return nil
		},
	}
}

func onlineCommand() Command {
	return Command{
		Names: []string{
			"online",
			"on",
		},
		Usage: offlineUsage,
		Niladic: func(ctx context.Context, driver *Driver) error {
			driver.memo.Offline = false
			return nil
		},
	}
}
