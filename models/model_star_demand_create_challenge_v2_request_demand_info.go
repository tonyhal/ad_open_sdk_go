/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2RequestDemandInfo
type StarDemandCreateChallengeV2RequestDemandInfo struct {
	AdSyncConf *StarDemandCreateChallengeV2RequestDemandInfoAdSyncConf `json:"ad_sync_conf,omitempty"`
	// 附件描述（补充说明，可包含网盘链接，建议用此项）
	AttachmentText *string `json:"attachment_text,omitempty"`
	// 附件（素材参考） 通过上传材料接口上传的文件名列表
	Attachments []string `json:"attachments,omitempty"`
	//
	AttatchmentsUrl []string `json:"attatchments_url,omitempty"`
	// 品牌名称 50字内。去品牌库中匹配，若无则返回错误。
	BrandName            string                                                            `json:"brand_name"`
	ChallengeRequirement *StarDemandCreateChallengeV2RequestDemandInfoChallengeRequirement `json:"challenge_requirement,omitempty"`
	// 组件点击监测链接 https开头的URL
	ComponentClickMonitorUrl *string                                                    `json:"component_click_monitor_url,omitempty"`
	ComponentInfo            *StarDemandCreateChallengeV2RequestDemandInfoComponentInfo `json:"component_info,omitempty"`
	// 联系人电话 电话号码校验
	ContactName string `json:"contact_name"`
	// 联系人姓名 20字内
	ContactPhone string `json:"contact_phone"`
	// 任务名称 60字内
	DemandName string `json:"demand_name"`
	//
	ExpirationTime *int64 `json:"expiration_time,omitempty"`
	//
	ExpirationTimeEnd *int64 `json:"expiration_time_end,omitempty"`
	// ip活动ID
	IpActId *int64 `json:"ip_act_id,omitempty"`
	// 视频曝光监测链接 https开头的URL
	ItemShowMonitorUrl *string `json:"item_show_monitor_url,omitempty"`
	// 所属类目 类目文本，特定品牌下需要
	ProductCategory []string `json:"product_category,omitempty"`
	// 产品所属行业 [一级行业, 二级行业]
	ProductIndustry []string `json:"product_industry"`
	// 产品介绍 500字以内
	ProductInformation string `json:"product_information"`
	// 商品链接
	ProductLink *string `json:"product_link,omitempty"`
	// 产品名称 40字内
	ProductName string `json:"product_name"`
	// 项目ID
	ProjectId *int64 `json:"project_id,omitempty"`
	// 营销目标 为以下之一 品牌传播 应用下载 电商卖货 影视宣传 门店推广 线索收集 破圈种草
	PromotionTarget *string `json:"promotion_target,omitempty"`
}
