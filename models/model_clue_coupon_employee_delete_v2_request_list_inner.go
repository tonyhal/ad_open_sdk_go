/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponEmployeeDeleteV2RequestListInner struct for ClueCouponEmployeeDeleteV2RequestListInner
type ClueCouponEmployeeDeleteV2RequestListInner struct {
	//
	StoreId *int64 `json:"store_id,omitempty"`
	//
	UserId   *int64                                  `json:"user_id,omitempty"`
	UserType *ClueCouponEmployeeDeleteV2ListUserType `json:"user_type,omitempty"`
}