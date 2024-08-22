/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeUpdateV2RequestCreative
type CreativeProceduralCreativeUpdateV2RequestCreative struct {
	//
	AbstractMaterials []*CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInner `json:"abstract_materials,omitempty"`
	//
	ComponentMaterials []*CreativeProceduralCreativeUpdateV2RequestCreativeComponentMaterialsInner `json:"component_materials,omitempty"`
	DecorationMaterial *CreativeProceduralCreativeUpdateV2RequestCreativeDecorationMaterial        `json:"decoration_material,omitempty"`
	//
	ImageMaterials   []*CreativeProceduralCreativeUpdateV2RequestCreativeImageMaterialsInner `json:"image_materials,omitempty"`
	SubTitleMaterial *CreativeProceduralCreativeUpdateV2RequestCreativeSubTitleMaterial      `json:"sub_title_material,omitempty"`
	//
	TitleMaterials []*CreativeProceduralCreativeUpdateV2RequestCreativeTitleMaterialsInner `json:"title_materials,omitempty"`
	//
	VideoMaterials []*CreativeProceduralCreativeUpdateV2RequestCreativeVideoMaterialsInner `json:"video_materials,omitempty"`
}
