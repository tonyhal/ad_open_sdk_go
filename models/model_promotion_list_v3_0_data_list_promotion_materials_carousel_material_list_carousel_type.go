/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType
type PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType string

// List of promotion_list_v3.0_data_list_promotion_materials_carousel_material_list_carousel_type
const (
	INFORMATION_FLOW_IMAGE_PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType      PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType = "INFORMATION_FLOW_IMAGE"
	SEARCH_DISPLAY_WINDOW_IMAGE_PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType = "SEARCH_DISPLAY_WINDOW_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType     PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType enum
var AllowedPromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselTypeEnumValues = []PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType{
	"INFORMATION_FLOW_IMAGE",
	"SEARCH_DISPLAY_WINDOW_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewPromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselTypeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselTypeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType, error) {
	ev := PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_carousel_material_list_carousel_type value
func (v PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType) Ptr() *PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType {
	return &v
}