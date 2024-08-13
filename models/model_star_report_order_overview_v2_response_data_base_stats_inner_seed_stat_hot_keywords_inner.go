/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInner struct for StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInner
type StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInner struct {
	// 全部
	All []*StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInnerAllInner `json:"all,omitempty"`
	// 品牌
	Brand []*StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInnerBrandInner `json:"brand,omitempty"`
	// 角色
	Character []*StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInnerCharacterInner `json:"character,omitempty"`
	// 痛点
	Pain []*StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInnerPainInner `json:"pain,omitempty"`
	// 产品
	Product []*StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInnerProductInner `json:"product,omitempty"`
	// 场景
	Scene []*StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInnerSceneInner `json:"scene,omitempty"`
}
