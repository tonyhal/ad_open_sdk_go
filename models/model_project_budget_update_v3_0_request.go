/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectBudgetUpdateV30Request struct for ProjectBudgetUpdateV30Request
type ProjectBudgetUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Data []*ProjectBudgetUpdateV30RequestDataInner `json:"data"`
}
