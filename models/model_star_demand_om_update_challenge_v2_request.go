/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmUpdateChallengeV2Request struct for StarDemandOmUpdateChallengeV2Request
type StarDemandOmUpdateChallengeV2Request struct {
	ChallengeInfo StarDemandOmUpdateChallengeV2RequestChallengeInfo `json:"challenge_info"`
	// 任务ID,19位数字
	ChallengeTaskId int64                                          `json:"challenge_task_id"`
	DemandInfo      StarDemandOmUpdateChallengeV2RequestDemandInfo `json:"demand_info"`
	//
	DeveloperId *string `json:"developer_id,omitempty"`
	// 客户星图ID
	StarId int64 `json:"star_id"`
}
