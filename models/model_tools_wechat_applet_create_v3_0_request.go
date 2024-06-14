/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletCreateV30Request struct for ToolsWechatAppletCreateV30Request
type ToolsWechatAppletCreateV30Request struct {
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
	Introduction *string `json:"introduction,omitempty"`
	//
	Labels []string `json:"labels,omitempty"`
	//
	Name string `json:"name"`
	//
	Path *string `json:"path,omitempty"`
	//
	RemarkMessage *string `json:"remark_message,omitempty"`
	//
	UserName string `json:"user_name"`
}
