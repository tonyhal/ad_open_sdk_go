/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30Filtering
type PromotionListV30Filtering struct {
	DeliveryMode *PromotionListV30FilteringDeliveryMode `json:"delivery_mode,omitempty"`
	//
	Ids []int64 `json:"ids,omitempty"`
	//
	LearningPhase []*PromotionListV30FilteringLearningPhase `json:"learning_phase,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	ProjectId *int64 `json:"project_id,omitempty"`
	//
	PromotionCreateTime *string `json:"promotion_create_time,omitempty"`
	//
	PromotionModifyTime *string `json:"promotion_modify_time,omitempty"`
	//
	RejectReasonType []*PromotionListV30FilteringRejectReasonType `json:"reject_reason_type,omitempty"`
	Status           *PromotionListV30FilteringStatus             `json:"status,omitempty"`
	StatusFirst      *PromotionListV30FilteringStatusFirst        `json:"status_first,omitempty"`
	StatusSecond     *PromotionListV30FilteringStatusSecond       `json:"status_second,omitempty"`
}
