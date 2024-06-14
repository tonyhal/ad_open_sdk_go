/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2Request struct for StarDemandCreateChallengeV2Request
type StarDemandCreateChallengeV2Request struct {
	ChallengeInfo StarDemandCreateChallengeV2RequestChallengeInfo `json:"challenge_info"`
	DemandInfo    StarDemandCreateChallengeV2RequestDemandInfo    `json:"demand_info"`
	// 星图客户ID
	StarId int64 `json:"star_id"`
}
