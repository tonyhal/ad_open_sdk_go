/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30ResponseDataErrorListInner struct for PromotionUpdateV30ResponseDataErrorListInner
type PromotionUpdateV30ResponseDataErrorListInner struct {
	//
	ErrorCode *int64 `json:"error_code,omitempty"`
	//
	ErrorMessage *string                                    `json:"error_message,omitempty"`
	ObjectType   *PromotionUpdateV30DataErrorListObjectType `json:"object_type,omitempty"`
}
