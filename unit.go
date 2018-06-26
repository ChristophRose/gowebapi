
package gowebapi

type Unit struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Abbreviation string `json:"Abbreviation,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	Factor float64 `json:"Factor,omitempty"`

	Offset float64 `json:"Offset,omitempty"`

	ReferenceFactor float64 `json:"ReferenceFactor,omitempty"`

	ReferenceOffset float64 `json:"ReferenceOffset,omitempty"`

	ReferenceUnitAbbreviation string `json:"ReferenceUnitAbbreviation,omitempty"`

	Links *UnitLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
