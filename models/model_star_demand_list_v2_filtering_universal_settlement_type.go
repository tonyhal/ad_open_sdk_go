/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarDemandListV2FilteringUniversalSettlementType
type StarDemandListV2FilteringUniversalSettlementType string

// List of star_demand_list_v2_filtering_universal_settlement_type
const (
	EFFECT_StarDemandListV2FilteringUniversalSettlementType       StarDemandListV2FilteringUniversalSettlementType = "EFFECT"
	FLOW_SHARE_StarDemandListV2FilteringUniversalSettlementType   StarDemandListV2FilteringUniversalSettlementType = "FLOW_SHARE"
	FIXED_PRICE_StarDemandListV2FilteringUniversalSettlementType  StarDemandListV2FilteringUniversalSettlementType = "FIXED_PRICE"
	CPM_StarDemandListV2FilteringUniversalSettlementType          StarDemandListV2FilteringUniversalSettlementType = "CPM"
	GIFT_StarDemandListV2FilteringUniversalSettlementType         StarDemandListV2FilteringUniversalSettlementType = "GIFT"
	MONEY_SHARE_StarDemandListV2FilteringUniversalSettlementType  StarDemandListV2FilteringUniversalSettlementType = "MONEY_SHARE"
	STAR_SUPPORT_StarDemandListV2FilteringUniversalSettlementType StarDemandListV2FilteringUniversalSettlementType = "STAR_SUPPORT"
	CUSTOMIZE_StarDemandListV2FilteringUniversalSettlementType    StarDemandListV2FilteringUniversalSettlementType = "CUSTOMIZE"
	ALL_StarDemandListV2FilteringUniversalSettlementType          StarDemandListV2FilteringUniversalSettlementType = "ALL"
	CPA_StarDemandListV2FilteringUniversalSettlementType          StarDemandListV2FilteringUniversalSettlementType = "CPA"
	EXCHANGE_StarDemandListV2FilteringUniversalSettlementType     StarDemandListV2FilteringUniversalSettlementType = "EXCHANGE"
	RANK_StarDemandListV2FilteringUniversalSettlementType         StarDemandListV2FilteringUniversalSettlementType = "RANK"
	DOU_PLUS_StarDemandListV2FilteringUniversalSettlementType     StarDemandListV2FilteringUniversalSettlementType = "DOU_PLUS"
)

// All allowed values of StarDemandListV2FilteringUniversalSettlementType enum
var AllowedStarDemandListV2FilteringUniversalSettlementTypeEnumValues = []StarDemandListV2FilteringUniversalSettlementType{
	"EFFECT",
	"FLOW_SHARE",
	"FIXED_PRICE",
	"CPM",
	"GIFT",
	"MONEY_SHARE",
	"STAR_SUPPORT",
	"CUSTOMIZE",
	"ALL",
	"CPA",
	"EXCHANGE",
	"RANK",
	"DOU_PLUS",
}

// NewStarDemandListV2FilteringUniversalSettlementTypeFromValue returns a pointer to a valid StarDemandListV2FilteringUniversalSettlementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarDemandListV2FilteringUniversalSettlementTypeFromValue(v string) (*StarDemandListV2FilteringUniversalSettlementType, error) {
	ev := StarDemandListV2FilteringUniversalSettlementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarDemandListV2FilteringUniversalSettlementType: valid values are %v", v, AllowedStarDemandListV2FilteringUniversalSettlementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarDemandListV2FilteringUniversalSettlementType) IsValid() bool {
	for _, existing := range AllowedStarDemandListV2FilteringUniversalSettlementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_demand_list_v2_filtering_universal_settlement_type value
func (v StarDemandListV2FilteringUniversalSettlementType) Ptr() *StarDemandListV2FilteringUniversalSettlementType {
	return &v
}
