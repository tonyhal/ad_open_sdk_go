/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionListV10Fields
type QianchuanUniPromotionListV10Fields string

// List of qianchuan_uni_promotion_list_v1.0_fields
const (
	STAT_COST_QianchuanUniPromotionListV10Fields                         QianchuanUniPromotionListV10Fields = "stat_cost"
	TOTAL_COST_PER_PAY_ORDER_FOR_ROI2_QianchuanUniPromotionListV10Fields QianchuanUniPromotionListV10Fields = "total_cost_per_pay_order_for_roi2"
	TOTAL_PAY_ORDER_COUNT_FOR_ROI2_QianchuanUniPromotionListV10Fields    QianchuanUniPromotionListV10Fields = "total_pay_order_count_for_roi2"
	TOTAL_PAY_ORDER_GMV_FOR_ROI2_QianchuanUniPromotionListV10Fields      QianchuanUniPromotionListV10Fields = "total_pay_order_gmv_for_roi2"
	TOTAL_PREPAY_AND_PAY_ORDER_ROI2_QianchuanUniPromotionListV10Fields   QianchuanUniPromotionListV10Fields = "total_prepay_and_pay_order_roi2"
	TOTAL_PREPAY_ORDER_COUNT_FOR_ROI2_QianchuanUniPromotionListV10Fields QianchuanUniPromotionListV10Fields = "total_prepay_order_count_for_roi2"
	TOTAL_PREPAY_ORDER_GMV_FOR_ROI2_QianchuanUniPromotionListV10Fields   QianchuanUniPromotionListV10Fields = "total_prepay_order_gmv_for_roi2"
)

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_fields value
func (v QianchuanUniPromotionListV10Fields) Ptr() *QianchuanUniPromotionListV10Fields {
	return &v
}
