/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueExtInfoCallbackV2Request struct for ToolsClueExtInfoCallbackV2Request
type ToolsClueExtInfoCallbackV2Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 回传时间，秒级时间戳
	CallbackTime int64                                `json:"callback_time"`
	ClueEvent    *ToolsClueExtInfoCallbackV2ClueEvent `json:"clue_event,omitempty"`
	// 线索id
	ClueId int64                            `json:"clue_id"`
	ExtKey ToolsClueExtInfoCallbackV2ExtKey `json:"ext_key"`
	// 扩展信息value
	ExtValue string `json:"ext_value"`
	// 完整阶段描述（严格有序）<br>当扩展信息中传【客户阶段】或【订单状态】时，将完整的客户阶段和订单状态传入该字段<br>例：客户完整阶段【未联系-产生意向-可跟进客户-邀约上门-成交】
	ExtValueDetail []string `json:"ext_value_detail,omitempty"`
}
