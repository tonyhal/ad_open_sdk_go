/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmExpandChallengeProviderV2RequestOmParticipateProviderRange
type StarDemandOmExpandChallengeProviderV2RequestOmParticipateProviderRange struct {
	OperationType *StarDemandOmExpandChallengeProviderV2OmParticipateProviderRangeOperationType `json:"operation_type,omitempty"`
	//
	ProviderIds []int64 `json:"provider_ids,omitempty"`
}
