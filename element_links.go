/*
 * PI Web API 2018 Swagger Spec
 *
 * Swagger Spec file that describes PI Web API
 *
 * API version: 1.11.0.640
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gowebapi

type ElementLinks struct {
	Self string `json:"Self,omitempty"`

	Analyses string `json:"Analyses,omitempty"`

	Attributes string `json:"Attributes,omitempty"`

	Elements string `json:"Elements,omitempty"`

	Database string `json:"Database,omitempty"`

	Parent string `json:"Parent,omitempty"`

	Template string `json:"Template,omitempty"`

	Categories string `json:"Categories,omitempty"`

	DefaultAttribute string `json:"DefaultAttribute,omitempty"`

	EventFrames string `json:"EventFrames,omitempty"`

	InterpolatedData string `json:"InterpolatedData,omitempty"`

	RecordedData string `json:"RecordedData,omitempty"`

	PlotData string `json:"PlotData,omitempty"`

	SummaryData string `json:"SummaryData,omitempty"`

	Value string `json:"Value,omitempty"`

	EndValue string `json:"EndValue,omitempty"`

	Security string `json:"Security,omitempty"`

	SecurityEntries string `json:"SecurityEntries,omitempty"`

	NotificationRules string `json:"NotificationRules,omitempty"`
}