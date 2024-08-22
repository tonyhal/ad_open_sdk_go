/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomAsyncTaskCreateV30DataTaskStatus
type ReportCustomAsyncTaskCreateV30DataTaskStatus string

// List of report_custom_async_task_create_v3.0_data_task_status
const (
	ASYNC_TASK_STATUS_COMPLETED_ReportCustomAsyncTaskCreateV30DataTaskStatus ReportCustomAsyncTaskCreateV30DataTaskStatus = "ASYNC_TASK_STATUS_COMPLETED"
	ASYNC_TASK_STATUS_CREATED_ReportCustomAsyncTaskCreateV30DataTaskStatus   ReportCustomAsyncTaskCreateV30DataTaskStatus = "ASYNC_TASK_STATUS_CREATED"
	ASYNC_TASK_STATUS_EXECUTING_ReportCustomAsyncTaskCreateV30DataTaskStatus ReportCustomAsyncTaskCreateV30DataTaskStatus = "ASYNC_TASK_STATUS_EXECUTING"
	ASYNC_TASK_STATUS_FAILED_ReportCustomAsyncTaskCreateV30DataTaskStatus    ReportCustomAsyncTaskCreateV30DataTaskStatus = "ASYNC_TASK_STATUS_FAILED"
)

// Ptr returns reference to report_custom_async_task_create_v3.0_data_task_status value
func (v ReportCustomAsyncTaskCreateV30DataTaskStatus) Ptr() *ReportCustomAsyncTaskCreateV30DataTaskStatus {
	return &v
}
