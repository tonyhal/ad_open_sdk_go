/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignGetV30Response struct for BrandCampaignGetV30Response
type BrandCampaignGetV30Response struct {
	//
	Code *int64                           `json:"code,omitempty"`
	Data *BrandCampaignGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
