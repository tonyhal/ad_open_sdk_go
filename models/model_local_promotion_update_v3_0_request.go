/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalPromotionUpdateV30Request struct for LocalPromotionUpdateV30Request
type LocalPromotionUpdateV30Request struct {
	// 抖音号id
	AwemeId *string `json:"aweme_id,omitempty"`
	//
	CustomerMaterialList []*LocalPromotionUpdateV30RequestCustomerMaterialListInner `json:"customer_material_list,omitempty"`
	//
	LocalAccountId int64 `json:"local_account_id"`
	//
	Name *string `json:"name,omitempty"`
	//
	PromotionId       int64                                     `json:"promotion_id"`
	VideoHpVisibility *LocalPromotionUpdateV30VideoHpVisibility `json:"video_hp_visibility,omitempty"`
}