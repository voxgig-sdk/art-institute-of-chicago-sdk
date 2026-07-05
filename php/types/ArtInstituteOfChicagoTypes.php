<?php
declare(strict_types=1);

// Typed models for the ArtInstituteOfChicago SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Agent entity data model. */
class Agent
{
    public mixed $alt_title = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $birth_date = null;
    public mixed $death_date = null;
    public ?string $description = null;
    public ?string $id = null;
    public ?bool $is_artist = null;
    public mixed $sort_title = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public ?string $ulan_id = null;
    public mixed $updated_at = null;
}

/** Request payload for Agent#load. */
class AgentLoadMatch
{
    public string $id;
}

/** Request payload for Agent#list. */
class AgentListMatch
{
    public mixed $alt_title = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $birth_date = null;
    public mixed $death_date = null;
    public ?string $description = null;
    public ?string $id = null;
    public ?bool $is_artist = null;
    public mixed $sort_title = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public ?string $ulan_id = null;
    public mixed $updated_at = null;
}

/** AgentRole entity data model. */
class AgentRole
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for AgentRole#load. */
class AgentRoleLoadMatch
{
    public string $id;
}

/** Request payload for AgentRole#list. */
class AgentRoleListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** AgentType entity data model. */
class AgentType
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for AgentType#load. */
class AgentTypeLoadMatch
{
    public string $id;
}

/** Request payload for AgentType#list. */
class AgentTypeListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Article entity data model. */
class Article
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for Article#load. */
class ArticleLoadMatch
{
    public string $id;
}

/** Request payload for Article#list. */
class ArticleListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Artwork entity data model. */
class Artwork
{
    public ?string $alt_artist_id = null;
    public ?string $alt_classification_id = null;
    public ?string $alt_image_id = null;
    public ?string $alt_material_id = null;
    public ?string $alt_style_id = null;
    public ?string $alt_subject_id = null;
    public ?string $alt_technique_id = null;
    public mixed $alt_title = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $artist_display = null;
    public ?string $artist_id = null;
    public mixed $artist_title = null;
    public ?string $artwork_type_id = null;
    public mixed $artwork_type_title = null;
    public mixed $boost_rank = null;
    public mixed $catalog_based_search_keyword_title = null;
    public mixed $catalogue_display = null;
    public ?string $category_id = null;
    public mixed $category_title = null;
    public ?string $classification_id = null;
    public mixed $classification_title = null;
    public mixed $color = null;
    public mixed $colorfulness = null;
    public mixed $copyright_notice = null;
    public mixed $credit_line = null;
    public mixed $date_display = null;
    public mixed $date_end = null;
    public ?string $date_qualifier_id = null;
    public mixed $date_qualifier_title = null;
    public mixed $date_start = null;
    public ?string $department_id = null;
    public mixed $department_title = null;
    public ?string $description = null;
    public mixed $dimension = null;
    public mixed $dimensions_detail = null;
    public ?string $document_id = null;
    public mixed $edition = null;
    public mixed $exhibition_history = null;
    public mixed $fiscal_year = null;
    public mixed $fiscal_year_deaccession = null;
    public ?string $gallery_id = null;
    public mixed $gallery_title = null;
    public ?bool $has_advanced_imaging = null;
    public ?bool $has_educational_resource = null;
    public ?bool $has_multimedia_resource = null;
    public ?bool $has_not_been_viewed_much = null;
    public ?string $id = null;
    public mixed $image_embedding = null;
    public ?string $image_id = null;
    public mixed $inscription = null;
    public ?string $internal_department_id = null;
    public ?bool $is_boosted = null;
    public ?bool $is_on_view = null;
    public ?bool $is_public_domain = null;
    public ?bool $is_zoomable = null;
    public ?float $latitude = null;
    public mixed $latlon = null;
    public ?float $longitude = null;
    public ?int $main_reference_number = null;
    public ?string $material_id = null;
    public mixed $material_title = null;
    public mixed $max_zoom_window_size = null;
    public mixed $medium_display = null;
    public ?string $nomisma_id = null;
    public mixed $on_loan_display = null;
    public mixed $pageview = null;
    public mixed $pageviews_recent = null;
    public mixed $place_of_origin = null;
    public mixed $provenance_text = null;
    public mixed $publication_history = null;
    public mixed $publishing_verification_level = null;
    public ?string $section_id = null;
    public mixed $section_title = null;
    public mixed $short_description = null;
    public ?string $site_id = null;
    public ?string $sound_id = null;
    public mixed $source_updated_at = null;
    public ?string $style_id = null;
    public mixed $style_title = null;
    public ?string $subject_id = null;
    public mixed $subject_title = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public ?string $technique_id = null;
    public mixed $technique_title = null;
    public mixed $term_title = null;
    public mixed $text_embedding = null;
    public ?string $text_id = null;
    public mixed $theme_title = null;
    public mixed $thumbnail = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public ?string $video_id = null;
}

/** Request payload for Artwork#load. */
class ArtworkLoadMatch
{
    public string $id;
}

/** Request payload for Artwork#list. */
class ArtworkListMatch
{
    public ?string $alt_artist_id = null;
    public ?string $alt_classification_id = null;
    public ?string $alt_image_id = null;
    public ?string $alt_material_id = null;
    public ?string $alt_style_id = null;
    public ?string $alt_subject_id = null;
    public ?string $alt_technique_id = null;
    public mixed $alt_title = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $artist_display = null;
    public ?string $artist_id = null;
    public mixed $artist_title = null;
    public ?string $artwork_type_id = null;
    public mixed $artwork_type_title = null;
    public mixed $boost_rank = null;
    public mixed $catalog_based_search_keyword_title = null;
    public mixed $catalogue_display = null;
    public ?string $category_id = null;
    public mixed $category_title = null;
    public ?string $classification_id = null;
    public mixed $classification_title = null;
    public mixed $color = null;
    public mixed $colorfulness = null;
    public mixed $copyright_notice = null;
    public mixed $credit_line = null;
    public mixed $date_display = null;
    public mixed $date_end = null;
    public ?string $date_qualifier_id = null;
    public mixed $date_qualifier_title = null;
    public mixed $date_start = null;
    public ?string $department_id = null;
    public mixed $department_title = null;
    public ?string $description = null;
    public mixed $dimension = null;
    public mixed $dimensions_detail = null;
    public ?string $document_id = null;
    public mixed $edition = null;
    public mixed $exhibition_history = null;
    public mixed $fiscal_year = null;
    public mixed $fiscal_year_deaccession = null;
    public ?string $gallery_id = null;
    public mixed $gallery_title = null;
    public ?bool $has_advanced_imaging = null;
    public ?bool $has_educational_resource = null;
    public ?bool $has_multimedia_resource = null;
    public ?bool $has_not_been_viewed_much = null;
    public ?string $id = null;
    public mixed $image_embedding = null;
    public ?string $image_id = null;
    public mixed $inscription = null;
    public ?string $internal_department_id = null;
    public ?bool $is_boosted = null;
    public ?bool $is_on_view = null;
    public ?bool $is_public_domain = null;
    public ?bool $is_zoomable = null;
    public ?float $latitude = null;
    public mixed $latlon = null;
    public ?float $longitude = null;
    public ?int $main_reference_number = null;
    public ?string $material_id = null;
    public mixed $material_title = null;
    public mixed $max_zoom_window_size = null;
    public mixed $medium_display = null;
    public ?string $nomisma_id = null;
    public mixed $on_loan_display = null;
    public mixed $pageview = null;
    public mixed $pageviews_recent = null;
    public mixed $place_of_origin = null;
    public mixed $provenance_text = null;
    public mixed $publication_history = null;
    public mixed $publishing_verification_level = null;
    public ?string $section_id = null;
    public mixed $section_title = null;
    public mixed $short_description = null;
    public ?string $site_id = null;
    public ?string $sound_id = null;
    public mixed $source_updated_at = null;
    public ?string $style_id = null;
    public mixed $style_title = null;
    public ?string $subject_id = null;
    public mixed $subject_title = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public ?string $technique_id = null;
    public mixed $technique_title = null;
    public mixed $term_title = null;
    public mixed $text_embedding = null;
    public ?string $text_id = null;
    public mixed $theme_title = null;
    public mixed $thumbnail = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public ?string $video_id = null;
}

/** ArtworkDateQualifier entity data model. */
class ArtworkDateQualifier
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for ArtworkDateQualifier#load. */
class ArtworkDateQualifierLoadMatch
{
    public string $id;
}

/** Request payload for ArtworkDateQualifier#list. */
class ArtworkDateQualifierListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** ArtworkPlaceQualifier entity data model. */
class ArtworkPlaceQualifier
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for ArtworkPlaceQualifier#load. */
class ArtworkPlaceQualifierLoadMatch
{
    public string $id;
}

/** Request payload for ArtworkPlaceQualifier#list. */
class ArtworkPlaceQualifierListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** ArtworkType entity data model. */
class ArtworkType
{
    public ?string $aat_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for ArtworkType#load. */
class ArtworkTypeLoadMatch
{
    public string $id;
}

/** Request payload for ArtworkType#list. */
class ArtworkTypeListMatch
{
    public ?string $aat_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** CategoryTerm entity data model. */
class CategoryTerm
{
    public ?string $aat_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?string $parent_id = null;
    public mixed $source_updated_at = null;
    public mixed $subtype = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for CategoryTerm#load. */
class CategoryTermLoadMatch
{
    public string $id;
}

/** Request payload for CategoryTerm#list. */
class CategoryTermListMatch
{
    public ?string $aat_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?string $parent_id = null;
    public mixed $source_updated_at = null;
    public mixed $subtype = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** DigitalPublication entity data model. */
class DigitalPublication
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for DigitalPublication#load. */
class DigitalPublicationLoadMatch
{
    public string $id;
}

/** Request payload for DigitalPublication#list. */
class DigitalPublicationListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** DigitalPublicationArticle entity data model. */
class DigitalPublicationArticle
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $author_display = null;
    public mixed $copy = null;
    public ?string $digital_publication_id = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for DigitalPublicationArticle#load. */
class DigitalPublicationArticleLoadMatch
{
    public string $id;
}

/** Request payload for DigitalPublicationArticle#list. */
class DigitalPublicationArticleListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $author_display = null;
    public mixed $copy = null;
    public ?string $digital_publication_id = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** EducatorResource entity data model. */
class EducatorResource
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for EducatorResource#load. */
class EducatorResourceLoadMatch
{
    public string $id;
}

/** Request payload for EducatorResource#list. */
class EducatorResourceListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Event entity data model. */
class Event
{
    public ?string $alt_audience_id = null;
    public ?string $alt_event_type_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $audience_id = null;
    public mixed $buy_button_caption = null;
    public mixed $buy_button_text = null;
    public mixed $date_display = null;
    public ?string $description = null;
    public mixed $door_time = null;
    public mixed $end_date = null;
    public mixed $end_time = null;
    public mixed $entrance = null;
    public ?string $event_host_id = null;
    public mixed $event_host_title = null;
    public ?string $event_type_id = null;
    public mixed $header_description = null;
    public mixed $hero_caption = null;
    public ?string $id = null;
    public mixed $image_url = null;
    public ?bool $is_admission_required = null;
    public ?bool $is_after_hour = null;
    public ?bool $is_free = null;
    public ?bool $is_member_exclusive = null;
    public ?bool $is_private = null;
    public ?bool $is_registration_required = null;
    public ?bool $is_sales_button_hidden = null;
    public ?bool $is_sold_out = null;
    public ?bool $is_ticketed = null;
    public ?bool $is_virtual_event = null;
    public mixed $join_url = null;
    public mixed $layout_type = null;
    public mixed $list_description = null;
    public mixed $location = null;
    public ?string $program_id = null;
    public mixed $program_title = null;
    public mixed $rsvp_link = null;
    public mixed $search_tag = null;
    public mixed $short_description = null;
    public ?string $slug = null;
    public mixed $source_updated_at = null;
    public mixed $start_date = null;
    public mixed $start_time = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $survey_url = null;
    public ?string $ticketed_event_id = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $title_display = null;
    public mixed $updated_at = null;
    public mixed $virtual_event_passcode = null;
    public mixed $virtual_event_url = null;
}

/** Request payload for Event#load. */
class EventLoadMatch
{
    public string $id;
}

/** Request payload for Event#list. */
class EventListMatch
{
    public ?string $alt_audience_id = null;
    public ?string $alt_event_type_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $audience_id = null;
    public mixed $buy_button_caption = null;
    public mixed $buy_button_text = null;
    public mixed $date_display = null;
    public ?string $description = null;
    public mixed $door_time = null;
    public mixed $end_date = null;
    public mixed $end_time = null;
    public mixed $entrance = null;
    public ?string $event_host_id = null;
    public mixed $event_host_title = null;
    public ?string $event_type_id = null;
    public mixed $header_description = null;
    public mixed $hero_caption = null;
    public ?string $id = null;
    public mixed $image_url = null;
    public ?bool $is_admission_required = null;
    public ?bool $is_after_hour = null;
    public ?bool $is_free = null;
    public ?bool $is_member_exclusive = null;
    public ?bool $is_private = null;
    public ?bool $is_registration_required = null;
    public ?bool $is_sales_button_hidden = null;
    public ?bool $is_sold_out = null;
    public ?bool $is_ticketed = null;
    public ?bool $is_virtual_event = null;
    public mixed $join_url = null;
    public mixed $layout_type = null;
    public mixed $list_description = null;
    public mixed $location = null;
    public ?string $program_id = null;
    public mixed $program_title = null;
    public mixed $rsvp_link = null;
    public mixed $search_tag = null;
    public mixed $short_description = null;
    public ?string $slug = null;
    public mixed $source_updated_at = null;
    public mixed $start_date = null;
    public mixed $start_time = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $survey_url = null;
    public ?string $ticketed_event_id = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $title_display = null;
    public mixed $updated_at = null;
    public mixed $virtual_event_passcode = null;
    public mixed $virtual_event_url = null;
}

/** EventOccurrence entity data model. */
class EventOccurrence
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $button_caption = null;
    public mixed $button_text = null;
    public mixed $button_url = null;
    public ?string $description = null;
    public mixed $end_at = null;
    public ?string $event_id = null;
    public ?string $id = null;
    public mixed $image_url = null;
    public ?bool $is_private = null;
    public ?bool $is_sales_button_hidden = null;
    public ?bool $is_ticketed = null;
    public mixed $location = null;
    public mixed $off_sale_at = null;
    public mixed $on_sale_at = null;
    public mixed $short_description = null;
    public mixed $source_updated_at = null;
    public mixed $start_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $title_display = null;
    public mixed $updated_at = null;
}

/** Request payload for EventOccurrence#load. */
class EventOccurrenceLoadMatch
{
    public string $id;
}

/** Request payload for EventOccurrence#list. */
class EventOccurrenceListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $button_caption = null;
    public mixed $button_text = null;
    public mixed $button_url = null;
    public ?string $description = null;
    public mixed $end_at = null;
    public ?string $event_id = null;
    public ?string $id = null;
    public mixed $image_url = null;
    public ?bool $is_private = null;
    public ?bool $is_sales_button_hidden = null;
    public ?bool $is_ticketed = null;
    public mixed $location = null;
    public mixed $off_sale_at = null;
    public mixed $on_sale_at = null;
    public mixed $short_description = null;
    public mixed $source_updated_at = null;
    public mixed $start_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $title_display = null;
    public mixed $updated_at = null;
}

/** EventProgram entity data model. */
class EventProgram
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?bool $is_affiliate_group = null;
    public ?bool $is_event_host = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for EventProgram#load. */
class EventProgramLoadMatch
{
    public string $id;
}

/** Request payload for EventProgram#list. */
class EventProgramListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?bool $is_affiliate_group = null;
    public ?bool $is_event_host = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Exhibition entity data model. */
class Exhibition
{
    public mixed $aic_end_at = null;
    public mixed $aic_start_at = null;
    public ?string $alt_image_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artist_id = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public ?string $document_id = null;
    public ?string $gallery_id = null;
    public mixed $gallery_title = null;
    public ?string $id = null;
    public ?string $image_id = null;
    public mixed $image_url = null;
    public ?bool $is_featured = null;
    public ?bool $is_published = null;
    public mixed $position = null;
    public mixed $short_description = null;
    public ?string $site_id = null;
    public mixed $source_updated_at = null;
    public mixed $status = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for Exhibition#load. */
class ExhibitionLoadMatch
{
    public string $id;
}

/** Request payload for Exhibition#list. */
class ExhibitionListMatch
{
    public mixed $aic_end_at = null;
    public mixed $aic_start_at = null;
    public ?string $alt_image_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artist_id = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public ?string $document_id = null;
    public ?string $gallery_id = null;
    public mixed $gallery_title = null;
    public ?string $id = null;
    public ?string $image_id = null;
    public mixed $image_url = null;
    public ?bool $is_featured = null;
    public ?bool $is_published = null;
    public mixed $position = null;
    public mixed $short_description = null;
    public ?string $site_id = null;
    public mixed $source_updated_at = null;
    public mixed $status = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Gallery entity data model. */
class Gallery
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $floor = null;
    public ?string $id = null;
    public ?bool $is_closed = null;
    public ?float $latitude = null;
    public mixed $latlon = null;
    public ?float $longitude = null;
    public mixed $number = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public ?string $tgn_id = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for Gallery#load. */
class GalleryLoadMatch
{
    public string $id;
}

/** Request payload for Gallery#list. */
class GalleryListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $floor = null;
    public ?string $id = null;
    public ?bool $is_closed = null;
    public ?float $latitude = null;
    public mixed $latlon = null;
    public ?float $longitude = null;
    public mixed $number = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public ?string $tgn_id = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** GenericPage entity data model. */
class GenericPage
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $search_tag = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for GenericPage#load. */
class GenericPageLoadMatch
{
    public string $id;
}

/** Request payload for GenericPage#list. */
class GenericPageListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $search_tag = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Highlight entity data model. */
class Highlight
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for Highlight#load. */
class HighlightLoadMatch
{
    public string $id;
}

/** Request payload for Highlight#list. */
class HighlightListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Hour entity data model. */
class Hour
{
    public mixed $additional_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $friday_is_closed = null;
    public mixed $friday_member_close = null;
    public mixed $friday_member_open = null;
    public mixed $friday_public_close = null;
    public mixed $friday_public_open = null;
    public ?string $id = null;
    public mixed $monday_is_closed = null;
    public mixed $monday_member_close = null;
    public mixed $monday_member_open = null;
    public mixed $monday_public_close = null;
    public mixed $monday_public_open = null;
    public mixed $saturday_is_closed = null;
    public mixed $saturday_member_close = null;
    public mixed $saturday_member_open = null;
    public mixed $saturday_public_close = null;
    public mixed $saturday_public_open = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $summary = null;
    public mixed $sunday_is_closed = null;
    public mixed $sunday_member_close = null;
    public mixed $sunday_member_open = null;
    public mixed $sunday_public_close = null;
    public mixed $sunday_public_open = null;
    public mixed $thursday_is_closed = null;
    public mixed $thursday_member_close = null;
    public mixed $thursday_member_open = null;
    public mixed $thursday_public_close = null;
    public mixed $thursday_public_open = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $tuesday_is_closed = null;
    public mixed $tuesday_member_close = null;
    public mixed $tuesday_member_open = null;
    public mixed $tuesday_public_close = null;
    public mixed $tuesday_public_open = null;
    public mixed $updated_at = null;
    public mixed $wednesday_is_closed = null;
    public mixed $wednesday_member_close = null;
    public mixed $wednesday_member_open = null;
    public mixed $wednesday_public_close = null;
    public mixed $wednesday_public_open = null;
}

/** Request payload for Hour#load. */
class HourLoadMatch
{
    public string $id;
}

/** Request payload for Hour#list. */
class HourListMatch
{
    public mixed $additional_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $friday_is_closed = null;
    public mixed $friday_member_close = null;
    public mixed $friday_member_open = null;
    public mixed $friday_public_close = null;
    public mixed $friday_public_open = null;
    public ?string $id = null;
    public mixed $monday_is_closed = null;
    public mixed $monday_member_close = null;
    public mixed $monday_member_open = null;
    public mixed $monday_public_close = null;
    public mixed $monday_public_open = null;
    public mixed $saturday_is_closed = null;
    public mixed $saturday_member_close = null;
    public mixed $saturday_member_open = null;
    public mixed $saturday_public_close = null;
    public mixed $saturday_public_open = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $summary = null;
    public mixed $sunday_is_closed = null;
    public mixed $sunday_member_close = null;
    public mixed $sunday_member_open = null;
    public mixed $sunday_public_close = null;
    public mixed $sunday_public_open = null;
    public mixed $thursday_is_closed = null;
    public mixed $thursday_member_close = null;
    public mixed $thursday_member_open = null;
    public mixed $thursday_public_close = null;
    public mixed $thursday_public_open = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $tuesday_is_closed = null;
    public mixed $tuesday_member_close = null;
    public mixed $tuesday_member_open = null;
    public mixed $tuesday_public_close = null;
    public mixed $tuesday_public_open = null;
    public mixed $updated_at = null;
    public mixed $wednesday_is_closed = null;
    public mixed $wednesday_member_close = null;
    public mixed $wednesday_member_open = null;
    public mixed $wednesday_public_close = null;
    public mixed $wednesday_public_open = null;
}

/** Image entity data model. */
class Image
{
    public mixed $ahash = null;
    public mixed $alt_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public mixed $color = null;
    public mixed $colorfulness = null;
    public mixed $content = null;
    public mixed $content_e_tag = null;
    public mixed $credit_line = null;
    public mixed $fingerprint = null;
    public ?float $height = null;
    public ?string $id = null;
    public mixed $iiif_url = null;
    public ?bool $is_educational_resource = null;
    public ?bool $is_multimedia_resource = null;
    public ?bool $is_teacher_resource = null;
    public mixed $lake_guid = null;
    public mixed $lqip = null;
    public mixed $phash = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $type = null;
    public mixed $updated_at = null;
    public ?float $width = null;
}

/** Request payload for Image#load. */
class ImageLoadMatch
{
    public string $id;
}

/** Request payload for Image#list. */
class ImageListMatch
{
    public mixed $ahash = null;
    public mixed $alt_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public mixed $color = null;
    public mixed $colorfulness = null;
    public mixed $content = null;
    public mixed $content_e_tag = null;
    public mixed $credit_line = null;
    public mixed $fingerprint = null;
    public ?float $height = null;
    public ?string $id = null;
    public mixed $iiif_url = null;
    public ?bool $is_educational_resource = null;
    public ?bool $is_multimedia_resource = null;
    public ?bool $is_teacher_resource = null;
    public mixed $lake_guid = null;
    public mixed $lqip = null;
    public mixed $phash = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $type = null;
    public mixed $updated_at = null;
    public ?float $width = null;
}

/** LandingPage entity data model. */
class LandingPage
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $search_tag = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for LandingPage#load. */
class LandingPageLoadMatch
{
    public string $id;
}

/** Request payload for LandingPage#list. */
class LandingPageListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $search_tag = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Place entity data model. */
class Place
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?float $latitude = null;
    public ?float $longitude = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public ?string $tgn_id = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** Request payload for Place#load. */
class PlaceLoadMatch
{
    public string $id;
}

/** Request payload for Place#list. */
class PlaceListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?float $latitude = null;
    public ?float $longitude = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public ?string $tgn_id = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
}

/** PressRelease entity data model. */
class PressRelease
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for PressRelease#load. */
class PressReleaseLoadMatch
{
    public string $id;
}

/** Request payload for PressRelease#list. */
class PressReleaseListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** PrintedPublication entity data model. */
class PrintedPublication
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for PrintedPublication#load. */
class PrintedPublicationLoadMatch
{
    public string $id;
}

/** Request payload for PrintedPublication#list. */
class PrintedPublicationListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $copy = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Product entity data model. */
class Product
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artist_id = null;
    public ?string $artwork_id = null;
    public ?string $description = null;
    public ?string $exhibition_id = null;
    public mixed $external_sku = null;
    public ?string $id = null;
    public mixed $image_url = null;
    public mixed $max_compare_at_price = null;
    public mixed $max_current_price = null;
    public mixed $min_compare_at_price = null;
    public mixed $min_current_price = null;
    public mixed $price_display = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for Product#load. */
class ProductLoadMatch
{
    public string $id;
}

/** Request payload for Product#list. */
class ProductListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artist_id = null;
    public ?string $artwork_id = null;
    public ?string $description = null;
    public ?string $exhibition_id = null;
    public mixed $external_sku = null;
    public ?string $id = null;
    public mixed $image_url = null;
    public mixed $max_compare_at_price = null;
    public mixed $max_current_price = null;
    public mixed $min_compare_at_price = null;
    public mixed $min_current_price = null;
    public mixed $price_display = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Publication entity data model. */
class Publication
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?string $section_id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for Publication#load. */
class PublicationLoadMatch
{
    public string $id;
}

/** Request payload for Publication#list. */
class PublicationListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?string $section_id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Search entity data model. */
class Search
{
    public ?string $api_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?bool $is_boosted = null;
    public ?float $score = null;
    public mixed $thumbnail = null;
    public mixed $timestamp = null;
    public ?string $title = null;
}

/** Request payload for Search#list. */
class SearchListMatch
{
    public ?string $api_id = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public ?bool $is_boosted = null;
    public ?float $score = null;
    public mixed $thumbnail = null;
    public mixed $timestamp = null;
    public ?string $title = null;
}

/** Section entity data model. */
class Section
{
    public mixed $accession = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $content = null;
    public ?string $generic_page_id = null;
    public ?string $id = null;
    public ?string $publication_id = null;
    public mixed $publication_title = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for Section#load. */
class SectionLoadMatch
{
    public string $id;
}

/** Request payload for Section#list. */
class SectionListMatch
{
    public mixed $accession = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $content = null;
    public ?string $generic_page_id = null;
    public ?string $id = null;
    public ?string $publication_id = null;
    public mixed $publication_title = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Site entity data model. */
class Site
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public ?string $description = null;
    public ?string $exhibition_id = null;
    public mixed $exhibition_title = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for Site#load. */
class SiteLoadMatch
{
    public string $id;
}

/** Request payload for Site#list. */
class SiteListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public ?string $description = null;
    public ?string $exhibition_id = null;
    public mixed $exhibition_title = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Sound entity data model. */
class Sound
{
    public mixed $alt_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public mixed $content = null;
    public mixed $content_e_tag = null;
    public mixed $credit_line = null;
    public ?string $id = null;
    public ?bool $is_educational_resource = null;
    public ?bool $is_multimedia_resource = null;
    public ?bool $is_teacher_resource = null;
    public mixed $lake_guid = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $transcript = null;
    public mixed $type = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for Sound#load. */
class SoundLoadMatch
{
    public string $id;
}

/** Request payload for Sound#list. */
class SoundListMatch
{
    public mixed $alt_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public mixed $content = null;
    public mixed $content_e_tag = null;
    public mixed $credit_line = null;
    public ?string $id = null;
    public ?bool $is_educational_resource = null;
    public ?bool $is_multimedia_resource = null;
    public ?bool $is_teacher_resource = null;
    public mixed $lake_guid = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $transcript = null;
    public mixed $type = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** StaticPage entity data model. */
class StaticPage
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Request payload for StaticPage#load. */
class StaticPageLoadMatch
{
    public string $id;
}

/** Request payload for StaticPage#list. */
class StaticPageListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $id = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public mixed $web_url = null;
}

/** Text entity data model. */
class Text
{
    public mixed $alt_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public mixed $content = null;
    public mixed $content_e_tag = null;
    public mixed $credit_line = null;
    public ?string $id = null;
    public ?bool $is_educational_resource = null;
    public ?bool $is_multimedia_resource = null;
    public ?bool $is_teacher_resource = null;
    public mixed $lake_guid = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $type = null;
    public mixed $updated_at = null;
}

/** Request payload for Text#load. */
class TextLoadMatch
{
    public string $id;
}

/** Request payload for Text#list. */
class TextListMatch
{
    public mixed $alt_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public mixed $content = null;
    public mixed $content_e_tag = null;
    public mixed $credit_line = null;
    public ?string $id = null;
    public ?bool $is_educational_resource = null;
    public ?bool $is_multimedia_resource = null;
    public ?bool $is_teacher_resource = null;
    public mixed $lake_guid = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $type = null;
    public mixed $updated_at = null;
}

/** Tour entity data model. */
class Tour
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $artist_title = null;
    public mixed $artwork_title = null;
    public ?string $description = null;
    public ?string $id = null;
    public mixed $image = null;
    public mixed $intro = null;
    public mixed $intro_link = null;
    public mixed $intro_transcript = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public ?float $weight = null;
}

/** Request payload for Tour#load. */
class TourLoadMatch
{
    public string $id;
}

/** Request payload for Tour#list. */
class TourListMatch
{
    public mixed $api_link = null;
    public mixed $api_model = null;
    public mixed $artist_title = null;
    public mixed $artwork_title = null;
    public ?string $description = null;
    public ?string $id = null;
    public mixed $image = null;
    public mixed $intro = null;
    public mixed $intro_link = null;
    public mixed $intro_transcript = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $updated_at = null;
    public ?float $weight = null;
}

/** Video entity data model. */
class Video
{
    public mixed $alt_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public mixed $content = null;
    public mixed $content_e_tag = null;
    public mixed $credit_line = null;
    public ?string $id = null;
    public ?bool $is_educational_resource = null;
    public ?bool $is_multimedia_resource = null;
    public ?bool $is_teacher_resource = null;
    public mixed $lake_guid = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $type = null;
    public mixed $updated_at = null;
}

/** Request payload for Video#load. */
class VideoLoadMatch
{
    public string $id;
}

/** Request payload for Video#list. */
class VideoListMatch
{
    public mixed $alt_text = null;
    public mixed $api_link = null;
    public mixed $api_model = null;
    public ?string $artwork_id = null;
    public mixed $artwork_title = null;
    public mixed $content = null;
    public mixed $content_e_tag = null;
    public mixed $credit_line = null;
    public ?string $id = null;
    public ?bool $is_educational_resource = null;
    public ?bool $is_multimedia_resource = null;
    public ?bool $is_teacher_resource = null;
    public mixed $lake_guid = null;
    public mixed $source_updated_at = null;
    public mixed $suggest_autocomplete_all = null;
    public mixed $suggest_autocomplete_boosted = null;
    public mixed $timestamp = null;
    public ?string $title = null;
    public mixed $type = null;
    public mixed $updated_at = null;
}

