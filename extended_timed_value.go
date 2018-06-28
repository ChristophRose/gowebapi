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

type ExtendedTimedValue struct {
	Annotations []StreamAnnotation `json:"Annotations,omitempty"`

	Timestamp time.Time `json:"Timestamp,omitempty"`

	UnitsAbbreviation string `json:"UnitsAbbreviation,omitempty"`

	Good bool `json:"Good,omitempty"`

	Questionable bool `json:"Questionable,omitempty"`

	Substituted bool `json:"Substituted,omitempty"`

	Annotated bool `json:"Annotated,omitempty"`

	Value *interface{} `json:"Value,omitempty"`

	Errors []PropertyError `json:"Errors,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
