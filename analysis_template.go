
package gowebapi

type AnalysisTemplate struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	AnalysisRulePlugInName string `json:"AnalysisRulePlugInName,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	CreateEnabled bool `json:"CreateEnabled,omitempty"`

	GroupId int32 `json:"GroupId,omitempty"`

	HasNotificationTemplate bool `json:"HasNotificationTemplate,omitempty"`

	HasTarget bool `json:"HasTarget,omitempty"`

	OutputTime string `json:"OutputTime,omitempty"`

	TargetName string `json:"TargetName,omitempty"`

	TimeRulePlugInName string `json:"TimeRulePlugInName,omitempty"`

	Links *AnalysisTemplateLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
