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

// ProjectListV30DataListAudienceCarrier
type ProjectListV30DataListAudienceCarrier string

// List of project_list_v3.0_data_list_audience_carrier
const (
	MOBILE_ProjectListV30DataListAudienceCarrier ProjectListV30DataListAudienceCarrier = "MOBILE"
	TELCOM_ProjectListV30DataListAudienceCarrier ProjectListV30DataListAudienceCarrier = "TELCOM"
	UNICOM_ProjectListV30DataListAudienceCarrier ProjectListV30DataListAudienceCarrier = "UNICOM"
)

// All allowed values of ProjectListV30DataListAudienceCarrier enum
var AllowedProjectListV30DataListAudienceCarrierEnumValues = []ProjectListV30DataListAudienceCarrier{
	"MOBILE",
	"TELCOM",
	"UNICOM",
}

// NewProjectListV30DataListAudienceCarrierFromValue returns a pointer to a valid ProjectListV30DataListAudienceCarrier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceCarrierFromValue(v string) (*ProjectListV30DataListAudienceCarrier, error) {
	ev := ProjectListV30DataListAudienceCarrier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceCarrier: valid values are %v", v, AllowedProjectListV30DataListAudienceCarrierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceCarrier) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceCarrierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_carrier value
func (v ProjectListV30DataListAudienceCarrier) Ptr() *ProjectListV30DataListAudienceCarrier {
	return &v
}
