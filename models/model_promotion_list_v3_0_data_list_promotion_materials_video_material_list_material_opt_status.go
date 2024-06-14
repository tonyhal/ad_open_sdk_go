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

// PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus
type PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus string

// List of promotion_list_v3.0_data_list_promotion_materials_video_material_list_material_opt_status
const (
	DISABLE_PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus = "DISABLE"
	ENABLE_PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus  PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus = "ENABLE"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus enum
var AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatusEnumValues = []PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus{
	"DISABLE",
	"ENABLE",
}

// NewPromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatusFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatusFromValue(v string) (*PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus, error) {
	ev := PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_video_material_list_material_opt_status value
func (v PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus) Ptr() *PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialOptStatus {
	return &v
}
