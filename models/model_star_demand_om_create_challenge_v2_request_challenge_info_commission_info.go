/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmCreateChallengeV2RequestChallengeInfoCommissionInfo 分佣信息
type StarDemandOmCreateChallengeV2RequestChallengeInfoCommissionInfo struct {
	// IAA分佣比例
	AdCommissionRate *int64 `json:"ad_commission_rate,omitempty"`
	// IAP分佣比例
	PayCommissionRate *int64 `json:"pay_commission_rate,omitempty"`
}
