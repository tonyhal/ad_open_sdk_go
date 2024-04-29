/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderCreateV30RequestTargeting 定向设置
type DouplusOrderCreateV30RequestTargeting struct {
	// 年龄，允许值：若不传值，则默认为不限 AGE_BETWEEN_18_23：18-23岁、 AGE_BETWEEN_24_30：24-30岁、 AGE_BETWEEN_31_40：31-40岁、 AGE_BETWEEN_41_49：41-49岁 注意：仅想吸引的观众类型为”CUSTOM”时，有效
	Age          []*DouplusOrderCreateV30TargetingAge       `json:"age,omitempty"`
	AudienceMode DouplusOrderCreateV30TargetingAudienceMode `json:"audience_mode"`
	// 具体定向包id 注意：想吸引的观众类型为”AUDIENCE_PKG”时，必填
	AudiencePkgId *int64 `json:"audience_pkg_id,omitempty"`
	// 定向人群包ID（需客户从网页版查询获取） 注意：仅想吸引的观众类型为”DMP_PKG”时，有效
	DmpPkgIds []int64                               `json:"dmp_pkg_ids,omitempty"`
	Gender    *DouplusOrderCreateV30TargetingGender `json:"gender,omitempty"`
}