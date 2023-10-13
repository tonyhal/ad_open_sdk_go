/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserInfoV2DataRole
type AdvertiserInfoV2DataRole string

// List of advertiser_info_v2_data_role
const (
	ROLE_ADVERTISER_AdvertiserInfoV2DataRole                   AdvertiserInfoV2DataRole = "ROLE_ADVERTISER"
	ROLE_HOTSOON_ADVERTISER_AdvertiserInfoV2DataRole           AdvertiserInfoV2DataRole = "ROLE_HOTSOON_ADVERTISER"
	ROLE_DOUDIAN_ADVERTISER_AdvertiserInfoV2DataRole           AdvertiserInfoV2DataRole = "ROLE_DOUDIAN_ADVERTISER"
	ROLE_CHILD_AGENT_AdvertiserInfoV2DataRole                  AdvertiserInfoV2DataRole = "ROLE_CHILD_AGENT"
	ROLE_ECP_CHILD_ADVERTISER_AdvertiserInfoV2DataRole         AdvertiserInfoV2DataRole = "ROLE_ECP_CHILD_ADVERTISER"
	ROLE_ECP_INTERNAL_ADVERTISER_AdvertiserInfoV2DataRole      AdvertiserInfoV2DataRole = "ROLE_ECP_INTERNAL_ADVERTISER"
	ROLE_ADVERTISER_SYSTEM_ACCOUNT_AdvertiserInfoV2DataRole    AdvertiserInfoV2DataRole = "ROLE_ADVERTISER_SYSTEM_ACCOUNT"
	ROLE_DSP_ADVERTISER_AdvertiserInfoV2DataRole               AdvertiserInfoV2DataRole = "ROLE_DSP_ADVERTISER"
	ROLE_AGENT_AdvertiserInfoV2DataRole                        AdvertiserInfoV2DataRole = "ROLE_AGENT"
	ROLE_INTERNAL_ADV_AdvertiserInfoV2DataRole                 AdvertiserInfoV2DataRole = "ROLE_INTERNAL_ADV"
	ROLE_HOTSOON_PROMOTION_ADVERTISER_AdvertiserInfoV2DataRole AdvertiserInfoV2DataRole = "ROLE_HOTSOON_PROMOTION_ADVERTISER"
	ROLE_AGENT_ABSTRACT_AdvertiserInfoV2DataRole               AdvertiserInfoV2DataRole = "ROLE_AGENT_ABSTRACT"
	ROLE_PGC_ADVERTISER_AdvertiserInfoV2DataRole               AdvertiserInfoV2DataRole = "ROLE_PGC_ADVERTISER"
	ROLE_ADVERTISER_ABSTRACT_AdvertiserInfoV2DataRole          AdvertiserInfoV2DataRole = "ROLE_ADVERTISER_ABSTRACT"
	ROLE_MAJORDOMO_AdvertiserInfoV2DataRole                    AdvertiserInfoV2DataRole = "ROLE_MAJORDOMO"
	ROLE_ADMIN_AdvertiserInfoV2DataRole                        AdvertiserInfoV2DataRole = "ROLE_ADMIN"
	ROLE_ECP_ADVERTISER_AdvertiserInfoV2DataRole               AdvertiserInfoV2DataRole = "ROLE_ECP_ADVERTISER"
	ROLE_CHILD_ADVERTISER_AdvertiserInfoV2DataRole             AdvertiserInfoV2DataRole = "ROLE_CHILD_ADVERTISER"
	ROLE_AWEME_PROMOTION_ADVERTISER_AdvertiserInfoV2DataRole   AdvertiserInfoV2DataRole = "ROLE_AWEME_PROMOTION_ADVERTISER"
	ROLE_LITE_ADVERTISER_AdvertiserInfoV2DataRole              AdvertiserInfoV2DataRole = "ROLE_LITE_ADVERTISER"
	ROLE_AGENT_SYSTEM_ACCOUNT_AdvertiserInfoV2DataRole         AdvertiserInfoV2DataRole = "ROLE_AGENT_SYSTEM_ACCOUNT"
)

// All allowed values of AdvertiserInfoV2DataRole enum
var AllowedAdvertiserInfoV2DataRoleEnumValues = []AdvertiserInfoV2DataRole{
	"ROLE_ADVERTISER",
	"ROLE_HOTSOON_ADVERTISER",
	"ROLE_DOUDIAN_ADVERTISER",
	"ROLE_CHILD_AGENT",
	"ROLE_ECP_CHILD_ADVERTISER",
	"ROLE_ECP_INTERNAL_ADVERTISER",
	"ROLE_ADVERTISER_SYSTEM_ACCOUNT",
	"ROLE_DSP_ADVERTISER",
	"ROLE_AGENT",
	"ROLE_INTERNAL_ADV",
	"ROLE_HOTSOON_PROMOTION_ADVERTISER",
	"ROLE_AGENT_ABSTRACT",
	"ROLE_PGC_ADVERTISER",
	"ROLE_ADVERTISER_ABSTRACT",
	"ROLE_MAJORDOMO",
	"ROLE_ADMIN",
	"ROLE_ECP_ADVERTISER",
	"ROLE_CHILD_ADVERTISER",
	"ROLE_AWEME_PROMOTION_ADVERTISER",
	"ROLE_LITE_ADVERTISER",
	"ROLE_AGENT_SYSTEM_ACCOUNT",
}

// NewAdvertiserInfoV2DataRoleFromValue returns a pointer to a valid AdvertiserInfoV2DataRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserInfoV2DataRoleFromValue(v string) (*AdvertiserInfoV2DataRole, error) {
	ev := AdvertiserInfoV2DataRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserInfoV2DataRole: valid values are %v", v, AllowedAdvertiserInfoV2DataRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserInfoV2DataRole) IsValid() bool {
	for _, existing := range AllowedAdvertiserInfoV2DataRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_info_v2_data_role value
func (v AdvertiserInfoV2DataRole) Ptr() *AdvertiserInfoV2DataRole {
	return &v
}
