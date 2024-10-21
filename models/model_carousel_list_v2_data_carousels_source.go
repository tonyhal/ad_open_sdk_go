/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselListV2DataCarouselsSource
type CarouselListV2DataCarouselsSource string

// List of carousel_list_v2_data_carousels_source
const (
	ACCOUNT_PUSH_CarouselListV2DataCarouselsSource       CarouselListV2DataCarouselsSource = "ACCOUNT_PUSH"
	AD_SITE_CarouselListV2DataCarouselsSource            CarouselListV2DataCarouselsSource = "AD_SITE"
	AIC_CarouselListV2DataCarouselsSource                CarouselListV2DataCarouselsSource = "AIC"
	BP_CarouselListV2DataCarouselsSource                 CarouselListV2DataCarouselsSource = "BP"
	CEWEBRITY_CAROUSEL_CarouselListV2DataCarouselsSource CarouselListV2DataCarouselsSource = "CEWEBRITY_CAROUSEL"
	OPEN_API_CarouselListV2DataCarouselsSource           CarouselListV2DataCarouselsSource = "OPEN_API"
)

// Ptr returns reference to carousel_list_v2_data_carousels_source value
func (v CarouselListV2DataCarouselsSource) Ptr() *CarouselListV2DataCarouselsSource {
	return &v
}