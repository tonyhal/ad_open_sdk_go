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

// PromotionAutoGenerateConfigGetV30DataVersion
type PromotionAutoGenerateConfigGetV30DataVersion string

// List of promotion_auto_generate_config_get_v3.0_data_version
const (
	STRATEGY_PromotionAutoGenerateConfigGetV30DataVersion PromotionAutoGenerateConfigGetV30DataVersion = "Strategy"
	TEMPLATE_PromotionAutoGenerateConfigGetV30DataVersion PromotionAutoGenerateConfigGetV30DataVersion = "Template"
)

// All allowed values of PromotionAutoGenerateConfigGetV30DataVersion enum
var AllowedPromotionAutoGenerateConfigGetV30DataVersionEnumValues = []PromotionAutoGenerateConfigGetV30DataVersion{
	"Strategy",
	"Template",
}

// NewPromotionAutoGenerateConfigGetV30DataVersionFromValue returns a pointer to a valid PromotionAutoGenerateConfigGetV30DataVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionAutoGenerateConfigGetV30DataVersionFromValue(v string) (*PromotionAutoGenerateConfigGetV30DataVersion, error) {
	ev := PromotionAutoGenerateConfigGetV30DataVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionAutoGenerateConfigGetV30DataVersion: valid values are %v", v, AllowedPromotionAutoGenerateConfigGetV30DataVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionAutoGenerateConfigGetV30DataVersion) IsValid() bool {
	for _, existing := range AllowedPromotionAutoGenerateConfigGetV30DataVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_auto_generate_config_get_v3.0_data_version value
func (v PromotionAutoGenerateConfigGetV30DataVersion) Ptr() *PromotionAutoGenerateConfigGetV30DataVersion {
	return &v
}
