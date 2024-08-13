/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletDailyStatGetV30ResponseDataResultsInner struct for SharedWalletDailyStatGetV30ResponseDataResultsInner
type SharedWalletDailyStatGetV30ResponseDataResultsInner struct {
	// 日终结余(单位元）
	Balance *float64 `json:"balance,omitempty"`
	// 总消耗(单位元)
	Cost *float64 `json:"cost,omitempty"`
	// 授信日终结余(单位元）
	CreditBalance *float64 `json:"credit_balance,omitempty"`
	// 总支出(单位元)
	Expenses *float64 `json:"expenses,omitempty"`
	// 总存入(单位元)
	Incomes *float64 `json:"incomes,omitempty"`
	// 预付日终结余(单位元）
	PrepayBalance *float64 `json:"prepay_balance,omitempty"`
	// 共享钱包ID
	SharedWalletId *int64 `json:"shared_wallet_id,omitempty"`
	// 共享钱包名称
	SharedWalletName *string `json:"shared_wallet_name,omitempty"`
	// 日期,精确到天，格式YYYY-MM-DD
	TransactionDate *string                                           `json:"transaction_date,omitempty"`
	WalletType      *SharedWalletDailyStatGetV30DataResultsWalletType `json:"wallet_type,omitempty"`
}
