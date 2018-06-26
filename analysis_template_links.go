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

type AnalysisTemplateLinks struct {
	Self string `json:"Self,omitempty"`

	Database string `json:"Database,omitempty"`

	Categories string `json:"Categories,omitempty"`

	AnalysisRule string `json:"AnalysisRule,omitempty"`

	AnalysisRulePlugIn string `json:"AnalysisRulePlugIn,omitempty"`

	TimeRule string `json:"TimeRule,omitempty"`

	TimeRulePlugIn string `json:"TimeRulePlugIn,omitempty"`

	Target string `json:"Target,omitempty"`

	Security string `json:"Security,omitempty"`

	SecurityEntries string `json:"SecurityEntries,omitempty"`
}
