/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupListV30ResponseDataGroupsInnerAdInfoAudience 定向信息
type AdlabGroupListV30ResponseDataGroupsInnerAdInfoAudience struct {
	// 年龄
	Age []*[]int64 `json:"age,omitempty"`
	// 定向包ID
	AudiencePackageId *int64 `json:"audience_package_id,omitempty"`
	// 是否启用智能放量
	AutoExtendEnabled *int64 `json:"auto_extend_enabled,omitempty"`
	// 可放开定向
	AutoExtendTargets []*AdlabGroupListV30DataGroupsAdInfoAudienceAutoExtendTargets `json:"auto_extend_targets,omitempty"`
	// 地域定向省市或者区县列表
	City                  []int64                                                         `json:"city,omitempty"`
	ConvertedTimeDuration *AdlabGroupListV30DataGroupsAdInfoAudienceConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	District              *AdlabGroupListV30DataGroupsAdInfoAudienceDistrict              `json:"district,omitempty"`
	// 排除定向逻辑
	ExcludeFlowPackage []int64 `json:"exclude_flow_package,omitempty"`
	// 定向逻辑
	FlowPackage []*[]int64                                       `json:"flow_package,omitempty"`
	Gender      *AdlabGroupListV30DataGroupsAdInfoAudienceGender `json:"gender,omitempty"`
	// 商圈定向
	Geolocation     []*AdlabGroupListV30ResponseDataGroupsInnerAdInfoAudienceGeolocationInner `json:"geolocation,omitempty"`
	HideIfConverted *AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted                 `json:"hide_if_converted,omitempty"`
	LocationType    *AdlabGroupListV30DataGroupsAdInfoAudienceLocationType                    `json:"location_type,omitempty"`
	// 受众平台
	Platform []int64 `json:"platform,omitempty"`
	// 地理位置版本号
	RegionVersion *string `json:"region_version,omitempty"`
	// 排除人群包列表（自定义人群）
	RetargetingTagsExclude []int64 `json:"retargeting_tags_exclude,omitempty"`
	// 定向人群包列表（自定义人群）
	RetargetingTagsInclude []int64                                                          `json:"retargeting_tags_include,omitempty"`
	SuperiorPopularityType *AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType `json:"superior_popularity_type,omitempty"`
}
