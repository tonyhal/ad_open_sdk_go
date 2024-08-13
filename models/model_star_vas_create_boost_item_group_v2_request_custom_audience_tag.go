/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasCreateBoostItemGroupV2RequestCustomAudienceTag 人群标签定向，仅boost_type为CUSTOM_TAG时传参
type StarVasCreateBoostItemGroupV2RequestCustomAudienceTag struct {
	// 年龄，空则为不限 可选值: 18_TO_23 18-23;     24_TO_30 24-30;     31_TO_40 31-40;     41_TO_50 41-50;     INFINITE 不限;
	Age    []*StarVasCreateBoostItemGroupV2CustomAudienceTagAge  `json:"age,omitempty"`
	Gender *StarVasCreateBoostItemGroupV2CustomAudienceTagGender `json:"gender,omitempty"`
}
