/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanTodayLiveRoomFlowPerformanceGetV10ResponseData
type QianchuanTodayLiveRoomFlowPerformanceGetV10ResponseData struct {
	//
	Activity *int64 `json:"activity,omitempty"`
	//
	DouyinShoppingCenter *int64 `json:"douyin_shopping_center,omitempty"`
	//
	EcomAdThirdQcBrand *int64 `json:"ecom_ad_third_qc_brand,omitempty"`
	//
	EcomAdThirdQcBrandOtherBidding *int64 `json:"ecom_ad_third_qc_brand_other_bidding,omitempty"`
	//
	EcomAdThirdQcOtherBrand *int64 `json:"ecom_ad_third_qc_other_brand,omitempty"`
	//
	EcomOther *int64 `json:"ecom_other,omitempty"`
	//
	GeneralSearch *int64 `json:"general_search,omitempty"`
	//
	HomepageFollow *int64 `json:"homepage_follow,omitempty"`
	//
	HomepageFresh *int64 `json:"homepage_fresh,omitempty"`
	//
	HomepageHot *int64 `json:"homepage_hot,omitempty"`
	//
	LiveMerge *int64 `json:"live_merge,omitempty"`
	//
	LivePcQianChuan *int64 `json:"live_pc_qian_chuan,omitempty"`
	//
	LiveXiaoDianQianChuan *int64 `json:"live_xiao_dian_qian_chuan,omitempty"`
	//
	OtherRecommendLive *int64 `json:"other_recommend_live,omitempty"`
	//
	OthersHomepage *int64 `json:"others_homepage,omitempty"`
	//
	TouxiSaas *int64 `json:"touxi_saas,omitempty"`
	//
	VideoToLive *int64 `json:"video_to_live,omitempty"`
}
