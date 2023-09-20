/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType
type AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType string

// List of adlab_group_detail_v3.0_data_data_ad_info_delivery_range_union_video_type
const (
	ORIGIN_VIDEO_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType = "ORIGIN_VIDEO"
	REWARD_VIDEO_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType = "REWARD_VIDEO"
	SPLASH_VIDEO_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType = "SPLASH_VIDEO"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType enum
var AllowedAdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoTypeEnumValues = []AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType{
	"ORIGIN_VIDEO",
	"REWARD_VIDEO",
	"SPLASH_VIDEO",
}

// NewAdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoTypeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoTypeFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_delivery_range_union_video_type value
func (v AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType) Ptr() *AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType {
	return &v
}
