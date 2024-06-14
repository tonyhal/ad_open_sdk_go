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

// DpaProductUpdateV2ProductInfoStock
type DpaProductUpdateV2ProductInfoStock int64

// List of dpa_product_update_v2_product_info_stock
const (
	Enum_0_DpaProductUpdateV2ProductInfoStock DpaProductUpdateV2ProductInfoStock = 0
	Enum_1_DpaProductUpdateV2ProductInfoStock DpaProductUpdateV2ProductInfoStock = 1
)

// All allowed values of DpaProductUpdateV2ProductInfoStock enum
var AllowedDpaProductUpdateV2ProductInfoStockEnumValues = []DpaProductUpdateV2ProductInfoStock{
	0,
	1,
}

// NewDpaProductUpdateV2ProductInfoStockFromValue returns a pointer to a valid DpaProductUpdateV2ProductInfoStock
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaProductUpdateV2ProductInfoStockFromValue(v int64) (*DpaProductUpdateV2ProductInfoStock, error) {
	ev := DpaProductUpdateV2ProductInfoStock(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaProductUpdateV2ProductInfoStock: valid values are %v", v, AllowedDpaProductUpdateV2ProductInfoStockEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaProductUpdateV2ProductInfoStock) IsValid() bool {
	for _, existing := range AllowedDpaProductUpdateV2ProductInfoStockEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_product_update_v2_product_info_stock value
func (v DpaProductUpdateV2ProductInfoStock) Ptr() *DpaProductUpdateV2ProductInfoStock {
	return &v
}
