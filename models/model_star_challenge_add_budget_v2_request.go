/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeAddBudgetV2Request struct for StarChallengeAddBudgetV2Request
type StarChallengeAddBudgetV2Request struct {
	// 追加预算金额（单位元，不低于10000）
	AddBudget int64 `json:"add_budget"`
	// 投稿任务ID
	ChallengeTaskId int64 `json:"challenge_task_id"`
	// 客户星图ID
	StarId int64 `json:"star_id"`
}
