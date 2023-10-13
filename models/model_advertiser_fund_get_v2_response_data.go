/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserFundGetV2ResponseData
type AdvertiserFundGetV2ResponseData struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	Balance *float64 `json:"balance,omitempty"`
	//
	Cash *float64 `json:"cash,omitempty"`
	//
	CommonGrant *float64 `json:"common_grant,omitempty"`
	//
	CompensationGrant *float64 `json:"compensation_grant,omitempty"`
	//
	CompensationValidGrant *float64 `json:"compensation_valid_grant,omitempty"`
	//
	DefaultGrant *float64 `json:"default_grant,omitempty"`
	//
	Email *string `json:"email,omitempty"`
	//
	Grant *float64 `json:"grant,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	ReturnGoodsAbs *float64 `json:"return_goods_abs,omitempty"`
	//
	ReturnGoodsCost *float64 `json:"return_goods_cost,omitempty"`
	//
	ReturnGoodsGrant *float64 `json:"return_goods_grant,omitempty"`
	//
	ReturnGoodsValidGrant *float64 `json:"return_goods_valid_grant,omitempty"`
	//
	SearchGrant *float64 `json:"search_grant,omitempty"`
	//
	UnionGrant *float64 `json:"union_grant,omitempty"`
	//
	ValidBalance *float64 `json:"valid_balance,omitempty"`
	//
	ValidCash *float64 `json:"valid_cash,omitempty"`
	//
	ValidGrant *float64 `json:"valid_grant,omitempty"`
	//
	ValidReturnGoodsAbs *float64 `json:"valid_return_goods_abs,omitempty"`
	//
	WalletId *string `json:"wallet_id,omitempty"`
	//
	WalletName *string `json:"wallet_name,omitempty"`
	//
	WalletTotalBalanceValid *string `json:"wallet_total_balance_valid,omitempty"`
}
