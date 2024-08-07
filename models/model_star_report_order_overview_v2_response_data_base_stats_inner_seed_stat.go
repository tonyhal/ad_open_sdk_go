/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStat 种草价值
type StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStat struct {
	EmotionDistribution *StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatEmotionDistribution `json:"emotion_distribution,omitempty"`
	// 热词分类
	HotKeywords []*StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatHotKeywordsInner `json:"hot_keywords,omitempty"`
	// 搜索指数
	SearchIndex *int64 `json:"search_index,omitempty"`
}
