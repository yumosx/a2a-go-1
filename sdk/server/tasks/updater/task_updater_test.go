// Copyright 2025 yeeaiclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package updater

import (
	"testing"

	"github.com/yeeaiclub/a2a-go/sdk/server/event"
	"github.com/yeeaiclub/a2a-go/sdk/types"
)

func TestUpdateStatus(t *testing.T) {
	testcases := []struct {
		name      string
		taskId    string
		contextId string
		state     types.TaskState
	}{
		{
			name: "update status",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			queue := event.NewQueue(0)
			updater := NewTaskUpdater(queue, tc.taskId, tc.contextId)
			updater.UpdateStatus(tc.state)
		})
	}
}
