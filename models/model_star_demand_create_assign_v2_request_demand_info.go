/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2RequestDemandInfo 需求信息
type StarDemandCreateAssignV2RequestDemandInfo struct {
	// 接收任务超时时间（天，1≤x≤365）
	AcceptExpirationDay *int64                                               `json:"accept_expiration_day,omitempty"`
	AdSyncConf          *StarDemandCreateAssignV2RequestDemandInfoAdSyncConf `json:"ad_sync_conf,omitempty"`
	// 附件描述（可包含网盘链接，建议用此项）
	AttachmentText *string `json:"attachment_text,omitempty"`
	// 附件 通过上传材料接口上传的文件名列表
	Attachments []string `json:"attachments,omitempty"`
	// 品牌名称(50字内)
	BrandName string `json:"brand_name"`
	// 组件点击监测链接 https开头的URL
	ComponentClickMonitorUrl *string                                                 `json:"component_click_monitor_url,omitempty"`
	ComponentInfo            *StarDemandCreateAssignV2RequestDemandInfoComponentInfo `json:"component_info,omitempty"`
	// 联系人姓名
	ContactName string `json:"contact_name"`
	// 联系人电话
	ContactPhone string `json:"contact_phone"`
	// 需求名称（任务名称60字内）
	DemandName        string                                                     `json:"demand_name"`
	DemandRequirement StarDemandCreateAssignV2RequestDemandInfoDemandRequirement `json:"demand_requirement"`
	// 期望（视频）保留时间（天） 0<x≤365整数
	ExpectRemainTime int64 `json:"expect_remain_time"`
	// 期望发布时间
	ExpirationTime int64 `json:"expiration_time"`
	// 期望最迟发布时间 unix timestamp，发布时间范围长度不可超过30天
	ExpirationTimeEnd int64 `json:"expiration_time_end"`
	// 是否跳过脚本环节 (0)或不传：不跳过，(1)：跳过
	IgnoreScript *int64 `json:"ignore_script,omitempty"`
	// 视频曝光监测链接 https开头的URL
	ItemShowMonitorUrl *string `json:"item_show_monitor_url,omitempty"`
	// 所属类目
	ProductCategory []string `json:"product_category,omitempty"`
	// 产品所属行业 [一级行业, 二级行业]
	ProductIndustry []string `json:"product_industry"`
	// 产品介绍(500字内)
	ProductInformation string `json:"product_information"`
	// 产品名称(40字内)
	ProductName string `json:"product_name"`
	// 星智投预审，0或不传为不开启，1为开启
	XingzhitouPreAudit *int32 `json:"xingzhitou_pre_audit,omitempty"`
}
