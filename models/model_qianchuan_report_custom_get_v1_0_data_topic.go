/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportCustomGetV10DataTopic
type QianchuanReportCustomGetV10DataTopic string

// List of qianchuan_report_custom_get_v1.0_data_topic
const (
	ECP_BASIC_DATA_QianchuanReportCustomGetV10DataTopic                 QianchuanReportCustomGetV10DataTopic = "ECP_BASIC_DATA"
	SITE_PROMOTION_POST_DATA_LIVE_QianchuanReportCustomGetV10DataTopic  QianchuanReportCustomGetV10DataTopic = "SITE_PROMOTION_POST_DATA_LIVE"
	SITE_PROMOTION_POST_DATA_OTHER_QianchuanReportCustomGetV10DataTopic QianchuanReportCustomGetV10DataTopic = "SITE_PROMOTION_POST_DATA_OTHER"
	SITE_PROMOTION_POST_DATA_TITLE_QianchuanReportCustomGetV10DataTopic QianchuanReportCustomGetV10DataTopic = "SITE_PROMOTION_POST_DATA_TITLE"
	SITE_PROMOTION_POST_DATA_VIDEO_QianchuanReportCustomGetV10DataTopic QianchuanReportCustomGetV10DataTopic = "SITE_PROMOTION_POST_DATA_VIDEO"
)

// Ptr returns reference to qianchuan_report_custom_get_v1.0_data_topic value
func (v QianchuanReportCustomGetV10DataTopic) Ptr() *QianchuanReportCustomGetV10DataTopic {
	return &v
}
