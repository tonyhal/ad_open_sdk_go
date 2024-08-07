/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestPromotionMaterialsProductInfo
type PromotionCreateV30RequestPromotionMaterialsProductInfo struct {
	//
	ImageIds []string `json:"image_ids,omitempty"`
	//
	ProductImageFields []string                                                         `json:"product_image_fields,omitempty"`
	ProductImageType   *PromotionCreateV30PromotionMaterialsProductInfoProductImageType `json:"product_image_type,omitempty"`
	//
	ProductNameFields []string                                                        `json:"product_name_fields,omitempty"`
	ProductNameType   *PromotionCreateV30PromotionMaterialsProductInfoProductNameType `json:"product_name_type,omitempty"`
	//
	ProductSellingPointFields []string                                                                `json:"product_selling_point_fields,omitempty"`
	ProductSellingPointType   *PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType `json:"product_selling_point_type,omitempty"`
	//
	SellingPoints []string `json:"selling_points,omitempty"`
	//
	Titles []string `json:"titles,omitempty"`
}
