/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateSiteCreateV2RequestBricksInner struct for ToolsSiteTemplateSiteCreateV2RequestBricksInner
type ToolsSiteTemplateSiteCreateV2RequestBricksInner struct {
	Button *ToolsSiteTemplateSiteCreateV2RequestBricksInnerButton `json:"button,omitempty"`
	Coupon *ToolsSiteTemplateSiteCreateV2RequestBricksInnerCoupon `json:"coupon,omitempty"`
	Form   *ToolsSiteTemplateSiteCreateV2RequestBricksInnerForm   `json:"form,omitempty"`
	// 组件在模板中的位置描述，需和【基于站点创建模板】接口返回的组件index保持一致
	Index        string                                                       `json:"index"`
	Picture      *ToolsSiteTemplateSiteCreateV2RequestBricksInnerPicture      `json:"picture,omitempty"`
	PictureGroup *ToolsSiteTemplateSiteCreateV2RequestBricksInnerPictureGroup `json:"picture_group,omitempty"`
	Text         *ToolsSiteTemplateSiteCreateV2RequestBricksInnerText         `json:"text,omitempty"`
	Type         ToolsSiteTemplateSiteCreateV2BricksType                      `json:"type"`
	Video        *ToolsSiteTemplateSiteCreateV2RequestBricksInnerVideo        `json:"video,omitempty"`
	WechatApplet *ToolsSiteTemplateSiteCreateV2RequestBricksInnerWechatApplet `json:"wechat_applet,omitempty"`
	WechatGame   *ToolsSiteTemplateSiteCreateV2RequestBricksInnerWechatGame   `json:"wechat_game,omitempty"`
}
