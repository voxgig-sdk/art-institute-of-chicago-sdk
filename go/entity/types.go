// Typed models for the ArtInstituteOfChicago SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Agent is the typed data model for the agent entity.
type Agent struct {
	AltTitle *any `json:"alt_title,omitempty"`
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

// AgentListMatch mirrors the agent fields as an all-optional match
// filter (Go analog of Partial<Agent>).
type AgentListMatch struct {
	AltTitle *any `json:"alt_title,omitempty"`
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

// AgentRoleListMatch mirrors the agent_role fields as an all-optional match
// filter (Go analog of Partial<AgentRole>).
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

// AgentTypeListMatch mirrors the agent_type fields as an all-optional match
// filter (Go analog of Partial<AgentType>).
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

// ArticleListMatch mirrors the article fields as an all-optional match
// filter (Go analog of Partial<Article>).
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
	AltArtistId *string `json:"alt_artist_id,omitempty"`
	AltClassificationId *string `json:"alt_classification_id,omitempty"`
	AltImageId *string `json:"alt_image_id,omitempty"`
	AltMaterialId *string `json:"alt_material_id,omitempty"`
	AltStyleId *string `json:"alt_style_id,omitempty"`
	AltSubjectId *string `json:"alt_subject_id,omitempty"`
	AltTechniqueId *string `json:"alt_technique_id,omitempty"`
	AltTitle *any `json:"alt_title,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistDisplay *any `json:"artist_display,omitempty"`
	ArtistId *string `json:"artist_id,omitempty"`
	ArtistTitle *any `json:"artist_title,omitempty"`
	ArtworkTypeId *string `json:"artwork_type_id,omitempty"`
	ArtworkTypeTitle *any `json:"artwork_type_title,omitempty"`
	BoostRank *any `json:"boost_rank,omitempty"`
	CatalogBasedSearchKeywordTitle *any `json:"catalog_based_search_keyword_title,omitempty"`
	CatalogueDisplay *any `json:"catalogue_display,omitempty"`
	CategoryId *string `json:"category_id,omitempty"`
	CategoryTitle *any `json:"category_title,omitempty"`
	ClassificationId *string `json:"classification_id,omitempty"`
	ClassificationTitle *any `json:"classification_title,omitempty"`
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
	Dimension *any `json:"dimension,omitempty"`
	DimensionsDetail *any `json:"dimensions_detail,omitempty"`
	DocumentId *string `json:"document_id,omitempty"`
	Edition *any `json:"edition,omitempty"`
	ExhibitionHistory *any `json:"exhibition_history,omitempty"`
	FiscalYear *any `json:"fiscal_year,omitempty"`
	FiscalYearDeaccession *any `json:"fiscal_year_deaccession,omitempty"`
	GalleryId *string `json:"gallery_id,omitempty"`
	GalleryTitle *any `json:"gallery_title,omitempty"`
	HasAdvancedImaging *bool `json:"has_advanced_imaging,omitempty"`
	HasEducationalResource *bool `json:"has_educational_resource,omitempty"`
	HasMultimediaResource *bool `json:"has_multimedia_resource,omitempty"`
	HasNotBeenViewedMuch *bool `json:"has_not_been_viewed_much,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageEmbedding *any `json:"image_embedding,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	Inscription *any `json:"inscription,omitempty"`
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
	MaterialTitle *any `json:"material_title,omitempty"`
	MaxZoomWindowSize *any `json:"max_zoom_window_size,omitempty"`
	MediumDisplay *any `json:"medium_display,omitempty"`
	NomismaId *string `json:"nomisma_id,omitempty"`
	OnLoanDisplay *any `json:"on_loan_display,omitempty"`
	Pageview *any `json:"pageview,omitempty"`
	PageviewsRecent *any `json:"pageviews_recent,omitempty"`
	PlaceOfOrigin *any `json:"place_of_origin,omitempty"`
	ProvenanceText *any `json:"provenance_text,omitempty"`
	PublicationHistory *any `json:"publication_history,omitempty"`
	PublishingVerificationLevel *any `json:"publishing_verification_level,omitempty"`
	SectionId *string `json:"section_id,omitempty"`
	SectionTitle *any `json:"section_title,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	SoundId *string `json:"sound_id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	StyleId *string `json:"style_id,omitempty"`
	StyleTitle *any `json:"style_title,omitempty"`
	SubjectId *string `json:"subject_id,omitempty"`
	SubjectTitle *any `json:"subject_title,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	TechniqueId *string `json:"technique_id,omitempty"`
	TechniqueTitle *any `json:"technique_title,omitempty"`
	TermTitle *any `json:"term_title,omitempty"`
	TextEmbedding *any `json:"text_embedding,omitempty"`
	TextId *string `json:"text_id,omitempty"`
	ThemeTitle *any `json:"theme_title,omitempty"`
	Thumbnail *any `json:"thumbnail,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	VideoId *string `json:"video_id,omitempty"`
}

// ArtworkLoadMatch is the typed request payload for Artwork.LoadTyped.
type ArtworkLoadMatch struct {
	Id string `json:"id"`
}

// ArtworkListMatch mirrors the artwork fields as an all-optional match
// filter (Go analog of Partial<Artwork>).
type ArtworkListMatch struct {
	AltArtistId *string `json:"alt_artist_id,omitempty"`
	AltClassificationId *string `json:"alt_classification_id,omitempty"`
	AltImageId *string `json:"alt_image_id,omitempty"`
	AltMaterialId *string `json:"alt_material_id,omitempty"`
	AltStyleId *string `json:"alt_style_id,omitempty"`
	AltSubjectId *string `json:"alt_subject_id,omitempty"`
	AltTechniqueId *string `json:"alt_technique_id,omitempty"`
	AltTitle *any `json:"alt_title,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistDisplay *any `json:"artist_display,omitempty"`
	ArtistId *string `json:"artist_id,omitempty"`
	ArtistTitle *any `json:"artist_title,omitempty"`
	ArtworkTypeId *string `json:"artwork_type_id,omitempty"`
	ArtworkTypeTitle *any `json:"artwork_type_title,omitempty"`
	BoostRank *any `json:"boost_rank,omitempty"`
	CatalogBasedSearchKeywordTitle *any `json:"catalog_based_search_keyword_title,omitempty"`
	CatalogueDisplay *any `json:"catalogue_display,omitempty"`
	CategoryId *string `json:"category_id,omitempty"`
	CategoryTitle *any `json:"category_title,omitempty"`
	ClassificationId *string `json:"classification_id,omitempty"`
	ClassificationTitle *any `json:"classification_title,omitempty"`
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
	Dimension *any `json:"dimension,omitempty"`
	DimensionsDetail *any `json:"dimensions_detail,omitempty"`
	DocumentId *string `json:"document_id,omitempty"`
	Edition *any `json:"edition,omitempty"`
	ExhibitionHistory *any `json:"exhibition_history,omitempty"`
	FiscalYear *any `json:"fiscal_year,omitempty"`
	FiscalYearDeaccession *any `json:"fiscal_year_deaccession,omitempty"`
	GalleryId *string `json:"gallery_id,omitempty"`
	GalleryTitle *any `json:"gallery_title,omitempty"`
	HasAdvancedImaging *bool `json:"has_advanced_imaging,omitempty"`
	HasEducationalResource *bool `json:"has_educational_resource,omitempty"`
	HasMultimediaResource *bool `json:"has_multimedia_resource,omitempty"`
	HasNotBeenViewedMuch *bool `json:"has_not_been_viewed_much,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageEmbedding *any `json:"image_embedding,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	Inscription *any `json:"inscription,omitempty"`
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
	MaterialTitle *any `json:"material_title,omitempty"`
	MaxZoomWindowSize *any `json:"max_zoom_window_size,omitempty"`
	MediumDisplay *any `json:"medium_display,omitempty"`
	NomismaId *string `json:"nomisma_id,omitempty"`
	OnLoanDisplay *any `json:"on_loan_display,omitempty"`
	Pageview *any `json:"pageview,omitempty"`
	PageviewsRecent *any `json:"pageviews_recent,omitempty"`
	PlaceOfOrigin *any `json:"place_of_origin,omitempty"`
	ProvenanceText *any `json:"provenance_text,omitempty"`
	PublicationHistory *any `json:"publication_history,omitempty"`
	PublishingVerificationLevel *any `json:"publishing_verification_level,omitempty"`
	SectionId *string `json:"section_id,omitempty"`
	SectionTitle *any `json:"section_title,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	SoundId *string `json:"sound_id,omitempty"`
	SourceUpdatedAt *any `json:"source_updated_at,omitempty"`
	StyleId *string `json:"style_id,omitempty"`
	StyleTitle *any `json:"style_title,omitempty"`
	SubjectId *string `json:"subject_id,omitempty"`
	SubjectTitle *any `json:"subject_title,omitempty"`
	SuggestAutocompleteAll *any `json:"suggest_autocomplete_all,omitempty"`
	SuggestAutocompleteBoosted *any `json:"suggest_autocomplete_boosted,omitempty"`
	TechniqueId *string `json:"technique_id,omitempty"`
	TechniqueTitle *any `json:"technique_title,omitempty"`
	TermTitle *any `json:"term_title,omitempty"`
	TextEmbedding *any `json:"text_embedding,omitempty"`
	TextId *string `json:"text_id,omitempty"`
	ThemeTitle *any `json:"theme_title,omitempty"`
	Thumbnail *any `json:"thumbnail,omitempty"`
	Timestamp *any `json:"timestamp,omitempty"`
	Title *string `json:"title,omitempty"`
	UpdatedAt *any `json:"updated_at,omitempty"`
	VideoId *string `json:"video_id,omitempty"`
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

// ArtworkDateQualifierListMatch mirrors the artwork_date_qualifier fields as an all-optional match
// filter (Go analog of Partial<ArtworkDateQualifier>).
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

// ArtworkPlaceQualifierListMatch mirrors the artwork_place_qualifier fields as an all-optional match
// filter (Go analog of Partial<ArtworkPlaceQualifier>).
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

// ArtworkTypeListMatch mirrors the artwork_type fields as an all-optional match
// filter (Go analog of Partial<ArtworkType>).
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

// CategoryTermListMatch mirrors the category_term fields as an all-optional match
// filter (Go analog of Partial<CategoryTerm>).
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

// DigitalPublicationListMatch mirrors the digital_publication fields as an all-optional match
// filter (Go analog of Partial<DigitalPublication>).
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

// DigitalPublicationArticleListMatch mirrors the digital_publication_article fields as an all-optional match
// filter (Go analog of Partial<DigitalPublicationArticle>).
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

// EducatorResourceListMatch mirrors the educator_resource fields as an all-optional match
// filter (Go analog of Partial<EducatorResource>).
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
	AltAudienceId *string `json:"alt_audience_id,omitempty"`
	AltEventTypeId *string `json:"alt_event_type_id,omitempty"`
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
	IsAfterHour *bool `json:"is_after_hour,omitempty"`
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
	ProgramId *string `json:"program_id,omitempty"`
	ProgramTitle *any `json:"program_title,omitempty"`
	RsvpLink *any `json:"rsvp_link,omitempty"`
	SearchTag *any `json:"search_tag,omitempty"`
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

// EventListMatch mirrors the event fields as an all-optional match
// filter (Go analog of Partial<Event>).
type EventListMatch struct {
	AltAudienceId *string `json:"alt_audience_id,omitempty"`
	AltEventTypeId *string `json:"alt_event_type_id,omitempty"`
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
	IsAfterHour *bool `json:"is_after_hour,omitempty"`
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
	ProgramId *string `json:"program_id,omitempty"`
	ProgramTitle *any `json:"program_title,omitempty"`
	RsvpLink *any `json:"rsvp_link,omitempty"`
	SearchTag *any `json:"search_tag,omitempty"`
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

// EventOccurrenceListMatch mirrors the event_occurrence fields as an all-optional match
// filter (Go analog of Partial<EventOccurrence>).
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

// EventProgramListMatch mirrors the event_program fields as an all-optional match
// filter (Go analog of Partial<EventProgram>).
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
	AltImageId *string `json:"alt_image_id,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistId *string `json:"artist_id,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
	DocumentId *string `json:"document_id,omitempty"`
	GalleryId *string `json:"gallery_id,omitempty"`
	GalleryTitle *any `json:"gallery_title,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	IsFeatured *bool `json:"is_featured,omitempty"`
	IsPublished *bool `json:"is_published,omitempty"`
	Position *any `json:"position,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
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

// ExhibitionListMatch mirrors the exhibition fields as an all-optional match
// filter (Go analog of Partial<Exhibition>).
type ExhibitionListMatch struct {
	AicEndAt *any `json:"aic_end_at,omitempty"`
	AicStartAt *any `json:"aic_start_at,omitempty"`
	AltImageId *string `json:"alt_image_id,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistId *string `json:"artist_id,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
	DocumentId *string `json:"document_id,omitempty"`
	GalleryId *string `json:"gallery_id,omitempty"`
	GalleryTitle *any `json:"gallery_title,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	ImageUrl *any `json:"image_url,omitempty"`
	IsFeatured *bool `json:"is_featured,omitempty"`
	IsPublished *bool `json:"is_published,omitempty"`
	Position *any `json:"position,omitempty"`
	ShortDescription *any `json:"short_description,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
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

// GalleryListMatch mirrors the gallery fields as an all-optional match
// filter (Go analog of Partial<Gallery>).
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
	SearchTag *any `json:"search_tag,omitempty"`
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

// GenericPageListMatch mirrors the generic_page fields as an all-optional match
// filter (Go analog of Partial<GenericPage>).
type GenericPageListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SearchTag *any `json:"search_tag,omitempty"`
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

// HighlightListMatch mirrors the highlight fields as an all-optional match
// filter (Go analog of Partial<Highlight>).
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

// HourListMatch mirrors the hour fields as an all-optional match
// filter (Go analog of Partial<Hour>).
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
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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

// ImageListMatch mirrors the image fields as an all-optional match
// filter (Go analog of Partial<Image>).
type ImageListMatch struct {
	Ahash *any `json:"ahash,omitempty"`
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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
	SearchTag *any `json:"search_tag,omitempty"`
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

// LandingPageListMatch mirrors the landing_page fields as an all-optional match
// filter (Go analog of Partial<LandingPage>).
type LandingPageListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Copy *any `json:"copy,omitempty"`
	Id *string `json:"id,omitempty"`
	SearchTag *any `json:"search_tag,omitempty"`
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

// PlaceListMatch mirrors the place fields as an all-optional match
// filter (Go analog of Partial<Place>).
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

// PressReleaseListMatch mirrors the press_release fields as an all-optional match
// filter (Go analog of Partial<PressRelease>).
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

// PrintedPublicationListMatch mirrors the printed_publication fields as an all-optional match
// filter (Go analog of Partial<PrintedPublication>).
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
	ArtistId *string `json:"artist_id,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	Description *string `json:"description,omitempty"`
	ExhibitionId *string `json:"exhibition_id,omitempty"`
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

// ProductListMatch mirrors the product fields as an all-optional match
// filter (Go analog of Partial<Product>).
type ProductListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistId *string `json:"artist_id,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	Description *string `json:"description,omitempty"`
	ExhibitionId *string `json:"exhibition_id,omitempty"`
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
	SectionId *string `json:"section_id,omitempty"`
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

// PublicationListMatch mirrors the publication fields as an all-optional match
// filter (Go analog of Partial<Publication>).
type PublicationListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	Id *string `json:"id,omitempty"`
	SectionId *string `json:"section_id,omitempty"`
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

// SearchListMatch mirrors the search fields as an all-optional match
// filter (Go analog of Partial<Search>).
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

// SectionListMatch mirrors the section fields as an all-optional match
// filter (Go analog of Partial<Section>).
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
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
	Description *string `json:"description,omitempty"`
	ExhibitionId *string `json:"exhibition_id,omitempty"`
	ExhibitionTitle *any `json:"exhibition_title,omitempty"`
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

// SiteListMatch mirrors the site fields as an all-optional match
// filter (Go analog of Partial<Site>).
type SiteListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
	Description *string `json:"description,omitempty"`
	ExhibitionId *string `json:"exhibition_id,omitempty"`
	ExhibitionTitle *any `json:"exhibition_title,omitempty"`
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
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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

// SoundListMatch mirrors the sound fields as an all-optional match
// filter (Go analog of Partial<Sound>).
type SoundListMatch struct {
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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

// StaticPageListMatch mirrors the static_page fields as an all-optional match
// filter (Go analog of Partial<StaticPage>).
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
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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

// TextListMatch mirrors the text fields as an all-optional match
// filter (Go analog of Partial<Text>).
type TextListMatch struct {
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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
	ArtistTitle *any `json:"artist_title,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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

// TourListMatch mirrors the tour fields as an all-optional match
// filter (Go analog of Partial<Tour>).
type TourListMatch struct {
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtistTitle *any `json:"artist_title,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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

// VideoListMatch mirrors the video fields as an all-optional match
// filter (Go analog of Partial<Video>).
type VideoListMatch struct {
	AltText *any `json:"alt_text,omitempty"`
	ApiLink *any `json:"api_link,omitempty"`
	ApiModel *any `json:"api_model,omitempty"`
	ArtworkId *string `json:"artwork_id,omitempty"`
	ArtworkTitle *any `json:"artwork_title,omitempty"`
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

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
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

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
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
