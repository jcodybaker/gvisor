// Copyright 2018 The containerd Authors.
// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package proc contains process-related utilities.
package proc

import (
	"fmt"
)

// RunscRoot is the path to the root runsc state directory.
const RunscRoot = "/run/containerd/runsc"

func stateName(v interface{}) string {
	switch v.(type) {
	case *runningState, *execRunningState:
		return "running"
	case *createdState, *execCreatedState:
		return "created"
	case *deletedState:
		return "deleted"
	case *stoppedState:
		return "stopped"
	}
	panic(fmt.Errorf("invalid state %v", v))
}
