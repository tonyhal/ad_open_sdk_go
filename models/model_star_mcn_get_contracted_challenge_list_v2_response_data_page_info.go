/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnGetContractedChallengeListV2ResponseDataPageInfo
type StarMcnGetContractedChallengeListV2ResponseDataPageInfo struct {
	//
	HasMore *bool `json:"has_more,omitempty"`
	//
	Limit int64 `json:"limit"`
	//
	Page int64 `json:"page"`
	//
	TotalCount int64 `json:"total_count"`
}
