/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInner struct for QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInner
type QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInner struct {
	//
	Ctr *float64 `json:"ctr,omitempty"`
	//
	DirectOrderPayRate *float64 `json:"direct_order_pay_rate,omitempty"`
	//
	DirectOrderPayRoi *float64 `json:"direct_order_pay_roi,omitempty"`
	//
	DirectOrderPrepayAndPayRoi *float64 `json:"direct_order_prepay_and_pay_roi,omitempty"`
	//
	DyComment *int64 `json:"dy_comment,omitempty"`
	//
	DyFollow *int64 `json:"dy_follow,omitempty"`
	//
	DyLike *int64 `json:"dy_like,omitempty"`
	//
	DyShare     *int64                                                                                       `json:"dy_share,omitempty"`
	ImageMode   *QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeImageMode                `json:"image_mode,omitempty"`
	ProductInfo *QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInnerProductInfo `json:"product_info,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
	//
	TitleMaterial []*QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInnerTitleMaterialInner `json:"title_material,omitempty"`
	//
	TotalPlay *int64 `json:"total_play,omitempty"`
	//
	VideoMaterial []*QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInnerVideoMaterialInner `json:"video_material,omitempty"`
}
