/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponEmployeeCreateV2ResponseDataSuccessListInner struct for ClueCouponEmployeeCreateV2ResponseDataSuccessListInner
type ClueCouponEmployeeCreateV2ResponseDataSuccessListInner struct {
	//
	StoreId *int64 `json:"store_id,omitempty"`
	//
	UserId   *int64                                             `json:"user_id,omitempty"`
	UserType *ClueCouponEmployeeCreateV2DataSuccessListUserType `json:"user_type,omitempty"`
}
