/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeCreateV2RequestAdDataSupplementsInner struct for CreativeCustomCreativeCreateV2RequestAdDataSupplementsInner
type CreativeCustomCreativeCreateV2RequestAdDataSupplementsInner struct {
	//
	Games          []*CreativeCustomCreativeCreateV2RequestAdDataSupplementsInnerGamesInner `json:"games,omitempty"`
	SupplementType *CreativeCustomCreativeCreateV2AdDataSupplementsSupplementType           `json:"supplement_type,omitempty"`
}
