/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30Request struct for PromotionUpdateV30Request
type PromotionUpdateV30Request struct {
	//
	AdvertiserId      int64                                `json:"advertiser_id"`
	AutoExtendTraffic *PromotionUpdateV30AutoExtendTraffic `json:"auto_extend_traffic,omitempty"`
	//
	Bid       *float32                            `json:"bid,omitempty"`
	BrandInfo *PromotionUpdateV30RequestBrandInfo `json:"brand_info,omitempty"`
	//
	Budget *float32 `json:"budget,omitempty"`
	//
	ConfigId *int64 `json:"config_id,omitempty"`
	//
	CpaBid                     *float32                                      `json:"cpa_bid,omitempty"`
	CreativeAutoGenerateSwitch *PromotionUpdateV30CreativeAutoGenerateSwitch `json:"creative_auto_generate_switch,omitempty"`
	//
	DeepCpabid       *float32                            `json:"deep_cpabid,omitempty"`
	IsCommentDisable *PromotionUpdateV30IsCommentDisable `json:"is_comment_disable,omitempty"`
	//
	Keywords []*PromotionUpdateV30RequestKeywordsInner `json:"keywords,omitempty"`
	//
	Name          string                                  `json:"name"`
	NativeSetting *PromotionUpdateV30RequestNativeSetting `json:"native_setting,omitempty"`
	//
	PromotionId        int64                                       `json:"promotion_id"`
	PromotionMaterials PromotionUpdateV30RequestPromotionMaterials `json:"promotion_materials"`
	//
	RoiGoal *float32 `json:"roi_goal,omitempty"`
	//
	Source *string `json:"source,omitempty"`
}
