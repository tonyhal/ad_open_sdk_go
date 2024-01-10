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

// EnterpriseCommentListGetV10DataCommentListSource
type EnterpriseCommentListGetV10DataCommentListSource string

// List of enterprise_comment_list_get_v1.0_data_comment_list_source
const (
	FROM_PERFORM_EnterpriseCommentListGetV10DataCommentListSource EnterpriseCommentListGetV10DataCommentListSource = "FROM_PERFORM"
	FROM_NATURAL_EnterpriseCommentListGetV10DataCommentListSource EnterpriseCommentListGetV10DataCommentListSource = "FROM_NATURAL"
	FROM_OTHER_EnterpriseCommentListGetV10DataCommentListSource   EnterpriseCommentListGetV10DataCommentListSource = "FROM_OTHER"
	FROM_DOUPLUS_EnterpriseCommentListGetV10DataCommentListSource EnterpriseCommentListGetV10DataCommentListSource = "FROM_DOUPLUS"
	FROM_BRAND_EnterpriseCommentListGetV10DataCommentListSource   EnterpriseCommentListGetV10DataCommentListSource = "FROM_BRAND"
)

// All allowed values of EnterpriseCommentListGetV10DataCommentListSource enum
var AllowedEnterpriseCommentListGetV10DataCommentListSourceEnumValues = []EnterpriseCommentListGetV10DataCommentListSource{
	"FROM_PERFORM",
	"FROM_NATURAL",
	"FROM_OTHER",
	"FROM_DOUPLUS",
	"FROM_BRAND",
}

// NewEnterpriseCommentListGetV10DataCommentListSourceFromValue returns a pointer to a valid EnterpriseCommentListGetV10DataCommentListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseCommentListGetV10DataCommentListSourceFromValue(v string) (*EnterpriseCommentListGetV10DataCommentListSource, error) {
	ev := EnterpriseCommentListGetV10DataCommentListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseCommentListGetV10DataCommentListSource: valid values are %v", v, AllowedEnterpriseCommentListGetV10DataCommentListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseCommentListGetV10DataCommentListSource) IsValid() bool {
	for _, existing := range AllowedEnterpriseCommentListGetV10DataCommentListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_comment_list_get_v1.0_data_comment_list_source value
func (v EnterpriseCommentListGetV10DataCommentListSource) Ptr() *EnterpriseCommentListGetV10DataCommentListSource {
	return &v
}
