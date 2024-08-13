/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryRebateBalanceV2ResponseDataRebatesInner struct for QueryRebateBalanceV2ResponseDataRebatesInner
type QueryRebateBalanceV2ResponseDataRebatesInner struct {
	// 代理商ID
	AgentId *int64 `json:"agent_id,omitempty"`
	// 代理商名称
	AgentName *string `json:"agent_name,omitempty"`
	// 返点总金额 单位 元 保留2位小数
	Amount *float64 `json:"amount,omitempty"`
	// 回收状态
	ArchiveStatus *int32 `json:"archive_status,omitempty"`
	// 回收状态名称
	ArchiveStatusName *string `json:"archive_status_name,omitempty"`
	// 合同名称
	ContractName *string `json:"contract_name,omitempty"`
	// 合同编号
	ContractSerial *string `json:"contract_serial,omitempty"`
	// 媒体签约主体
	ContractSubjectName *string `json:"contract_subject_name,omitempty"`
	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 当前审批节点名称
	CurrentApprovalStatusName *string `json:"current_approval_status_name,omitempty"`
	// 冻结金额 单位 元 保留2位小数
	FrozenAmount *float64 `json:"frozen_amount,omitempty"`
	// 实际开票金额 单位 元 保留2位小数
	InvoiceAmount *float64 `json:"invoice_amount,omitempty"`
	// 可开票金额 单位 元 保留2位小数
	InvoiceAmountRemain *float64 `json:"invoice_amount_remain,omitempty"`
	// 已申请开票金额 单位 元 保留2位小数
	InvoiceFrozenAmount *float64 `json:"invoice_frozen_amount,omitempty"`
	// 开票状态
	InvoiceStatus *int32 `json:"invoice_status,omitempty"`
	// 开票状态名称
	InvoiceStatusName *string `json:"invoice_status_name,omitempty"`
	// 签章类型名称
	IsOnlineStamp *string `json:"is_online_stamp,omitempty"`
	// 所属季度／月
	MonthQuarter *int32 `json:"month_quarter,omitempty"`
	// 所属季度／月名称
	MonthQuarterName *string `json:"month_quarter_name,omitempty"`
	// 平台/业务线
	Platform *int64 `json:"platform,omitempty"`
	// 平台/业务线名称
	PlatformName *string `json:"platform_name,omitempty"`
	// 返点流水ID
	RebateBalanceId *int64 `json:"rebate_balance_id,omitempty"`
	// 返点流水编号 （与平台披露编号一致，建议使用）
	RebateBalanceSerial *string `json:"rebate_balance_serial,omitempty"`
	// 盖章状态
	StampStatus *int32 `json:"stamp_status,omitempty"`
	// 盖章状态名称
	StampStatusName *string `json:"stamp_status_name,omitempty"`
	// 是否标准
	Standard *int32 `json:"standard,omitempty"`
	// 是否标准名称
	StandardName *string `json:"standard_name,omitempty"`
	// 返点状态
	Status *int32 `json:"status,omitempty"`
	// 返点状态名称
	StatusName *string `json:"status_name,omitempty"`
	// 使用类型名称
	UseTypeNames []string `json:"use_type_names,omitempty"`
	// 使用类型
	UseTypes []int32 `json:"use_types,omitempty"`
	// 已使用金额 单位 元 保留2位小数
	UsedAmount *float64 `json:"used_amount,omitempty"`
	// 结算周期年
	Year *int32 `json:"year,omitempty"`
}
