/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// MaterialStatusUpdateV30ResponseDataErrorsInner struct for MaterialStatusUpdateV30ResponseDataErrorsInner
type MaterialStatusUpdateV30ResponseDataErrorsInner struct {
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	// 失败原因
	Message *string `json:"message,omitempty"`
}
