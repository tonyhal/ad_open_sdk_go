/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseListV10ResponseDataProductListInner struct for QianchuanProductAnalyseListV10ResponseDataProductListInner
type QianchuanProductAnalyseListV10ResponseDataProductListInner struct {
	//
	CategoryName *string `json:"category_name,omitempty"`
	//
	Ctr *float64 `json:"ctr,omitempty"`
	//
	DirectOrderPayCost *float64 `json:"direct_order_pay_cost,omitempty"`
	//
	DirectOrderPayCostPerOrder *float64 `json:"direct_order_pay_cost_per_order,omitempty"`
	//
	DirectOrderPayCount *int64 `json:"direct_order_pay_count,omitempty"`
	//
	DirectOrderPayGmv *float64 `json:"direct_order_pay_gmv,omitempty"`
	//
	DirectOrderPayRate *float64 `json:"direct_order_pay_rate,omitempty"`
	//
	DirectOrderPrepayAndPayRoi *float64 `json:"direct_order_prepay_and_pay_roi,omitempty"`
	//
	DiscountPrice *float64 `json:"discount_price,omitempty"`
	//
	MarketPrice *float64 `json:"market_price,omitempty"`
	//
	PayOrderCouponAmount *float64 `json:"pay_order_coupon_amount,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
	//
	ProductImg []string `json:"product_img,omitempty"`
	//
	ProductName *string `json:"product_name,omitempty"`
	//
	ProductRate *float64 `json:"product_rate,omitempty"`
	//
	Recommendation *string `json:"recommendation,omitempty"`
	//
	SaleTime *int64 `json:"sale_time,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
}
