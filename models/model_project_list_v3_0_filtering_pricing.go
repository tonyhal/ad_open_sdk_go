/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30FilteringPricing
type ProjectListV30FilteringPricing string

// List of project_list_v3.0_filtering_pricing
const (
	PRICING_CPC_ProjectListV30FilteringPricing  ProjectListV30FilteringPricing = "PRICING_CPC"
	PRICING_CPM_ProjectListV30FilteringPricing  ProjectListV30FilteringPricing = "PRICING_CPM"
	PRICING_OCPC_ProjectListV30FilteringPricing ProjectListV30FilteringPricing = "PRICING_OCPC"
	PRICING_OCPM_ProjectListV30FilteringPricing ProjectListV30FilteringPricing = "PRICING_OCPM"
)

// All allowed values of ProjectListV30FilteringPricing enum
var AllowedProjectListV30FilteringPricingEnumValues = []ProjectListV30FilteringPricing{
	"PRICING_CPC",
	"PRICING_CPM",
	"PRICING_OCPC",
	"PRICING_OCPM",
}

// NewProjectListV30FilteringPricingFromValue returns a pointer to a valid ProjectListV30FilteringPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringPricingFromValue(v string) (*ProjectListV30FilteringPricing, error) {
	ev := ProjectListV30FilteringPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringPricing: valid values are %v", v, AllowedProjectListV30FilteringPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringPricing) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_pricing value
func (v ProjectListV30FilteringPricing) Ptr() *ProjectListV30FilteringPricing {
	return &v
}
