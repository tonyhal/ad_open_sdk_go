/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicMixcutTaskResultGetV30DataStatus
type AicMixcutTaskResultGetV30DataStatus string

// List of aic_mixcut_task_result_get_v3.0_data_status
const (
	FAILED_AicMixcutTaskResultGetV30DataStatus      AicMixcutTaskResultGetV30DataStatus = "FAILED"
	PARTSUCCESS_AicMixcutTaskResultGetV30DataStatus AicMixcutTaskResultGetV30DataStatus = "PARTSUCCESS"
	PROCESSING_AicMixcutTaskResultGetV30DataStatus  AicMixcutTaskResultGetV30DataStatus = "PROCESSING"
	SUCCESS_AicMixcutTaskResultGetV30DataStatus     AicMixcutTaskResultGetV30DataStatus = "SUCCESS"
)

// Ptr returns reference to aic_mixcut_task_result_get_v3.0_data_status value
func (v AicMixcutTaskResultGetV30DataStatus) Ptr() *AicMixcutTaskResultGetV30DataStatus {
	return &v
}