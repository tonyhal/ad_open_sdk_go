/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeCreateV2RequestAdDataMiniProgramInfo
type CreativeCustomCreativeCreateV2RequestAdDataMiniProgramInfo struct {
	//
	AppId string `json:"app_id"`
	//
	Params *string `json:"params,omitempty"`
	//
	StartPath *string                                                 `json:"start_path,omitempty"`
	Type      CreativeCustomCreativeCreateV2AdDataMiniProgramInfoType `json:"type"`
	//
	Url *string `json:"url,omitempty"`
}
