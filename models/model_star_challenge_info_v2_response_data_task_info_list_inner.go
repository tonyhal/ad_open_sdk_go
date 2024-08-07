/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeInfoV2ResponseDataTaskInfoListInner struct for StarChallengeInfoV2ResponseDataTaskInfoListInner
type StarChallengeInfoV2ResponseDataTaskInfoListInner struct {
	//
	AuditInfo *string `json:"audit_info,omitempty"`
	//
	ChallengeAuditStatus *int64                                                             `json:"challenge_audit_status,omitempty"`
	ChallengeBillInfo    *StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeBillInfo `json:"challenge_bill_info,omitempty"`
	ChallengeInfo        *StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfo     `json:"challenge_info,omitempty"`
	//
	ChallengeTaskId *int64 `json:"challenge_task_id,omitempty"`
	//
	ChallengeTaskStatus *int64 `json:"challenge_task_status,omitempty"`
	//
	CreateTime   *int64                                                        `json:"create_time,omitempty"`
	DemandInfo   *StarChallengeInfoV2ResponseDataTaskInfoListInnerDemandInfo   `json:"demand_info,omitempty"`
	ProgressInfo *StarChallengeInfoV2ResponseDataTaskInfoListInnerProgressInfo `json:"progress_info,omitempty"`
}
