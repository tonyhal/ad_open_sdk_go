/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons
type ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons string

// List of project_cost_protect_status_get_v3.0_data_compensate_status_info_list_compensate_invalid_reasons
const (
	ANTI_SPAM_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons         ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons = "ANTI_SPAM"
	AUD_BID_CHANGES_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons   ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons = "AUD_BID_CHANGES"
	AUD_CHANGES_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons       ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons = "AUD_CHANGES"
	AUD_ROI_CHANGES_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons   ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons = "AUD_ROI_CHANGES"
	BID_CHANGES_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons       ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons = "BID_CHANGES"
	BID_TYPE_EXPIRED_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons  ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons = "BID_TYPE_EXPIRED"
	MANUAL_JUDGE_SPAM_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons = "MANUAL_JUDGE_SPAM"
	ROI_CHANGES_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons       ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons = "ROI_CHANGES"
	TRANSFER_ACCOUNT_ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons  ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons = "TRANSFER_ACCOUNT"
)

// All allowed values of ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons enum
var AllowedProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasonsEnumValues = []ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons{
	"ANTI_SPAM",
	"AUD_BID_CHANGES",
	"AUD_CHANGES",
	"AUD_ROI_CHANGES",
	"BID_CHANGES",
	"BID_TYPE_EXPIRED",
	"MANUAL_JUDGE_SPAM",
	"ROI_CHANGES",
	"TRANSFER_ACCOUNT",
}

// NewProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasonsFromValue returns a pointer to a valid ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasonsFromValue(v string) (*ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons, error) {
	ev := ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons: valid values are %v", v, AllowedProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasonsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons) IsValid() bool {
	for _, existing := range AllowedProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasonsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_cost_protect_status_get_v3.0_data_compensate_status_info_list_compensate_invalid_reasons value
func (v ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons) Ptr() *ProjectCostProtectStatusGetV30DataCompensateStatusInfoListCompensateInvalidReasons {
	return &v
}