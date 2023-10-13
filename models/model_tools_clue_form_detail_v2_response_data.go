/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueFormDetailV2ResponseData
type ToolsClueFormDetailV2ResponseData struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	Elements []*ToolsClueFormDetailV2ResponseDataElementsInner `json:"elements,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SubTitle *string `json:"sub_title,omitempty"`
	//
	SubmitText *string `json:"submit_text,omitempty"`
	//
	Title *string `json:"title,omitempty"`
}
