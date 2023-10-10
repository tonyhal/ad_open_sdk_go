/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserUpdateV2Request struct for AgentAdvertiserUpdateV2Request
type AgentAdvertiserUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Contacter *string `json:"contacter,omitempty"`
	//
	Email *string `json:"email,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Note *string `json:"note,omitempty"`
	//
	Phonenumber *string `json:"phonenumber,omitempty"`
}
