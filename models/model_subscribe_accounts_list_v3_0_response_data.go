/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SubscribeAccountsListV30ResponseData
type SubscribeAccountsListV30ResponseData struct {
	//
	Advertisers []*SubscribeAccountsListV30ResponseDataAdvertisersInner `json:"advertisers,omitempty"`
	//
	Count *int64 `json:"count,omitempty"`
	//
	NextCursor *int64 `json:"next_cursor,omitempty"`
}
