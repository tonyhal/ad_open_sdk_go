/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseDataBricksInner struct for ToolsSiteTemplateCreateV2ResponseDataBricksInner
type ToolsSiteTemplateCreateV2ResponseDataBricksInner struct {
	Button *ToolsSiteTemplateCreateV2ResponseDataBricksInnerButton `json:"button,omitempty"`
	Coupon *ToolsSiteTemplateCreateV2ResponseDataBricksInnerCoupon `json:"coupon,omitempty"`
	Form   *ToolsSiteTemplateCreateV2ResponseDataBricksInnerForm   `json:"form,omitempty"`
	// 组件在模板中的位置描述
	Index        string                                                        `json:"index"`
	Picture      *ToolsSiteTemplateCreateV2ResponseDataBricksInnerPicture      `json:"picture,omitempty"`
	PictureGroup *ToolsSiteTemplateCreateV2ResponseDataBricksInnerPictureGroup `json:"picture_group,omitempty"`
	Text         *ToolsSiteTemplateCreateV2ResponseDataBricksInnerText         `json:"text,omitempty"`
	Type         ToolsSiteTemplateCreateV2DataBricksType                       `json:"type"`
	Video        *ToolsSiteTemplateCreateV2ResponseDataBricksInnerVideo        `json:"video,omitempty"`
	WechatApplet *ToolsSiteTemplateCreateV2ResponseDataBricksInnerWechatApplet `json:"wechat_applet,omitempty"`
	WechatGame   *ToolsSiteTemplateCreateV2ResponseDataBricksInnerWechatGame   `json:"wechat_game,omitempty"`
}
