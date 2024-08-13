/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeEstimateProfitV10AudienceAudienceMode
type QianchuanAwemeEstimateProfitV10AudienceAudienceMode string

// List of qianchuan_aweme_estimate_profit_v1.0_audience_audience_mode
const (
	ATUO_QianchuanAwemeEstimateProfitV10AudienceAudienceMode     QianchuanAwemeEstimateProfitV10AudienceAudienceMode = "ATUO"
	CUSTOM_QianchuanAwemeEstimateProfitV10AudienceAudienceMode   QianchuanAwemeEstimateProfitV10AudienceAudienceMode = "CUSTOM"
	FANS_QianchuanAwemeEstimateProfitV10AudienceAudienceMode     QianchuanAwemeEstimateProfitV10AudienceAudienceMode = "FANS"
	LIVEFANS_QianchuanAwemeEstimateProfitV10AudienceAudienceMode QianchuanAwemeEstimateProfitV10AudienceAudienceMode = "LIVEFANS"
)

// All allowed values of QianchuanAwemeEstimateProfitV10AudienceAudienceMode enum
var AllowedQianchuanAwemeEstimateProfitV10AudienceAudienceModeEnumValues = []QianchuanAwemeEstimateProfitV10AudienceAudienceMode{
	"ATUO",
	"CUSTOM",
	"FANS",
	"LIVEFANS",
}

// NewQianchuanAwemeEstimateProfitV10AudienceAudienceModeFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10AudienceAudienceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10AudienceAudienceModeFromValue(v string) (*QianchuanAwemeEstimateProfitV10AudienceAudienceMode, error) {
	ev := QianchuanAwemeEstimateProfitV10AudienceAudienceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10AudienceAudienceMode: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10AudienceAudienceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10AudienceAudienceMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10AudienceAudienceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_audience_audience_mode value
func (v QianchuanAwemeEstimateProfitV10AudienceAudienceMode) Ptr() *QianchuanAwemeEstimateProfitV10AudienceAudienceMode {
	return &v
}
