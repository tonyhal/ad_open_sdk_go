/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfo
type StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfo struct {
	//
	AcceptExpirationDay *int64                                                                `json:"accept_expiration_day,omitempty"`
	AdSyncConf          *StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfoAdSyncConf `json:"ad_sync_conf,omitempty"`
	//
	AttachmentText *string `json:"attachment_text,omitempty"`
	//
	Attachments []string `json:"attachments,omitempty"`
	//
	AttatchmentsUrl []string `json:"attatchments_url,omitempty"`
	//
	BrandName            *string                                                                         `json:"brand_name,omitempty"`
	ChallengeRequirement *StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfoChallengeRequirement `json:"challenge_requirement,omitempty"`
	//
	ComponentClickMonitorUrl *string                                                                  `json:"component_click_monitor_url,omitempty"`
	ComponentInfo            *StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfoComponentInfo `json:"component_info,omitempty"`
	//
	ContactName *string `json:"contact_name,omitempty"`
	//
	ContactPhone  *string                                                                  `json:"contact_phone,omitempty"`
	CustomizeInfo *StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfoCustomizeInfo `json:"customize_info,omitempty"`
	//
	DemandName        string                                                                       `json:"demand_name"`
	DemandRequirement *StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfoDemandRequirement `json:"demand_requirement,omitempty"`
	//
	ExpectRemainTime *int64 `json:"expect_remain_time,omitempty"`
	//
	GoodsLink *string `json:"goods_link,omitempty"`
	//
	IgnoreScript *int64 `json:"ignore_script,omitempty"`
	//
	ItemShowMonitorUrl *string `json:"item_show_monitor_url,omitempty"`
	//
	ProductCategory []string `json:"product_category,omitempty"`
	//
	ProductIndustry []string `json:"product_industry,omitempty"`
	//
	ProductInformation *string `json:"product_information,omitempty"`
	//
	ProductName string `json:"product_name"`
	//
	PromotionTarget *string `json:"promotion_target,omitempty"`
}
