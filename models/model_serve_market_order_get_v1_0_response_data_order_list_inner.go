/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ServeMarketOrderGetV10ResponseDataOrderListInner struct for ServeMarketOrderGetV10ResponseDataOrderListInner
type ServeMarketOrderGetV10ResponseDataOrderListInner struct {
	//
	AppActiveDays *int64 `json:"app_active_days,omitempty"`
	//
	AppAvailableUserIds []int64 `json:"app_available_user_ids,omitempty"`
	//
	AppLimitUserCount *int64 `json:"app_limit_user_count,omitempty"`
	//
	BeginTime *int64 `json:"begin_time,omitempty"`
	//
	CreateTime *int64 `json:"create_time,omitempty"`
	//
	EndTime *int64 `json:"end_time,omitempty"`
	//
	Fee      *int64                                                    `json:"fee,omitempty"`
	Function *ServeMarketOrderGetV10ResponseDataOrderListInnerFunction `json:"function,omitempty"`
	// 是否是付费功能点类型的sku
	IsFunc *bool `json:"is_func,omitempty"`
	//
	OrderId     *int64                                          `json:"order_id,omitempty"`
	OrderStatus *ServeMarketOrderGetV10DataOrderListOrderStatus `json:"order_status,omitempty"`
	//
	OriginPrice *int64 `json:"origin_price,omitempty"`
	//
	PaidUserId *int64 `json:"paid_user_id,omitempty"`
	//
	PayTime *int64 `json:"pay_time,omitempty"`
	//
	SkuDescription *string `json:"sku_description,omitempty"`
	//
	SkuId   *int64                                      `json:"sku_id,omitempty"`
	SkuType *ServeMarketOrderGetV10DataOrderListSkuType `json:"sku_type,omitempty"`
	//
	Specs []*ServeMarketOrderGetV10ResponseDataOrderListInnerSpecsInner `json:"specs,omitempty"`
}
