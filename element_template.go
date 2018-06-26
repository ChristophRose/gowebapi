
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
