/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataAdDataSupplementsInner struct for CreativeDetailGetV30ResponseDataAdDataSupplementsInner
type CreativeDetailGetV30ResponseDataAdDataSupplementsInner struct {
	//
	Games          []*CreativeDetailGetV30ResponseDataAdDataSupplementsInnerGamesInner `json:"games,omitempty"`
	SupplementType *CreativeDetailGetV30DataAdDataSupplementsSupplementType            `json:"supplement_type,omitempty"`
}
