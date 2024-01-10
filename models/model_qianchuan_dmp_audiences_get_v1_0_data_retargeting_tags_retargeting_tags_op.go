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

// QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp
type QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp string

// List of qianchuan_dmp_audiences_get_v1.0_data_retargeting_tags_retargeting_tags_op
const (
	ALL_QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp     QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp = "ALL"
	EXCLUDE_QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp = "EXCLUDE"
	INCLUDE_QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp = "INCLUDE"
)

// All allowed values of QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp enum
var AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOpEnumValues = []QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp{
	"ALL",
	"EXCLUDE",
	"INCLUDE",
}

// NewQianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOpFromValue returns a pointer to a valid QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOpFromValue(v string) (*QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp, error) {
	ev := QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp: valid values are %v", v, AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp) IsValid() bool {
	for _, existing := range AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_dmp_audiences_get_v1.0_data_retargeting_tags_retargeting_tags_op value
func (v QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp) Ptr() *QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp {
	return &v
}
