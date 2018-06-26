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

type TimeRule struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	ConfigString string `json:"ConfigString,omitempty"`

	ConfigStringStored string `json:"ConfigStringStored,omitempty"`

	DisplayString string `json:"DisplayString,omitempty"`

	EditorType string `json:"EditorType,omitempty"`

	IsConfigured bool `json:"IsConfigured,omitempty"`

	IsInitializing bool `json:"IsInitializing,omitempty"`

	MergeDuplicatedItems bool `json:"MergeDuplicatedItems,omitempty"`

	PlugInName string `json:"PlugInName,omitempty"`

	Links *TimeRuleLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
