/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponEmployeeCreateV2RequestListInner struct for ClueCouponEmployeeCreateV2RequestListInner
type ClueCouponEmployeeCreateV2RequestListInner struct {
	//
	StoreId *int64 `json:"store_id,omitempty"`
	//
	UserId   *int64                                  `json:"user_id,omitempty"`
	UserType *ClueCouponEmployeeCreateV2ListUserType `json:"user_type,omitempty"`
}
