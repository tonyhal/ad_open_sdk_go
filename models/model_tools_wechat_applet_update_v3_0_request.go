/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletUpdateV30Request struct for ToolsWechatAppletUpdateV30Request
type ToolsWechatAppletUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	GuideText *string `json:"guide_text,omitempty"`
	//
	HeaderImageUrl *string `json:"header_image_url,omitempty"`
	//
	IconImageUrl *string `json:"icon_image_url,omitempty"`
	//
	ImagesHorizontalUrl []string `json:"images_horizontal_url,omitempty"`
	//
	ImagesVerticalUrl []string `json:"images_vertical_url,omitempty"`
	//
	InstanceId int64 `json:"instance_id"`
	//
	Introduction *string `json:"introduction,omitempty"`
	//
	Labels []string `json:"labels,omitempty"`
	//
	RemarkMessage *string `json:"remark_message,omitempty"`
}
