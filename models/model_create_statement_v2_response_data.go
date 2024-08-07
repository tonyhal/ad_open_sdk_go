/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreateStatementV2ResponseData
type CreateStatementV2ResponseData struct {
	// 创建失败原因
	ErrorReason *string                      `json:"error_reason,omitempty"`
	Result      *CreateStatementV2DataResult `json:"result,omitempty"`
	// 结算单编号
	StatementSerial *string `json:"statement_serial,omitempty"`
}
