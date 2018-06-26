
package gowebapi

import (
	"time"
)

type AnalysisRulePlugIn struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	AssemblyFileName string `json:"AssemblyFileName,omitempty"`

	AssemblyID string `json:"AssemblyID,omitempty"`

	AssemblyLoadProperties []string `json:"AssemblyLoadProperties,omitempty"`

	AssemblyTime time.Time `json:"AssemblyTime,omitempty"`

	CompatibilityVersion int32 `json:"CompatibilityVersion,omitempty"`

	IsBrowsable bool `json:"IsBrowsable,omitempty"`

	IsNonEditableConfig bool `json:"IsNonEditableConfig,omitempty"`

	LoadedAssemblyTime time.Time `json:"LoadedAssemblyTime,omitempty"`

	LoadedVersion string `json:"LoadedVersion,omitempty"`

	Version string `json:"Version,omitempty"`

	Links *AnalysisRulePlugInLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
