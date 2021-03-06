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

type EventFrameLinks struct {
	Self string `json:"Self,omitempty"`

	Attributes string `json:"Attributes,omitempty"`

	EventFrames string `json:"EventFrames,omitempty"`

	Database string `json:"Database,omitempty"`

	ReferencedElements string `json:"ReferencedElements,omitempty"`

	PrimaryReferencedElement string `json:"PrimaryReferencedElement,omitempty"`

	Parent string `json:"Parent,omitempty"`

	Template string `json:"Template,omitempty"`

	DefaultAttribute string `json:"DefaultAttribute,omitempty"`

	Categories string `json:"Categories,omitempty"`

	Annotations string `json:"Annotations,omitempty"`

	InterpolatedData string `json:"InterpolatedData,omitempty"`

	RecordedData string `json:"RecordedData,omitempty"`

	PlotData string `json:"PlotData,omitempty"`

	SummaryData string `json:"SummaryData,omitempty"`

	Value string `json:"Value,omitempty"`

	EndValue string `json:"EndValue,omitempty"`

	Security string `json:"Security,omitempty"`

	SecurityEntries string `json:"SecurityEntries,omitempty"`
}
