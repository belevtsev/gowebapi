/*
 * PI Web API 2018 Swagger Spec
 *
 * Swagger Spec file that describes PI Web API
 *
 * API version: 1.11.0.640
 * Contact: techsupport@osisoft.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gowebapi

type SecurityEntry struct {
	Name string `json:"Name,omitempty"`

	SecurityIdentityName string `json:"SecurityIdentityName,omitempty"`

	AllowRights []string `json:"AllowRights,omitempty"`

	DenyRights []string `json:"DenyRights,omitempty"`

	Links *SecurityEntryLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}