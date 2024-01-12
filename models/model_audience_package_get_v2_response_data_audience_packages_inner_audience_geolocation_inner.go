/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceGeolocationInner struct for AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceGeolocationInner
type AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudienceGeolocationInner struct {
	//
	City *string `json:"city,omitempty"`
	//
	District *string `json:"district,omitempty"`
	//
	Lat *float64 `json:"lat,omitempty"`
	//
	Long *float64 `json:"long,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Province *string `json:"province,omitempty"`
	//
	Radius *float64 `json:"radius,omitempty"`
	//
	Street *string `json:"street,omitempty"`
	//
	StreetNumber *string `json:"street_number,omitempty"`
}