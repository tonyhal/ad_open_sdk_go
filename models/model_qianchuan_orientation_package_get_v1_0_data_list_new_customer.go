/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanOrientationPackageGetV10DataListNewCustomer
type QianchuanOrientationPackageGetV10DataListNewCustomer string

// List of qianchuan_orientation_package_get_v1.0_data_list_new_customer
const (
	NONE_QianchuanOrientationPackageGetV10DataListNewCustomer   QianchuanOrientationPackageGetV10DataListNewCustomer = "NONE"
	NO_BUY_QianchuanOrientationPackageGetV10DataListNewCustomer QianchuanOrientationPackageGetV10DataListNewCustomer = "NO_BUY"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListNewCustomer enum
var AllowedQianchuanOrientationPackageGetV10DataListNewCustomerEnumValues = []QianchuanOrientationPackageGetV10DataListNewCustomer{
	"NONE",
	"NO_BUY",
}

// NewQianchuanOrientationPackageGetV10DataListNewCustomerFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListNewCustomer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListNewCustomerFromValue(v string) (*QianchuanOrientationPackageGetV10DataListNewCustomer, error) {
	ev := QianchuanOrientationPackageGetV10DataListNewCustomer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListNewCustomer: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListNewCustomerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListNewCustomer) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListNewCustomerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_new_customer value
func (v QianchuanOrientationPackageGetV10DataListNewCustomer) Ptr() *QianchuanOrientationPackageGetV10DataListNewCustomer {
	return &v
}
