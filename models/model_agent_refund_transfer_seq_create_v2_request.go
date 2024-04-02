/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentRefundTransferSeqCreateV2Request struct for AgentRefundTransferSeqCreateV2Request
type AgentRefundTransferSeqCreateV2Request struct {
	// 广告主账户ID
	AccountId int64 `json:"account_id"`
	// 代理商账户ID
	AgentId int64 `json:"agent_id"`
	// 转账金额，单位为分
	Amount       float64                                    `json:"amount"`
	TransferType AgentRefundTransferSeqCreateV2TransferType `json:"transfer_type"`
}
