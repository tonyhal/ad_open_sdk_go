/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgSubmitV30Response struct for AdvertiserDeliveryPkgSubmitV30Response
type AdvertiserDeliveryPkgSubmitV30Response struct {
	//
	Code *int64                                      `json:"code,omitempty"`
	Data *AdvertiserDeliveryPkgSubmitV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
