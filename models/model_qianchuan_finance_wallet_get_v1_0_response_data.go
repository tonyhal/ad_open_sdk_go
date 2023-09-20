/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFinanceWalletGetV10ResponseData
type QianchuanFinanceWalletGetV10ResponseData struct {
	// 品牌余额
	BrandBalance *int64 `json:"brand_balance,omitempty"`
	// 品牌余额-不可用余额
	BrandBalanceInvalid *int64 `json:"brand_balance_invalid,omitempty"`
	// 品牌余额-不可用余额-冻结
	BrandBalanceInvalidFrozen *int64 `json:"brand_balance_invalid_frozen,omitempty"`
	// 品牌余额-可用余额
	BrandBalanceValid *int64 `json:"brand_balance_valid,omitempty"`
	// 品牌余额-可用余额-赠款
	BrandBalanceValidGrant *int64 `json:"brand_balance_valid_grant,omitempty"`
	// 品牌余额-可用余额-非赠款
	BrandBalanceValidNonGrant *int64 `json:"brand_balance_valid_non_grant,omitempty"`
	// 赠款余额-巨量信息流广告-可用
	CommonValidGrantBalance *int64 `json:"common_valid_grant_balance,omitempty"`
	// 消返红包余额
	DeductionCouponBalance *int64 `json:"deduction_coupon_balance,omitempty"`
	// 消返红包余额（通用）
	DeductionCouponBalanceAll *int64 `json:"deduction_coupon_balance_all,omitempty"`
	// 消返红包余额（代投）
	DeductionCouponBalanceOther *int64 `json:"deduction_coupon_balance_other,omitempty"`
	// 消返红包余额（自投）
	DeductionCouponBalanceSelf *int64 `json:"deduction_coupon_balance_self,omitempty"`
	// 赠款余额-通用-可用
	DefaultValidGrantBalance *int64 `json:"default_valid_grant_balance,omitempty"`
	// 通用余额-不可用余额
	GeneralBalanceInvalid *int64 `json:"general_balance_invalid,omitempty"`
	// 通用余额-不可用余额-冻结
	GeneralBalanceInvalidFrozen *int64 `json:"general_balance_invalid_frozen,omitempty"`
	// 通用余额-不可用余额-随心推已下单
	GeneralBalanceInvalidOrder *int64 `json:"general_balance_invalid_order,omitempty"`
	// 通用余额-可用余额
	GeneralBalanceValid *int64 `json:"general_balance_valid,omitempty"`
	// 通用余额-可用余额-赠款-巨量信息流广告
	GeneralBalanceValidGrantCommon *int64 `json:"general_balance_valid_grant_common,omitempty"`
	// 通用余额-可用余额-赠款-通用
	GeneralBalanceValidGrantDefault *int64 `json:"general_balance_valid_grant_default,omitempty"`
	// 通用余额-可用余额-赠款-巨量搜索广告
	GeneralBalanceValidGrantSearch *int64 `json:"general_balance_valid_grant_search,omitempty"`
	// 通用余额-可用余额-赠款-穿山甲
	GeneralBalanceValidGrantUnion *int64 `json:"general_balance_valid_grant_union,omitempty"`
	// 通用余额-可用余额-非赠款
	GeneralBalanceValidNonGrant *int64 `json:"general_balance_valid_non_grant,omitempty"`
	// 通用余额
	GeneralTotalBalance *int64 `json:"general_total_balance,omitempty"`
	// 赠款余额
	GrantBalance *int64 `json:"grant_balance,omitempty"`
	// 15天内赠款到期金额
	GrantExpiring *int64 `json:"grant_expiring,omitempty"`
	// 赠款余额-穿山甲-可用
	SearchValidGrantBalance *int64 `json:"search_valid_grant_balance,omitempty"`
	// 共享赠款余额
	ShareBalance *int64 `json:"share_balance,omitempty"`
	// 共享赠款余额-30天内到期余额
	ShareBalanceExpiring *int64 `json:"share_balance_expiring,omitempty"`
	// 共享赠款余额-可用余额
	ShareBalanceValid *int64 `json:"share_balance_valid,omitempty"`
	// 共享赠款余额-可用余额-赠款-巨量信息流广告
	ShareBalanceValidGrantCommon *int64 `json:"share_balance_valid_grant_common,omitempty"`
	// 共享赠款余额-可用余额-赠款-通用
	ShareBalanceValidGrantDefault *int64 `json:"share_balance_valid_grant_default,omitempty"`
	// 共享赠款余额-可用余额-赠款-巨量搜索广告
	ShareBalanceValidGrantSearch *int64 `json:"share_balance_valid_grant_search,omitempty"`
	// 共享赠款余额-可用余额-赠款-穿山甲
	ShareBalanceValidGrantUnion *int64 `json:"share_balance_valid_grant_union,omitempty"`
	// 共享赠款余额到期详情
	ShareExpiringDetailList []*QianchuanFinanceWalletGetV10ResponseDataShareExpiringDetailListInner `json:"share_expiring_detail_list,omitempty"`
	// 账户总余额
	TotalBalanceAbs *int64 `json:"total_balance_abs,omitempty"`
	// 赠款余额-穿山甲-可用
	UnionValidGrantBalance *int64 `json:"union_valid_grant_balance,omitempty"`
}