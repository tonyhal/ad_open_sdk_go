/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceCityV2ResponseData
type ReportAudienceCityV2ResponseData struct {
	//
	CityName    *string                                     `json:"city_name,omitempty"`
	MetricsDict *ReportAudienceAgeV2ResponseDataMetricsDict `json:"metrics_dict,omitempty"`
	//
	ProvinceName *string `json:"province_name,omitempty"`
}
