
package gowebapi

type Element struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	TemplateName string `json:"TemplateName,omitempty"`

	HasChildren bool `json:"HasChildren,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	ExtendedProperties map[string]Value `json:"ExtendedProperties,omitempty"`

	Links *ElementLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
