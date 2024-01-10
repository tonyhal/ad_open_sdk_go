/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderListV30ResponseDataOrderListInnerAdListInnerAudience
type DouplusOrderListV30ResponseDataOrderListInnerAdListInnerAudience struct {
	//
	Age []string `json:"age,omitempty"`
	//
	AuthorPkgs []string `json:"author_pkgs,omitempty"`
	//
	Business []string `json:"business,omitempty"`
	//
	City []string `json:"city,omitempty"`
	//
	DeliveryType *string                                                 `json:"delivery_type,omitempty"`
	District     *DouplusOrderListV30DataOrderListAdListAudienceDistrict `json:"district,omitempty"`
	//
	Gender *string `json:"gender,omitempty"`
	//
	InterestCategories []string `json:"interest_categories,omitempty"`
	//
	Platform []string `json:"platform,omitempty"`
	//
	Province []string `json:"province,omitempty"`
}
