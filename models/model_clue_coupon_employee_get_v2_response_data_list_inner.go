/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponEmployeeGetV2ResponseDataListInner struct for ClueCouponEmployeeGetV2ResponseDataListInner
type ClueCouponEmployeeGetV2ResponseDataListInner struct {
	//
	StoreId *int64 `json:"store_id,omitempty"`
	//
	UserId   *int64                                   `json:"user_id,omitempty"`
	UserType *ClueCouponEmployeeGetV2DataListUserType `json:"user_type,omitempty"`
}
