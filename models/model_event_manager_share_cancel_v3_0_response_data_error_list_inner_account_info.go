/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerShareCancelV30ResponseDataErrorListInnerAccountInfo
type EventManagerShareCancelV30ResponseDataErrorListInnerAccountInfo struct {
	//
	AccountId   *int64                                                         `json:"account_id,omitempty"`
	AccountType *EventManagerShareCancelV30DataErrorListAccountInfoAccountType `json:"account_type,omitempty"`
}
