/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType
type QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType string

// List of qianchuan_ad_update_v1.0_programmatic_creative_title_list_title_type
const (
	AWEME_CAROUSEL_QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType = "AWEME_CAROUSEL"
	COMMODITY_CARD_QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType = "COMMODITY_CARD"
	CUSTOM_QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType         QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType = "CUSTOM"
)

// All allowed values of QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType enum
var AllowedQianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleTypeEnumValues = []QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType{
	"AWEME_CAROUSEL",
	"COMMODITY_CARD",
	"CUSTOM",
}

// NewQianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleTypeFromValue returns a pointer to a valid QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleTypeFromValue(v string) (*QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType, error) {
	ev := QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType: valid values are %v", v, AllowedQianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_programmatic_creative_title_list_title_type value
func (v QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType) Ptr() *QianchuanAdUpdateV10ProgrammaticCreativeTitleListTitleType {
	return &v
}
