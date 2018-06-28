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

type NotificationRule struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	AutoCreated bool `json:"AutoCreated,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	Criteria string `json:"Criteria,omitempty"`

	MultiTriggerEventOption string `json:"MultiTriggerEventOption,omitempty"`

	NonrepetitionInterval string `json:"NonrepetitionInterval,omitempty"`

	ResendInterval string `json:"ResendInterval,omitempty"`

	Status string `json:"Status,omitempty"`

	TemplateName string `json:"TemplateName,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}