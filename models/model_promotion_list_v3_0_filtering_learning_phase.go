/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30FilteringLearningPhase
type PromotionListV30FilteringLearningPhase string

// List of promotion_list_v3.0_filtering_learning_phase
const (
	LEARNING_PromotionListV30FilteringLearningPhase     PromotionListV30FilteringLearningPhase = "LEARNING"
	LEARN_FAILED_PromotionListV30FilteringLearningPhase PromotionListV30FilteringLearningPhase = "LEARN_FAILED"
	LEARNED_PromotionListV30FilteringLearningPhase      PromotionListV30FilteringLearningPhase = "LEARNED"
)

// All allowed values of PromotionListV30FilteringLearningPhase enum
var AllowedPromotionListV30FilteringLearningPhaseEnumValues = []PromotionListV30FilteringLearningPhase{
	"LEARNING",
	"LEARN_FAILED",
	"LEARNED",
}

// NewPromotionListV30FilteringLearningPhaseFromValue returns a pointer to a valid PromotionListV30FilteringLearningPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30FilteringLearningPhaseFromValue(v string) (*PromotionListV30FilteringLearningPhase, error) {
	ev := PromotionListV30FilteringLearningPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30FilteringLearningPhase: valid values are %v", v, AllowedPromotionListV30FilteringLearningPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30FilteringLearningPhase) IsValid() bool {
	for _, existing := range AllowedPromotionListV30FilteringLearningPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_filtering_learning_phase value
func (v PromotionListV30FilteringLearningPhase) Ptr() *PromotionListV30FilteringLearningPhase {
	return &v
}
