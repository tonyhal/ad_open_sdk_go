/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceCarrier
type ProjectUpdateV30AudienceCarrier string

// List of project_update_v3.0_audience_carrier
const (
	MOBILE_ProjectUpdateV30AudienceCarrier ProjectUpdateV30AudienceCarrier = "MOBILE"
	TELCOM_ProjectUpdateV30AudienceCarrier ProjectUpdateV30AudienceCarrier = "TELCOM"
	UNICOM_ProjectUpdateV30AudienceCarrier ProjectUpdateV30AudienceCarrier = "UNICOM"
)

// All allowed values of ProjectUpdateV30AudienceCarrier enum
var AllowedProjectUpdateV30AudienceCarrierEnumValues = []ProjectUpdateV30AudienceCarrier{
	"MOBILE",
	"TELCOM",
	"UNICOM",
}

// NewProjectUpdateV30AudienceCarrierFromValue returns a pointer to a valid ProjectUpdateV30AudienceCarrier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceCarrierFromValue(v string) (*ProjectUpdateV30AudienceCarrier, error) {
	ev := ProjectUpdateV30AudienceCarrier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceCarrier: valid values are %v", v, AllowedProjectUpdateV30AudienceCarrierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceCarrier) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceCarrierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_carrier value
func (v ProjectUpdateV30AudienceCarrier) Ptr() *ProjectUpdateV30AudienceCarrier {
	return &v
}