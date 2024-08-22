/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomAsyncTaskCreateV30DataTopic
type ReportCustomAsyncTaskCreateV30DataTopic string

// List of report_custom_async_task_create_v3.0_data_topic
const (
	BASIC_DATA_ReportCustomAsyncTaskCreateV30DataTopic    ReportCustomAsyncTaskCreateV30DataTopic = "BASIC_DATA"
	BIDWORD_DATA_ReportCustomAsyncTaskCreateV30DataTopic  ReportCustomAsyncTaskCreateV30DataTopic = "BIDWORD_DATA"
	MATERIAL_DATA_ReportCustomAsyncTaskCreateV30DataTopic ReportCustomAsyncTaskCreateV30DataTopic = "MATERIAL_DATA"
	PRODUCT_DATA_ReportCustomAsyncTaskCreateV30DataTopic  ReportCustomAsyncTaskCreateV30DataTopic = "PRODUCT_DATA"
	QUERY_DATA_ReportCustomAsyncTaskCreateV30DataTopic    ReportCustomAsyncTaskCreateV30DataTopic = "QUERY_DATA"
)

// Ptr returns reference to report_custom_async_task_create_v3.0_data_topic value
func (v ReportCustomAsyncTaskCreateV30DataTopic) Ptr() *ReportCustomAsyncTaskCreateV30DataTopic {
	return &v
}
