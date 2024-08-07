/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreateStatementV2Request struct for CreateStatementV2Request
type CreateStatementV2Request struct {
	// 若传入，则校验项目账户ID是否一致
	AccountId *int64 `json:"account_id,omitempty"`
	// 代理商账户ID列表
	AgentIds []int64 `json:"agent_ids"`
	// 合同编号
	ContractSerial string `json:"contract_serial"`
	// 文件ID，当结算单模版类型为非标，则必须传入文件
	FileId *int64 `json:"file_id,omitempty"`
	// 结算单名称
	Name     *string                   `json:"name,omitempty"`
	Platform CreateStatementV2Platform `json:"platform"`
	// 项目编号列表 200以内
	ProjectSerialList []string `json:"project_serial_list"`
	// 请求ID，要求唯一，幂等用，必传
	RequestId    string                        `json:"request_id"`
	TemplateType CreateStatementV2TemplateType `json:"template_type"`
}
