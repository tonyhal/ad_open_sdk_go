/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeCard
type QianchuanAdCreateV10RequestMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeCard struct {
	//
	PromotionCardActionButton            *string                                                                                                                       `json:"promotion_card_action_button,omitempty"`
	PromotionCardButtonSmartOptimization *QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeCardPromotionCardButtonSmartOptimization `json:"promotion_card_button_smart_optimization,omitempty"`
	//
	PromotionCardImageId *string `json:"promotion_card_image_id,omitempty"`
	//
	PromotionCardSellingPoints []string `json:"promotion_card_selling_points,omitempty"`
	//
	PromotionCardTitle *string `json:"promotion_card_title,omitempty"`
}