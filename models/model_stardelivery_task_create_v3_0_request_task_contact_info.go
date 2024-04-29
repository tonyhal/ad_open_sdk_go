/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskCreateV30RequestTaskContactInfo 联系人姓名及手机号，请传入可接通的手机号 达人接单后，您可以通过虚拟电话与达人沟通
type StardeliveryTaskCreateV30RequestTaskContactInfo struct {
	// 联系人ID，可通过「获取星广联投任务详情」接口查询，如传入联系人ID优先取ID（同时传入联系人姓名+手机号不生效）
	ContactId *int64 `json:"contact_id,omitempty"`
	// 联系人姓名，1-50个字符，此信息将被用于XXXXX - 姓名需与手机号同时传入，只传入姓名会报错 - 达人接单后，您可以通过虚拟电话与达人沟通，传入手机号信息前请阅读《个人信息保护声明》
	ContactName *string `json:"contact_name,omitempty"`
	// 联系人手机号，请传入可以接通的手机号。达人接单后，您可以通过虚拟电话与达人沟通，传入手机号信息前请阅读《个人信息保护声明》 - 姓名需与手机号同时传入，只传入手机号会报错
	ContactPhone *string `json:"contact_phone,omitempty"`
}