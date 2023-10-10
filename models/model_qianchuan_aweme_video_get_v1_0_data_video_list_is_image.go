/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeVideoGetV10DataVideoListIsImage
type QianchuanAwemeVideoGetV10DataVideoListIsImage int64

// List of qianchuan_aweme_video_get_v1.0_data_video_list_is_image
const (
	Enum_0_QianchuanAwemeVideoGetV10DataVideoListIsImage QianchuanAwemeVideoGetV10DataVideoListIsImage = 0
	Enum_1_QianchuanAwemeVideoGetV10DataVideoListIsImage QianchuanAwemeVideoGetV10DataVideoListIsImage = 1
)

// All allowed values of QianchuanAwemeVideoGetV10DataVideoListIsImage enum
var AllowedQianchuanAwemeVideoGetV10DataVideoListIsImageEnumValues = []QianchuanAwemeVideoGetV10DataVideoListIsImage{
	0,
	1,
}

// NewQianchuanAwemeVideoGetV10DataVideoListIsImageFromValue returns a pointer to a valid QianchuanAwemeVideoGetV10DataVideoListIsImage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeVideoGetV10DataVideoListIsImageFromValue(v int64) (*QianchuanAwemeVideoGetV10DataVideoListIsImage, error) {
	ev := QianchuanAwemeVideoGetV10DataVideoListIsImage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeVideoGetV10DataVideoListIsImage: valid values are %v", v, AllowedQianchuanAwemeVideoGetV10DataVideoListIsImageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeVideoGetV10DataVideoListIsImage) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeVideoGetV10DataVideoListIsImageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_video_get_v1.0_data_video_list_is_image value
func (v QianchuanAwemeVideoGetV10DataVideoListIsImage) Ptr() *QianchuanAwemeVideoGetV10DataVideoListIsImage {
	return &v
}
