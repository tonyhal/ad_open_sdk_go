/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryProjectV30ResponseDataProjectInfoListInner struct for QueryProjectV30ResponseDataProjectInfoListInner
type QueryProjectV30ResponseDataProjectInfoListInner struct {
	// 广告发布协议ID
	AdcId *int64 `json:"adc_id,omitempty"`
	// 广告发布协议编号
	AdcSerial *string `json:"adc_serial,omitempty"`
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 广告主名称
	AdvertiserName *string `json:"advertiser_name,omitempty"`
	// 项目金额（元）
	Amount *string `json:"amount,omitempty"`
	// 确认状态
	ConfirmStatus *int64 `json:"confirm_status,omitempty"`
	// 确认状态名称
	ConfirmStatusName *string `json:"confirm_status_name,omitempty"`
	// 主合同编号
	ContractSerial *string `json:"contract_serial,omitempty"`
	// 客户ID
	CostCustomerId *int64 `json:"cost_customer_id,omitempty"`
	// 客户名称
	CostCustomerName *string `json:"cost_customer_name,omitempty"`
	// 授信消耗
	CreditCost *string `json:"credit_cost,omitempty"`
	// 应回款时间
	Deadline *string `json:"deadline,omitempty"`
	// 赠款消耗
	GrantCost *string `json:"grant_cost,omitempty"`
	// 项目ID
	Id *int64 `json:"id,omitempty"`
	// 已开票金额
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// 项目名称
	Name *string `json:"name,omitempty"`
	// 主订单编号
	OrderIds []int64 `json:"order_ids,omitempty"`
	// 付款方式
	PayMethodName *string `json:"pay_method_name,omitempty"`
	// 平台
	Platform *int64 `json:"platform,omitempty"`
	// 平台名称
	PlatformName *string `json:"platform_name,omitempty"`
	// 预付消耗
	PrepayCost *string `json:"prepay_cost,omitempty"`
	// 项目结束时间 例如 2023-07-06
	ProjectEndDate *string `json:"project_end_date,omitempty"`
	// 项目类型分组
	ProjectGroupType *int64 `json:"project_group_type,omitempty"`
	// 项目类型分组名称
	ProjectGroupTypeName *string `json:"project_group_type_name,omitempty"`
	// 项目开始时间 例如 2023-07-06
	ProjectStartDate *string `json:"project_start_date,omitempty"`
	// 项目状态
	ProjectStatus *int64 `json:"project_status,omitempty"`
	// 项目状态名称
	ProjectStatusName *string `json:"project_status_name,omitempty"`
	// 回款状态
	ReceiptStatus *int64 `json:"receipt_status,omitempty"`
	// 回款状态名称
	ReceiptStatusName *string `json:"receipt_status_name,omitempty"`
	// 我方主体名称
	ReceiptSubjectName *string `json:"receipt_subject_name,omitempty"`
	// 现金消耗
	RechargeCost *string `json:"recharge_cost,omitempty"`
	// 待回款金额（元）
	RemainVerificationAmount *string `json:"remain_verification_amount,omitempty"`
	// 项目编号
	Serial *string `json:"serial,omitempty"`
	// 投放方式
	ServingType *int64 `json:"serving_type,omitempty"`
	// 投放方式名称
	ServingTypeName *string `json:"serving_type_name,omitempty"`
	// 代理商ID
	SettlementCustomerId *int64 `json:"settlement_customer_id,omitempty"`
	// 代理商名称
	SettlementCustomerName *string `json:"settlement_customer_name,omitempty"`
}