/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisAdGetV2V2ResponseDataListInnerBidResult
type ToolsDiagnosisAdGetV2V2ResponseDataListInnerBidResult struct {
	//
	BidConclusionDetail *string `json:"bid_conclusion_detail,omitempty"`
	//
	BidConclusionTag *string `json:"bid_conclusion_tag,omitempty"`
	//
	BidSuggestion []string `json:"bid_suggestion,omitempty"`
}
