/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteUpdateV2RequestBricksInnerSetting
type ToolsSiteUpdateV2RequestBricksInnerSetting struct {
	Avatar     *ToolsSiteUpdateV2RequestBricksInnerSettingAvatar     `json:"avatar,omitempty"`
	Background *ToolsSiteUpdateV2RequestBricksInnerSettingBackground `json:"background,omitempty"`
	Button     *ToolsSiteUpdateV2RequestBricksInnerSettingButton     `json:"button,omitempty"`
	Image      *ToolsSiteUpdateV2RequestBricksInnerSettingImage      `json:"image,omitempty"`
	//
	IntroduceType *string `json:"introduce_type,omitempty"`
	//
	Items []string                                         `json:"items,omitempty"`
	Label *ToolsSiteUpdateV2RequestBricksInnerSettingLabel `json:"label,omitempty"`
	//
	StyleType *int64                                           `json:"style_type,omitempty"`
	Title     *ToolsSiteUpdateV2RequestBricksInnerSettingTitle `json:"title,omitempty"`
}
