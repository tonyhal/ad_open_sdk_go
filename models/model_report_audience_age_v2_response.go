/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceAgeV2Response struct for ReportAudienceAgeV2Response
type ReportAudienceAgeV2Response struct {
	//
	Code *int64                           `json:"code,omitempty"`
	Data *ReportAudienceAgeV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
