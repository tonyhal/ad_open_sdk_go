/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisAdGetV2V2ResponseDataListInnerPotentialResult
type ToolsDiagnosisAdGetV2V2ResponseDataListInnerPotentialResult struct {
	//
	BidIncrEstimate []*ToolsDiagnosisAdGetV2V2ResponseDataListInnerPotentialResultBidIncrEstimateInner `json:"bid_incr_estimate,omitempty"`
	//
	PotentialAdLevel *int64 `json:"potential_ad_level,omitempty"`
}
