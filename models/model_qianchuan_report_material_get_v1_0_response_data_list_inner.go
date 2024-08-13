/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportMaterialGetV10ResponseDataListInner struct for QianchuanReportMaterialGetV10ResponseDataListInner
type QianchuanReportMaterialGetV10ResponseDataListInner struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AnalysisType   []string                                             `json:"analysis_type,omitempty"`
	CarouselSource *QianchuanReportMaterialGetV10DataListCarouselSource `json:"carousel_source,omitempty"`
	//
	CreateData *string `json:"create_data,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	Duration    *float64                                                  `json:"duration,omitempty"`
	Fields      *QianchuanReportMaterialGetV10ResponseDataListInnerFields `json:"fields,omitempty"`
	ImageSource *QianchuanReportMaterialGetV10DataListImageSource         `json:"image_source,omitempty"`
	//
	MaterialId   *int64                                             `json:"material_id,omitempty"`
	MaterialMode *QianchuanReportMaterialGetV10DataListMaterialMode `json:"material_mode,omitempty"`
	MaterialType QianchuanReportMaterialGetV10DataListMaterialType  `json:"material_type"`
	//
	RelatedAdCnt *int64 `json:"related_ad_cnt,omitempty"`
	//
	RelatedAdIds []int64 `json:"related_ad_ids,omitempty"`
	//
	RelatedCreativeCnt *int64 `json:"related_creative_cnt,omitempty"`
	//
	RelatedCreativeIds []int64 `json:"related_creative_ids,omitempty"`
	//
	Tags        *string                                           `json:"tags,omitempty"`
	VideoSource *QianchuanReportMaterialGetV10DataListVideoSource `json:"video_source,omitempty"`
}
