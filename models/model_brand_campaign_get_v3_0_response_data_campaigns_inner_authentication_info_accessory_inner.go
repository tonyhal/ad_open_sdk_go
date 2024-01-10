/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignGetV30ResponseDataCampaignsInnerAuthenticationInfoAccessoryInner struct for BrandCampaignGetV30ResponseDataCampaignsInnerAuthenticationInfoAccessoryInner
type BrandCampaignGetV30ResponseDataCampaignsInnerAuthenticationInfoAccessoryInner struct {
	// 附件类型
	AccessoryType int64 `json:"accessory_type"`
	// 创建人的账号名称
	CreateUserName *string `json:"create_user_name,omitempty"`
	// 文件的下载链接供外网使用
	DownloadUrl *string `json:"download_url,omitempty"`
	// 附件名称
	Name string `json:"name"`
	// 附件key
	ObjectKey string `json:"object_key"`
	// 附件备注
	Remark *string `json:"remark,omitempty"`
	// 附件地址
	Url string `json:"url"`
}
