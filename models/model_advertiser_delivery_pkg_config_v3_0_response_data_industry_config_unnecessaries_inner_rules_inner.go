/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigUnnecessariesInnerRulesInner struct for AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigUnnecessariesInnerRulesInner
type AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigUnnecessariesInnerRulesInner struct {
	// 原子规则描述，用于引导用户提交
	Description string `json:"description"`
	// 资质类型
	QualTypes []*AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigUnnecessariesInnerRulesInnerQualTypesInner `json:"qual_types"`
	// 原子规则id
	RuleId int64                                                                  `json:"rule_id"`
	Type   AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType `json:"type"`
}
