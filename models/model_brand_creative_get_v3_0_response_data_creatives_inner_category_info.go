/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCategoryInfo 分类信息
type BrandCreativeGetV30ResponseDataCreativesInnerCategoryInfo struct {
	// 创意分类
	AdCategory *int64 `json:"ad_category,omitempty"`
	// 创意标签
	AdKeywords []string `json:"ad_keywords,omitempty"`
	// 品牌名称
	BrandName *string `json:"brand_name,omitempty"`
	// 品牌名称Id
	BrandNameId *int64 `json:"brand_name_id,omitempty"`
	// cdp返回的品牌名称id
	CdpBrandId *int64 `json:"cdp_brand_id,omitempty"`
	// 创意分类（行业3.0版）
	IndustryV3 *int64 `json:"industry_v3,omitempty"`
	// 产品名称
	ProductName []string `json:"product_name,omitempty"`
	// 品牌二级名称Ids
	SubBrandNameIds []int64 `json:"sub_brand_name_ids,omitempty"`
	// 品牌二级名称
	SubBrandNames []string `json:"sub_brand_names,omitempty"`
	// 云图分类
	YuntuCategoryId *int64 `json:"yuntu_category_id,omitempty"`
	// 云图一级分类ID
	YuntuMainIndustryId *int64 `json:"yuntu_main_industry_id,omitempty"`
}
