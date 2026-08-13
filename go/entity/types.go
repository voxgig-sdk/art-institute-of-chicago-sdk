// Typed models for the ArtInstituteOfChicago SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/art-institute-of-chicago-sdk/go/core"
)

// Agent is the typed data model for the agent entity.
type Agent struct {
	AltTitles *any `json:"alt_titles,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	BirthDate *any `json:"birth_date,omitempty"`
	DeathDate *any `json:"death_date,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	IsArtist *bool `json:"is_artist,omitempty"`
	SortTitle *any `json:"sort_title,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UlanId *string `json:"ulan_id,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// AgentLoadMatch is the typed request payload for Agent.LoadTyped.
type AgentLoadMatch struct {
	Id string `json:"id"`
}

// AgentListMatch is the typed request payload for Agent.ListTyped.
type AgentListMatch struct {
	AltTitles *any `json:"alt_titles,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	BirthDate *any `json:"birth_date,omitempty"`
	DeathDate *any `json:"death_date,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	IsArtist *bool `json:"is_artist,omitempty"`
	SortTitle *any `json:"sort_title,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UlanId *string `json:"ulan_id,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// AgentRole is the typed data model for the agent_role entity.
type AgentRole struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// AgentRoleLoadMatch is the typed request payload for AgentRole.LoadTyped.
type AgentRoleLoadMatch struct {
	Id string `json:"id"`
}

// AgentRoleListMatch is the typed request payload for AgentRole.ListTyped.
type AgentRoleListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// AgentType is the typed data model for the agent_type entity.
type AgentType struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// AgentTypeLoadMatch is the typed request payload for AgentType.LoadTyped.
type AgentTypeLoadMatch struct {
	Id string `json:"id"`
}

// AgentTypeListMatch is the typed request payload for AgentType.ListTyped.
type AgentTypeListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// Article is the typed data model for the article entity.
type Article struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// ArticleLoadMatch is the typed request payload for Article.LoadTyped.
type ArticleLoadMatch struct {
	Id string `json:"id"`
}

// ArticleListMatch is the typed request payload for Article.ListTyped.
type ArticleListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// Artwork is the typed data model for the artwork entity.
type Artwork struct {
	AltArtistIds *any `json:"alt_artist_ids,omitempty"`
	AltClassificationIds *any `json:"alt_classification_ids,omitempty"`
	AltImageIds *any `json:"alt_image_ids,omitempty"`
	AltMaterialIds *any `json:"alt_material_ids,omitempty"`
	AltStyleIds *any `json:"alt_style_ids,omitempty"`
	AltSubjectIds *any `json:"alt_subject_ids,omitempty"`
	AltTechniqueIds *any `json:"alt_technique_ids,omitempty"`
	AltTitles *any `json:"alt_titles,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistDisplay *any `json:"artist_display,omitempty"`
	ArtistId *string `json:"artist_id,omitempty"`
	ArtistIds *any `json:"artist_ids,omitempty"`
	ArtistTitle *any `json:"artist_title,omitempty"`
	ArtistTitles *any `json:"artist_titles,omitempty"`
	ArtworkTypeId *string `json:"artwork_type_id,omitempty"`
	ArtworkTypeTitle *any `json:"artwork_type_title,omitempty"`
	BoostRank *any `json:"boost_rank,omitempty"`
	CatalogBasedSearchKeywordTitles *any `json:"catalog_based_search_keyword_titles,omitempty"`
	CatalogueDisplay *any `json:"catalogue_display,omitempty"`
	CategoryIds *any `json:"category_ids,omitempty"`
	CategoryTitles *any `json:"category_titles,omitempty"`
	ClassificationId *string `json:"classification_id,omitempty"`
	ClassificationIds *any `json:"classification_ids,omitempty"`
	ClassificationTitle *any `json:"classification_title,omitempty"`
	ClassificationTitles *any `json:"classification_titles,omitempty"`
	Color *any `json:"color,omitempty"`
	Colorfulness *any `json:"colorfulness,omitempty"`
	CopyrightNotice *any `json:"copyright_notice,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	DateDisplay *any `json:"date_display,omitempty"`
	DateEnd *any `json:"date_end,omitempty"`
	DateQualifierId *string `json:"date_qualifier_id,omitempty"`
	DateQualifierTitle *any `json:"date_qualifier_title,omitempty"`
	DateStart *any `json:"date_start,omitempty"`
	DepartmentId *string `json:"department_id,omitempty"`
	DepartmentTitle *any `json:"department_title,omitempty"`
	Description *string `json:"description,omitempty"`
	Dimensions *any `json:"dimensions,omitempty"`
	DimensionsDetail *any `json:"dimensions_detail,omitempty"`
	DocumentIds *any `json:"document_ids,omitempty"`
	Edition *any `json:"edition,omitempty"`
	ExhibitionHistory *any `json:"exhibition_history,omitempty"`
	FiscalYear *any `json:"fiscal_year,omitempty"`
	FiscalYearDeaccession *any `json:"fiscal_year_deaccession,omitempty"`
	GalleryId *string `json:"gallery_id,omitempty"`
	GalleryTitle *any `json:"gallery_title,omitempty"`
	HasAdvancedImaging *bool `json:"has_advanced_imaging,omitempty"`
	HasEducationalResources *bool `json:"has_educational_resources,omitempty"`
	HasMultimediaResources *bool `json:"has_multimedia_resources,omitempty"`
	HasNotBeenViewedMuch *bool `json:"has_not_been_viewed_much,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageEmbedding *any `json:"image_embedding,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	Inscriptions *any `json:"inscriptions,omitempty"`
	InternalDepartmentId *string `json:"internal_department_id,omitempty"`
	IsBoosted *bool `json:"is_boosted,omitempty"`
	IsOnView *bool `json:"is_on_view,omitempty"`
	IsPublicDomain *bool `json:"is_public_domain,omitempty"`
	IsZoomable *bool `json:"is_zoomable,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Latlon *any `json:"latlon,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	MainReferenceNumber *int `json:"main_reference_number,omitempty"`
	MaterialId *string `json:"material_id,omitempty"`
	MaterialIds *any `json:"material_ids,omitempty"`
	MaterialTitles *any `json:"material_titles,omitempty"`
	MaxZoomWindowSize *any `json:"max_zoom_window_size,omitempty"`
	MediumDisplay *any `json:"medium_display,omitempty"`
	NomismaId *string `json:"nomisma_id,omitempty"`
	OnLoanDisplay *any `json:"on_loan_display,omitempty"`
	Pageviews *any `json:"pageviews,omitempty"`
	PageviewsRecent *any `json:"pageviews_recent,omitempty"`
	PlaceOfOrigin *any `json:"place_of_origin,omitempty"`
	ProvenanceText *any `json:"provenance_text,omitempty"`
	PublicationHistory *any `json:"publication_history,omitempty"`
	PublishingVerificationLevel *any `json:"publishing_verification_level,omitempty"`
	SectionIds *any `json:"section_ids,omitempty"`
	SectionTitles *any `json:"section_titles,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SiteIds *any `json:"site_ids,omitempty"`
	SoundIds *any `json:"sound_ids,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	StyleId *string `json:"style_id,omitempty"`
	StyleIds *any `json:"style_ids,omitempty"`
	StyleTitle *any `json:"style_title,omitempty"`
	StyleTitles *any `json:"style_titles,omitempty"`
	SubjectId *string `json:"subject_id,omitempty"`
	SubjectIds *any `json:"subject_ids,omitempty"`
	SubjectTitles *any `json:"subject_titles,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	TechniqueId *string `json:"technique_id,omitempty"`
	TechniqueIds *any `json:"technique_ids,omitempty"`
	TechniqueTitles *any `json:"technique_titles,omitempty"`
	TermTitles *any `json:"term_titles,omitempty"`
	TextEmbedding *any `json:"text_embedding,omitempty"`
	TextIds *any `json:"text_ids,omitempty"`
	ThemeTitles *any `json:"theme_titles,omitempty"`
	Thumbnail *any `json:"thumbnail,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	VideoIds *any `json:"video_ids,omitempty"`
}

// ArtworkLoadMatch is the typed request payload for Artwork.LoadTyped.
type ArtworkLoadMatch struct {
	Id string `json:"id"`
}

// ArtworkListMatch is the typed request payload for Artwork.ListTyped.
type ArtworkListMatch struct {
	AltArtistIds *any `json:"alt_artist_ids,omitempty"`
	AltClassificationIds *any `json:"alt_classification_ids,omitempty"`
	AltImageIds *any `json:"alt_image_ids,omitempty"`
	AltMaterialIds *any `json:"alt_material_ids,omitempty"`
	AltStyleIds *any `json:"alt_style_ids,omitempty"`
	AltSubjectIds *any `json:"alt_subject_ids,omitempty"`
	AltTechniqueIds *any `json:"alt_technique_ids,omitempty"`
	AltTitles *any `json:"alt_titles,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistDisplay *any `json:"artist_display,omitempty"`
	ArtistId *string `json:"artist_id,omitempty"`
	ArtistIds *any `json:"artist_ids,omitempty"`
	ArtistTitle *any `json:"artist_title,omitempty"`
	ArtistTitles *any `json:"artist_titles,omitempty"`
	ArtworkTypeId *string `json:"artwork_type_id,omitempty"`
	ArtworkTypeTitle *any `json:"artwork_type_title,omitempty"`
	BoostRank *any `json:"boost_rank,omitempty"`
	CatalogBasedSearchKeywordTitles *any `json:"catalog_based_search_keyword_titles,omitempty"`
	CatalogueDisplay *any `json:"catalogue_display,omitempty"`
	CategoryIds *any `json:"category_ids,omitempty"`
	CategoryTitles *any `json:"category_titles,omitempty"`
	ClassificationId *string `json:"classification_id,omitempty"`
	ClassificationIds *any `json:"classification_ids,omitempty"`
	ClassificationTitle *any `json:"classification_title,omitempty"`
	ClassificationTitles *any `json:"classification_titles,omitempty"`
	Color *any `json:"color,omitempty"`
	Colorfulness *any `json:"colorfulness,omitempty"`
	CopyrightNotice *any `json:"copyright_notice,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	DateDisplay *any `json:"date_display,omitempty"`
	DateEnd *any `json:"date_end,omitempty"`
	DateQualifierId *string `json:"date_qualifier_id,omitempty"`
	DateQualifierTitle *any `json:"date_qualifier_title,omitempty"`
	DateStart *any `json:"date_start,omitempty"`
	DepartmentId *string `json:"department_id,omitempty"`
	DepartmentTitle *any `json:"department_title,omitempty"`
	Description *string `json:"description,omitempty"`
	Dimensions *any `json:"dimensions,omitempty"`
	DimensionsDetail *any `json:"dimensions_detail,omitempty"`
	DocumentIds *any `json:"document_ids,omitempty"`
	Edition *any `json:"edition,omitempty"`
	ExhibitionHistory *any `json:"exhibition_history,omitempty"`
	FiscalYear *any `json:"fiscal_year,omitempty"`
	FiscalYearDeaccession *any `json:"fiscal_year_deaccession,omitempty"`
	GalleryId *string `json:"gallery_id,omitempty"`
	GalleryTitle *any `json:"gallery_title,omitempty"`
	HasAdvancedImaging *bool `json:"has_advanced_imaging,omitempty"`
	HasEducationalResources *bool `json:"has_educational_resources,omitempty"`
	HasMultimediaResources *bool `json:"has_multimedia_resources,omitempty"`
	HasNotBeenViewedMuch *bool `json:"has_not_been_viewed_much,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageEmbedding *any `json:"image_embedding,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	Inscriptions *any `json:"inscriptions,omitempty"`
	InternalDepartmentId *string `json:"internal_department_id,omitempty"`
	IsBoosted *bool `json:"is_boosted,omitempty"`
	IsOnView *bool `json:"is_on_view,omitempty"`
	IsPublicDomain *bool `json:"is_public_domain,omitempty"`
	IsZoomable *bool `json:"is_zoomable,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Latlon *any `json:"latlon,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	MainReferenceNumber *int `json:"main_reference_number,omitempty"`
	MaterialId *string `json:"material_id,omitempty"`
	MaterialIds *any `json:"material_ids,omitempty"`
	MaterialTitles *any `json:"material_titles,omitempty"`
	MaxZoomWindowSize *any `json:"max_zoom_window_size,omitempty"`
	MediumDisplay *any `json:"medium_display,omitempty"`
	NomismaId *string `json:"nomisma_id,omitempty"`
	OnLoanDisplay *any `json:"on_loan_display,omitempty"`
	Pageviews *any `json:"pageviews,omitempty"`
	PageviewsRecent *any `json:"pageviews_recent,omitempty"`
	PlaceOfOrigin *any `json:"place_of_origin,omitempty"`
	ProvenanceText *any `json:"provenance_text,omitempty"`
	PublicationHistory *any `json:"publication_history,omitempty"`
	PublishingVerificationLevel *any `json:"publishing_verification_level,omitempty"`
	SectionIds *any `json:"section_ids,omitempty"`
	SectionTitles *any `json:"section_titles,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SiteIds *any `json:"site_ids,omitempty"`
	SoundIds *any `json:"sound_ids,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	StyleId *string `json:"style_id,omitempty"`
	StyleIds *any `json:"style_ids,omitempty"`
	StyleTitle *any `json:"style_title,omitempty"`
	StyleTitles *any `json:"style_titles,omitempty"`
	SubjectId *string `json:"subject_id,omitempty"`
	SubjectIds *any `json:"subject_ids,omitempty"`
	SubjectTitles *any `json:"subject_titles,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	TechniqueId *string `json:"technique_id,omitempty"`
	TechniqueIds *any `json:"technique_ids,omitempty"`
	TechniqueTitles *any `json:"technique_titles,omitempty"`
	TermTitles *any `json:"term_titles,omitempty"`
	TextEmbedding *any `json:"text_embedding,omitempty"`
	TextIds *any `json:"text_ids,omitempty"`
	ThemeTitles *any `json:"theme_titles,omitempty"`
	Thumbnail *any `json:"thumbnail,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	VideoIds *any `json:"video_ids,omitempty"`
}

// ArtworkDateQualifier is the typed data model for the artwork_date_qualifier entity.
type ArtworkDateQualifier struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// ArtworkDateQualifierLoadMatch is the typed request payload for ArtworkDateQualifier.LoadTyped.
type ArtworkDateQualifierLoadMatch struct {
	Id string `json:"id"`
}

// ArtworkDateQualifierListMatch is the typed request payload for ArtworkDateQualifier.ListTyped.
type ArtworkDateQualifierListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// ArtworkPlaceQualifier is the typed data model for the artwork_place_qualifier entity.
type ArtworkPlaceQualifier struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// ArtworkPlaceQualifierLoadMatch is the typed request payload for ArtworkPlaceQualifier.LoadTyped.
type ArtworkPlaceQualifierLoadMatch struct {
	Id string `json:"id"`
}

// ArtworkPlaceQualifierListMatch is the typed request payload for ArtworkPlaceQualifier.ListTyped.
type ArtworkPlaceQualifierListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// ArtworkType is the typed data model for the artwork_type entity.
type ArtworkType struct {
	AatId *string `json:"aat_id,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// ArtworkTypeLoadMatch is the typed request payload for ArtworkType.LoadTyped.
type ArtworkTypeLoadMatch struct {
	Id string `json:"id"`
}

// ArtworkTypeListMatch is the typed request payload for ArtworkType.ListTyped.
type ArtworkTypeListMatch struct {
	AatId *string `json:"aat_id,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// CategoryTerm is the typed data model for the category_term entity.
type CategoryTerm struct {
	AatId *string `json:"aat_id,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	ParentId *string `json:"parent_id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	Subtype *any `json:"subtype,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// CategoryTermLoadMatch is the typed request payload for CategoryTerm.LoadTyped.
type CategoryTermLoadMatch struct {
	Id string `json:"id"`
}

// CategoryTermListMatch is the typed request payload for CategoryTerm.ListTyped.
type CategoryTermListMatch struct {
	AatId *string `json:"aat_id,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	ParentId *string `json:"parent_id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	Subtype *any `json:"subtype,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// DigitalPublication is the typed data model for the digital_publication entity.
type DigitalPublication struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// DigitalPublicationLoadMatch is the typed request payload for DigitalPublication.LoadTyped.
type DigitalPublicationLoadMatch struct {
	Id string `json:"id"`
}

// DigitalPublicationListMatch is the typed request payload for DigitalPublication.ListTyped.
type DigitalPublicationListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// DigitalPublicationArticle is the typed data model for the digital_publication_article entity.
type DigitalPublicationArticle struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	AuthorDisplay *any `json:"author_display,omitempty"`
	Copy *any `json:"copy,omitempty"`
	DigitalPublicationId *string `json:"digital_publication_id,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// DigitalPublicationArticleLoadMatch is the typed request payload for DigitalPublicationArticle.LoadTyped.
type DigitalPublicationArticleLoadMatch struct {
	Id string `json:"id"`
}

// DigitalPublicationArticleListMatch is the typed request payload for DigitalPublicationArticle.ListTyped.
type DigitalPublicationArticleListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	AuthorDisplay *any `json:"author_display,omitempty"`
	Copy *any `json:"copy,omitempty"`
	DigitalPublicationId *string `json:"digital_publication_id,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// EducatorResource is the typed data model for the educator_resource entity.
type EducatorResource struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// EducatorResourceLoadMatch is the typed request payload for EducatorResource.LoadTyped.
type EducatorResourceLoadMatch struct {
	Id string `json:"id"`
}

// EducatorResourceListMatch is the typed request payload for EducatorResource.ListTyped.
type EducatorResourceListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Event is the typed data model for the event entity.
type Event struct {
	AltAudienceIds *any `json:"alt_audience_ids,omitempty"`
	AltEventTypeIds *any `json:"alt_event_type_ids,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	AudienceId *string `json:"audience_id,omitempty"`
	BuyButtonCaption *any `json:"buy_button_caption,omitempty"`
	BuyButtonText *any `json:"buy_button_text,omitempty"`
	DateDisplay *any `json:"date_display,omitempty"`
	Description *string `json:"description,omitempty"`
	DoorTime *any `json:"door_time,omitempty"`
	EndDate *any `json:"end_date,omitempty"`
	EndTime *any `json:"end_time,omitempty"`
	Entrance *any `json:"entrance,omitempty"`
	EventHostId *string `json:"event_host_id,omitempty"`
	EventHostTitle *any `json:"event_host_title,omitempty"`
	EventTypeId *string `json:"event_type_id,omitempty"`
	HeaderDescription *any `json:"header_description,omitempty"`
	HeroCaption *any `json:"hero_caption,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	IsAdmissionRequired *bool `json:"is_admission_required,omitempty"`
	IsAfterHours *bool `json:"is_after_hours,omitempty"`
	IsFree *bool `json:"is_free,omitempty"`
	IsMemberExclusive *bool `json:"is_member_exclusive,omitempty"`
	IsPrivate *bool `json:"is_private,omitempty"`
	IsRegistrationRequired *bool `json:"is_registration_required,omitempty"`
	IsSalesButtonHidden *bool `json:"is_sales_button_hidden,omitempty"`
	IsSoldOut *bool `json:"is_sold_out,omitempty"`
	IsTicketed *bool `json:"is_ticketed,omitempty"`
	IsVirtualEvent *bool `json:"is_virtual_event,omitempty"`
	JoinUrl *any `json:"join_url,omitempty"`
	LayoutType *any `json:"layout_type,omitempty"`
	ListDescription *any `json:"list_description,omitempty"`
	Location *any `json:"location,omitempty"`
	ProgramIds *any `json:"program_ids,omitempty"`
	ProgramTitles *any `json:"program_titles,omitempty"`
	RsvpLink *any `json:"rsvp_link,omitempty"`
	SearchTags *any `json:"search_tags,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	StartDate *any `json:"start_date,omitempty"`
	StartTime *any `json:"start_time,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	SurveyUrl *any `json:"survey_url,omitempty"`
	TicketedEventId *string `json:"ticketed_event_id,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	TitleDisplay *any `json:"title_display,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	VirtualEventPasscode *any `json:"virtual_event_passcode,omitempty"`
	VirtualEventUrl *any `json:"virtual_event_url,omitempty"`
}

// EventLoadMatch is the typed request payload for Event.LoadTyped.
type EventLoadMatch struct {
	Id string `json:"id"`
}

// EventListMatch is the typed request payload for Event.ListTyped.
type EventListMatch struct {
	AltAudienceIds *any `json:"alt_audience_ids,omitempty"`
	AltEventTypeIds *any `json:"alt_event_type_ids,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	AudienceId *string `json:"audience_id,omitempty"`
	BuyButtonCaption *any `json:"buy_button_caption,omitempty"`
	BuyButtonText *any `json:"buy_button_text,omitempty"`
	DateDisplay *any `json:"date_display,omitempty"`
	Description *string `json:"description,omitempty"`
	DoorTime *any `json:"door_time,omitempty"`
	EndDate *any `json:"end_date,omitempty"`
	EndTime *any `json:"end_time,omitempty"`
	Entrance *any `json:"entrance,omitempty"`
	EventHostId *string `json:"event_host_id,omitempty"`
	EventHostTitle *any `json:"event_host_title,omitempty"`
	EventTypeId *string `json:"event_type_id,omitempty"`
	HeaderDescription *any `json:"header_description,omitempty"`
	HeroCaption *any `json:"hero_caption,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	IsAdmissionRequired *bool `json:"is_admission_required,omitempty"`
	IsAfterHours *bool `json:"is_after_hours,omitempty"`
	IsFree *bool `json:"is_free,omitempty"`
	IsMemberExclusive *bool `json:"is_member_exclusive,omitempty"`
	IsPrivate *bool `json:"is_private,omitempty"`
	IsRegistrationRequired *bool `json:"is_registration_required,omitempty"`
	IsSalesButtonHidden *bool `json:"is_sales_button_hidden,omitempty"`
	IsSoldOut *bool `json:"is_sold_out,omitempty"`
	IsTicketed *bool `json:"is_ticketed,omitempty"`
	IsVirtualEvent *bool `json:"is_virtual_event,omitempty"`
	JoinUrl *any `json:"join_url,omitempty"`
	LayoutType *any `json:"layout_type,omitempty"`
	ListDescription *any `json:"list_description,omitempty"`
	Location *any `json:"location,omitempty"`
	ProgramIds *any `json:"program_ids,omitempty"`
	ProgramTitles *any `json:"program_titles,omitempty"`
	RsvpLink *any `json:"rsvp_link,omitempty"`
	SearchTags *any `json:"search_tags,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	StartDate *any `json:"start_date,omitempty"`
	StartTime *any `json:"start_time,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	SurveyUrl *any `json:"survey_url,omitempty"`
	TicketedEventId *string `json:"ticketed_event_id,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	TitleDisplay *any `json:"title_display,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	VirtualEventPasscode *any `json:"virtual_event_passcode,omitempty"`
	VirtualEventUrl *any `json:"virtual_event_url,omitempty"`
}

// EventOccurrence is the typed data model for the event_occurrence entity.
type EventOccurrence struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ButtonCaption *any `json:"button_caption,omitempty"`
	ButtonText *any `json:"button_text,omitempty"`
	ButtonUrl *any `json:"button_url,omitempty"`
	Description *string `json:"description,omitempty"`
	EndAt *any `json:"end_at,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	IsPrivate *bool `json:"is_private,omitempty"`
	IsSalesButtonHidden *bool `json:"is_sales_button_hidden,omitempty"`
	IsTicketed *bool `json:"is_ticketed,omitempty"`
	Location *any `json:"location,omitempty"`
	OffSaleAt *any `json:"off_sale_at,omitempty"`
	OnSaleAt *any `json:"on_sale_at,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	StartAt *any `json:"start_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	TitleDisplay *any `json:"title_display,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// EventOccurrenceLoadMatch is the typed request payload for EventOccurrence.LoadTyped.
type EventOccurrenceLoadMatch struct {
	Id string `json:"id"`
}

// EventOccurrenceListMatch is the typed request payload for EventOccurrence.ListTyped.
type EventOccurrenceListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ButtonCaption *any `json:"button_caption,omitempty"`
	ButtonText *any `json:"button_text,omitempty"`
	ButtonUrl *any `json:"button_url,omitempty"`
	Description *string `json:"description,omitempty"`
	EndAt *any `json:"end_at,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	IsPrivate *bool `json:"is_private,omitempty"`
	IsSalesButtonHidden *bool `json:"is_sales_button_hidden,omitempty"`
	IsTicketed *bool `json:"is_ticketed,omitempty"`
	Location *any `json:"location,omitempty"`
	OffSaleAt *any `json:"off_sale_at,omitempty"`
	OnSaleAt *any `json:"on_sale_at,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	StartAt *any `json:"start_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	TitleDisplay *any `json:"title_display,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// EventProgram is the typed data model for the event_program entity.
type EventProgram struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	IsAffiliateGroup *bool `json:"is_affiliate_group,omitempty"`
	IsEventHost *bool `json:"is_event_host,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// EventProgramLoadMatch is the typed request payload for EventProgram.LoadTyped.
type EventProgramLoadMatch struct {
	Id string `json:"id"`
}

// EventProgramListMatch is the typed request payload for EventProgram.ListTyped.
type EventProgramListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	IsAffiliateGroup *bool `json:"is_affiliate_group,omitempty"`
	IsEventHost *bool `json:"is_event_host,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// Exhibition is the typed data model for the exhibition entity.
type Exhibition struct {
	AicEndAt *any `json:"aic_end_at,omitempty"`
	AicStartAt *any `json:"aic_start_at,omitempty"`
	AltImageIds *any `json:"alt_image_ids,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistIds *any `json:"artist_ids,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	DocumentIds *any `json:"document_ids,omitempty"`
	GalleryId *string `json:"gallery_id,omitempty"`
	GalleryTitle *any `json:"gallery_title,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	IsFeatured *bool `json:"is_featured,omitempty"`
	IsPublished *bool `json:"is_published,omitempty"`
	Position *any `json:"position,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SiteIds *any `json:"site_ids,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	Status *any `json:"status,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// ExhibitionLoadMatch is the typed request payload for Exhibition.LoadTyped.
type ExhibitionLoadMatch struct {
	Id string `json:"id"`
}

// ExhibitionListMatch is the typed request payload for Exhibition.ListTyped.
type ExhibitionListMatch struct {
	AicEndAt *any `json:"aic_end_at,omitempty"`
	AicStartAt *any `json:"aic_start_at,omitempty"`
	AltImageIds *any `json:"alt_image_ids,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistIds *any `json:"artist_ids,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	DocumentIds *any `json:"document_ids,omitempty"`
	GalleryId *string `json:"gallery_id,omitempty"`
	GalleryTitle *any `json:"gallery_title,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	IsFeatured *bool `json:"is_featured,omitempty"`
	IsPublished *bool `json:"is_published,omitempty"`
	Position *any `json:"position,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SiteIds *any `json:"site_ids,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	Status *any `json:"status,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Gallery is the typed data model for the gallery entity.
type Gallery struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Floor *any `json:"floor,omitempty"`
	Id *string `json:"id,omitempty"`
	IsClosed *bool `json:"is_closed,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Latlon *any `json:"latlon,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Number *any `json:"number,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	TgnId *string `json:"tgn_id,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// GalleryLoadMatch is the typed request payload for Gallery.LoadTyped.
type GalleryLoadMatch struct {
	Id string `json:"id"`
}

// GalleryListMatch is the typed request payload for Gallery.ListTyped.
type GalleryListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Floor *any `json:"floor,omitempty"`
	Id *string `json:"id,omitempty"`
	IsClosed *bool `json:"is_closed,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Latlon *any `json:"latlon,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Number *any `json:"number,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	TgnId *string `json:"tgn_id,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// GenericPage is the typed data model for the generic_page entity.
type GenericPage struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SearchTags *any `json:"search_tags,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// GenericPageLoadMatch is the typed request payload for GenericPage.LoadTyped.
type GenericPageLoadMatch struct {
	Id string `json:"id"`
}

// GenericPageListMatch is the typed request payload for GenericPage.ListTyped.
type GenericPageListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SearchTags *any `json:"search_tags,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Highlight is the typed data model for the highlight entity.
type Highlight struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// HighlightLoadMatch is the typed request payload for Highlight.LoadTyped.
type HighlightLoadMatch struct {
	Id string `json:"id"`
}

// HighlightListMatch is the typed request payload for Highlight.ListTyped.
type HighlightListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// Hour is the typed data model for the hour entity.
type Hour struct {
	AdditionalText *any `json:"additional_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	FridayIsClosed *any `json:"friday_is_closed,omitempty"`
	FridayMemberClose *any `json:"friday_member_close,omitempty"`
	FridayMemberOpen *any `json:"friday_member_open,omitempty"`
	FridayPublicClose *any `json:"friday_public_close,omitempty"`
	FridayPublicOpen *any `json:"friday_public_open,omitempty"`
	Id *string `json:"id,omitempty"`
	MondayIsClosed *any `json:"monday_is_closed,omitempty"`
	MondayMemberClose *any `json:"monday_member_close,omitempty"`
	MondayMemberOpen *any `json:"monday_member_open,omitempty"`
	MondayPublicClose *any `json:"monday_public_close,omitempty"`
	MondayPublicOpen *any `json:"monday_public_open,omitempty"`
	SaturdayIsClosed *any `json:"saturday_is_closed,omitempty"`
	SaturdayMemberClose *any `json:"saturday_member_close,omitempty"`
	SaturdayMemberOpen *any `json:"saturday_member_open,omitempty"`
	SaturdayPublicClose *any `json:"saturday_public_close,omitempty"`
	SaturdayPublicOpen *any `json:"saturday_public_open,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Summary *any `json:"summary,omitempty"`
	SundayIsClosed *any `json:"sunday_is_closed,omitempty"`
	SundayMemberClose *any `json:"sunday_member_close,omitempty"`
	SundayMemberOpen *any `json:"sunday_member_open,omitempty"`
	SundayPublicClose *any `json:"sunday_public_close,omitempty"`
	SundayPublicOpen *any `json:"sunday_public_open,omitempty"`
	ThursdayIsClosed *any `json:"thursday_is_closed,omitempty"`
	ThursdayMemberClose *any `json:"thursday_member_close,omitempty"`
	ThursdayMemberOpen *any `json:"thursday_member_open,omitempty"`
	ThursdayPublicClose *any `json:"thursday_public_close,omitempty"`
	ThursdayPublicOpen *any `json:"thursday_public_open,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	TuesdayIsClosed *any `json:"tuesday_is_closed,omitempty"`
	TuesdayMemberClose *any `json:"tuesday_member_close,omitempty"`
	TuesdayMemberOpen *any `json:"tuesday_member_open,omitempty"`
	TuesdayPublicClose *any `json:"tuesday_public_close,omitempty"`
	TuesdayPublicOpen *any `json:"tuesday_public_open,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WednesdayIsClosed *any `json:"wednesday_is_closed,omitempty"`
	WednesdayMemberClose *any `json:"wednesday_member_close,omitempty"`
	WednesdayMemberOpen *any `json:"wednesday_member_open,omitempty"`
	WednesdayPublicClose *any `json:"wednesday_public_close,omitempty"`
	WednesdayPublicOpen *any `json:"wednesday_public_open,omitempty"`
}

// HourLoadMatch is the typed request payload for Hour.LoadTyped.
type HourLoadMatch struct {
	Id string `json:"id"`
}

// HourListMatch is the typed request payload for Hour.ListTyped.
type HourListMatch struct {
	AdditionalText *any `json:"additional_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	FridayIsClosed *any `json:"friday_is_closed,omitempty"`
	FridayMemberClose *any `json:"friday_member_close,omitempty"`
	FridayMemberOpen *any `json:"friday_member_open,omitempty"`
	FridayPublicClose *any `json:"friday_public_close,omitempty"`
	FridayPublicOpen *any `json:"friday_public_open,omitempty"`
	Id *string `json:"id,omitempty"`
	MondayIsClosed *any `json:"monday_is_closed,omitempty"`
	MondayMemberClose *any `json:"monday_member_close,omitempty"`
	MondayMemberOpen *any `json:"monday_member_open,omitempty"`
	MondayPublicClose *any `json:"monday_public_close,omitempty"`
	MondayPublicOpen *any `json:"monday_public_open,omitempty"`
	SaturdayIsClosed *any `json:"saturday_is_closed,omitempty"`
	SaturdayMemberClose *any `json:"saturday_member_close,omitempty"`
	SaturdayMemberOpen *any `json:"saturday_member_open,omitempty"`
	SaturdayPublicClose *any `json:"saturday_public_close,omitempty"`
	SaturdayPublicOpen *any `json:"saturday_public_open,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Summary *any `json:"summary,omitempty"`
	SundayIsClosed *any `json:"sunday_is_closed,omitempty"`
	SundayMemberClose *any `json:"sunday_member_close,omitempty"`
	SundayMemberOpen *any `json:"sunday_member_open,omitempty"`
	SundayPublicClose *any `json:"sunday_public_close,omitempty"`
	SundayPublicOpen *any `json:"sunday_public_open,omitempty"`
	ThursdayIsClosed *any `json:"thursday_is_closed,omitempty"`
	ThursdayMemberClose *any `json:"thursday_member_close,omitempty"`
	ThursdayMemberOpen *any `json:"thursday_member_open,omitempty"`
	ThursdayPublicClose *any `json:"thursday_public_close,omitempty"`
	ThursdayPublicOpen *any `json:"thursday_public_open,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	TuesdayIsClosed *any `json:"tuesday_is_closed,omitempty"`
	TuesdayMemberClose *any `json:"tuesday_member_close,omitempty"`
	TuesdayMemberOpen *any `json:"tuesday_member_open,omitempty"`
	TuesdayPublicClose *any `json:"tuesday_public_close,omitempty"`
	TuesdayPublicOpen *any `json:"tuesday_public_open,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WednesdayIsClosed *any `json:"wednesday_is_closed,omitempty"`
	WednesdayMemberClose *any `json:"wednesday_member_close,omitempty"`
	WednesdayMemberOpen *any `json:"wednesday_member_open,omitempty"`
	WednesdayPublicClose *any `json:"wednesday_public_close,omitempty"`
	WednesdayPublicOpen *any `json:"wednesday_public_open,omitempty"`
}

// Image is the typed data model for the image entity.
type Image struct {
	Ahash *any `json:"ahash,omitempty"`
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Color *any `json:"color,omitempty"`
	Colorfulness *any `json:"colorfulness,omitempty"`
	Content *any `json:"content,omitempty"`
	ContentETag *any `json:"content_e_tag,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	Fingerprint *any `json:"fingerprint,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Id *string `json:"id,omitempty"`
	IiifUrl *any `json:"iiif_url,omitempty"`
	IsEducationalResource *bool `json:"is_educational_resource,omitempty"`
	IsMultimediaResource *bool `json:"is_multimedia_resource,omitempty"`
	IsTeacherResource *bool `json:"is_teacher_resource,omitempty"`
	LakeGuid *any `json:"lake_guid,omitempty"`
	Lqip *any `json:"lqip,omitempty"`
	Phash *any `json:"phash,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *any `json:"type,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	Width *float64 `json:"width,omitempty"`
}

// ImageLoadMatch is the typed request payload for Image.LoadTyped.
type ImageLoadMatch struct {
	Id string `json:"id"`
}

// ImageListMatch is the typed request payload for Image.ListTyped.
type ImageListMatch struct {
	Ahash *any `json:"ahash,omitempty"`
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Color *any `json:"color,omitempty"`
	Colorfulness *any `json:"colorfulness,omitempty"`
	Content *any `json:"content,omitempty"`
	ContentETag *any `json:"content_e_tag,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	Fingerprint *any `json:"fingerprint,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Id *string `json:"id,omitempty"`
	IiifUrl *any `json:"iiif_url,omitempty"`
	IsEducationalResource *bool `json:"is_educational_resource,omitempty"`
	IsMultimediaResource *bool `json:"is_multimedia_resource,omitempty"`
	IsTeacherResource *bool `json:"is_teacher_resource,omitempty"`
	LakeGuid *any `json:"lake_guid,omitempty"`
	Lqip *any `json:"lqip,omitempty"`
	Phash *any `json:"phash,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *any `json:"type,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	Width *float64 `json:"width,omitempty"`
}

// LandingPage is the typed data model for the landing_page entity.
type LandingPage struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SearchTags *any `json:"search_tags,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// LandingPageLoadMatch is the typed request payload for LandingPage.LoadTyped.
type LandingPageLoadMatch struct {
	Id string `json:"id"`
}

// LandingPageListMatch is the typed request payload for LandingPage.ListTyped.
type LandingPageListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SearchTags *any `json:"search_tags,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Place is the typed data model for the place entity.
type Place struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	TgnId *string `json:"tgn_id,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// PlaceLoadMatch is the typed request payload for Place.LoadTyped.
type PlaceLoadMatch struct {
	Id string `json:"id"`
}

// PlaceListMatch is the typed request payload for Place.ListTyped.
type PlaceListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	TgnId *string `json:"tgn_id,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// PressRelease is the typed data model for the press_release entity.
type PressRelease struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// PressReleaseLoadMatch is the typed request payload for PressRelease.LoadTyped.
type PressReleaseLoadMatch struct {
	Id string `json:"id"`
}

// PressReleaseListMatch is the typed request payload for PressRelease.ListTyped.
type PressReleaseListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// PrintedPublication is the typed data model for the printed_publication entity.
type PrintedPublication struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// PrintedPublicationLoadMatch is the typed request payload for PrintedPublication.LoadTyped.
type PrintedPublicationLoadMatch struct {
	Id string `json:"id"`
}

// PrintedPublicationListMatch is the typed request payload for PrintedPublication.ListTyped.
type PrintedPublicationListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Product is the typed data model for the product entity.
type Product struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistIds *any `json:"artist_ids,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	Description *string `json:"description,omitempty"`
	ExhibitionIds *any `json:"exhibition_ids,omitempty"`
	ExternalSku *any `json:"external_sku,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	MaxCompareAtPrice *any `json:"max_compare_at_price,omitempty"`
	MaxCurrentPrice *any `json:"max_current_price,omitempty"`
	MinCompareAtPrice *any `json:"min_compare_at_price,omitempty"`
	MinCurrentPrice *any `json:"min_current_price,omitempty"`
	PriceDisplay *any `json:"price_display,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// ProductLoadMatch is the typed request payload for Product.LoadTyped.
type ProductLoadMatch struct {
	Id string `json:"id"`
}

// ProductListMatch is the typed request payload for Product.ListTyped.
type ProductListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistIds *any `json:"artist_ids,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	Description *string `json:"description,omitempty"`
	ExhibitionIds *any `json:"exhibition_ids,omitempty"`
	ExternalSku *any `json:"external_sku,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	MaxCompareAtPrice *any `json:"max_compare_at_price,omitempty"`
	MaxCurrentPrice *any `json:"max_current_price,omitempty"`
	MinCompareAtPrice *any `json:"min_compare_at_price,omitempty"`
	MinCurrentPrice *any `json:"min_current_price,omitempty"`
	PriceDisplay *any `json:"price_display,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Publication is the typed data model for the publication entity.
type Publication struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SectionIds *any `json:"section_ids,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// PublicationLoadMatch is the typed request payload for Publication.LoadTyped.
type PublicationLoadMatch struct {
	Id string `json:"id"`
}

// PublicationListMatch is the typed request payload for Publication.ListTyped.
type PublicationListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SectionIds *any `json:"section_ids,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Search is the typed data model for the search entity.
type Search struct {
	ApiId *string `json:"api_id,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	IsBoosted *bool `json:"is_boosted,omitempty"`
	Score *float64 `json:"score,omitempty"`
	Thumbnail *any `json:"thumbnail,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
}

// SearchListMatch is the typed request payload for Search.ListTyped.
type SearchListMatch struct {
	ApiId *string `json:"api_id,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	IsBoosted *bool `json:"is_boosted,omitempty"`
	Score *float64 `json:"score,omitempty"`
	Thumbnail *any `json:"thumbnail,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
}

// Section is the typed data model for the section entity.
type Section struct {
	Accession *any `json:"accession,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	Content *any `json:"content,omitempty"`
	GenericPageId *string `json:"generic_page_id,omitempty"`
	Id *string `json:"id,omitempty"`
	PublicationId *string `json:"publication_id,omitempty"`
	PublicationTitle *any `json:"publication_title,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// SectionLoadMatch is the typed request payload for Section.LoadTyped.
type SectionLoadMatch struct {
	Id string `json:"id"`
}

// SectionListMatch is the typed request payload for Section.ListTyped.
type SectionListMatch struct {
	Accession *any `json:"accession,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	Content *any `json:"content,omitempty"`
	GenericPageId *string `json:"generic_page_id,omitempty"`
	Id *string `json:"id,omitempty"`
	PublicationId *string `json:"publication_id,omitempty"`
	PublicationTitle *any `json:"publication_title,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Site is the typed data model for the site entity.
type Site struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Description *string `json:"description,omitempty"`
	ExhibitionIds *any `json:"exhibition_ids,omitempty"`
	ExhibitionTitles *any `json:"exhibition_titles,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// SiteLoadMatch is the typed request payload for Site.LoadTyped.
type SiteLoadMatch struct {
	Id string `json:"id"`
}

// SiteListMatch is the typed request payload for Site.ListTyped.
type SiteListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Description *string `json:"description,omitempty"`
	ExhibitionIds *any `json:"exhibition_ids,omitempty"`
	ExhibitionTitles *any `json:"exhibition_titles,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Sound is the typed data model for the sound entity.
type Sound struct {
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Content *any `json:"content,omitempty"`
	ContentETag *any `json:"content_e_tag,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	Id *string `json:"id,omitempty"`
	IsEducationalResource *bool `json:"is_educational_resource,omitempty"`
	IsMultimediaResource *bool `json:"is_multimedia_resource,omitempty"`
	IsTeacherResource *bool `json:"is_teacher_resource,omitempty"`
	LakeGuid *any `json:"lake_guid,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	Transcript *any `json:"transcript,omitempty"`
	Type *any `json:"type,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// SoundLoadMatch is the typed request payload for Sound.LoadTyped.
type SoundLoadMatch struct {
	Id string `json:"id"`
}

// SoundListMatch is the typed request payload for Sound.ListTyped.
type SoundListMatch struct {
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Content *any `json:"content,omitempty"`
	ContentETag *any `json:"content_e_tag,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	Id *string `json:"id,omitempty"`
	IsEducationalResource *bool `json:"is_educational_resource,omitempty"`
	IsMultimediaResource *bool `json:"is_multimedia_resource,omitempty"`
	IsTeacherResource *bool `json:"is_teacher_resource,omitempty"`
	LakeGuid *any `json:"lake_guid,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	Transcript *any `json:"transcript,omitempty"`
	Type *any `json:"type,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// StaticPage is the typed data model for the static_page entity.
type StaticPage struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// StaticPageLoadMatch is the typed request payload for StaticPage.LoadTyped.
type StaticPageLoadMatch struct {
	Id string `json:"id"`
}

// StaticPageListMatch is the typed request payload for StaticPage.ListTyped.
type StaticPageListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	WebUrl *any `json:"web_url,omitempty"`
}

// Text is the typed data model for the text entity.
type Text struct {
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Content *any `json:"content,omitempty"`
	ContentETag *any `json:"content_e_tag,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	Id *string `json:"id,omitempty"`
	IsEducationalResource *bool `json:"is_educational_resource,omitempty"`
	IsMultimediaResource *bool `json:"is_multimedia_resource,omitempty"`
	IsTeacherResource *bool `json:"is_teacher_resource,omitempty"`
	LakeGuid *any `json:"lake_guid,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *any `json:"type,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// TextLoadMatch is the typed request payload for Text.LoadTyped.
type TextLoadMatch struct {
	Id string `json:"id"`
}

// TextListMatch is the typed request payload for Text.ListTyped.
type TextListMatch struct {
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Content *any `json:"content,omitempty"`
	ContentETag *any `json:"content_e_tag,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	Id *string `json:"id,omitempty"`
	IsEducationalResource *bool `json:"is_educational_resource,omitempty"`
	IsMultimediaResource *bool `json:"is_multimedia_resource,omitempty"`
	IsTeacherResource *bool `json:"is_teacher_resource,omitempty"`
	LakeGuid *any `json:"lake_guid,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *any `json:"type,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// Tour is the typed data model for the tour entity.
type Tour struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistTitles *any `json:"artist_titles,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Image *any `json:"image,omitempty"`
	Intro *any `json:"intro,omitempty"`
	IntroLink *any `json:"intro_link,omitempty"`
	IntroTranscript *any `json:"intro_transcript,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	Weight *float64 `json:"weight,omitempty"`
}

// TourLoadMatch is the typed request payload for Tour.LoadTyped.
type TourLoadMatch struct {
	Id string `json:"id"`
}

// TourListMatch is the typed request payload for Tour.ListTyped.
type TourListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistTitles *any `json:"artist_titles,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Image *any `json:"image,omitempty"`
	Intro *any `json:"intro,omitempty"`
	IntroLink *any `json:"intro_link,omitempty"`
	IntroTranscript *any `json:"intro_transcript,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	Weight *float64 `json:"weight,omitempty"`
}

// Video is the typed data model for the video entity.
type Video struct {
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Content *any `json:"content,omitempty"`
	ContentETag *any `json:"content_e_tag,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	Id *string `json:"id,omitempty"`
	IsEducationalResource *bool `json:"is_educational_resource,omitempty"`
	IsMultimediaResource *bool `json:"is_multimedia_resource,omitempty"`
	IsTeacherResource *bool `json:"is_teacher_resource,omitempty"`
	LakeGuid *any `json:"lake_guid,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *any `json:"type,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// VideoLoadMatch is the typed request payload for Video.LoadTyped.
type VideoLoadMatch struct {
	Id string `json:"id"`
}

// VideoListMatch is the typed request payload for Video.ListTyped.
type VideoListMatch struct {
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkIds *any `json:"artwork_ids,omitempty"`
	ArtworkTitles *any `json:"artwork_titles,omitempty"`
	Content *any `json:"content,omitempty"`
	ContentETag *any `json:"content_e_tag,omitempty"`
	CreditLine *any `json:"credit_line,omitempty"`
	Id *string `json:"id,omitempty"`
	IsEducationalResource *bool `json:"is_educational_resource,omitempty"`
	IsMultimediaResource *bool `json:"is_multimedia_resource,omitempty"`
	IsTeacherResource *bool `json:"is_teacher_resource,omitempty"`
	LakeGuid *any `json:"lake_guid,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *any `json:"type,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
