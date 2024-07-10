/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCostProtectStatusGetV30ResponseDataCompensateStatusInfoListInner struct for ProjectCostProtectStatusGetV30ResponseDataCompensateStatusInfoListInner
type ProjectCostProtectStatusGetV30ResponseDataCompensateStatusInfoListInner struct {
	// 赔付金额
	CompensateAmount *float64 `json:"compensate_amount,omitempty"`
	// 成本保障结束原因
	CompensateEndReasons []*ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateEndReasons `json:"compensate_end_reasons,omitempty"`
	// 成本保障失效原因
	CompensateInvalidReasons []*ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons `json:"compensate_invalid_reasons,omitempty"`
	CompensateStatus         *ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateStatus           `json:"compensate_status,omitempty"`
	// 赔付规则链接
	CompensateUrl *string `json:"compensate_url,omitempty"`
	// 项目id
	ProjectId *int64 `json:"project_id,omitempty"`
}