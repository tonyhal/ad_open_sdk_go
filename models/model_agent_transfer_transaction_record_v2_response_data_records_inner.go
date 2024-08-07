/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentTransferTransactionRecordV2ResponseDataRecordsInner struct for AgentTransferTransactionRecordV2ResponseDataRecordsInner
type AgentTransferTransactionRecordV2ResponseDataRecordsInner struct {
	// 总金额（元）
	Amount *float64 `json:"amount,omitempty"`
	// 非赠款金额(元)
	Cash *float64 `json:"cash,omitempty"`
	// 授信金额额（元）
	CreditAmount *float64 `json:"credit_amount,omitempty"`
	// 赠款金额(元)
	Grants *float64 `json:"grants,omitempty"`
	// 转账时间
	ModifyTime *string `json:"modify_time,omitempty"`
	// 操作人名称
	Operator *string `json:"operator,omitempty"`
	// 操作人id
	OperatorId *string `json:"operator_id,omitempty"`
	// 转入方账户ID
	Payee *string `json:"payee,omitempty"`
	// 转入方客户ID
	PayeeCustomerId *string `json:"payee_customer_id,omitempty"`
	// 转入方客户名称
	PayeeCustomerName *string `json:"payee_customer_name,omitempty"`
	// 转入方一代账户ID
	PayeeFirstAdAgentId *string `json:"payee_first_ad_agent_id,omitempty"`
	// 转入方一代账户名称
	PayeeFirstAdAgentName *string `json:"payee_first_ad_agent_name,omitempty"`
	// 转入方账户名称
	PayeeName *string                                               `json:"payee_name,omitempty"`
	PayeeRole *AgentTransferTransactionRecordV2DataRecordsPayeeRole `json:"payee_role,omitempty"`
	// 转入方二代账户ID
	PayeeSecondAdAgentId *string                                              `json:"payee_second_ad_agent_id,omitempty"`
	Platform             *AgentTransferTransactionRecordV2DataRecordsPlatform `json:"platform,omitempty"`
	// 预付金额（元）
	PrepayAmount *float64 `json:"prepay_amount,omitempty"`
	// 转账备注
	Remark *string `json:"remark,omitempty"`
	// 转账时间转出方账户ID
	Remitter *string `json:"remitter,omitempty"`
	// 转出方客户ID
	RemitterCustomerId *string `json:"remitter_customer_id,omitempty"`
	// 转出方客户名称
	RemitterCustomerName *string `json:"remitter_customer_name,omitempty"`
	// 转出方一代账户ID
	RemitterFirstAdAgentId *string `json:"remitter_first_ad_agent_id,omitempty"`
	// 转出方一代账户名称
	RemitterFirstAdAgentName *string `json:"remitter_first_ad_agent_name,omitempty"`
	// 转出方账户名称
	RemitterName *string                                                  `json:"remitter_name,omitempty"`
	RemitterRole *AgentTransferTransactionRecordV2DataRecordsRemitterRole `json:"remitter_role,omitempty"`
	// 转出方二代账户ID
	RemitterSecondAdAgentId *string `json:"remitter_second_ad_agent_id,omitempty"`
	// 转账编号
	TransferOrderSerial     *string                                                             `json:"transfer_order_serial,omitempty"`
	TransferTargetPayStatus *AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus `json:"transfer_target_pay_status,omitempty"`
	TransferType            *AgentTransferTransactionRecordV2DataRecordsTransferType            `json:"transfer_type,omitempty"`
}
