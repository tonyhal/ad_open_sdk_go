/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// MaterialStatusUpdateV30RequestDataInner struct for MaterialStatusUpdateV30RequestDataInner
type MaterialStatusUpdateV30RequestDataInner struct {
	//
	MaterialId int64                                `json:"material_id"`
	OptStatus  MaterialStatusUpdateV30DataOptStatus `json:"opt_status"`
}