/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportCustomDataTopicDailyReportV2ResponseDataStatsInnerDataInner struct for StarReportCustomDataTopicDailyReportV2ResponseDataStatsInnerDataInner
type StarReportCustomDataTopicDailyReportV2ResponseDataStatsInnerDataInner struct {
	DataTopic StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic `json:"data_topic"`
	// 数据主题下的数据指标，一个数据主题对应一个或多个数据指标。不存在的数据指标将缺省。
	Metrics []*StarReportCustomDataTopicDailyReportV2ResponseDataStatsInnerDataInnerMetricsInner `json:"metrics"`
}