/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsUnionFlowPackageCreateV2Request struct for ToolsUnionFlowPackageCreateV2Request
type ToolsUnionFlowPackageCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Name string `json:"name"`
	//
	Rit []int64 `json:"rit"`
}