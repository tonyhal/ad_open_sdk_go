/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserPublicInfoV2Response struct for AdvertiserPublicInfoV2Response
type AdvertiserPublicInfoV2Response struct {
	//
	Code *int64 `json:"code,omitempty"`
	//
	Data []*AdvertiserPublicInfoV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
