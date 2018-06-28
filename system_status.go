// ************************************************************************
// * Copyright 2018 OSIsoft, LLC
// * Licensed under the Apache License, Version 2.0 (the "License");
// * you may not use this file except in compliance with the License.
// * You may obtain a copy of the License at
// * 
// *   <http://www.apache.org/licenses/LICENSE-2.0>
// * 
// * Unless required by applicable law or agreed to in writing, software
// * distributed under the License is distributed on an "AS IS" BASIS,
// * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// * See the License for the specific language governing permissions and
// * limitations under the License.
// ************************************************************************

package gowebapi

import (
	"time"
)

type SystemStatus struct {
	UpTimeInMinutes float64 `json:"UpTimeInMinutes,omitempty"`

	State string `json:"State,omitempty"`

	CacheInstances int32 `json:"CacheInstances,omitempty"`

	ServerTime time.Time `json:"ServerTime,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
