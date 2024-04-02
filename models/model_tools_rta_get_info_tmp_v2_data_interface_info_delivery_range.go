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

// ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange
type ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange string

// List of tools_rta_get_info_tmp_v2_data_interface_info_delivery_range
const (
	LOCAL_ONLY_ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange         ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange = "LOCAL_ONLY"
	UNION_ONLY_ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange         ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange = "UNION_ONLY"
	UNIVERSAL_DELIVERY_ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange = "UNIVERSAL_DELIVERY"
)

// All allowed values of ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange enum
var AllowedToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRangeEnumValues = []ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange{
	"LOCAL_ONLY",
	"UNION_ONLY",
	"UNIVERSAL_DELIVERY",
}

// NewToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRangeFromValue returns a pointer to a valid ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRangeFromValue(v string) (*ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange, error) {
	ev := ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange: valid values are %v", v, AllowedToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange) IsValid() bool {
	for _, existing := range AllowedToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rta_get_info_tmp_v2_data_interface_info_delivery_range value
func (v ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange) Ptr() *ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange {
	return &v
}
