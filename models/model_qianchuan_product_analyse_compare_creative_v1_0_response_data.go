/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseCompareCreativeV10ResponseData
type QianchuanProductAnalyseCompareCreativeV10ResponseData struct {
	//
	OwnProductCreative []*QianchuanProductAnalyseCompareCreativeV10ResponseDataOwnProductCreativeInner `json:"own_product_creative,omitempty"`
	//
	SimilarProductCreative []*QianchuanProductAnalyseCompareCreativeV10ResponseDataSimilarProductCreativeInner `json:"similar_product_creative,omitempty"`
}
