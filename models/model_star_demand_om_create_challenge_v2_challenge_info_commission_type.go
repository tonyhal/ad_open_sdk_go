/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarDemandOmCreateChallengeV2ChallengeInfoCommissionType
type StarDemandOmCreateChallengeV2ChallengeInfoCommissionType int64

// List of star_demand_om_create_challenge_v2_challenge_info_commission_type
const (
	Enum_0_StarDemandOmCreateChallengeV2ChallengeInfoCommissionType StarDemandOmCreateChallengeV2ChallengeInfoCommissionType = 0
	Enum_1_StarDemandOmCreateChallengeV2ChallengeInfoCommissionType StarDemandOmCreateChallengeV2ChallengeInfoCommissionType = 1
	Enum_2_StarDemandOmCreateChallengeV2ChallengeInfoCommissionType StarDemandOmCreateChallengeV2ChallengeInfoCommissionType = 2
	Enum_3_StarDemandOmCreateChallengeV2ChallengeInfoCommissionType StarDemandOmCreateChallengeV2ChallengeInfoCommissionType = 3
)

// All allowed values of StarDemandOmCreateChallengeV2ChallengeInfoCommissionType enum
var AllowedStarDemandOmCreateChallengeV2ChallengeInfoCommissionTypeEnumValues = []StarDemandOmCreateChallengeV2ChallengeInfoCommissionType{
	0,
	1,
	2,
	3,
}

// NewStarDemandOmCreateChallengeV2ChallengeInfoCommissionTypeFromValue returns a pointer to a valid StarDemandOmCreateChallengeV2ChallengeInfoCommissionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandOmCreateChallengeV2ChallengeInfoCommissionTypeFromValue(v int64) (*StarDemandOmCreateChallengeV2ChallengeInfoCommissionType, error) {
	ev := StarDemandOmCreateChallengeV2ChallengeInfoCommissionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandOmCreateChallengeV2ChallengeInfoCommissionType: valid values are %v", v, AllowedStarDemandOmCreateChallengeV2ChallengeInfoCommissionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandOmCreateChallengeV2ChallengeInfoCommissionType) IsValid() bool {
	for _, existing := range AllowedStarDemandOmCreateChallengeV2ChallengeInfoCommissionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_om_create_challenge_v2_challenge_info_commission_type value
func (v StarDemandOmCreateChallengeV2ChallengeInfoCommissionType) Ptr() *StarDemandOmCreateChallengeV2ChallengeInfoCommissionType {
	return &v
}
