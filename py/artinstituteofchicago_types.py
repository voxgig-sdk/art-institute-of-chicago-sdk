# Typed models for the ArtInstituteOfChicago SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Agent(TypedDict, total=False):
    alt_title: Any
    api_link: Any
    api_model: Any
    birth_date: Any
    death_date: Any
    description: str
    id: str
    is_artist: bool
    sort_title: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    ulan_id: str
    updated_at: Any


class AgentLoadMatch(TypedDict):
    id: str


class AgentListMatch(TypedDict, total=False):
    alt_title: Any
    api_link: Any
    api_model: Any
    birth_date: Any
    death_date: Any
    description: str
    id: str
    is_artist: bool
    sort_title: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    ulan_id: str
    updated_at: Any


class AgentRole(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class AgentRoleLoadMatch(TypedDict):
    id: str


class AgentRoleListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class AgentType(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class AgentTypeLoadMatch(TypedDict):
    id: str


class AgentTypeListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class Article(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class ArticleLoadMatch(TypedDict):
    id: str


class ArticleListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class Artwork(TypedDict, total=False):
    alt_artist_id: str
    alt_classification_id: str
    alt_image_id: str
    alt_material_id: str
    alt_style_id: str
    alt_subject_id: str
    alt_technique_id: str
    alt_title: Any
    api_link: Any
    api_model: Any
    artist_display: Any
    artist_id: str
    artist_title: Any
    artwork_type_id: str
    artwork_type_title: Any
    boost_rank: Any
    catalog_based_search_keyword_title: Any
    catalogue_display: Any
    category_id: str
    category_title: Any
    classification_id: str
    classification_title: Any
    color: Any
    colorfulness: Any
    copyright_notice: Any
    credit_line: Any
    date_display: Any
    date_end: Any
    date_qualifier_id: str
    date_qualifier_title: Any
    date_start: Any
    department_id: str
    department_title: Any
    description: str
    dimension: Any
    dimensions_detail: Any
    document_id: str
    edition: Any
    exhibition_history: Any
    fiscal_year: Any
    fiscal_year_deaccession: Any
    gallery_id: str
    gallery_title: Any
    has_advanced_imaging: bool
    has_educational_resource: bool
    has_multimedia_resource: bool
    has_not_been_viewed_much: bool
    id: str
    image_embedding: Any
    image_id: str
    inscription: Any
    internal_department_id: str
    is_boosted: bool
    is_on_view: bool
    is_public_domain: bool
    is_zoomable: bool
    latitude: float
    latlon: Any
    longitude: float
    main_reference_number: int
    material_id: str
    material_title: Any
    max_zoom_window_size: Any
    medium_display: Any
    nomisma_id: str
    on_loan_display: Any
    pageview: Any
    pageviews_recent: Any
    place_of_origin: Any
    provenance_text: Any
    publication_history: Any
    publishing_verification_level: Any
    section_id: str
    section_title: Any
    short_description: Any
    site_id: str
    sound_id: str
    source_updated_at: Any
    style_id: str
    style_title: Any
    subject_id: str
    subject_title: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    technique_id: str
    technique_title: Any
    term_title: Any
    text_embedding: Any
    text_id: str
    theme_title: Any
    thumbnail: Any
    timestamp: Any
    title: str
    updated_at: Any
    video_id: str


class ArtworkLoadMatch(TypedDict):
    id: str


class ArtworkListMatch(TypedDict, total=False):
    alt_artist_id: str
    alt_classification_id: str
    alt_image_id: str
    alt_material_id: str
    alt_style_id: str
    alt_subject_id: str
    alt_technique_id: str
    alt_title: Any
    api_link: Any
    api_model: Any
    artist_display: Any
    artist_id: str
    artist_title: Any
    artwork_type_id: str
    artwork_type_title: Any
    boost_rank: Any
    catalog_based_search_keyword_title: Any
    catalogue_display: Any
    category_id: str
    category_title: Any
    classification_id: str
    classification_title: Any
    color: Any
    colorfulness: Any
    copyright_notice: Any
    credit_line: Any
    date_display: Any
    date_end: Any
    date_qualifier_id: str
    date_qualifier_title: Any
    date_start: Any
    department_id: str
    department_title: Any
    description: str
    dimension: Any
    dimensions_detail: Any
    document_id: str
    edition: Any
    exhibition_history: Any
    fiscal_year: Any
    fiscal_year_deaccession: Any
    gallery_id: str
    gallery_title: Any
    has_advanced_imaging: bool
    has_educational_resource: bool
    has_multimedia_resource: bool
    has_not_been_viewed_much: bool
    id: str
    image_embedding: Any
    image_id: str
    inscription: Any
    internal_department_id: str
    is_boosted: bool
    is_on_view: bool
    is_public_domain: bool
    is_zoomable: bool
    latitude: float
    latlon: Any
    longitude: float
    main_reference_number: int
    material_id: str
    material_title: Any
    max_zoom_window_size: Any
    medium_display: Any
    nomisma_id: str
    on_loan_display: Any
    pageview: Any
    pageviews_recent: Any
    place_of_origin: Any
    provenance_text: Any
    publication_history: Any
    publishing_verification_level: Any
    section_id: str
    section_title: Any
    short_description: Any
    site_id: str
    sound_id: str
    source_updated_at: Any
    style_id: str
    style_title: Any
    subject_id: str
    subject_title: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    technique_id: str
    technique_title: Any
    term_title: Any
    text_embedding: Any
    text_id: str
    theme_title: Any
    thumbnail: Any
    timestamp: Any
    title: str
    updated_at: Any
    video_id: str


class ArtworkDateQualifier(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class ArtworkDateQualifierLoadMatch(TypedDict):
    id: str


class ArtworkDateQualifierListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class ArtworkPlaceQualifier(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class ArtworkPlaceQualifierLoadMatch(TypedDict):
    id: str


class ArtworkPlaceQualifierListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class ArtworkType(TypedDict, total=False):
    aat_id: str
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class ArtworkTypeLoadMatch(TypedDict):
    id: str


class ArtworkTypeListMatch(TypedDict, total=False):
    aat_id: str
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class CategoryTerm(TypedDict, total=False):
    aat_id: str
    api_link: Any
    api_model: Any
    id: str
    parent_id: str
    source_updated_at: Any
    subtype: Any
    suggest_autocomplete_all: Any
    timestamp: Any
    title: str
    updated_at: Any


class CategoryTermLoadMatch(TypedDict):
    id: str


class CategoryTermListMatch(TypedDict, total=False):
    aat_id: str
    api_link: Any
    api_model: Any
    id: str
    parent_id: str
    source_updated_at: Any
    subtype: Any
    suggest_autocomplete_all: Any
    timestamp: Any
    title: str
    updated_at: Any


class DigitalPublication(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class DigitalPublicationLoadMatch(TypedDict):
    id: str


class DigitalPublicationListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class DigitalPublicationArticle(TypedDict, total=False):
    api_link: Any
    api_model: Any
    author_display: Any
    copy: Any
    digital_publication_id: str
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class DigitalPublicationArticleLoadMatch(TypedDict):
    id: str


class DigitalPublicationArticleListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    author_display: Any
    copy: Any
    digital_publication_id: str
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class EducatorResource(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class EducatorResourceLoadMatch(TypedDict):
    id: str


class EducatorResourceListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Event(TypedDict, total=False):
    alt_audience_id: str
    alt_event_type_id: str
    api_link: Any
    api_model: Any
    audience_id: str
    buy_button_caption: Any
    buy_button_text: Any
    date_display: Any
    description: str
    door_time: Any
    end_date: Any
    end_time: Any
    entrance: Any
    event_host_id: str
    event_host_title: Any
    event_type_id: str
    header_description: Any
    hero_caption: Any
    id: str
    image_url: Any
    is_admission_required: bool
    is_after_hour: bool
    is_free: bool
    is_member_exclusive: bool
    is_private: bool
    is_registration_required: bool
    is_sales_button_hidden: bool
    is_sold_out: bool
    is_ticketed: bool
    is_virtual_event: bool
    join_url: Any
    layout_type: Any
    list_description: Any
    location: Any
    program_id: str
    program_title: Any
    rsvp_link: Any
    search_tag: Any
    short_description: Any
    slug: str
    source_updated_at: Any
    start_date: Any
    start_time: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    survey_url: Any
    ticketed_event_id: str
    timestamp: Any
    title: str
    title_display: Any
    updated_at: Any
    virtual_event_passcode: Any
    virtual_event_url: Any


class EventLoadMatch(TypedDict):
    id: str


class EventListMatch(TypedDict, total=False):
    alt_audience_id: str
    alt_event_type_id: str
    api_link: Any
    api_model: Any
    audience_id: str
    buy_button_caption: Any
    buy_button_text: Any
    date_display: Any
    description: str
    door_time: Any
    end_date: Any
    end_time: Any
    entrance: Any
    event_host_id: str
    event_host_title: Any
    event_type_id: str
    header_description: Any
    hero_caption: Any
    id: str
    image_url: Any
    is_admission_required: bool
    is_after_hour: bool
    is_free: bool
    is_member_exclusive: bool
    is_private: bool
    is_registration_required: bool
    is_sales_button_hidden: bool
    is_sold_out: bool
    is_ticketed: bool
    is_virtual_event: bool
    join_url: Any
    layout_type: Any
    list_description: Any
    location: Any
    program_id: str
    program_title: Any
    rsvp_link: Any
    search_tag: Any
    short_description: Any
    slug: str
    source_updated_at: Any
    start_date: Any
    start_time: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    survey_url: Any
    ticketed_event_id: str
    timestamp: Any
    title: str
    title_display: Any
    updated_at: Any
    virtual_event_passcode: Any
    virtual_event_url: Any


class EventOccurrence(TypedDict, total=False):
    api_link: Any
    api_model: Any
    button_caption: Any
    button_text: Any
    button_url: Any
    description: str
    end_at: Any
    event_id: str
    id: str
    image_url: Any
    is_private: bool
    is_sales_button_hidden: bool
    is_ticketed: bool
    location: Any
    off_sale_at: Any
    on_sale_at: Any
    short_description: Any
    source_updated_at: Any
    start_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    title_display: Any
    updated_at: Any


class EventOccurrenceLoadMatch(TypedDict):
    id: str


class EventOccurrenceListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    button_caption: Any
    button_text: Any
    button_url: Any
    description: str
    end_at: Any
    event_id: str
    id: str
    image_url: Any
    is_private: bool
    is_sales_button_hidden: bool
    is_ticketed: bool
    location: Any
    off_sale_at: Any
    on_sale_at: Any
    short_description: Any
    source_updated_at: Any
    start_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    title_display: Any
    updated_at: Any


class EventProgram(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    is_affiliate_group: bool
    is_event_host: bool
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class EventProgramLoadMatch(TypedDict):
    id: str


class EventProgramListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    is_affiliate_group: bool
    is_event_host: bool
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class Exhibition(TypedDict, total=False):
    aic_end_at: Any
    aic_start_at: Any
    alt_image_id: str
    api_link: Any
    api_model: Any
    artist_id: str
    artwork_id: str
    artwork_title: Any
    document_id: str
    gallery_id: str
    gallery_title: Any
    id: str
    image_id: str
    image_url: Any
    is_featured: bool
    is_published: bool
    position: Any
    short_description: Any
    site_id: str
    source_updated_at: Any
    status: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class ExhibitionLoadMatch(TypedDict):
    id: str


class ExhibitionListMatch(TypedDict, total=False):
    aic_end_at: Any
    aic_start_at: Any
    alt_image_id: str
    api_link: Any
    api_model: Any
    artist_id: str
    artwork_id: str
    artwork_title: Any
    document_id: str
    gallery_id: str
    gallery_title: Any
    id: str
    image_id: str
    image_url: Any
    is_featured: bool
    is_published: bool
    position: Any
    short_description: Any
    site_id: str
    source_updated_at: Any
    status: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Gallery(TypedDict, total=False):
    api_link: Any
    api_model: Any
    floor: Any
    id: str
    is_closed: bool
    latitude: float
    latlon: Any
    longitude: float
    number: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    tgn_id: str
    timestamp: Any
    title: str
    updated_at: Any


class GalleryLoadMatch(TypedDict):
    id: str


class GalleryListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    floor: Any
    id: str
    is_closed: bool
    latitude: float
    latlon: Any
    longitude: float
    number: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    tgn_id: str
    timestamp: Any
    title: str
    updated_at: Any


class GenericPage(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    search_tag: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class GenericPageLoadMatch(TypedDict):
    id: str


class GenericPageListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    search_tag: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Highlight(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class HighlightLoadMatch(TypedDict):
    id: str


class HighlightListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any


class Hour(TypedDict, total=False):
    additional_text: Any
    api_link: Any
    api_model: Any
    friday_is_closed: Any
    friday_member_close: Any
    friday_member_open: Any
    friday_public_close: Any
    friday_public_open: Any
    id: str
    monday_is_closed: Any
    monday_member_close: Any
    monday_member_open: Any
    monday_public_close: Any
    monday_public_open: Any
    saturday_is_closed: Any
    saturday_member_close: Any
    saturday_member_open: Any
    saturday_public_close: Any
    saturday_public_open: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    summary: Any
    sunday_is_closed: Any
    sunday_member_close: Any
    sunday_member_open: Any
    sunday_public_close: Any
    sunday_public_open: Any
    thursday_is_closed: Any
    thursday_member_close: Any
    thursday_member_open: Any
    thursday_public_close: Any
    thursday_public_open: Any
    timestamp: Any
    title: str
    tuesday_is_closed: Any
    tuesday_member_close: Any
    tuesday_member_open: Any
    tuesday_public_close: Any
    tuesday_public_open: Any
    updated_at: Any
    wednesday_is_closed: Any
    wednesday_member_close: Any
    wednesday_member_open: Any
    wednesday_public_close: Any
    wednesday_public_open: Any


class HourLoadMatch(TypedDict):
    id: str


class HourListMatch(TypedDict, total=False):
    additional_text: Any
    api_link: Any
    api_model: Any
    friday_is_closed: Any
    friday_member_close: Any
    friday_member_open: Any
    friday_public_close: Any
    friday_public_open: Any
    id: str
    monday_is_closed: Any
    monday_member_close: Any
    monday_member_open: Any
    monday_public_close: Any
    monday_public_open: Any
    saturday_is_closed: Any
    saturday_member_close: Any
    saturday_member_open: Any
    saturday_public_close: Any
    saturday_public_open: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    summary: Any
    sunday_is_closed: Any
    sunday_member_close: Any
    sunday_member_open: Any
    sunday_public_close: Any
    sunday_public_open: Any
    thursday_is_closed: Any
    thursday_member_close: Any
    thursday_member_open: Any
    thursday_public_close: Any
    thursday_public_open: Any
    timestamp: Any
    title: str
    tuesday_is_closed: Any
    tuesday_member_close: Any
    tuesday_member_open: Any
    tuesday_public_close: Any
    tuesday_public_open: Any
    updated_at: Any
    wednesday_is_closed: Any
    wednesday_member_close: Any
    wednesday_member_open: Any
    wednesday_public_close: Any
    wednesday_public_open: Any


class Image(TypedDict, total=False):
    ahash: Any
    alt_text: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    color: Any
    colorfulness: Any
    content: Any
    content_e_tag: Any
    credit_line: Any
    fingerprint: Any
    height: float
    id: str
    iiif_url: Any
    is_educational_resource: bool
    is_multimedia_resource: bool
    is_teacher_resource: bool
    lake_guid: Any
    lqip: Any
    phash: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    type: Any
    updated_at: Any
    width: float


class ImageLoadMatch(TypedDict):
    id: str


class ImageListMatch(TypedDict, total=False):
    ahash: Any
    alt_text: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    color: Any
    colorfulness: Any
    content: Any
    content_e_tag: Any
    credit_line: Any
    fingerprint: Any
    height: float
    id: str
    iiif_url: Any
    is_educational_resource: bool
    is_multimedia_resource: bool
    is_teacher_resource: bool
    lake_guid: Any
    lqip: Any
    phash: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    type: Any
    updated_at: Any
    width: float


class LandingPage(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    search_tag: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class LandingPageLoadMatch(TypedDict):
    id: str


class LandingPageListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    search_tag: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Place(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    latitude: float
    longitude: float
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    tgn_id: str
    timestamp: Any
    title: str
    updated_at: Any


class PlaceLoadMatch(TypedDict):
    id: str


class PlaceListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    latitude: float
    longitude: float
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    tgn_id: str
    timestamp: Any
    title: str
    updated_at: Any


class PressRelease(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class PressReleaseLoadMatch(TypedDict):
    id: str


class PressReleaseListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class PrintedPublication(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class PrintedPublicationLoadMatch(TypedDict):
    id: str


class PrintedPublicationListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    copy: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Product(TypedDict, total=False):
    api_link: Any
    api_model: Any
    artist_id: str
    artwork_id: str
    description: str
    exhibition_id: str
    external_sku: Any
    id: str
    image_url: Any
    max_compare_at_price: Any
    max_current_price: Any
    min_compare_at_price: Any
    min_current_price: Any
    price_display: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class ProductLoadMatch(TypedDict):
    id: str


class ProductListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    artist_id: str
    artwork_id: str
    description: str
    exhibition_id: str
    external_sku: Any
    id: str
    image_url: Any
    max_compare_at_price: Any
    max_current_price: Any
    min_compare_at_price: Any
    min_current_price: Any
    price_display: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Publication(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    section_id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class PublicationLoadMatch(TypedDict):
    id: str


class PublicationListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    section_id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Search(TypedDict, total=False):
    api_id: str
    api_link: Any
    api_model: Any
    id: str
    is_boosted: bool
    score: float
    thumbnail: Any
    timestamp: Any
    title: str


class SearchListMatch(TypedDict, total=False):
    api_id: str
    api_link: Any
    api_model: Any
    id: str
    is_boosted: bool
    score: float
    thumbnail: Any
    timestamp: Any
    title: str


class Section(TypedDict, total=False):
    accession: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    content: Any
    generic_page_id: str
    id: str
    publication_id: str
    publication_title: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class SectionLoadMatch(TypedDict):
    id: str


class SectionListMatch(TypedDict, total=False):
    accession: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    content: Any
    generic_page_id: str
    id: str
    publication_id: str
    publication_title: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Site(TypedDict, total=False):
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    description: str
    exhibition_id: str
    exhibition_title: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class SiteLoadMatch(TypedDict):
    id: str


class SiteListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    description: str
    exhibition_id: str
    exhibition_title: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Sound(TypedDict, total=False):
    alt_text: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    content: Any
    content_e_tag: Any
    credit_line: Any
    id: str
    is_educational_resource: bool
    is_multimedia_resource: bool
    is_teacher_resource: bool
    lake_guid: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    transcript: Any
    type: Any
    updated_at: Any
    web_url: Any


class SoundLoadMatch(TypedDict):
    id: str


class SoundListMatch(TypedDict, total=False):
    alt_text: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    content: Any
    content_e_tag: Any
    credit_line: Any
    id: str
    is_educational_resource: bool
    is_multimedia_resource: bool
    is_teacher_resource: bool
    lake_guid: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    transcript: Any
    type: Any
    updated_at: Any
    web_url: Any


class StaticPage(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class StaticPageLoadMatch(TypedDict):
    id: str


class StaticPageListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    id: str
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    web_url: Any


class Text(TypedDict, total=False):
    alt_text: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    content: Any
    content_e_tag: Any
    credit_line: Any
    id: str
    is_educational_resource: bool
    is_multimedia_resource: bool
    is_teacher_resource: bool
    lake_guid: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    type: Any
    updated_at: Any


class TextLoadMatch(TypedDict):
    id: str


class TextListMatch(TypedDict, total=False):
    alt_text: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    content: Any
    content_e_tag: Any
    credit_line: Any
    id: str
    is_educational_resource: bool
    is_multimedia_resource: bool
    is_teacher_resource: bool
    lake_guid: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    type: Any
    updated_at: Any


class Tour(TypedDict, total=False):
    api_link: Any
    api_model: Any
    artist_title: Any
    artwork_title: Any
    description: str
    id: str
    image: Any
    intro: Any
    intro_link: Any
    intro_transcript: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    weight: float


class TourLoadMatch(TypedDict):
    id: str


class TourListMatch(TypedDict, total=False):
    api_link: Any
    api_model: Any
    artist_title: Any
    artwork_title: Any
    description: str
    id: str
    image: Any
    intro: Any
    intro_link: Any
    intro_transcript: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    updated_at: Any
    weight: float


class Video(TypedDict, total=False):
    alt_text: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    content: Any
    content_e_tag: Any
    credit_line: Any
    id: str
    is_educational_resource: bool
    is_multimedia_resource: bool
    is_teacher_resource: bool
    lake_guid: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    type: Any
    updated_at: Any


class VideoLoadMatch(TypedDict):
    id: str


class VideoListMatch(TypedDict, total=False):
    alt_text: Any
    api_link: Any
    api_model: Any
    artwork_id: str
    artwork_title: Any
    content: Any
    content_e_tag: Any
    credit_line: Any
    id: str
    is_educational_resource: bool
    is_multimedia_resource: bool
    is_teacher_resource: bool
    lake_guid: Any
    source_updated_at: Any
    suggest_autocomplete_all: Any
    suggest_autocomplete_boosted: Any
    timestamp: Any
    title: str
    type: Any
    updated_at: Any
