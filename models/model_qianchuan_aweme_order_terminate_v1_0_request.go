/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderTerminateV10Request struct for QianchuanAwemeOrderTerminateV10Request
type QianchuanAwemeOrderTerminateV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	OrderIds []int64 `json:"order_ids"`
}
