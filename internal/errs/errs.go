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

package errs

import "errors"

var (
	ErrUnsupportedOperation = errors.New("this operation is not supported")
	ErrTaskNotFound         = errors.New("task not found")
	ErrInValidResponse      = errors.New("agent did not return valid response for cancel")
	ErrAuthRequired         = errors.New("authentication required")
	ErrTaskIdMissingMatch   = errors.New("task Id mismatch in agent response")
	ErrBadTaskId            = errors.New("bad task id: task id in request does not match the task object")
	ErrNilMessage           = errors.New("message is nil")
)
