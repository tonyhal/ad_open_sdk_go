/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataFeedDeliverySearch
type AdGetV2DataFeedDeliverySearch string

// List of ad_get_v2_data_feed_delivery_search
const (
	HAS_OPEN_AdGetV2DataFeedDeliverySearch AdGetV2DataFeedDeliverySearch = "HAS_OPEN"
	DISABLED_AdGetV2DataFeedDeliverySearch AdGetV2DataFeedDeliverySearch = "DISABLED"
)

// All allowed values of AdGetV2DataFeedDeliverySearch enum
var AllowedAdGetV2DataFeedDeliverySearchEnumValues = []AdGetV2DataFeedDeliverySearch{
	"HAS_OPEN",
	"DISABLED",
}

// NewAdGetV2DataFeedDeliverySearchFromValue returns a pointer to a valid AdGetV2DataFeedDeliverySearch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataFeedDeliverySearchFromValue(v string) (*AdGetV2DataFeedDeliverySearch, error) {
	ev := AdGetV2DataFeedDeliverySearch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataFeedDeliverySearch: valid values are %v", v, AllowedAdGetV2DataFeedDeliverySearchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataFeedDeliverySearch) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataFeedDeliverySearchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_feed_delivery_search value
func (v AdGetV2DataFeedDeliverySearch) Ptr() *AdGetV2DataFeedDeliverySearch {
	return &v
}
