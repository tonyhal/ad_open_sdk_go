/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmExpandChallengeV2RequestOmParticipateAuthorRange
type StarDemandOmExpandChallengeV2RequestOmParticipateAuthorRange struct {
	// 增派达人抖音号 专属任务，list，最多200
	AddOrCancelAuthorUids []int64                                                            `json:"add_or_cancel_author_uids"`
	OperationType         StarDemandOmExpandChallengeV2OmParticipateAuthorRangeOperationType `json:"operation_type"`
}