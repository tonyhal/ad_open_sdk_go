/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButton 按钮组件描述
type ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButton struct {
	AppointEvent *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButtonAppointEvent `json:"appoint_event,omitempty"`
	// linkEvent事件自定义图片链接
	BgImageUrl    *string                                                                    `json:"bg_image_url,omitempty"`
	DownloadEvent *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButtonDownloadEvent `json:"download_event,omitempty"`
	EventType     ToolsSiteTemplateGetV2DataListBricksButtonEventType                        `json:"event_type"`
	LinkEvent     *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButtonLinkEvent     `json:"link_event,omitempty"`
	PhoneEvent    *ToolsSiteTemplateGetV2ResponseDataListInnerBricksInnerButtonPhoneEvent    `json:"phone_event,omitempty"`
}
