# Typed models for the ArtInstituteOfChicago SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Agent:
    alt_title: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    birth_date: Optional[Any] = None
    death_date: Optional[Any] = None
    description: Optional[str] = None
    id: Optional[str] = None
    is_artist: Optional[bool] = None
    sort_title: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    ulan_id: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class AgentLoadMatch:
    id: str


@dataclass
class AgentListMatch:
    alt_title: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    birth_date: Optional[Any] = None
    death_date: Optional[Any] = None
    description: Optional[str] = None
    id: Optional[str] = None
    is_artist: Optional[bool] = None
    sort_title: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    ulan_id: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class AgentRole:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class AgentRoleLoadMatch:
    id: str


@dataclass
class AgentRoleListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class AgentType:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class AgentTypeLoadMatch:
    id: str


@dataclass
class AgentTypeListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class Article:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class ArticleLoadMatch:
    id: str


@dataclass
class ArticleListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class Artwork:
    alt_artist_id: Optional[str] = None
    alt_classification_id: Optional[str] = None
    alt_image_id: Optional[str] = None
    alt_material_id: Optional[str] = None
    alt_style_id: Optional[str] = None
    alt_subject_id: Optional[str] = None
    alt_technique_id: Optional[str] = None
    alt_title: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artist_display: Optional[Any] = None
    artist_id: Optional[str] = None
    artist_title: Optional[Any] = None
    artwork_type_id: Optional[str] = None
    artwork_type_title: Optional[Any] = None
    boost_rank: Optional[Any] = None
    catalog_based_search_keyword_title: Optional[Any] = None
    catalogue_display: Optional[Any] = None
    category_id: Optional[str] = None
    category_title: Optional[Any] = None
    classification_id: Optional[str] = None
    classification_title: Optional[Any] = None
    color: Optional[Any] = None
    colorfulness: Optional[Any] = None
    copyright_notice: Optional[Any] = None
    credit_line: Optional[Any] = None
    date_display: Optional[Any] = None
    date_end: Optional[Any] = None
    date_qualifier_id: Optional[str] = None
    date_qualifier_title: Optional[Any] = None
    date_start: Optional[Any] = None
    department_id: Optional[str] = None
    department_title: Optional[Any] = None
    description: Optional[str] = None
    dimension: Optional[Any] = None
    dimensions_detail: Optional[Any] = None
    document_id: Optional[str] = None
    edition: Optional[Any] = None
    exhibition_history: Optional[Any] = None
    fiscal_year: Optional[Any] = None
    fiscal_year_deaccession: Optional[Any] = None
    gallery_id: Optional[str] = None
    gallery_title: Optional[Any] = None
    has_advanced_imaging: Optional[bool] = None
    has_educational_resource: Optional[bool] = None
    has_multimedia_resource: Optional[bool] = None
    has_not_been_viewed_much: Optional[bool] = None
    id: Optional[str] = None
    image_embedding: Optional[Any] = None
    image_id: Optional[str] = None
    inscription: Optional[Any] = None
    internal_department_id: Optional[str] = None
    is_boosted: Optional[bool] = None
    is_on_view: Optional[bool] = None
    is_public_domain: Optional[bool] = None
    is_zoomable: Optional[bool] = None
    latitude: Optional[float] = None
    latlon: Optional[Any] = None
    longitude: Optional[float] = None
    main_reference_number: Optional[int] = None
    material_id: Optional[str] = None
    material_title: Optional[Any] = None
    max_zoom_window_size: Optional[Any] = None
    medium_display: Optional[Any] = None
    nomisma_id: Optional[str] = None
    on_loan_display: Optional[Any] = None
    pageview: Optional[Any] = None
    pageviews_recent: Optional[Any] = None
    place_of_origin: Optional[Any] = None
    provenance_text: Optional[Any] = None
    publication_history: Optional[Any] = None
    publishing_verification_level: Optional[Any] = None
    section_id: Optional[str] = None
    section_title: Optional[Any] = None
    short_description: Optional[Any] = None
    site_id: Optional[str] = None
    sound_id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    style_id: Optional[str] = None
    style_title: Optional[Any] = None
    subject_id: Optional[str] = None
    subject_title: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    technique_id: Optional[str] = None
    technique_title: Optional[Any] = None
    term_title: Optional[Any] = None
    text_embedding: Optional[Any] = None
    text_id: Optional[str] = None
    theme_title: Optional[Any] = None
    thumbnail: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    video_id: Optional[str] = None


@dataclass
class ArtworkLoadMatch:
    id: str


@dataclass
class ArtworkListMatch:
    alt_artist_id: Optional[str] = None
    alt_classification_id: Optional[str] = None
    alt_image_id: Optional[str] = None
    alt_material_id: Optional[str] = None
    alt_style_id: Optional[str] = None
    alt_subject_id: Optional[str] = None
    alt_technique_id: Optional[str] = None
    alt_title: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artist_display: Optional[Any] = None
    artist_id: Optional[str] = None
    artist_title: Optional[Any] = None
    artwork_type_id: Optional[str] = None
    artwork_type_title: Optional[Any] = None
    boost_rank: Optional[Any] = None
    catalog_based_search_keyword_title: Optional[Any] = None
    catalogue_display: Optional[Any] = None
    category_id: Optional[str] = None
    category_title: Optional[Any] = None
    classification_id: Optional[str] = None
    classification_title: Optional[Any] = None
    color: Optional[Any] = None
    colorfulness: Optional[Any] = None
    copyright_notice: Optional[Any] = None
    credit_line: Optional[Any] = None
    date_display: Optional[Any] = None
    date_end: Optional[Any] = None
    date_qualifier_id: Optional[str] = None
    date_qualifier_title: Optional[Any] = None
    date_start: Optional[Any] = None
    department_id: Optional[str] = None
    department_title: Optional[Any] = None
    description: Optional[str] = None
    dimension: Optional[Any] = None
    dimensions_detail: Optional[Any] = None
    document_id: Optional[str] = None
    edition: Optional[Any] = None
    exhibition_history: Optional[Any] = None
    fiscal_year: Optional[Any] = None
    fiscal_year_deaccession: Optional[Any] = None
    gallery_id: Optional[str] = None
    gallery_title: Optional[Any] = None
    has_advanced_imaging: Optional[bool] = None
    has_educational_resource: Optional[bool] = None
    has_multimedia_resource: Optional[bool] = None
    has_not_been_viewed_much: Optional[bool] = None
    id: Optional[str] = None
    image_embedding: Optional[Any] = None
    image_id: Optional[str] = None
    inscription: Optional[Any] = None
    internal_department_id: Optional[str] = None
    is_boosted: Optional[bool] = None
    is_on_view: Optional[bool] = None
    is_public_domain: Optional[bool] = None
    is_zoomable: Optional[bool] = None
    latitude: Optional[float] = None
    latlon: Optional[Any] = None
    longitude: Optional[float] = None
    main_reference_number: Optional[int] = None
    material_id: Optional[str] = None
    material_title: Optional[Any] = None
    max_zoom_window_size: Optional[Any] = None
    medium_display: Optional[Any] = None
    nomisma_id: Optional[str] = None
    on_loan_display: Optional[Any] = None
    pageview: Optional[Any] = None
    pageviews_recent: Optional[Any] = None
    place_of_origin: Optional[Any] = None
    provenance_text: Optional[Any] = None
    publication_history: Optional[Any] = None
    publishing_verification_level: Optional[Any] = None
    section_id: Optional[str] = None
    section_title: Optional[Any] = None
    short_description: Optional[Any] = None
    site_id: Optional[str] = None
    sound_id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    style_id: Optional[str] = None
    style_title: Optional[Any] = None
    subject_id: Optional[str] = None
    subject_title: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    technique_id: Optional[str] = None
    technique_title: Optional[Any] = None
    term_title: Optional[Any] = None
    text_embedding: Optional[Any] = None
    text_id: Optional[str] = None
    theme_title: Optional[Any] = None
    thumbnail: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    video_id: Optional[str] = None


@dataclass
class ArtworkDateQualifier:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class ArtworkDateQualifierLoadMatch:
    id: str


@dataclass
class ArtworkDateQualifierListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class ArtworkPlaceQualifier:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class ArtworkPlaceQualifierLoadMatch:
    id: str


@dataclass
class ArtworkPlaceQualifierListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class ArtworkType:
    aat_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class ArtworkTypeLoadMatch:
    id: str


@dataclass
class ArtworkTypeListMatch:
    aat_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class CategoryTerm:
    aat_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    parent_id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    subtype: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class CategoryTermLoadMatch:
    id: str


@dataclass
class CategoryTermListMatch:
    aat_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    parent_id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    subtype: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class DigitalPublication:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class DigitalPublicationLoadMatch:
    id: str


@dataclass
class DigitalPublicationListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class DigitalPublicationArticle:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    author_display: Optional[Any] = None
    copy: Optional[Any] = None
    digital_publication_id: Optional[str] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class DigitalPublicationArticleLoadMatch:
    id: str


@dataclass
class DigitalPublicationArticleListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    author_display: Optional[Any] = None
    copy: Optional[Any] = None
    digital_publication_id: Optional[str] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class EducatorResource:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class EducatorResourceLoadMatch:
    id: str


@dataclass
class EducatorResourceListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Event:
    alt_audience_id: Optional[str] = None
    alt_event_type_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    audience_id: Optional[str] = None
    buy_button_caption: Optional[Any] = None
    buy_button_text: Optional[Any] = None
    date_display: Optional[Any] = None
    description: Optional[str] = None
    door_time: Optional[Any] = None
    end_date: Optional[Any] = None
    end_time: Optional[Any] = None
    entrance: Optional[Any] = None
    event_host_id: Optional[str] = None
    event_host_title: Optional[Any] = None
    event_type_id: Optional[str] = None
    header_description: Optional[Any] = None
    hero_caption: Optional[Any] = None
    id: Optional[str] = None
    image_url: Optional[Any] = None
    is_admission_required: Optional[bool] = None
    is_after_hour: Optional[bool] = None
    is_free: Optional[bool] = None
    is_member_exclusive: Optional[bool] = None
    is_private: Optional[bool] = None
    is_registration_required: Optional[bool] = None
    is_sales_button_hidden: Optional[bool] = None
    is_sold_out: Optional[bool] = None
    is_ticketed: Optional[bool] = None
    is_virtual_event: Optional[bool] = None
    join_url: Optional[Any] = None
    layout_type: Optional[Any] = None
    list_description: Optional[Any] = None
    location: Optional[Any] = None
    program_id: Optional[str] = None
    program_title: Optional[Any] = None
    rsvp_link: Optional[Any] = None
    search_tag: Optional[Any] = None
    short_description: Optional[Any] = None
    slug: Optional[str] = None
    source_updated_at: Optional[Any] = None
    start_date: Optional[Any] = None
    start_time: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    survey_url: Optional[Any] = None
    ticketed_event_id: Optional[str] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    title_display: Optional[Any] = None
    updated_at: Optional[Any] = None
    virtual_event_passcode: Optional[Any] = None
    virtual_event_url: Optional[Any] = None


@dataclass
class EventLoadMatch:
    id: str


@dataclass
class EventListMatch:
    alt_audience_id: Optional[str] = None
    alt_event_type_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    audience_id: Optional[str] = None
    buy_button_caption: Optional[Any] = None
    buy_button_text: Optional[Any] = None
    date_display: Optional[Any] = None
    description: Optional[str] = None
    door_time: Optional[Any] = None
    end_date: Optional[Any] = None
    end_time: Optional[Any] = None
    entrance: Optional[Any] = None
    event_host_id: Optional[str] = None
    event_host_title: Optional[Any] = None
    event_type_id: Optional[str] = None
    header_description: Optional[Any] = None
    hero_caption: Optional[Any] = None
    id: Optional[str] = None
    image_url: Optional[Any] = None
    is_admission_required: Optional[bool] = None
    is_after_hour: Optional[bool] = None
    is_free: Optional[bool] = None
    is_member_exclusive: Optional[bool] = None
    is_private: Optional[bool] = None
    is_registration_required: Optional[bool] = None
    is_sales_button_hidden: Optional[bool] = None
    is_sold_out: Optional[bool] = None
    is_ticketed: Optional[bool] = None
    is_virtual_event: Optional[bool] = None
    join_url: Optional[Any] = None
    layout_type: Optional[Any] = None
    list_description: Optional[Any] = None
    location: Optional[Any] = None
    program_id: Optional[str] = None
    program_title: Optional[Any] = None
    rsvp_link: Optional[Any] = None
    search_tag: Optional[Any] = None
    short_description: Optional[Any] = None
    slug: Optional[str] = None
    source_updated_at: Optional[Any] = None
    start_date: Optional[Any] = None
    start_time: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    survey_url: Optional[Any] = None
    ticketed_event_id: Optional[str] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    title_display: Optional[Any] = None
    updated_at: Optional[Any] = None
    virtual_event_passcode: Optional[Any] = None
    virtual_event_url: Optional[Any] = None


@dataclass
class EventOccurrence:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    button_caption: Optional[Any] = None
    button_text: Optional[Any] = None
    button_url: Optional[Any] = None
    description: Optional[str] = None
    end_at: Optional[Any] = None
    event_id: Optional[str] = None
    id: Optional[str] = None
    image_url: Optional[Any] = None
    is_private: Optional[bool] = None
    is_sales_button_hidden: Optional[bool] = None
    is_ticketed: Optional[bool] = None
    location: Optional[Any] = None
    off_sale_at: Optional[Any] = None
    on_sale_at: Optional[Any] = None
    short_description: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    start_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    title_display: Optional[Any] = None
    updated_at: Optional[Any] = None


@dataclass
class EventOccurrenceLoadMatch:
    id: str


@dataclass
class EventOccurrenceListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    button_caption: Optional[Any] = None
    button_text: Optional[Any] = None
    button_url: Optional[Any] = None
    description: Optional[str] = None
    end_at: Optional[Any] = None
    event_id: Optional[str] = None
    id: Optional[str] = None
    image_url: Optional[Any] = None
    is_private: Optional[bool] = None
    is_sales_button_hidden: Optional[bool] = None
    is_ticketed: Optional[bool] = None
    location: Optional[Any] = None
    off_sale_at: Optional[Any] = None
    on_sale_at: Optional[Any] = None
    short_description: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    start_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    title_display: Optional[Any] = None
    updated_at: Optional[Any] = None


@dataclass
class EventProgram:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    is_affiliate_group: Optional[bool] = None
    is_event_host: Optional[bool] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class EventProgramLoadMatch:
    id: str


@dataclass
class EventProgramListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    is_affiliate_group: Optional[bool] = None
    is_event_host: Optional[bool] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class Exhibition:
    aic_end_at: Optional[Any] = None
    aic_start_at: Optional[Any] = None
    alt_image_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artist_id: Optional[str] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    document_id: Optional[str] = None
    gallery_id: Optional[str] = None
    gallery_title: Optional[Any] = None
    id: Optional[str] = None
    image_id: Optional[str] = None
    image_url: Optional[Any] = None
    is_featured: Optional[bool] = None
    is_published: Optional[bool] = None
    position: Optional[Any] = None
    short_description: Optional[Any] = None
    site_id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    status: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class ExhibitionLoadMatch:
    id: str


@dataclass
class ExhibitionListMatch:
    aic_end_at: Optional[Any] = None
    aic_start_at: Optional[Any] = None
    alt_image_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artist_id: Optional[str] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    document_id: Optional[str] = None
    gallery_id: Optional[str] = None
    gallery_title: Optional[Any] = None
    id: Optional[str] = None
    image_id: Optional[str] = None
    image_url: Optional[Any] = None
    is_featured: Optional[bool] = None
    is_published: Optional[bool] = None
    position: Optional[Any] = None
    short_description: Optional[Any] = None
    site_id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    status: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Gallery:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    floor: Optional[Any] = None
    id: Optional[str] = None
    is_closed: Optional[bool] = None
    latitude: Optional[float] = None
    latlon: Optional[Any] = None
    longitude: Optional[float] = None
    number: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    tgn_id: Optional[str] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class GalleryLoadMatch:
    id: str


@dataclass
class GalleryListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    floor: Optional[Any] = None
    id: Optional[str] = None
    is_closed: Optional[bool] = None
    latitude: Optional[float] = None
    latlon: Optional[Any] = None
    longitude: Optional[float] = None
    number: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    tgn_id: Optional[str] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class GenericPage:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    search_tag: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class GenericPageLoadMatch:
    id: str


@dataclass
class GenericPageListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    search_tag: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Highlight:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class HighlightLoadMatch:
    id: str


@dataclass
class HighlightListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class Hour:
    additional_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    friday_is_closed: Optional[Any] = None
    friday_member_close: Optional[Any] = None
    friday_member_open: Optional[Any] = None
    friday_public_close: Optional[Any] = None
    friday_public_open: Optional[Any] = None
    id: Optional[str] = None
    monday_is_closed: Optional[Any] = None
    monday_member_close: Optional[Any] = None
    monday_member_open: Optional[Any] = None
    monday_public_close: Optional[Any] = None
    monday_public_open: Optional[Any] = None
    saturday_is_closed: Optional[Any] = None
    saturday_member_close: Optional[Any] = None
    saturday_member_open: Optional[Any] = None
    saturday_public_close: Optional[Any] = None
    saturday_public_open: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    summary: Optional[Any] = None
    sunday_is_closed: Optional[Any] = None
    sunday_member_close: Optional[Any] = None
    sunday_member_open: Optional[Any] = None
    sunday_public_close: Optional[Any] = None
    sunday_public_open: Optional[Any] = None
    thursday_is_closed: Optional[Any] = None
    thursday_member_close: Optional[Any] = None
    thursday_member_open: Optional[Any] = None
    thursday_public_close: Optional[Any] = None
    thursday_public_open: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    tuesday_is_closed: Optional[Any] = None
    tuesday_member_close: Optional[Any] = None
    tuesday_member_open: Optional[Any] = None
    tuesday_public_close: Optional[Any] = None
    tuesday_public_open: Optional[Any] = None
    updated_at: Optional[Any] = None
    wednesday_is_closed: Optional[Any] = None
    wednesday_member_close: Optional[Any] = None
    wednesday_member_open: Optional[Any] = None
    wednesday_public_close: Optional[Any] = None
    wednesday_public_open: Optional[Any] = None


@dataclass
class HourLoadMatch:
    id: str


@dataclass
class HourListMatch:
    additional_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    friday_is_closed: Optional[Any] = None
    friday_member_close: Optional[Any] = None
    friday_member_open: Optional[Any] = None
    friday_public_close: Optional[Any] = None
    friday_public_open: Optional[Any] = None
    id: Optional[str] = None
    monday_is_closed: Optional[Any] = None
    monday_member_close: Optional[Any] = None
    monday_member_open: Optional[Any] = None
    monday_public_close: Optional[Any] = None
    monday_public_open: Optional[Any] = None
    saturday_is_closed: Optional[Any] = None
    saturday_member_close: Optional[Any] = None
    saturday_member_open: Optional[Any] = None
    saturday_public_close: Optional[Any] = None
    saturday_public_open: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    summary: Optional[Any] = None
    sunday_is_closed: Optional[Any] = None
    sunday_member_close: Optional[Any] = None
    sunday_member_open: Optional[Any] = None
    sunday_public_close: Optional[Any] = None
    sunday_public_open: Optional[Any] = None
    thursday_is_closed: Optional[Any] = None
    thursday_member_close: Optional[Any] = None
    thursday_member_open: Optional[Any] = None
    thursday_public_close: Optional[Any] = None
    thursday_public_open: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    tuesday_is_closed: Optional[Any] = None
    tuesday_member_close: Optional[Any] = None
    tuesday_member_open: Optional[Any] = None
    tuesday_public_close: Optional[Any] = None
    tuesday_public_open: Optional[Any] = None
    updated_at: Optional[Any] = None
    wednesday_is_closed: Optional[Any] = None
    wednesday_member_close: Optional[Any] = None
    wednesday_member_open: Optional[Any] = None
    wednesday_public_close: Optional[Any] = None
    wednesday_public_open: Optional[Any] = None


@dataclass
class Image:
    ahash: Optional[Any] = None
    alt_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    color: Optional[Any] = None
    colorfulness: Optional[Any] = None
    content: Optional[Any] = None
    content_e_tag: Optional[Any] = None
    credit_line: Optional[Any] = None
    fingerprint: Optional[Any] = None
    height: Optional[float] = None
    id: Optional[str] = None
    iiif_url: Optional[Any] = None
    is_educational_resource: Optional[bool] = None
    is_multimedia_resource: Optional[bool] = None
    is_teacher_resource: Optional[bool] = None
    lake_guid: Optional[Any] = None
    lqip: Optional[Any] = None
    phash: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    type: Optional[Any] = None
    updated_at: Optional[Any] = None
    width: Optional[float] = None


@dataclass
class ImageLoadMatch:
    id: str


@dataclass
class ImageListMatch:
    ahash: Optional[Any] = None
    alt_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    color: Optional[Any] = None
    colorfulness: Optional[Any] = None
    content: Optional[Any] = None
    content_e_tag: Optional[Any] = None
    credit_line: Optional[Any] = None
    fingerprint: Optional[Any] = None
    height: Optional[float] = None
    id: Optional[str] = None
    iiif_url: Optional[Any] = None
    is_educational_resource: Optional[bool] = None
    is_multimedia_resource: Optional[bool] = None
    is_teacher_resource: Optional[bool] = None
    lake_guid: Optional[Any] = None
    lqip: Optional[Any] = None
    phash: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    type: Optional[Any] = None
    updated_at: Optional[Any] = None
    width: Optional[float] = None


@dataclass
class LandingPage:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    search_tag: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class LandingPageLoadMatch:
    id: str


@dataclass
class LandingPageListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    search_tag: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Place:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    latitude: Optional[float] = None
    longitude: Optional[float] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    tgn_id: Optional[str] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class PlaceLoadMatch:
    id: str


@dataclass
class PlaceListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    latitude: Optional[float] = None
    longitude: Optional[float] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    tgn_id: Optional[str] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None


@dataclass
class PressRelease:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class PressReleaseLoadMatch:
    id: str


@dataclass
class PressReleaseListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class PrintedPublication:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class PrintedPublicationLoadMatch:
    id: str


@dataclass
class PrintedPublicationListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    copy: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Product:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artist_id: Optional[str] = None
    artwork_id: Optional[str] = None
    description: Optional[str] = None
    exhibition_id: Optional[str] = None
    external_sku: Optional[Any] = None
    id: Optional[str] = None
    image_url: Optional[Any] = None
    max_compare_at_price: Optional[Any] = None
    max_current_price: Optional[Any] = None
    min_compare_at_price: Optional[Any] = None
    min_current_price: Optional[Any] = None
    price_display: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class ProductLoadMatch:
    id: str


@dataclass
class ProductListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artist_id: Optional[str] = None
    artwork_id: Optional[str] = None
    description: Optional[str] = None
    exhibition_id: Optional[str] = None
    external_sku: Optional[Any] = None
    id: Optional[str] = None
    image_url: Optional[Any] = None
    max_compare_at_price: Optional[Any] = None
    max_current_price: Optional[Any] = None
    min_compare_at_price: Optional[Any] = None
    min_current_price: Optional[Any] = None
    price_display: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Publication:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    section_id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class PublicationLoadMatch:
    id: str


@dataclass
class PublicationListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    section_id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Search:
    api_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    is_boosted: Optional[bool] = None
    score: Optional[float] = None
    thumbnail: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None


@dataclass
class SearchListMatch:
    api_id: Optional[str] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    is_boosted: Optional[bool] = None
    score: Optional[float] = None
    thumbnail: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None


@dataclass
class Section:
    accession: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    content: Optional[Any] = None
    generic_page_id: Optional[str] = None
    id: Optional[str] = None
    publication_id: Optional[str] = None
    publication_title: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class SectionLoadMatch:
    id: str


@dataclass
class SectionListMatch:
    accession: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    content: Optional[Any] = None
    generic_page_id: Optional[str] = None
    id: Optional[str] = None
    publication_id: Optional[str] = None
    publication_title: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Site:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    description: Optional[str] = None
    exhibition_id: Optional[str] = None
    exhibition_title: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class SiteLoadMatch:
    id: str


@dataclass
class SiteListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    description: Optional[str] = None
    exhibition_id: Optional[str] = None
    exhibition_title: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Sound:
    alt_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    content: Optional[Any] = None
    content_e_tag: Optional[Any] = None
    credit_line: Optional[Any] = None
    id: Optional[str] = None
    is_educational_resource: Optional[bool] = None
    is_multimedia_resource: Optional[bool] = None
    is_teacher_resource: Optional[bool] = None
    lake_guid: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    transcript: Optional[Any] = None
    type: Optional[Any] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class SoundLoadMatch:
    id: str


@dataclass
class SoundListMatch:
    alt_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    content: Optional[Any] = None
    content_e_tag: Optional[Any] = None
    credit_line: Optional[Any] = None
    id: Optional[str] = None
    is_educational_resource: Optional[bool] = None
    is_multimedia_resource: Optional[bool] = None
    is_teacher_resource: Optional[bool] = None
    lake_guid: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    transcript: Optional[Any] = None
    type: Optional[Any] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class StaticPage:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class StaticPageLoadMatch:
    id: str


@dataclass
class StaticPageListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    id: Optional[str] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    web_url: Optional[Any] = None


@dataclass
class Text:
    alt_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    content: Optional[Any] = None
    content_e_tag: Optional[Any] = None
    credit_line: Optional[Any] = None
    id: Optional[str] = None
    is_educational_resource: Optional[bool] = None
    is_multimedia_resource: Optional[bool] = None
    is_teacher_resource: Optional[bool] = None
    lake_guid: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    type: Optional[Any] = None
    updated_at: Optional[Any] = None


@dataclass
class TextLoadMatch:
    id: str


@dataclass
class TextListMatch:
    alt_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    content: Optional[Any] = None
    content_e_tag: Optional[Any] = None
    credit_line: Optional[Any] = None
    id: Optional[str] = None
    is_educational_resource: Optional[bool] = None
    is_multimedia_resource: Optional[bool] = None
    is_teacher_resource: Optional[bool] = None
    lake_guid: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    type: Optional[Any] = None
    updated_at: Optional[Any] = None


@dataclass
class Tour:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artist_title: Optional[Any] = None
    artwork_title: Optional[Any] = None
    description: Optional[str] = None
    id: Optional[str] = None
    image: Optional[Any] = None
    intro: Optional[Any] = None
    intro_link: Optional[Any] = None
    intro_transcript: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    weight: Optional[float] = None


@dataclass
class TourLoadMatch:
    id: str


@dataclass
class TourListMatch:
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artist_title: Optional[Any] = None
    artwork_title: Optional[Any] = None
    description: Optional[str] = None
    id: Optional[str] = None
    image: Optional[Any] = None
    intro: Optional[Any] = None
    intro_link: Optional[Any] = None
    intro_transcript: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    updated_at: Optional[Any] = None
    weight: Optional[float] = None


@dataclass
class Video:
    alt_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    content: Optional[Any] = None
    content_e_tag: Optional[Any] = None
    credit_line: Optional[Any] = None
    id: Optional[str] = None
    is_educational_resource: Optional[bool] = None
    is_multimedia_resource: Optional[bool] = None
    is_teacher_resource: Optional[bool] = None
    lake_guid: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    type: Optional[Any] = None
    updated_at: Optional[Any] = None


@dataclass
class VideoLoadMatch:
    id: str


@dataclass
class VideoListMatch:
    alt_text: Optional[Any] = None
    api_link: Optional[Any] = None
    api_model: Optional[Any] = None
    artwork_id: Optional[str] = None
    artwork_title: Optional[Any] = None
    content: Optional[Any] = None
    content_e_tag: Optional[Any] = None
    credit_line: Optional[Any] = None
    id: Optional[str] = None
    is_educational_resource: Optional[bool] = None
    is_multimedia_resource: Optional[bool] = None
    is_teacher_resource: Optional[bool] = None
    lake_guid: Optional[Any] = None
    source_updated_at: Optional[Any] = None
    suggest_autocomplete_all: Optional[Any] = None
    suggest_autocomplete_boosted: Optional[Any] = None
    timestamp: Optional[Any] = None
    title: Optional[str] = None
    type: Optional[Any] = None
    updated_at: Optional[Any] = None

