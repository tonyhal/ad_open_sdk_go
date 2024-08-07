/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30Request struct for ProjectUpdateV30Request
type ProjectUpdateV30Request struct {
	//
	AdvertiserId              int64                                      `json:"advertiser_id"`
	AigcDynamicCreativeSwitch *ProjectUpdateV30AigcDynamicCreativeSwitch `json:"aigc_dynamic_creative_switch,omitempty"`
	Audience                  *ProjectUpdateV30RequestAudience           `json:"audience,omitempty"`
	AudienceExtend            *ProjectUpdateV30AudienceExtend            `json:"audience_extend,omitempty"`
	DeliverySetting           *ProjectUpdateV30RequestDeliverySetting    `json:"delivery_setting,omitempty"`
	DownloadMode              *ProjectUpdateV30DownloadMode              `json:"download_mode,omitempty"`
	//
	DpaCategories []int64 `json:"dpa_categories,omitempty"`
	//
	DpaProductTarget []*ProjectUpdateV30RequestDpaProductTargetInner `json:"dpa_product_target,omitempty"`
	//
	Keywords []*ProjectUpdateV30RequestKeywordsInner `json:"keywords,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	OpenUrl *string `json:"open_url,omitempty"`
	//
	OpenUrlField *string `json:"open_url_field,omitempty"`
	//
	OpenUrlParams *string `json:"open_url_params,omitempty"`
	//
	ProjectId int64 `json:"project_id"`
	//
	SearchBidRatio  *float64                                `json:"search_bid_ratio,omitempty"`
	TrackUrlSetting *ProjectUpdateV30RequestTrackUrlSetting `json:"track_url_setting,omitempty"`
	//
	UlinkUrl *string `json:"ulink_url,omitempty"`
}
