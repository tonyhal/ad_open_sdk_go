/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2DataFormElementsElementType
type ClueCouponDetailV2DataFormElementsElementType string

// List of clue_coupon_detail_v2_data_form_elements_element_type
const (
	NUMBER_ClueCouponDetailV2DataFormElementsElementType     ClueCouponDetailV2DataFormElementsElementType = "NUMBER"
	EMAIL_ClueCouponDetailV2DataFormElementsElementType      ClueCouponDetailV2DataFormElementsElementType = "EMAIL"
	TEXTAREA_ClueCouponDetailV2DataFormElementsElementType   ClueCouponDetailV2DataFormElementsElementType = "TEXTAREA"
	RADIO_ClueCouponDetailV2DataFormElementsElementType      ClueCouponDetailV2DataFormElementsElementType = "RADIO"
	CITY_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "CITY"
	NAME_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "NAME"
	TELEPHONE_ClueCouponDetailV2DataFormElementsElementType  ClueCouponDetailV2DataFormElementsElementType = "TELEPHONE"
	CALCULATOR_ClueCouponDetailV2DataFormElementsElementType ClueCouponDetailV2DataFormElementsElementType = "CALCULATOR"
	SEX_ClueCouponDetailV2DataFormElementsElementType        ClueCouponDetailV2DataFormElementsElementType = "SEX"
	SELECT_ClueCouponDetailV2DataFormElementsElementType     ClueCouponDetailV2DataFormElementsElementType = "SELECT"
	TEXT_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "TEXT"
	DATE_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "DATE"
	CHECKBOX_ClueCouponDetailV2DataFormElementsElementType   ClueCouponDetailV2DataFormElementsElementType = "CHECKBOX"
)

// Ptr returns reference to clue_coupon_detail_v2_data_form_elements_element_type value
func (v ClueCouponDetailV2DataFormElementsElementType) Ptr() *ClueCouponDetailV2DataFormElementsElementType {
	return &v
}
