/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceGenderV2ResponseData
type ReportAudienceGenderV2ResponseData struct {
	//
	GenderName  *string                                        `json:"gender_name,omitempty"`
	MetricsDict *ReportAudienceGenderV2ResponseDataMetricsDict `json:"metrics_dict,omitempty"`
}
