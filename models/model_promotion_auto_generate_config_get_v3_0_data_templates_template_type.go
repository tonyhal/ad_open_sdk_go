/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType
type PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType string

// List of promotion_auto_generate_config_get_v3.0_data_templates_template_type
const (
	HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO_PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType = "HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO_PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType = "VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO_PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType  PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType = "VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO"
)

// All allowed values of PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType enum
var AllowedPromotionAutoGenerateConfigGetV30DataTemplatesTemplateTypeEnumValues = []PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType{
	"HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO",
}

// NewPromotionAutoGenerateConfigGetV30DataTemplatesTemplateTypeFromValue returns a pointer to a valid PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionAutoGenerateConfigGetV30DataTemplatesTemplateTypeFromValue(v string) (*PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType, error) {
	ev := PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType: valid values are %v", v, AllowedPromotionAutoGenerateConfigGetV30DataTemplatesTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType) IsValid() bool {
	for _, existing := range AllowedPromotionAutoGenerateConfigGetV30DataTemplatesTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_auto_generate_config_get_v3.0_data_templates_template_type value
func (v PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType) Ptr() *PromotionAutoGenerateConfigGetV30DataTemplatesTemplateType {
	return &v
}
