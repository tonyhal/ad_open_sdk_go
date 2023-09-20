/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListInternalAdvertiserInfoClassify
type ProjectListV30DataListInternalAdvertiserInfoClassify string

// List of project_list_v3.0_data_list_internal_advertiser_info_classify
const (
	CLASSIFY_APPORTION_ProjectListV30DataListInternalAdvertiserInfoClassify  ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_APPORTION"
	CLASSIFY_EXCHANGE_ProjectListV30DataListInternalAdvertiserInfoClassify   ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_EXCHANGE"
	CLASSIFY_INTERNAL_ProjectListV30DataListInternalAdvertiserInfoClassify   ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_INTERNAL"
	CLASSIFY_SALE_ProjectListV30DataListInternalAdvertiserInfoClassify       ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_SALE"
	CLASSIFY_SUPPLEMENT_ProjectListV30DataListInternalAdvertiserInfoClassify ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_SUPPLEMENT"
)

// All allowed values of ProjectListV30DataListInternalAdvertiserInfoClassify enum
var AllowedProjectListV30DataListInternalAdvertiserInfoClassifyEnumValues = []ProjectListV30DataListInternalAdvertiserInfoClassify{
	"CLASSIFY_APPORTION",
	"CLASSIFY_EXCHANGE",
	"CLASSIFY_INTERNAL",
	"CLASSIFY_SALE",
	"CLASSIFY_SUPPLEMENT",
}

// NewProjectListV30DataListInternalAdvertiserInfoClassifyFromValue returns a pointer to a valid ProjectListV30DataListInternalAdvertiserInfoClassify
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListInternalAdvertiserInfoClassifyFromValue(v string) (*ProjectListV30DataListInternalAdvertiserInfoClassify, error) {
	ev := ProjectListV30DataListInternalAdvertiserInfoClassify(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListInternalAdvertiserInfoClassify: valid values are %v", v, AllowedProjectListV30DataListInternalAdvertiserInfoClassifyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListInternalAdvertiserInfoClassify) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListInternalAdvertiserInfoClassifyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_internal_advertiser_info_classify value
func (v ProjectListV30DataListInternalAdvertiserInfoClassify) Ptr() *ProjectListV30DataListInternalAdvertiserInfoClassify {
	return &v
}
