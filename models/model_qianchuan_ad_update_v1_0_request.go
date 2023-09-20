/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10Request struct for QianchuanAdUpdateV10Request
type QianchuanAdUpdateV10Request struct {
	//
	AdId int64 `json:"ad_id"`
	//
	AdKeywords []string `json:"ad_keywords,omitempty"`
	//
	AdvertiserId         int64                                     `json:"advertiser_id"`
	Audience             *QianchuanAdUpdateV10RequestAudience      `json:"audience,omitempty"`
	CreativeAutoGenerate *QianchuanAdUpdateV10CreativeAutoGenerate `json:"creative_auto_generate,omitempty"`
	//
	CreativeList    []*QianchuanAdUpdateV10RequestCreativeListInner `json:"creative_list,omitempty"`
	DeliverySetting QianchuanAdUpdateV10RequestDeliverySetting      `json:"delivery_setting"`
	DynamicCreative *QianchuanAdUpdateV10DynamicCreative            `json:"dynamic_creative,omitempty"`
	//
	FirstIndustryId *int64                              `json:"first_industry_id,omitempty"`
	IsHomepageHide  *QianchuanAdUpdateV10IsHomepageHide `json:"is_homepage_hide,omitempty"`
	IsIntelligent   *QianchuanAdUpdateV10IsIntelligent  `json:"is_intelligent,omitempty"`
	//
	Keywords []*QianchuanAdUpdateV10RequestKeywordsInner `json:"keywords,omitempty"`
	//
	Name                     *string                                              `json:"name,omitempty"`
	ProgrammaticCreativeCard *QianchuanAdUpdateV10RequestProgrammaticCreativeCard `json:"programmatic_creative_card,omitempty"`
	//
	ProgrammaticCreativeMediaList []*QianchuanAdUpdateV10RequestProgrammaticCreativeMediaListInner `json:"programmatic_creative_media_list,omitempty"`
	//
	ProgrammaticCreativeTitleList []*QianchuanAdUpdateV10RequestProgrammaticCreativeTitleListInner `json:"programmatic_creative_title_list,omitempty"`
	//
	SecondIndustryId *int64 `json:"second_industry_id,omitempty"`
	//
	ThirdIndustryId *int64                               `json:"third_industry_id,omitempty"`
	TrackUrl        *QianchuanAdUpdateV10RequestTrackUrl `json:"track_url,omitempty"`
}
