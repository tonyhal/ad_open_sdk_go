/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestCreateV2RequestObjectListInner struct for ToolsAbTestCreateV2RequestObjectListInner
type ToolsAbTestCreateV2RequestObjectListInner struct {
	// 实验对象ID，当test_type为CAMPAIGN时，传入广告组ID，当test_type为AD时，传入广告计划ID。
	ObjectId int64 `json:"object_id"`
}
