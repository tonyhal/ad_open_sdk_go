/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10Request struct for QianchuanAdCreateV10Request
type QianchuanAdCreateV10Request struct {
	//
	AdKeywords []string `json:"ad_keywords,omitempty"`
	//
	AdvertiserId int64                                `json:"advertiser_id"`
	Audience     *QianchuanAdCreateV10RequestAudience `json:"audience,omitempty"`
	//
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	BrandId *int64 `json:"brand_id,omitempty"`
	//
	CampaignId    *int64                             `json:"campaign_id,omitempty"`
	CampaignScene *QianchuanAdCreateV10CampaignScene `json:"campaign_scene,omitempty"`
	//
	ChannelProductInfos  []*QianchuanAdCreateV10RequestChannelProductInfosInner `json:"channel_product_infos,omitempty"`
	CreativeAutoGenerate *QianchuanAdCreateV10CreativeAutoGenerate              `json:"creative_auto_generate,omitempty"`
	//
	CreativeList         []*QianchuanAdCreateV10RequestCreativeListInner `json:"creative_list,omitempty"`
	CreativeMaterialMode QianchuanAdCreateV10CreativeMaterialMode        `json:"creative_material_mode"`
	DeliverySetting      QianchuanAdCreateV10RequestDeliverySetting      `json:"delivery_setting"`
	DynamicCreative      *QianchuanAdCreateV10DynamicCreative            `json:"dynamic_creative,omitempty"`
	//
	FirstIndustryId *int64                              `json:"first_industry_id,omitempty"`
	IsHomepageHide  *QianchuanAdCreateV10IsHomepageHide `json:"is_homepage_hide,omitempty"`
	IsIntelligent   *QianchuanAdCreateV10IsIntelligent  `json:"is_intelligent,omitempty"`
	//
	Keywords       []*QianchuanAdCreateV10RequestKeywordsInner `json:"keywords,omitempty"`
	LabAdType      *QianchuanAdCreateV10LabAdType              `json:"lab_ad_type,omitempty"`
	MarketingGoal  QianchuanAdCreateV10MarketingGoal           `json:"marketing_goal"`
	MarketingScene *QianchuanAdCreateV10MarketingScene         `json:"marketing_scene,omitempty"`
	//
	Name string `json:"name"`
	//
	ProductIds               []int64                                              `json:"product_ids,omitempty"`
	ProgrammaticCreativeCard *QianchuanAdCreateV10RequestProgrammaticCreativeCard `json:"programmatic_creative_card,omitempty"`
	//
	ProgrammaticCreativeMediaList []*QianchuanAdCreateV10RequestProgrammaticCreativeMediaListInner `json:"programmatic_creative_media_list,omitempty"`
	//
	ProgrammaticCreativeTitleList []*QianchuanAdCreateV10RequestProgrammaticCreativeTitleListInner `json:"programmatic_creative_title_list,omitempty"`
	//
	SecondIndustryId *int64 `json:"second_industry_id,omitempty"`
	//
	ShopId *int64 `json:"shop_id,omitempty"`
	//
	ThirdIndustryId *int64                               `json:"third_industry_id,omitempty"`
	TrackUrl        *QianchuanAdCreateV10RequestTrackUrl `json:"track_url,omitempty"`
}