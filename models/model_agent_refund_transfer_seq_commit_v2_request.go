/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentRefundTransferSeqCommitV2Request struct for AgentRefundTransferSeqCommitV2Request
type AgentRefundTransferSeqCommitV2Request struct {
	// 代理商ID
	AgentId int64 `json:"agent_id"`
	// 退款流水号，第一步请求会返回
	RefundSeq int64 `json:"refund_seq"`
}
