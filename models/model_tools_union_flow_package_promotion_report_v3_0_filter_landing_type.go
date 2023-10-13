/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsUnionFlowPackagePromotionReportV30FilterLandingType
type ToolsUnionFlowPackagePromotionReportV30FilterLandingType string

// List of tools_union_flow_package_promotion_report_v3.0_filter_landing_type
const (
	APP_ToolsUnionFlowPackagePromotionReportV30FilterLandingType        ToolsUnionFlowPackagePromotionReportV30FilterLandingType = "APP"
	DPA_ToolsUnionFlowPackagePromotionReportV30FilterLandingType        ToolsUnionFlowPackagePromotionReportV30FilterLandingType = "DPA"
	LINK_ToolsUnionFlowPackagePromotionReportV30FilterLandingType       ToolsUnionFlowPackagePromotionReportV30FilterLandingType = "LINK"
	MICRO_GAME_ToolsUnionFlowPackagePromotionReportV30FilterLandingType ToolsUnionFlowPackagePromotionReportV30FilterLandingType = "MICRO_GAME"
	QUICK_APP_ToolsUnionFlowPackagePromotionReportV30FilterLandingType  ToolsUnionFlowPackagePromotionReportV30FilterLandingType = "QUICK_APP"
	SHOP_ToolsUnionFlowPackagePromotionReportV30FilterLandingType       ToolsUnionFlowPackagePromotionReportV30FilterLandingType = "SHOP"
)

// All allowed values of ToolsUnionFlowPackagePromotionReportV30FilterLandingType enum
var AllowedToolsUnionFlowPackagePromotionReportV30FilterLandingTypeEnumValues = []ToolsUnionFlowPackagePromotionReportV30FilterLandingType{
	"APP",
	"DPA",
	"LINK",
	"MICRO_GAME",
	"QUICK_APP",
	"SHOP",
}

// NewToolsUnionFlowPackagePromotionReportV30FilterLandingTypeFromValue returns a pointer to a valid ToolsUnionFlowPackagePromotionReportV30FilterLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsUnionFlowPackagePromotionReportV30FilterLandingTypeFromValue(v string) (*ToolsUnionFlowPackagePromotionReportV30FilterLandingType, error) {
	ev := ToolsUnionFlowPackagePromotionReportV30FilterLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsUnionFlowPackagePromotionReportV30FilterLandingType: valid values are %v", v, AllowedToolsUnionFlowPackagePromotionReportV30FilterLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsUnionFlowPackagePromotionReportV30FilterLandingType) IsValid() bool {
	for _, existing := range AllowedToolsUnionFlowPackagePromotionReportV30FilterLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_union_flow_package_promotion_report_v3.0_filter_landing_type value
func (v ToolsUnionFlowPackagePromotionReportV30FilterLandingType) Ptr() *ToolsUnionFlowPackagePromotionReportV30FilterLandingType {
	return &v
}