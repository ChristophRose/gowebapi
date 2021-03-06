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

type ElementTemplate struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	AllowElementToExtend bool `json:"AllowElementToExtend,omitempty"`

	BaseTemplate string `json:"BaseTemplate,omitempty"`

	InstanceType string `json:"InstanceType,omitempty"`

	NamingPattern string `json:"NamingPattern,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	ExtendedProperties map[string]Value `json:"ExtendedProperties,omitempty"`

	Severity string `json:"Severity,omitempty"`

	CanBeAcknowledged bool `json:"CanBeAcknowledged,omitempty"`

	Links *ElementTemplateLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
