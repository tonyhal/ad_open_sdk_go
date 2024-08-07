/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueBridgeCallCreateV2Request struct for ToolsClueBridgeCallCreateV2Request
type ToolsClueBridgeCallCreateV2Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 被叫号码，即线索号码
	CalleeNumber string `json:"callee_number"`
	// 主叫号码，必须为11位手机号码，否则呼叫失败。
	CallerNumber string `json:"caller_number"`
	// 线索id
	ClueId int64 `json:"clue_id"`
}
