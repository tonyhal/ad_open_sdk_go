/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmUpdateChallengeV2Response struct for StarDemandOmUpdateChallengeV2Response
type StarDemandOmUpdateChallengeV2Response struct {
	//
	Code *int64                                     `json:"code,omitempty"`
	Data *StarDemandOmUpdateChallengeV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
