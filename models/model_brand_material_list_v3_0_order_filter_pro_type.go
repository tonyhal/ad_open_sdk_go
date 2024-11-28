/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandMaterialListV30OrderFilterProType
type BrandMaterialListV30OrderFilterProType string

// List of brand_material_list_v3.0_order_filter_pro_type
const (
	FEED_BrandMaterialListV30OrderFilterProType        BrandMaterialListV30OrderFilterProType = "FEED"
	MALL_BrandMaterialListV30OrderFilterProType        BrandMaterialListV30OrderFilterProType = "MALL"
	PLANT_GRASS_BrandMaterialListV30OrderFilterProType BrandMaterialListV30OrderFilterProType = "PLANT_GRASS"
	SEARCH_BrandMaterialListV30OrderFilterProType      BrandMaterialListV30OrderFilterProType = "SEARCH"
	SPLASH_BrandMaterialListV30OrderFilterProType      BrandMaterialListV30OrderFilterProType = "SPLASH"
)

// Ptr returns reference to brand_material_list_v3.0_order_filter_pro_type value
func (v BrandMaterialListV30OrderFilterProType) Ptr() *BrandMaterialListV30OrderFilterProType {
	return &v
}