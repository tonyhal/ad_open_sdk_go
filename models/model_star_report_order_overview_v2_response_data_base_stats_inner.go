/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderOverviewV2ResponseDataBaseStatsInner struct for StarReportOrderOverviewV2ResponseDataBaseStatsInner
type StarReportOrderOverviewV2ResponseDataBaseStatsInner struct {
	ConvertStat *StarReportOrderOverviewV2ResponseDataBaseStatsInnerConvertStat `json:"convert_stat,omitempty"`
	// 订单号
	OrderId        *int64                                                             `json:"order_id,omitempty"`
	SeedStat       *StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStat       `json:"seed_stat,omitempty"`
	SpreadStat     *StarReportOrderOverviewV2ResponseDataBaseStatsInnerSpreadStat     `json:"spread_stat,omitempty"`
	WorthScoreStat *StarReportOrderOverviewV2ResponseDataBaseStatsInnerWorthScoreStat `json:"worth_score_stat,omitempty"`
}
