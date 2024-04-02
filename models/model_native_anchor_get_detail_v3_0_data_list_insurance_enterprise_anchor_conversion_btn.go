/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn
type NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn string

// List of native_anchor_get_detail_v3.0_data_list_insurance_enterprise_anchor_conversion_btn
const (
	BUY_NOW_NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn      NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn = "BUY_NOW"
	FREE_RECEIVE_NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn = "FREE_RECEIVE"
	IMPROVE_NOW_NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn  NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn = "IMPROVE_NOW"
	VIEW_DETIALS_NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn = "VIEW_DETIALS"
)

// All allowed values of NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn enum
var AllowedNativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtnEnumValues = []NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn{
	"BUY_NOW",
	"FREE_RECEIVE",
	"IMPROVE_NOW",
	"VIEW_DETIALS",
}

// NewNativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtnFromValue returns a pointer to a valid NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtnFromValue(v string) (*NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn, error) {
	ev := NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn: valid values are %v", v, AllowedNativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_detail_v3.0_data_list_insurance_enterprise_anchor_conversion_btn value
func (v NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn) Ptr() *NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn {
	return &v
}
