/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30FilterAdForm
type BrandOrderListV30FilterAdForm string

// List of brand_order_list_v3.0_filter_ad_form
const (
	AWEME_LBS_BrandOrderListV30FilterAdForm                 BrandOrderListV30FilterAdForm = "AWEME_LBS"
	CONTENTRECOMMEND_BrandOrderListV30FilterAdForm          BrandOrderListV30FilterAdForm = "CONTENTRECOMMEND"
	CONTENTRECOMMEND_NO_GRASS_BrandOrderListV30FilterAdForm BrandOrderListV30FilterAdForm = "CONTENTRECOMMEND_NO_GRASS"
	CONTENTSERVICE_BrandOrderListV30FilterAdForm            BrandOrderListV30FilterAdForm = "CONTENTSERVICE"
	CONTENTSERVICE_NO_GRASS_BrandOrderListV30FilterAdForm   BrandOrderListV30FilterAdForm = "CONTENTSERVICE_NO_GRASS"
	FEED_BrandOrderListV30FilterAdForm                      BrandOrderListV30FilterAdForm = "FEED"
	FEEDSLIVE_BrandOrderListV30FilterAdForm                 BrandOrderListV30FilterAdForm = "FEEDSLIVE"
	FEEDSLIVE_AD_BrandOrderListV30FilterAdForm              BrandOrderListV30FilterAdForm = "FEEDSLIVE_AD"
	FEEDSLIVE_SERVICE_BrandOrderListV30FilterAdForm         BrandOrderListV30FilterAdForm = "FEEDSLIVE_SERVICE"
	SEARCH_BrandOrderListV30FilterAdForm                    BrandOrderListV30FilterAdForm = "SEARCH"
	SPLASH_BrandOrderListV30FilterAdForm                    BrandOrderListV30FilterAdForm = "SPLASH"
	VIDEOLIVE_AD_BrandOrderListV30FilterAdForm              BrandOrderListV30FilterAdForm = "VIDEOLIVE_AD"
	VIDEOLIVE_SERVICE_BrandOrderListV30FilterAdForm         BrandOrderListV30FilterAdForm = "VIDEOLIVE_SERVICE"
)

// Ptr returns reference to brand_order_list_v3.0_filter_ad_form value
func (v BrandOrderListV30FilterAdForm) Ptr() *BrandOrderListV30FilterAdForm {
	return &v
}