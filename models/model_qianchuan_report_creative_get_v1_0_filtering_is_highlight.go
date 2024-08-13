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

// QianchuanReportCreativeGetV10FilteringIsHighlight
type QianchuanReportCreativeGetV10FilteringIsHighlight string

// List of qianchuan_report_creative_get_v1.0_filtering_is_highlight
const (
	ALL_QianchuanReportCreativeGetV10FilteringIsHighlight            QianchuanReportCreativeGetV10FilteringIsHighlight = "ALL"
	NO_HIGHLIGHT_QianchuanReportCreativeGetV10FilteringIsHighlight   QianchuanReportCreativeGetV10FilteringIsHighlight = "NO_HIGHLIGHT"
	ONLY_HIGHLIGHT_QianchuanReportCreativeGetV10FilteringIsHighlight QianchuanReportCreativeGetV10FilteringIsHighlight = "ONLY_HIGHLIGHT"
)

// All allowed values of QianchuanReportCreativeGetV10FilteringIsHighlight enum
var AllowedQianchuanReportCreativeGetV10FilteringIsHighlightEnumValues = []QianchuanReportCreativeGetV10FilteringIsHighlight{
	"ALL",
	"NO_HIGHLIGHT",
	"ONLY_HIGHLIGHT",
}

// NewQianchuanReportCreativeGetV10FilteringIsHighlightFromValue returns a pointer to a valid QianchuanReportCreativeGetV10FilteringIsHighlight
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10FilteringIsHighlightFromValue(v string) (*QianchuanReportCreativeGetV10FilteringIsHighlight, error) {
	ev := QianchuanReportCreativeGetV10FilteringIsHighlight(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10FilteringIsHighlight: valid values are %v", v, AllowedQianchuanReportCreativeGetV10FilteringIsHighlightEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10FilteringIsHighlight) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10FilteringIsHighlightEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_is_highlight value
func (v QianchuanReportCreativeGetV10FilteringIsHighlight) Ptr() *QianchuanReportCreativeGetV10FilteringIsHighlight {
	return &v
}
