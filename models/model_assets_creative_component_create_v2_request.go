/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AssetsCreativeComponentCreateV2Request struct for AssetsCreativeComponentCreateV2Request
type AssetsCreativeComponentCreateV2Request struct {
	//
	AdvertiserId  int64                                               `json:"advertiser_id"`
	ComponentInfo AssetsCreativeComponentCreateV2RequestComponentInfo `json:"component_info"`
}
