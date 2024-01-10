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

// CampaignCreateV2CampaignBudgetOptimization
type CampaignCreateV2CampaignBudgetOptimization string

// List of campaign_create_v2_campaign_budget_optimization
const (
	OFF_CampaignCreateV2CampaignBudgetOptimization CampaignCreateV2CampaignBudgetOptimization = "OFF"
	ON_CampaignCreateV2CampaignBudgetOptimization  CampaignCreateV2CampaignBudgetOptimization = "ON"
)

// All allowed values of CampaignCreateV2CampaignBudgetOptimization enum
var AllowedCampaignCreateV2CampaignBudgetOptimizationEnumValues = []CampaignCreateV2CampaignBudgetOptimization{
	"OFF",
	"ON",
}

// NewCampaignCreateV2CampaignBudgetOptimizationFromValue returns a pointer to a valid CampaignCreateV2CampaignBudgetOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignCreateV2CampaignBudgetOptimizationFromValue(v string) (*CampaignCreateV2CampaignBudgetOptimization, error) {
	ev := CampaignCreateV2CampaignBudgetOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignCreateV2CampaignBudgetOptimization: valid values are %v", v, AllowedCampaignCreateV2CampaignBudgetOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignCreateV2CampaignBudgetOptimization) IsValid() bool {
	for _, existing := range AllowedCampaignCreateV2CampaignBudgetOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_create_v2_campaign_budget_optimization value
func (v CampaignCreateV2CampaignBudgetOptimization) Ptr() *CampaignCreateV2CampaignBudgetOptimization {
	return &v
}
