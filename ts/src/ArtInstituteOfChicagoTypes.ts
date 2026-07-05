// Typed models for the ArtInstituteOfChicago SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Agent {
  alt_title?: any
  api_link?: any
  api_model?: any
  birth_date?: any
  death_date?: any
  description?: string
  id?: string
  is_artist?: boolean
  sort_title?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  ulan_id?: string
  updated_at?: any
}

export interface AgentLoadMatch {
  id: string
}

export interface AgentListMatch {
  alt_title?: any
  api_link?: any
  api_model?: any
  birth_date?: any
  death_date?: any
  description?: string
  id?: string
  is_artist?: boolean
  sort_title?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  ulan_id?: string
  updated_at?: any
}

export interface AgentRole {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface AgentRoleLoadMatch {
  id: string
}

export interface AgentRoleListMatch {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface AgentType {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface AgentTypeLoadMatch {
  id: string
}

export interface AgentTypeListMatch {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface Article {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface ArticleLoadMatch {
  id: string
}

export interface ArticleListMatch {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface Artwork {
  alt_artist_id?: string
  alt_classification_id?: string
  alt_image_id?: string
  alt_material_id?: string
  alt_style_id?: string
  alt_subject_id?: string
  alt_technique_id?: string
  alt_title?: any
  api_link?: any
  api_model?: any
  artist_display?: any
  artist_id?: string
  artist_title?: any
  artwork_type_id?: string
  artwork_type_title?: any
  boost_rank?: any
  catalog_based_search_keyword_title?: any
  catalogue_display?: any
  category_id?: string
  category_title?: any
  classification_id?: string
  classification_title?: any
  color?: any
  colorfulness?: any
  copyright_notice?: any
  credit_line?: any
  date_display?: any
  date_end?: any
  date_qualifier_id?: string
  date_qualifier_title?: any
  date_start?: any
  department_id?: string
  department_title?: any
  description?: string
  dimension?: any
  dimensions_detail?: any
  document_id?: string
  edition?: any
  exhibition_history?: any
  fiscal_year?: any
  fiscal_year_deaccession?: any
  gallery_id?: string
  gallery_title?: any
  has_advanced_imaging?: boolean
  has_educational_resource?: boolean
  has_multimedia_resource?: boolean
  has_not_been_viewed_much?: boolean
  id?: string
  image_embedding?: any
  image_id?: string
  inscription?: any
  internal_department_id?: string
  is_boosted?: boolean
  is_on_view?: boolean
  is_public_domain?: boolean
  is_zoomable?: boolean
  latitude?: number
  latlon?: any
  longitude?: number
  main_reference_number?: number
  material_id?: string
  material_title?: any
  max_zoom_window_size?: any
  medium_display?: any
  nomisma_id?: string
  on_loan_display?: any
  pageview?: any
  pageviews_recent?: any
  place_of_origin?: any
  provenance_text?: any
  publication_history?: any
  publishing_verification_level?: any
  section_id?: string
  section_title?: any
  short_description?: any
  site_id?: string
  sound_id?: string
  source_updated_at?: any
  style_id?: string
  style_title?: any
  subject_id?: string
  subject_title?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  technique_id?: string
  technique_title?: any
  term_title?: any
  text_embedding?: any
  text_id?: string
  theme_title?: any
  thumbnail?: any
  timestamp?: any
  title?: string
  updated_at?: any
  video_id?: string
}

export interface ArtworkLoadMatch {
  id: string
}

export interface ArtworkListMatch {
  alt_artist_id?: string
  alt_classification_id?: string
  alt_image_id?: string
  alt_material_id?: string
  alt_style_id?: string
  alt_subject_id?: string
  alt_technique_id?: string
  alt_title?: any
  api_link?: any
  api_model?: any
  artist_display?: any
  artist_id?: string
  artist_title?: any
  artwork_type_id?: string
  artwork_type_title?: any
  boost_rank?: any
  catalog_based_search_keyword_title?: any
  catalogue_display?: any
  category_id?: string
  category_title?: any
  classification_id?: string
  classification_title?: any
  color?: any
  colorfulness?: any
  copyright_notice?: any
  credit_line?: any
  date_display?: any
  date_end?: any
  date_qualifier_id?: string
  date_qualifier_title?: any
  date_start?: any
  department_id?: string
  department_title?: any
  description?: string
  dimension?: any
  dimensions_detail?: any
  document_id?: string
  edition?: any
  exhibition_history?: any
  fiscal_year?: any
  fiscal_year_deaccession?: any
  gallery_id?: string
  gallery_title?: any
  has_advanced_imaging?: boolean
  has_educational_resource?: boolean
  has_multimedia_resource?: boolean
  has_not_been_viewed_much?: boolean
  id?: string
  image_embedding?: any
  image_id?: string
  inscription?: any
  internal_department_id?: string
  is_boosted?: boolean
  is_on_view?: boolean
  is_public_domain?: boolean
  is_zoomable?: boolean
  latitude?: number
  latlon?: any
  longitude?: number
  main_reference_number?: number
  material_id?: string
  material_title?: any
  max_zoom_window_size?: any
  medium_display?: any
  nomisma_id?: string
  on_loan_display?: any
  pageview?: any
  pageviews_recent?: any
  place_of_origin?: any
  provenance_text?: any
  publication_history?: any
  publishing_verification_level?: any
  section_id?: string
  section_title?: any
  short_description?: any
  site_id?: string
  sound_id?: string
  source_updated_at?: any
  style_id?: string
  style_title?: any
  subject_id?: string
  subject_title?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  technique_id?: string
  technique_title?: any
  term_title?: any
  text_embedding?: any
  text_id?: string
  theme_title?: any
  thumbnail?: any
  timestamp?: any
  title?: string
  updated_at?: any
  video_id?: string
}

export interface ArtworkDateQualifier {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface ArtworkDateQualifierLoadMatch {
  id: string
}

export interface ArtworkDateQualifierListMatch {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface ArtworkPlaceQualifier {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface ArtworkPlaceQualifierLoadMatch {
  id: string
}

export interface ArtworkPlaceQualifierListMatch {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface ArtworkType {
  aat_id?: string
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface ArtworkTypeLoadMatch {
  id: string
}

export interface ArtworkTypeListMatch {
  aat_id?: string
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface CategoryTerm {
  aat_id?: string
  api_link?: any
  api_model?: any
  id?: string
  parent_id?: string
  source_updated_at?: any
  subtype?: any
  suggest_autocomplete_all?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface CategoryTermLoadMatch {
  id: string
}

export interface CategoryTermListMatch {
  aat_id?: string
  api_link?: any
  api_model?: any
  id?: string
  parent_id?: string
  source_updated_at?: any
  subtype?: any
  suggest_autocomplete_all?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface DigitalPublication {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface DigitalPublicationLoadMatch {
  id: string
}

export interface DigitalPublicationListMatch {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface DigitalPublicationArticle {
  api_link?: any
  api_model?: any
  author_display?: any
  copy?: any
  digital_publication_id?: string
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface DigitalPublicationArticleLoadMatch {
  id: string
}

export interface DigitalPublicationArticleListMatch {
  api_link?: any
  api_model?: any
  author_display?: any
  copy?: any
  digital_publication_id?: string
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface EducatorResource {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface EducatorResourceLoadMatch {
  id: string
}

export interface EducatorResourceListMatch {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Event {
  alt_audience_id?: string
  alt_event_type_id?: string
  api_link?: any
  api_model?: any
  audience_id?: string
  buy_button_caption?: any
  buy_button_text?: any
  date_display?: any
  description?: string
  door_time?: any
  end_date?: any
  end_time?: any
  entrance?: any
  event_host_id?: string
  event_host_title?: any
  event_type_id?: string
  header_description?: any
  hero_caption?: any
  id?: string
  image_url?: any
  is_admission_required?: boolean
  is_after_hour?: boolean
  is_free?: boolean
  is_member_exclusive?: boolean
  is_private?: boolean
  is_registration_required?: boolean
  is_sales_button_hidden?: boolean
  is_sold_out?: boolean
  is_ticketed?: boolean
  is_virtual_event?: boolean
  join_url?: any
  layout_type?: any
  list_description?: any
  location?: any
  program_id?: string
  program_title?: any
  rsvp_link?: any
  search_tag?: any
  short_description?: any
  slug?: string
  source_updated_at?: any
  start_date?: any
  start_time?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  survey_url?: any
  ticketed_event_id?: string
  timestamp?: any
  title?: string
  title_display?: any
  updated_at?: any
  virtual_event_passcode?: any
  virtual_event_url?: any
}

export interface EventLoadMatch {
  id: string
}

export interface EventListMatch {
  alt_audience_id?: string
  alt_event_type_id?: string
  api_link?: any
  api_model?: any
  audience_id?: string
  buy_button_caption?: any
  buy_button_text?: any
  date_display?: any
  description?: string
  door_time?: any
  end_date?: any
  end_time?: any
  entrance?: any
  event_host_id?: string
  event_host_title?: any
  event_type_id?: string
  header_description?: any
  hero_caption?: any
  id?: string
  image_url?: any
  is_admission_required?: boolean
  is_after_hour?: boolean
  is_free?: boolean
  is_member_exclusive?: boolean
  is_private?: boolean
  is_registration_required?: boolean
  is_sales_button_hidden?: boolean
  is_sold_out?: boolean
  is_ticketed?: boolean
  is_virtual_event?: boolean
  join_url?: any
  layout_type?: any
  list_description?: any
  location?: any
  program_id?: string
  program_title?: any
  rsvp_link?: any
  search_tag?: any
  short_description?: any
  slug?: string
  source_updated_at?: any
  start_date?: any
  start_time?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  survey_url?: any
  ticketed_event_id?: string
  timestamp?: any
  title?: string
  title_display?: any
  updated_at?: any
  virtual_event_passcode?: any
  virtual_event_url?: any
}

export interface EventOccurrence {
  api_link?: any
  api_model?: any
  button_caption?: any
  button_text?: any
  button_url?: any
  description?: string
  end_at?: any
  event_id?: string
  id?: string
  image_url?: any
  is_private?: boolean
  is_sales_button_hidden?: boolean
  is_ticketed?: boolean
  location?: any
  off_sale_at?: any
  on_sale_at?: any
  short_description?: any
  source_updated_at?: any
  start_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  title_display?: any
  updated_at?: any
}

export interface EventOccurrenceLoadMatch {
  id: string
}

export interface EventOccurrenceListMatch {
  api_link?: any
  api_model?: any
  button_caption?: any
  button_text?: any
  button_url?: any
  description?: string
  end_at?: any
  event_id?: string
  id?: string
  image_url?: any
  is_private?: boolean
  is_sales_button_hidden?: boolean
  is_ticketed?: boolean
  location?: any
  off_sale_at?: any
  on_sale_at?: any
  short_description?: any
  source_updated_at?: any
  start_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  title_display?: any
  updated_at?: any
}

export interface EventProgram {
  api_link?: any
  api_model?: any
  id?: string
  is_affiliate_group?: boolean
  is_event_host?: boolean
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface EventProgramLoadMatch {
  id: string
}

export interface EventProgramListMatch {
  api_link?: any
  api_model?: any
  id?: string
  is_affiliate_group?: boolean
  is_event_host?: boolean
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface Exhibition {
  aic_end_at?: any
  aic_start_at?: any
  alt_image_id?: string
  api_link?: any
  api_model?: any
  artist_id?: string
  artwork_id?: string
  artwork_title?: any
  document_id?: string
  gallery_id?: string
  gallery_title?: any
  id?: string
  image_id?: string
  image_url?: any
  is_featured?: boolean
  is_published?: boolean
  position?: any
  short_description?: any
  site_id?: string
  source_updated_at?: any
  status?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface ExhibitionLoadMatch {
  id: string
}

export interface ExhibitionListMatch {
  aic_end_at?: any
  aic_start_at?: any
  alt_image_id?: string
  api_link?: any
  api_model?: any
  artist_id?: string
  artwork_id?: string
  artwork_title?: any
  document_id?: string
  gallery_id?: string
  gallery_title?: any
  id?: string
  image_id?: string
  image_url?: any
  is_featured?: boolean
  is_published?: boolean
  position?: any
  short_description?: any
  site_id?: string
  source_updated_at?: any
  status?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Gallery {
  api_link?: any
  api_model?: any
  floor?: any
  id?: string
  is_closed?: boolean
  latitude?: number
  latlon?: any
  longitude?: number
  number?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  tgn_id?: string
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface GalleryLoadMatch {
  id: string
}

export interface GalleryListMatch {
  api_link?: any
  api_model?: any
  floor?: any
  id?: string
  is_closed?: boolean
  latitude?: number
  latlon?: any
  longitude?: number
  number?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  tgn_id?: string
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface GenericPage {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  search_tag?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface GenericPageLoadMatch {
  id: string
}

export interface GenericPageListMatch {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  search_tag?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Highlight {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface HighlightLoadMatch {
  id: string
}

export interface HighlightListMatch {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface Hour {
  additional_text?: any
  api_link?: any
  api_model?: any
  friday_is_closed?: any
  friday_member_close?: any
  friday_member_open?: any
  friday_public_close?: any
  friday_public_open?: any
  id?: string
  monday_is_closed?: any
  monday_member_close?: any
  monday_member_open?: any
  monday_public_close?: any
  monday_public_open?: any
  saturday_is_closed?: any
  saturday_member_close?: any
  saturday_member_open?: any
  saturday_public_close?: any
  saturday_public_open?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  summary?: any
  sunday_is_closed?: any
  sunday_member_close?: any
  sunday_member_open?: any
  sunday_public_close?: any
  sunday_public_open?: any
  thursday_is_closed?: any
  thursday_member_close?: any
  thursday_member_open?: any
  thursday_public_close?: any
  thursday_public_open?: any
  timestamp?: any
  title?: string
  tuesday_is_closed?: any
  tuesday_member_close?: any
  tuesday_member_open?: any
  tuesday_public_close?: any
  tuesday_public_open?: any
  updated_at?: any
  wednesday_is_closed?: any
  wednesday_member_close?: any
  wednesday_member_open?: any
  wednesday_public_close?: any
  wednesday_public_open?: any
}

export interface HourLoadMatch {
  id: string
}

export interface HourListMatch {
  additional_text?: any
  api_link?: any
  api_model?: any
  friday_is_closed?: any
  friday_member_close?: any
  friday_member_open?: any
  friday_public_close?: any
  friday_public_open?: any
  id?: string
  monday_is_closed?: any
  monday_member_close?: any
  monday_member_open?: any
  monday_public_close?: any
  monday_public_open?: any
  saturday_is_closed?: any
  saturday_member_close?: any
  saturday_member_open?: any
  saturday_public_close?: any
  saturday_public_open?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  summary?: any
  sunday_is_closed?: any
  sunday_member_close?: any
  sunday_member_open?: any
  sunday_public_close?: any
  sunday_public_open?: any
  thursday_is_closed?: any
  thursday_member_close?: any
  thursday_member_open?: any
  thursday_public_close?: any
  thursday_public_open?: any
  timestamp?: any
  title?: string
  tuesday_is_closed?: any
  tuesday_member_close?: any
  tuesday_member_open?: any
  tuesday_public_close?: any
  tuesday_public_open?: any
  updated_at?: any
  wednesday_is_closed?: any
  wednesday_member_close?: any
  wednesday_member_open?: any
  wednesday_public_close?: any
  wednesday_public_open?: any
}

export interface Image {
  ahash?: any
  alt_text?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  color?: any
  colorfulness?: any
  content?: any
  content_e_tag?: any
  credit_line?: any
  fingerprint?: any
  height?: number
  id?: string
  iiif_url?: any
  is_educational_resource?: boolean
  is_multimedia_resource?: boolean
  is_teacher_resource?: boolean
  lake_guid?: any
  lqip?: any
  phash?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  type?: any
  updated_at?: any
  width?: number
}

export interface ImageLoadMatch {
  id: string
}

export interface ImageListMatch {
  ahash?: any
  alt_text?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  color?: any
  colorfulness?: any
  content?: any
  content_e_tag?: any
  credit_line?: any
  fingerprint?: any
  height?: number
  id?: string
  iiif_url?: any
  is_educational_resource?: boolean
  is_multimedia_resource?: boolean
  is_teacher_resource?: boolean
  lake_guid?: any
  lqip?: any
  phash?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  type?: any
  updated_at?: any
  width?: number
}

export interface LandingPage {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  search_tag?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface LandingPageLoadMatch {
  id: string
}

export interface LandingPageListMatch {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  search_tag?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Place {
  api_link?: any
  api_model?: any
  id?: string
  latitude?: number
  longitude?: number
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  tgn_id?: string
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface PlaceLoadMatch {
  id: string
}

export interface PlaceListMatch {
  api_link?: any
  api_model?: any
  id?: string
  latitude?: number
  longitude?: number
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  tgn_id?: string
  timestamp?: any
  title?: string
  updated_at?: any
}

export interface PressRelease {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface PressReleaseLoadMatch {
  id: string
}

export interface PressReleaseListMatch {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface PrintedPublication {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface PrintedPublicationLoadMatch {
  id: string
}

export interface PrintedPublicationListMatch {
  api_link?: any
  api_model?: any
  copy?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Product {
  api_link?: any
  api_model?: any
  artist_id?: string
  artwork_id?: string
  description?: string
  exhibition_id?: string
  external_sku?: any
  id?: string
  image_url?: any
  max_compare_at_price?: any
  max_current_price?: any
  min_compare_at_price?: any
  min_current_price?: any
  price_display?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface ProductLoadMatch {
  id: string
}

export interface ProductListMatch {
  api_link?: any
  api_model?: any
  artist_id?: string
  artwork_id?: string
  description?: string
  exhibition_id?: string
  external_sku?: any
  id?: string
  image_url?: any
  max_compare_at_price?: any
  max_current_price?: any
  min_compare_at_price?: any
  min_current_price?: any
  price_display?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Publication {
  api_link?: any
  api_model?: any
  id?: string
  section_id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface PublicationLoadMatch {
  id: string
}

export interface PublicationListMatch {
  api_link?: any
  api_model?: any
  id?: string
  section_id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Search {
  api_id?: string
  api_link?: any
  api_model?: any
  id?: string
  is_boosted?: boolean
  score?: number
  thumbnail?: any
  timestamp?: any
  title?: string
}

export interface SearchListMatch {
  api_id?: string
  api_link?: any
  api_model?: any
  id?: string
  is_boosted?: boolean
  score?: number
  thumbnail?: any
  timestamp?: any
  title?: string
}

export interface Section {
  accession?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  content?: any
  generic_page_id?: string
  id?: string
  publication_id?: string
  publication_title?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface SectionLoadMatch {
  id: string
}

export interface SectionListMatch {
  accession?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  content?: any
  generic_page_id?: string
  id?: string
  publication_id?: string
  publication_title?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Site {
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  description?: string
  exhibition_id?: string
  exhibition_title?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface SiteLoadMatch {
  id: string
}

export interface SiteListMatch {
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  description?: string
  exhibition_id?: string
  exhibition_title?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Sound {
  alt_text?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  content?: any
  content_e_tag?: any
  credit_line?: any
  id?: string
  is_educational_resource?: boolean
  is_multimedia_resource?: boolean
  is_teacher_resource?: boolean
  lake_guid?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  transcript?: any
  type?: any
  updated_at?: any
  web_url?: any
}

export interface SoundLoadMatch {
  id: string
}

export interface SoundListMatch {
  alt_text?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  content?: any
  content_e_tag?: any
  credit_line?: any
  id?: string
  is_educational_resource?: boolean
  is_multimedia_resource?: boolean
  is_teacher_resource?: boolean
  lake_guid?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  transcript?: any
  type?: any
  updated_at?: any
  web_url?: any
}

export interface StaticPage {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface StaticPageLoadMatch {
  id: string
}

export interface StaticPageListMatch {
  api_link?: any
  api_model?: any
  id?: string
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  web_url?: any
}

export interface Text {
  alt_text?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  content?: any
  content_e_tag?: any
  credit_line?: any
  id?: string
  is_educational_resource?: boolean
  is_multimedia_resource?: boolean
  is_teacher_resource?: boolean
  lake_guid?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  type?: any
  updated_at?: any
}

export interface TextLoadMatch {
  id: string
}

export interface TextListMatch {
  alt_text?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  content?: any
  content_e_tag?: any
  credit_line?: any
  id?: string
  is_educational_resource?: boolean
  is_multimedia_resource?: boolean
  is_teacher_resource?: boolean
  lake_guid?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  type?: any
  updated_at?: any
}

export interface Tour {
  api_link?: any
  api_model?: any
  artist_title?: any
  artwork_title?: any
  description?: string
  id?: string
  image?: any
  intro?: any
  intro_link?: any
  intro_transcript?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  weight?: number
}

export interface TourLoadMatch {
  id: string
}

export interface TourListMatch {
  api_link?: any
  api_model?: any
  artist_title?: any
  artwork_title?: any
  description?: string
  id?: string
  image?: any
  intro?: any
  intro_link?: any
  intro_transcript?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  updated_at?: any
  weight?: number
}

export interface Video {
  alt_text?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  content?: any
  content_e_tag?: any
  credit_line?: any
  id?: string
  is_educational_resource?: boolean
  is_multimedia_resource?: boolean
  is_teacher_resource?: boolean
  lake_guid?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  type?: any
  updated_at?: any
}

export interface VideoLoadMatch {
  id: string
}

export interface VideoListMatch {
  alt_text?: any
  api_link?: any
  api_model?: any
  artwork_id?: string
  artwork_title?: any
  content?: any
  content_e_tag?: any
  credit_line?: any
  id?: string
  is_educational_resource?: boolean
  is_multimedia_resource?: boolean
  is_teacher_resource?: boolean
  lake_guid?: any
  source_updated_at?: any
  suggest_autocomplete_all?: any
  suggest_autocomplete_boosted?: any
  timestamp?: any
  title?: string
  type?: any
  updated_at?: any
}

