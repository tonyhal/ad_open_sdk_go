/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceGroupGetV10Response struct for QianchuanAudienceGroupGetV10Response
type QianchuanAudienceGroupGetV10Response struct {
	//
	Code *int64                                    `json:"code,omitempty"`
	Data *QianchuanAudienceGroupGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
