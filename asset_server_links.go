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

type AssetServerLinks struct {
	Self string `json:"Self,omitempty"`

	Databases string `json:"Databases,omitempty"`

	NotificationContactTemplates string `json:"NotificationContactTemplates,omitempty"`

	SecurityIdentities string `json:"SecurityIdentities,omitempty"`

	SecurityMappings string `json:"SecurityMappings,omitempty"`

	UnitClasses string `json:"UnitClasses,omitempty"`

	AnalysisRulePlugIns string `json:"AnalysisRulePlugIns,omitempty"`

	TimeRulePlugIns string `json:"TimeRulePlugIns,omitempty"`

	Security string `json:"Security,omitempty"`

	SecurityEntries string `json:"SecurityEntries,omitempty"`
}
