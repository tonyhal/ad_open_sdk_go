/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestListGetV2ResponseDataTestListInner struct for ToolsAbTestListGetV2ResponseDataTestListInner
type ToolsAbTestListGetV2ResponseDataTestListInner struct {
	Conclusion *ToolsAbTestListGetV2DataTestListConclusion `json:"conclusion,omitempty"`
	// 实验创建时间，格式：\"2020-12-25 07:12:08\"
	CreateTime *string `json:"create_time,omitempty"`
	// 实验报告更新日期，格式：\"2020-12-25 07:12:08\"
	ReportUpdateDate *string `json:"report_update_date,omitempty"`
	// 实验报告更新时间，格式：\"2020-12-25 07:12:08\"
	ReportUpdateTime *string                                 `json:"report_update_time,omitempty"`
	Status           *ToolsAbTestListGetV2DataTestListStatus `json:"status,omitempty"`
	// 实验结束时间
	TestEndTime *string `json:"test_end_time,omitempty"`
	// 实验ID
	TestId *int64 `json:"test_id,omitempty"`
	// 实验名称
	TestName *string `json:"test_name,omitempty"`
	// 实验开始时间
	TestStartTime *string                                      `json:"test_start_time,omitempty"`
	TestType      *ToolsAbTestListGetV2DataTestListTestType    `json:"test_type,omitempty"`
	TestVersion   *ToolsAbTestListGetV2DataTestListTestVersion `json:"test_version,omitempty"`
}
