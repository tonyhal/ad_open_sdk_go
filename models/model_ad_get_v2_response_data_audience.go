/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2ResponseDataAudience
type AdGetV2ResponseDataAudience struct {
	//
	Ac     []*AdGetV2DataAudienceAc           `json:"ac,omitempty"`
	Action *AdGetV2ResponseDataAudienceAction `json:"action,omitempty"`
	//
	ActivateType []*AdGetV2DataAudienceActivateType `json:"activate_type,omitempty"`
	//
	AdTag []int64 `json:"ad_tag,omitempty"`
	//
	Age               []*AdGetV2DataAudienceAge             `json:"age,omitempty"`
	AndroidOsv        *AdGetV2DataAudienceAndroidOsv        `json:"android_osv,omitempty"`
	AppBehaviorTarget *AdGetV2DataAudienceAppBehaviorTarget `json:"app_behavior_target,omitempty"`
	//
	AppCategory []int64 `json:"app_category,omitempty"`
	//
	AppIds []int64 `json:"app_ids,omitempty"`
	//
	ArticleCategory []*AdGetV2DataAudienceArticleCategory `json:"article_category,omitempty"`
	//
	AutoExtendTargets []*AdGetV2DataAudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	//
	AwemeFanAccounts []int64 `json:"aweme_fan_accounts,omitempty"`
	//
	AwemeFanBehaviors []*AdGetV2DataAudienceAwemeFanBehaviors `json:"aweme_fan_behaviors,omitempty"`
	//
	AwemeFanCategories []int64                               `json:"aweme_fan_categories,omitempty"`
	AwemeFanTimeScope  *AdGetV2DataAudienceAwemeFanTimeScope `json:"aweme_fan_time_scope,omitempty"`
	//
	AwemeFansNumbers []int64 `json:"aweme_fans_numbers,omitempty"`
	//
	BusinessIds []int64 `json:"business_ids,omitempty"`
	//
	Career []*AdGetV2DataAudienceCareer `json:"career,omitempty"`
	//
	Carrier []*AdGetV2DataAudienceCarrier `json:"carrier,omitempty"`
	//
	City []int64 `json:"city,omitempty"`
	//
	DeviceBrand []*AdGetV2DataAudienceDeviceBrand `json:"device_brand,omitempty"`
	//
	DeviceType          []*AdGetV2DataAudienceDeviceType        `json:"device_type,omitempty"`
	District            *AdGetV2DataAudienceDistrict            `json:"district,omitempty"`
	DpaRtaRecommendType *AdGetV2DataAudienceDpaRtaRecommendType `json:"dpa_rta_recommend_type,omitempty"`
	DpaRtaSwitch        *AdGetV2DataAudienceDpaRtaSwitch        `json:"dpa_rta_switch,omitempty"`
	//
	ExcludeFlowPackage []int64 `json:"exclude_flow_package,omitempty"`
	//
	FilterAwemeAbnormalActive *int64 `json:"filter_aweme_abnormal_active,omitempty"`
	//
	FilterAwemeFansCount *int64 `json:"filter_aweme_fans_count,omitempty"`
	//
	FilterOwnAwemeFans *int64 `json:"filter_own_aweme_fans,omitempty"`
	//
	FlowPackage []int64                    `json:"flow_package,omitempty"`
	Gender      *AdGetV2DataAudienceGender `json:"gender,omitempty"`
	//
	Geolocation        []*AdGetV2ResponseDataAudienceGeolocationInner `json:"geolocation,omitempty"`
	InterestActionMode *AdGetV2DataAudienceInterestActionMode         `json:"interest_action_mode,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
	//
	InterestTags []int64 `json:"interest_tags,omitempty"`
	//
	InterestWords []int64                    `json:"interest_words,omitempty"`
	IosOsv        *AdGetV2DataAudienceIosOsv `json:"ios_osv,omitempty"`
	//
	LaunchPrice  []int64                          `json:"launch_price,omitempty"`
	LocationType *AdGetV2DataAudienceLocationType `json:"location_type,omitempty"`
	//
	OwnAwemeNumber *int64 `json:"own_aweme_number,omitempty"`
	//
	Platform []*AdGetV2DataAudiencePlatform `json:"platform,omitempty"`
	//
	RegionVersion *string `json:"region_version,omitempty"`
	//
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	//
	RetargetingTagsInclude []int64                                    `json:"retargeting_tags_include,omitempty"`
	SuperiorPopularityType *AdGetV2DataAudienceSuperiorPopularityType `json:"superior_popularity_type,omitempty"`
	//
	UserType []int64 `json:"user_type,omitempty"`
}
