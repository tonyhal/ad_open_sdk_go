/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30RequestAudienceGeolocationInner struct for ProjectUpdateV30RequestAudienceGeolocationInner
type ProjectUpdateV30RequestAudienceGeolocationInner struct {
	//
	Lat *float64 `json:"lat,omitempty"`
	//
	Long *float64 `json:"long,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Radius *int64 `json:"radius,omitempty"`
}
