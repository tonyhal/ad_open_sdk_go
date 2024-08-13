/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageCreateV2V2Request struct for ToolsAppManagementExtendPackageCreateV2V2Request
type ToolsAppManagementExtendPackageCreateV2V2Request struct {
	// 账户id，如广告主id、bpid等
	AccountId   int64                                                `json:"account_id"`
	AccountType ToolsAppManagementExtendPackageCreateV2V2AccountType `json:"account_type"`
	// 创建数量，（mode=Auto时需指定）单次分包取值范围1~100
	ChannelCount *int64 `json:"channel_count,omitempty"`
	// 自定义渠道号信息，（mode=Manual时需指定），一次调用，list的size取值范围1~100
	ChannelList []*ToolsAppManagementExtendPackageCreateV2V2RequestChannelListInner `json:"channel_list,omitempty"`
	Mode        ToolsAppManagementExtendPackageCreateV2V2Mode                       `json:"mode"`
	// 应用包ID，从“查询应用信息”接口中获取
	PackageId string `json:"package_id"`
}
