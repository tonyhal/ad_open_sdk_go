/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeCreateV2RequestCreativeImageMaterialsInner struct for CreativeProceduralCreativeCreateV2RequestCreativeImageMaterialsInner
type CreativeProceduralCreativeCreateV2RequestCreativeImageMaterialsInner struct {
	//
	ImageInfo []*CreativeProceduralCreativeCreateV2RequestCreativeImageMaterialsInnerImageInfoInner `json:"image_info,omitempty"`
	ImageMode *CreativeProceduralCreativeCreateV2CreativeImageMaterialsImageMode                    `json:"image_mode,omitempty"`
	//
	MaterialId    *int64                                                                             `json:"material_id,omitempty"`
	TemplateImage *CreativeProceduralCreativeCreateV2RequestCreativeImageMaterialsInnerTemplateImage `json:"template_image,omitempty"`
}
