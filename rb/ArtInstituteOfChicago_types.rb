# frozen_string_literal: true

# Typed models for the ArtInstituteOfChicago SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Agent entity data model.
#
# @!attribute [rw] alt_title
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] birth_date
#   @return [Object, nil]
#
# @!attribute [rw] death_date
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_artist
#   @return [Boolean, nil]
#
# @!attribute [rw] sort_title
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] ulan_id
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
Agent = Struct.new(
  :alt_title,
  :api_link,
  :api_model,
  :birth_date,
  :death_date,
  :description,
  :id,
  :is_artist,
  :sort_title,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :ulan_id,
  :updated_at,
  keyword_init: true
)

# Request payload for Agent#load.
#
# @!attribute [rw] id
#   @return [String]
AgentLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Agent#list (any subset of Agent fields).
#
# @!attribute [rw] alt_title
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] birth_date
#   @return [Object, nil]
#
# @!attribute [rw] death_date
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_artist
#   @return [Boolean, nil]
#
# @!attribute [rw] sort_title
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] ulan_id
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
AgentListMatch = Struct.new(
  :alt_title,
  :api_link,
  :api_model,
  :birth_date,
  :death_date,
  :description,
  :id,
  :is_artist,
  :sort_title,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :ulan_id,
  :updated_at,
  keyword_init: true
)

# AgentRole entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
AgentRole = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for AgentRole#load.
#
# @!attribute [rw] id
#   @return [String]
AgentRoleLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for AgentRole#list (any subset of AgentRole fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
AgentRoleListMatch = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# AgentType entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
AgentType = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for AgentType#load.
#
# @!attribute [rw] id
#   @return [String]
AgentTypeLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for AgentType#list (any subset of AgentType fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
AgentTypeListMatch = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Article entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
Article = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for Article#load.
#
# @!attribute [rw] id
#   @return [String]
ArticleLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Article#list (any subset of Article fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
ArticleListMatch = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Artwork entity data model.
#
# @!attribute [rw] alt_artist_id
#   @return [String, nil]
#
# @!attribute [rw] alt_classification_id
#   @return [String, nil]
#
# @!attribute [rw] alt_image_id
#   @return [String, nil]
#
# @!attribute [rw] alt_material_id
#   @return [String, nil]
#
# @!attribute [rw] alt_style_id
#   @return [String, nil]
#
# @!attribute [rw] alt_subject_id
#   @return [String, nil]
#
# @!attribute [rw] alt_technique_id
#   @return [String, nil]
#
# @!attribute [rw] alt_title
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artist_display
#   @return [Object, nil]
#
# @!attribute [rw] artist_id
#   @return [String, nil]
#
# @!attribute [rw] artist_title
#   @return [Object, nil]
#
# @!attribute [rw] artwork_type_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_type_title
#   @return [Object, nil]
#
# @!attribute [rw] boost_rank
#   @return [Object, nil]
#
# @!attribute [rw] catalog_based_search_keyword_title
#   @return [Object, nil]
#
# @!attribute [rw] catalogue_display
#   @return [Object, nil]
#
# @!attribute [rw] category_id
#   @return [String, nil]
#
# @!attribute [rw] category_title
#   @return [Object, nil]
#
# @!attribute [rw] classification_id
#   @return [String, nil]
#
# @!attribute [rw] classification_title
#   @return [Object, nil]
#
# @!attribute [rw] color
#   @return [Object, nil]
#
# @!attribute [rw] colorfulness
#   @return [Object, nil]
#
# @!attribute [rw] copyright_notice
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] date_display
#   @return [Object, nil]
#
# @!attribute [rw] date_end
#   @return [Object, nil]
#
# @!attribute [rw] date_qualifier_id
#   @return [String, nil]
#
# @!attribute [rw] date_qualifier_title
#   @return [Object, nil]
#
# @!attribute [rw] date_start
#   @return [Object, nil]
#
# @!attribute [rw] department_id
#   @return [String, nil]
#
# @!attribute [rw] department_title
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] dimension
#   @return [Object, nil]
#
# @!attribute [rw] dimensions_detail
#   @return [Object, nil]
#
# @!attribute [rw] document_id
#   @return [String, nil]
#
# @!attribute [rw] edition
#   @return [Object, nil]
#
# @!attribute [rw] exhibition_history
#   @return [Object, nil]
#
# @!attribute [rw] fiscal_year
#   @return [Object, nil]
#
# @!attribute [rw] fiscal_year_deaccession
#   @return [Object, nil]
#
# @!attribute [rw] gallery_id
#   @return [String, nil]
#
# @!attribute [rw] gallery_title
#   @return [Object, nil]
#
# @!attribute [rw] has_advanced_imaging
#   @return [Boolean, nil]
#
# @!attribute [rw] has_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] has_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] has_not_been_viewed_much
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_embedding
#   @return [Object, nil]
#
# @!attribute [rw] image_id
#   @return [String, nil]
#
# @!attribute [rw] inscription
#   @return [Object, nil]
#
# @!attribute [rw] internal_department_id
#   @return [String, nil]
#
# @!attribute [rw] is_boosted
#   @return [Boolean, nil]
#
# @!attribute [rw] is_on_view
#   @return [Boolean, nil]
#
# @!attribute [rw] is_public_domain
#   @return [Boolean, nil]
#
# @!attribute [rw] is_zoomable
#   @return [Boolean, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] latlon
#   @return [Object, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] main_reference_number
#   @return [Integer, nil]
#
# @!attribute [rw] material_id
#   @return [String, nil]
#
# @!attribute [rw] material_title
#   @return [Object, nil]
#
# @!attribute [rw] max_zoom_window_size
#   @return [Object, nil]
#
# @!attribute [rw] medium_display
#   @return [Object, nil]
#
# @!attribute [rw] nomisma_id
#   @return [String, nil]
#
# @!attribute [rw] on_loan_display
#   @return [Object, nil]
#
# @!attribute [rw] pageview
#   @return [Object, nil]
#
# @!attribute [rw] pageviews_recent
#   @return [Object, nil]
#
# @!attribute [rw] place_of_origin
#   @return [Object, nil]
#
# @!attribute [rw] provenance_text
#   @return [Object, nil]
#
# @!attribute [rw] publication_history
#   @return [Object, nil]
#
# @!attribute [rw] publishing_verification_level
#   @return [Object, nil]
#
# @!attribute [rw] section_id
#   @return [String, nil]
#
# @!attribute [rw] section_title
#   @return [Object, nil]
#
# @!attribute [rw] short_description
#   @return [Object, nil]
#
# @!attribute [rw] site_id
#   @return [String, nil]
#
# @!attribute [rw] sound_id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] style_id
#   @return [String, nil]
#
# @!attribute [rw] style_title
#   @return [Object, nil]
#
# @!attribute [rw] subject_id
#   @return [String, nil]
#
# @!attribute [rw] subject_title
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] technique_id
#   @return [String, nil]
#
# @!attribute [rw] technique_title
#   @return [Object, nil]
#
# @!attribute [rw] term_title
#   @return [Object, nil]
#
# @!attribute [rw] text_embedding
#   @return [Object, nil]
#
# @!attribute [rw] text_id
#   @return [String, nil]
#
# @!attribute [rw] theme_title
#   @return [Object, nil]
#
# @!attribute [rw] thumbnail
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] video_id
#   @return [String, nil]
Artwork = Struct.new(
  :alt_artist_id,
  :alt_classification_id,
  :alt_image_id,
  :alt_material_id,
  :alt_style_id,
  :alt_subject_id,
  :alt_technique_id,
  :alt_title,
  :api_link,
  :api_model,
  :artist_display,
  :artist_id,
  :artist_title,
  :artwork_type_id,
  :artwork_type_title,
  :boost_rank,
  :catalog_based_search_keyword_title,
  :catalogue_display,
  :category_id,
  :category_title,
  :classification_id,
  :classification_title,
  :color,
  :colorfulness,
  :copyright_notice,
  :credit_line,
  :date_display,
  :date_end,
  :date_qualifier_id,
  :date_qualifier_title,
  :date_start,
  :department_id,
  :department_title,
  :description,
  :dimension,
  :dimensions_detail,
  :document_id,
  :edition,
  :exhibition_history,
  :fiscal_year,
  :fiscal_year_deaccession,
  :gallery_id,
  :gallery_title,
  :has_advanced_imaging,
  :has_educational_resource,
  :has_multimedia_resource,
  :has_not_been_viewed_much,
  :id,
  :image_embedding,
  :image_id,
  :inscription,
  :internal_department_id,
  :is_boosted,
  :is_on_view,
  :is_public_domain,
  :is_zoomable,
  :latitude,
  :latlon,
  :longitude,
  :main_reference_number,
  :material_id,
  :material_title,
  :max_zoom_window_size,
  :medium_display,
  :nomisma_id,
  :on_loan_display,
  :pageview,
  :pageviews_recent,
  :place_of_origin,
  :provenance_text,
  :publication_history,
  :publishing_verification_level,
  :section_id,
  :section_title,
  :short_description,
  :site_id,
  :sound_id,
  :source_updated_at,
  :style_id,
  :style_title,
  :subject_id,
  :subject_title,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :technique_id,
  :technique_title,
  :term_title,
  :text_embedding,
  :text_id,
  :theme_title,
  :thumbnail,
  :timestamp,
  :title,
  :updated_at,
  :video_id,
  keyword_init: true
)

# Request payload for Artwork#load.
#
# @!attribute [rw] id
#   @return [String]
ArtworkLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Artwork#list (any subset of Artwork fields).
#
# @!attribute [rw] alt_artist_id
#   @return [String, nil]
#
# @!attribute [rw] alt_classification_id
#   @return [String, nil]
#
# @!attribute [rw] alt_image_id
#   @return [String, nil]
#
# @!attribute [rw] alt_material_id
#   @return [String, nil]
#
# @!attribute [rw] alt_style_id
#   @return [String, nil]
#
# @!attribute [rw] alt_subject_id
#   @return [String, nil]
#
# @!attribute [rw] alt_technique_id
#   @return [String, nil]
#
# @!attribute [rw] alt_title
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artist_display
#   @return [Object, nil]
#
# @!attribute [rw] artist_id
#   @return [String, nil]
#
# @!attribute [rw] artist_title
#   @return [Object, nil]
#
# @!attribute [rw] artwork_type_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_type_title
#   @return [Object, nil]
#
# @!attribute [rw] boost_rank
#   @return [Object, nil]
#
# @!attribute [rw] catalog_based_search_keyword_title
#   @return [Object, nil]
#
# @!attribute [rw] catalogue_display
#   @return [Object, nil]
#
# @!attribute [rw] category_id
#   @return [String, nil]
#
# @!attribute [rw] category_title
#   @return [Object, nil]
#
# @!attribute [rw] classification_id
#   @return [String, nil]
#
# @!attribute [rw] classification_title
#   @return [Object, nil]
#
# @!attribute [rw] color
#   @return [Object, nil]
#
# @!attribute [rw] colorfulness
#   @return [Object, nil]
#
# @!attribute [rw] copyright_notice
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] date_display
#   @return [Object, nil]
#
# @!attribute [rw] date_end
#   @return [Object, nil]
#
# @!attribute [rw] date_qualifier_id
#   @return [String, nil]
#
# @!attribute [rw] date_qualifier_title
#   @return [Object, nil]
#
# @!attribute [rw] date_start
#   @return [Object, nil]
#
# @!attribute [rw] department_id
#   @return [String, nil]
#
# @!attribute [rw] department_title
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] dimension
#   @return [Object, nil]
#
# @!attribute [rw] dimensions_detail
#   @return [Object, nil]
#
# @!attribute [rw] document_id
#   @return [String, nil]
#
# @!attribute [rw] edition
#   @return [Object, nil]
#
# @!attribute [rw] exhibition_history
#   @return [Object, nil]
#
# @!attribute [rw] fiscal_year
#   @return [Object, nil]
#
# @!attribute [rw] fiscal_year_deaccession
#   @return [Object, nil]
#
# @!attribute [rw] gallery_id
#   @return [String, nil]
#
# @!attribute [rw] gallery_title
#   @return [Object, nil]
#
# @!attribute [rw] has_advanced_imaging
#   @return [Boolean, nil]
#
# @!attribute [rw] has_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] has_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] has_not_been_viewed_much
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_embedding
#   @return [Object, nil]
#
# @!attribute [rw] image_id
#   @return [String, nil]
#
# @!attribute [rw] inscription
#   @return [Object, nil]
#
# @!attribute [rw] internal_department_id
#   @return [String, nil]
#
# @!attribute [rw] is_boosted
#   @return [Boolean, nil]
#
# @!attribute [rw] is_on_view
#   @return [Boolean, nil]
#
# @!attribute [rw] is_public_domain
#   @return [Boolean, nil]
#
# @!attribute [rw] is_zoomable
#   @return [Boolean, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] latlon
#   @return [Object, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] main_reference_number
#   @return [Integer, nil]
#
# @!attribute [rw] material_id
#   @return [String, nil]
#
# @!attribute [rw] material_title
#   @return [Object, nil]
#
# @!attribute [rw] max_zoom_window_size
#   @return [Object, nil]
#
# @!attribute [rw] medium_display
#   @return [Object, nil]
#
# @!attribute [rw] nomisma_id
#   @return [String, nil]
#
# @!attribute [rw] on_loan_display
#   @return [Object, nil]
#
# @!attribute [rw] pageview
#   @return [Object, nil]
#
# @!attribute [rw] pageviews_recent
#   @return [Object, nil]
#
# @!attribute [rw] place_of_origin
#   @return [Object, nil]
#
# @!attribute [rw] provenance_text
#   @return [Object, nil]
#
# @!attribute [rw] publication_history
#   @return [Object, nil]
#
# @!attribute [rw] publishing_verification_level
#   @return [Object, nil]
#
# @!attribute [rw] section_id
#   @return [String, nil]
#
# @!attribute [rw] section_title
#   @return [Object, nil]
#
# @!attribute [rw] short_description
#   @return [Object, nil]
#
# @!attribute [rw] site_id
#   @return [String, nil]
#
# @!attribute [rw] sound_id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] style_id
#   @return [String, nil]
#
# @!attribute [rw] style_title
#   @return [Object, nil]
#
# @!attribute [rw] subject_id
#   @return [String, nil]
#
# @!attribute [rw] subject_title
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] technique_id
#   @return [String, nil]
#
# @!attribute [rw] technique_title
#   @return [Object, nil]
#
# @!attribute [rw] term_title
#   @return [Object, nil]
#
# @!attribute [rw] text_embedding
#   @return [Object, nil]
#
# @!attribute [rw] text_id
#   @return [String, nil]
#
# @!attribute [rw] theme_title
#   @return [Object, nil]
#
# @!attribute [rw] thumbnail
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] video_id
#   @return [String, nil]
ArtworkListMatch = Struct.new(
  :alt_artist_id,
  :alt_classification_id,
  :alt_image_id,
  :alt_material_id,
  :alt_style_id,
  :alt_subject_id,
  :alt_technique_id,
  :alt_title,
  :api_link,
  :api_model,
  :artist_display,
  :artist_id,
  :artist_title,
  :artwork_type_id,
  :artwork_type_title,
  :boost_rank,
  :catalog_based_search_keyword_title,
  :catalogue_display,
  :category_id,
  :category_title,
  :classification_id,
  :classification_title,
  :color,
  :colorfulness,
  :copyright_notice,
  :credit_line,
  :date_display,
  :date_end,
  :date_qualifier_id,
  :date_qualifier_title,
  :date_start,
  :department_id,
  :department_title,
  :description,
  :dimension,
  :dimensions_detail,
  :document_id,
  :edition,
  :exhibition_history,
  :fiscal_year,
  :fiscal_year_deaccession,
  :gallery_id,
  :gallery_title,
  :has_advanced_imaging,
  :has_educational_resource,
  :has_multimedia_resource,
  :has_not_been_viewed_much,
  :id,
  :image_embedding,
  :image_id,
  :inscription,
  :internal_department_id,
  :is_boosted,
  :is_on_view,
  :is_public_domain,
  :is_zoomable,
  :latitude,
  :latlon,
  :longitude,
  :main_reference_number,
  :material_id,
  :material_title,
  :max_zoom_window_size,
  :medium_display,
  :nomisma_id,
  :on_loan_display,
  :pageview,
  :pageviews_recent,
  :place_of_origin,
  :provenance_text,
  :publication_history,
  :publishing_verification_level,
  :section_id,
  :section_title,
  :short_description,
  :site_id,
  :sound_id,
  :source_updated_at,
  :style_id,
  :style_title,
  :subject_id,
  :subject_title,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :technique_id,
  :technique_title,
  :term_title,
  :text_embedding,
  :text_id,
  :theme_title,
  :thumbnail,
  :timestamp,
  :title,
  :updated_at,
  :video_id,
  keyword_init: true
)

# ArtworkDateQualifier entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
ArtworkDateQualifier = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for ArtworkDateQualifier#load.
#
# @!attribute [rw] id
#   @return [String]
ArtworkDateQualifierLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for ArtworkDateQualifier#list (any subset of ArtworkDateQualifier fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
ArtworkDateQualifierListMatch = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# ArtworkPlaceQualifier entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
ArtworkPlaceQualifier = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for ArtworkPlaceQualifier#load.
#
# @!attribute [rw] id
#   @return [String]
ArtworkPlaceQualifierLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for ArtworkPlaceQualifier#list (any subset of ArtworkPlaceQualifier fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
ArtworkPlaceQualifierListMatch = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# ArtworkType entity data model.
#
# @!attribute [rw] aat_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
ArtworkType = Struct.new(
  :aat_id,
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for ArtworkType#load.
#
# @!attribute [rw] id
#   @return [String]
ArtworkTypeLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for ArtworkType#list (any subset of ArtworkType fields).
#
# @!attribute [rw] aat_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
ArtworkTypeListMatch = Struct.new(
  :aat_id,
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# CategoryTerm entity data model.
#
# @!attribute [rw] aat_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] parent_id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] subtype
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
CategoryTerm = Struct.new(
  :aat_id,
  :api_link,
  :api_model,
  :id,
  :parent_id,
  :source_updated_at,
  :subtype,
  :suggest_autocomplete_all,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for CategoryTerm#load.
#
# @!attribute [rw] id
#   @return [String]
CategoryTermLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for CategoryTerm#list (any subset of CategoryTerm fields).
#
# @!attribute [rw] aat_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] parent_id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] subtype
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
CategoryTermListMatch = Struct.new(
  :aat_id,
  :api_link,
  :api_model,
  :id,
  :parent_id,
  :source_updated_at,
  :subtype,
  :suggest_autocomplete_all,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# DigitalPublication entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
DigitalPublication = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for DigitalPublication#load.
#
# @!attribute [rw] id
#   @return [String]
DigitalPublicationLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for DigitalPublication#list (any subset of DigitalPublication fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
DigitalPublicationListMatch = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# DigitalPublicationArticle entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] author_display
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] digital_publication_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
DigitalPublicationArticle = Struct.new(
  :api_link,
  :api_model,
  :author_display,
  :copy,
  :digital_publication_id,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for DigitalPublicationArticle#load.
#
# @!attribute [rw] id
#   @return [String]
DigitalPublicationArticleLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for DigitalPublicationArticle#list (any subset of DigitalPublicationArticle fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] author_display
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] digital_publication_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
DigitalPublicationArticleListMatch = Struct.new(
  :api_link,
  :api_model,
  :author_display,
  :copy,
  :digital_publication_id,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# EducatorResource entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
EducatorResource = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for EducatorResource#load.
#
# @!attribute [rw] id
#   @return [String]
EducatorResourceLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for EducatorResource#list (any subset of EducatorResource fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
EducatorResourceListMatch = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Event entity data model.
#
# @!attribute [rw] alt_audience_id
#   @return [String, nil]
#
# @!attribute [rw] alt_event_type_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] audience_id
#   @return [String, nil]
#
# @!attribute [rw] buy_button_caption
#   @return [Object, nil]
#
# @!attribute [rw] buy_button_text
#   @return [Object, nil]
#
# @!attribute [rw] date_display
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] door_time
#   @return [Object, nil]
#
# @!attribute [rw] end_date
#   @return [Object, nil]
#
# @!attribute [rw] end_time
#   @return [Object, nil]
#
# @!attribute [rw] entrance
#   @return [Object, nil]
#
# @!attribute [rw] event_host_id
#   @return [String, nil]
#
# @!attribute [rw] event_host_title
#   @return [Object, nil]
#
# @!attribute [rw] event_type_id
#   @return [String, nil]
#
# @!attribute [rw] header_description
#   @return [Object, nil]
#
# @!attribute [rw] hero_caption
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [Object, nil]
#
# @!attribute [rw] is_admission_required
#   @return [Boolean, nil]
#
# @!attribute [rw] is_after_hour
#   @return [Boolean, nil]
#
# @!attribute [rw] is_free
#   @return [Boolean, nil]
#
# @!attribute [rw] is_member_exclusive
#   @return [Boolean, nil]
#
# @!attribute [rw] is_private
#   @return [Boolean, nil]
#
# @!attribute [rw] is_registration_required
#   @return [Boolean, nil]
#
# @!attribute [rw] is_sales_button_hidden
#   @return [Boolean, nil]
#
# @!attribute [rw] is_sold_out
#   @return [Boolean, nil]
#
# @!attribute [rw] is_ticketed
#   @return [Boolean, nil]
#
# @!attribute [rw] is_virtual_event
#   @return [Boolean, nil]
#
# @!attribute [rw] join_url
#   @return [Object, nil]
#
# @!attribute [rw] layout_type
#   @return [Object, nil]
#
# @!attribute [rw] list_description
#   @return [Object, nil]
#
# @!attribute [rw] location
#   @return [Object, nil]
#
# @!attribute [rw] program_id
#   @return [String, nil]
#
# @!attribute [rw] program_title
#   @return [Object, nil]
#
# @!attribute [rw] rsvp_link
#   @return [Object, nil]
#
# @!attribute [rw] search_tag
#   @return [Object, nil]
#
# @!attribute [rw] short_description
#   @return [Object, nil]
#
# @!attribute [rw] slug
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] start_date
#   @return [Object, nil]
#
# @!attribute [rw] start_time
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] survey_url
#   @return [Object, nil]
#
# @!attribute [rw] ticketed_event_id
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] title_display
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] virtual_event_passcode
#   @return [Object, nil]
#
# @!attribute [rw] virtual_event_url
#   @return [Object, nil]
Event = Struct.new(
  :alt_audience_id,
  :alt_event_type_id,
  :api_link,
  :api_model,
  :audience_id,
  :buy_button_caption,
  :buy_button_text,
  :date_display,
  :description,
  :door_time,
  :end_date,
  :end_time,
  :entrance,
  :event_host_id,
  :event_host_title,
  :event_type_id,
  :header_description,
  :hero_caption,
  :id,
  :image_url,
  :is_admission_required,
  :is_after_hour,
  :is_free,
  :is_member_exclusive,
  :is_private,
  :is_registration_required,
  :is_sales_button_hidden,
  :is_sold_out,
  :is_ticketed,
  :is_virtual_event,
  :join_url,
  :layout_type,
  :list_description,
  :location,
  :program_id,
  :program_title,
  :rsvp_link,
  :search_tag,
  :short_description,
  :slug,
  :source_updated_at,
  :start_date,
  :start_time,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :survey_url,
  :ticketed_event_id,
  :timestamp,
  :title,
  :title_display,
  :updated_at,
  :virtual_event_passcode,
  :virtual_event_url,
  keyword_init: true
)

# Request payload for Event#load.
#
# @!attribute [rw] id
#   @return [String]
EventLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Event#list (any subset of Event fields).
#
# @!attribute [rw] alt_audience_id
#   @return [String, nil]
#
# @!attribute [rw] alt_event_type_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] audience_id
#   @return [String, nil]
#
# @!attribute [rw] buy_button_caption
#   @return [Object, nil]
#
# @!attribute [rw] buy_button_text
#   @return [Object, nil]
#
# @!attribute [rw] date_display
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] door_time
#   @return [Object, nil]
#
# @!attribute [rw] end_date
#   @return [Object, nil]
#
# @!attribute [rw] end_time
#   @return [Object, nil]
#
# @!attribute [rw] entrance
#   @return [Object, nil]
#
# @!attribute [rw] event_host_id
#   @return [String, nil]
#
# @!attribute [rw] event_host_title
#   @return [Object, nil]
#
# @!attribute [rw] event_type_id
#   @return [String, nil]
#
# @!attribute [rw] header_description
#   @return [Object, nil]
#
# @!attribute [rw] hero_caption
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [Object, nil]
#
# @!attribute [rw] is_admission_required
#   @return [Boolean, nil]
#
# @!attribute [rw] is_after_hour
#   @return [Boolean, nil]
#
# @!attribute [rw] is_free
#   @return [Boolean, nil]
#
# @!attribute [rw] is_member_exclusive
#   @return [Boolean, nil]
#
# @!attribute [rw] is_private
#   @return [Boolean, nil]
#
# @!attribute [rw] is_registration_required
#   @return [Boolean, nil]
#
# @!attribute [rw] is_sales_button_hidden
#   @return [Boolean, nil]
#
# @!attribute [rw] is_sold_out
#   @return [Boolean, nil]
#
# @!attribute [rw] is_ticketed
#   @return [Boolean, nil]
#
# @!attribute [rw] is_virtual_event
#   @return [Boolean, nil]
#
# @!attribute [rw] join_url
#   @return [Object, nil]
#
# @!attribute [rw] layout_type
#   @return [Object, nil]
#
# @!attribute [rw] list_description
#   @return [Object, nil]
#
# @!attribute [rw] location
#   @return [Object, nil]
#
# @!attribute [rw] program_id
#   @return [String, nil]
#
# @!attribute [rw] program_title
#   @return [Object, nil]
#
# @!attribute [rw] rsvp_link
#   @return [Object, nil]
#
# @!attribute [rw] search_tag
#   @return [Object, nil]
#
# @!attribute [rw] short_description
#   @return [Object, nil]
#
# @!attribute [rw] slug
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] start_date
#   @return [Object, nil]
#
# @!attribute [rw] start_time
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] survey_url
#   @return [Object, nil]
#
# @!attribute [rw] ticketed_event_id
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] title_display
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] virtual_event_passcode
#   @return [Object, nil]
#
# @!attribute [rw] virtual_event_url
#   @return [Object, nil]
EventListMatch = Struct.new(
  :alt_audience_id,
  :alt_event_type_id,
  :api_link,
  :api_model,
  :audience_id,
  :buy_button_caption,
  :buy_button_text,
  :date_display,
  :description,
  :door_time,
  :end_date,
  :end_time,
  :entrance,
  :event_host_id,
  :event_host_title,
  :event_type_id,
  :header_description,
  :hero_caption,
  :id,
  :image_url,
  :is_admission_required,
  :is_after_hour,
  :is_free,
  :is_member_exclusive,
  :is_private,
  :is_registration_required,
  :is_sales_button_hidden,
  :is_sold_out,
  :is_ticketed,
  :is_virtual_event,
  :join_url,
  :layout_type,
  :list_description,
  :location,
  :program_id,
  :program_title,
  :rsvp_link,
  :search_tag,
  :short_description,
  :slug,
  :source_updated_at,
  :start_date,
  :start_time,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :survey_url,
  :ticketed_event_id,
  :timestamp,
  :title,
  :title_display,
  :updated_at,
  :virtual_event_passcode,
  :virtual_event_url,
  keyword_init: true
)

# EventOccurrence entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] button_caption
#   @return [Object, nil]
#
# @!attribute [rw] button_text
#   @return [Object, nil]
#
# @!attribute [rw] button_url
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] end_at
#   @return [Object, nil]
#
# @!attribute [rw] event_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [Object, nil]
#
# @!attribute [rw] is_private
#   @return [Boolean, nil]
#
# @!attribute [rw] is_sales_button_hidden
#   @return [Boolean, nil]
#
# @!attribute [rw] is_ticketed
#   @return [Boolean, nil]
#
# @!attribute [rw] location
#   @return [Object, nil]
#
# @!attribute [rw] off_sale_at
#   @return [Object, nil]
#
# @!attribute [rw] on_sale_at
#   @return [Object, nil]
#
# @!attribute [rw] short_description
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] start_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] title_display
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
EventOccurrence = Struct.new(
  :api_link,
  :api_model,
  :button_caption,
  :button_text,
  :button_url,
  :description,
  :end_at,
  :event_id,
  :id,
  :image_url,
  :is_private,
  :is_sales_button_hidden,
  :is_ticketed,
  :location,
  :off_sale_at,
  :on_sale_at,
  :short_description,
  :source_updated_at,
  :start_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :title_display,
  :updated_at,
  keyword_init: true
)

# Request payload for EventOccurrence#load.
#
# @!attribute [rw] id
#   @return [String]
EventOccurrenceLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for EventOccurrence#list (any subset of EventOccurrence fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] button_caption
#   @return [Object, nil]
#
# @!attribute [rw] button_text
#   @return [Object, nil]
#
# @!attribute [rw] button_url
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] end_at
#   @return [Object, nil]
#
# @!attribute [rw] event_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [Object, nil]
#
# @!attribute [rw] is_private
#   @return [Boolean, nil]
#
# @!attribute [rw] is_sales_button_hidden
#   @return [Boolean, nil]
#
# @!attribute [rw] is_ticketed
#   @return [Boolean, nil]
#
# @!attribute [rw] location
#   @return [Object, nil]
#
# @!attribute [rw] off_sale_at
#   @return [Object, nil]
#
# @!attribute [rw] on_sale_at
#   @return [Object, nil]
#
# @!attribute [rw] short_description
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] start_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] title_display
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
EventOccurrenceListMatch = Struct.new(
  :api_link,
  :api_model,
  :button_caption,
  :button_text,
  :button_url,
  :description,
  :end_at,
  :event_id,
  :id,
  :image_url,
  :is_private,
  :is_sales_button_hidden,
  :is_ticketed,
  :location,
  :off_sale_at,
  :on_sale_at,
  :short_description,
  :source_updated_at,
  :start_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :title_display,
  :updated_at,
  keyword_init: true
)

# EventProgram entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_affiliate_group
#   @return [Boolean, nil]
#
# @!attribute [rw] is_event_host
#   @return [Boolean, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
EventProgram = Struct.new(
  :api_link,
  :api_model,
  :id,
  :is_affiliate_group,
  :is_event_host,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for EventProgram#load.
#
# @!attribute [rw] id
#   @return [String]
EventProgramLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for EventProgram#list (any subset of EventProgram fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_affiliate_group
#   @return [Boolean, nil]
#
# @!attribute [rw] is_event_host
#   @return [Boolean, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
EventProgramListMatch = Struct.new(
  :api_link,
  :api_model,
  :id,
  :is_affiliate_group,
  :is_event_host,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Exhibition entity data model.
#
# @!attribute [rw] aic_end_at
#   @return [Object, nil]
#
# @!attribute [rw] aic_start_at
#   @return [Object, nil]
#
# @!attribute [rw] alt_image_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artist_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] document_id
#   @return [String, nil]
#
# @!attribute [rw] gallery_id
#   @return [String, nil]
#
# @!attribute [rw] gallery_title
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_id
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [Object, nil]
#
# @!attribute [rw] is_featured
#   @return [Boolean, nil]
#
# @!attribute [rw] is_published
#   @return [Boolean, nil]
#
# @!attribute [rw] position
#   @return [Object, nil]
#
# @!attribute [rw] short_description
#   @return [Object, nil]
#
# @!attribute [rw] site_id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] status
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
Exhibition = Struct.new(
  :aic_end_at,
  :aic_start_at,
  :alt_image_id,
  :api_link,
  :api_model,
  :artist_id,
  :artwork_id,
  :artwork_title,
  :document_id,
  :gallery_id,
  :gallery_title,
  :id,
  :image_id,
  :image_url,
  :is_featured,
  :is_published,
  :position,
  :short_description,
  :site_id,
  :source_updated_at,
  :status,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for Exhibition#load.
#
# @!attribute [rw] id
#   @return [String]
ExhibitionLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Exhibition#list (any subset of Exhibition fields).
#
# @!attribute [rw] aic_end_at
#   @return [Object, nil]
#
# @!attribute [rw] aic_start_at
#   @return [Object, nil]
#
# @!attribute [rw] alt_image_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artist_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] document_id
#   @return [String, nil]
#
# @!attribute [rw] gallery_id
#   @return [String, nil]
#
# @!attribute [rw] gallery_title
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_id
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [Object, nil]
#
# @!attribute [rw] is_featured
#   @return [Boolean, nil]
#
# @!attribute [rw] is_published
#   @return [Boolean, nil]
#
# @!attribute [rw] position
#   @return [Object, nil]
#
# @!attribute [rw] short_description
#   @return [Object, nil]
#
# @!attribute [rw] site_id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] status
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
ExhibitionListMatch = Struct.new(
  :aic_end_at,
  :aic_start_at,
  :alt_image_id,
  :api_link,
  :api_model,
  :artist_id,
  :artwork_id,
  :artwork_title,
  :document_id,
  :gallery_id,
  :gallery_title,
  :id,
  :image_id,
  :image_url,
  :is_featured,
  :is_published,
  :position,
  :short_description,
  :site_id,
  :source_updated_at,
  :status,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Gallery entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] floor
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_closed
#   @return [Boolean, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] latlon
#   @return [Object, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] number
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] tgn_id
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
Gallery = Struct.new(
  :api_link,
  :api_model,
  :floor,
  :id,
  :is_closed,
  :latitude,
  :latlon,
  :longitude,
  :number,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :tgn_id,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for Gallery#load.
#
# @!attribute [rw] id
#   @return [String]
GalleryLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Gallery#list (any subset of Gallery fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] floor
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_closed
#   @return [Boolean, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] latlon
#   @return [Object, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] number
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] tgn_id
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
GalleryListMatch = Struct.new(
  :api_link,
  :api_model,
  :floor,
  :id,
  :is_closed,
  :latitude,
  :latlon,
  :longitude,
  :number,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :tgn_id,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# GenericPage entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] search_tag
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
GenericPage = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :search_tag,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for GenericPage#load.
#
# @!attribute [rw] id
#   @return [String]
GenericPageLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for GenericPage#list (any subset of GenericPage fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] search_tag
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
GenericPageListMatch = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :search_tag,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Highlight entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
Highlight = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for Highlight#load.
#
# @!attribute [rw] id
#   @return [String]
HighlightLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Highlight#list (any subset of Highlight fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
HighlightListMatch = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Hour entity data model.
#
# @!attribute [rw] additional_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] friday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] friday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] friday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] friday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] friday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] monday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] monday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] monday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] monday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] monday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] saturday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] saturday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] saturday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] saturday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] saturday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] summary
#   @return [Object, nil]
#
# @!attribute [rw] sunday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] sunday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] sunday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] sunday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] sunday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] thursday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] thursday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] thursday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] thursday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] thursday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tuesday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] tuesday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] tuesday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] tuesday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] tuesday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_public_open
#   @return [Object, nil]
Hour = Struct.new(
  :additional_text,
  :api_link,
  :api_model,
  :friday_is_closed,
  :friday_member_close,
  :friday_member_open,
  :friday_public_close,
  :friday_public_open,
  :id,
  :monday_is_closed,
  :monday_member_close,
  :monday_member_open,
  :monday_public_close,
  :monday_public_open,
  :saturday_is_closed,
  :saturday_member_close,
  :saturday_member_open,
  :saturday_public_close,
  :saturday_public_open,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :summary,
  :sunday_is_closed,
  :sunday_member_close,
  :sunday_member_open,
  :sunday_public_close,
  :sunday_public_open,
  :thursday_is_closed,
  :thursday_member_close,
  :thursday_member_open,
  :thursday_public_close,
  :thursday_public_open,
  :timestamp,
  :title,
  :tuesday_is_closed,
  :tuesday_member_close,
  :tuesday_member_open,
  :tuesday_public_close,
  :tuesday_public_open,
  :updated_at,
  :wednesday_is_closed,
  :wednesday_member_close,
  :wednesday_member_open,
  :wednesday_public_close,
  :wednesday_public_open,
  keyword_init: true
)

# Request payload for Hour#load.
#
# @!attribute [rw] id
#   @return [String]
HourLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Hour#list (any subset of Hour fields).
#
# @!attribute [rw] additional_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] friday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] friday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] friday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] friday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] friday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] monday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] monday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] monday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] monday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] monday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] saturday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] saturday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] saturday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] saturday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] saturday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] summary
#   @return [Object, nil]
#
# @!attribute [rw] sunday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] sunday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] sunday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] sunday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] sunday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] thursday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] thursday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] thursday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] thursday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] thursday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tuesday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] tuesday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] tuesday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] tuesday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] tuesday_public_open
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_is_closed
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_member_close
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_member_open
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_public_close
#   @return [Object, nil]
#
# @!attribute [rw] wednesday_public_open
#   @return [Object, nil]
HourListMatch = Struct.new(
  :additional_text,
  :api_link,
  :api_model,
  :friday_is_closed,
  :friday_member_close,
  :friday_member_open,
  :friday_public_close,
  :friday_public_open,
  :id,
  :monday_is_closed,
  :monday_member_close,
  :monday_member_open,
  :monday_public_close,
  :monday_public_open,
  :saturday_is_closed,
  :saturday_member_close,
  :saturday_member_open,
  :saturday_public_close,
  :saturday_public_open,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :summary,
  :sunday_is_closed,
  :sunday_member_close,
  :sunday_member_open,
  :sunday_public_close,
  :sunday_public_open,
  :thursday_is_closed,
  :thursday_member_close,
  :thursday_member_open,
  :thursday_public_close,
  :thursday_public_open,
  :timestamp,
  :title,
  :tuesday_is_closed,
  :tuesday_member_close,
  :tuesday_member_open,
  :tuesday_public_close,
  :tuesday_public_open,
  :updated_at,
  :wednesday_is_closed,
  :wednesday_member_close,
  :wednesday_member_open,
  :wednesday_public_close,
  :wednesday_public_open,
  keyword_init: true
)

# Image entity data model.
#
# @!attribute [rw] ahash
#   @return [Object, nil]
#
# @!attribute [rw] alt_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] color
#   @return [Object, nil]
#
# @!attribute [rw] colorfulness
#   @return [Object, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] content_e_tag
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] fingerprint
#   @return [Object, nil]
#
# @!attribute [rw] height
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] iiif_url
#   @return [Object, nil]
#
# @!attribute [rw] is_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_teacher_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] lake_guid
#   @return [Object, nil]
#
# @!attribute [rw] lqip
#   @return [Object, nil]
#
# @!attribute [rw] phash
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] width
#   @return [Float, nil]
Image = Struct.new(
  :ahash,
  :alt_text,
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :color,
  :colorfulness,
  :content,
  :content_e_tag,
  :credit_line,
  :fingerprint,
  :height,
  :id,
  :iiif_url,
  :is_educational_resource,
  :is_multimedia_resource,
  :is_teacher_resource,
  :lake_guid,
  :lqip,
  :phash,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :type,
  :updated_at,
  :width,
  keyword_init: true
)

# Request payload for Image#load.
#
# @!attribute [rw] id
#   @return [String]
ImageLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Image#list (any subset of Image fields).
#
# @!attribute [rw] ahash
#   @return [Object, nil]
#
# @!attribute [rw] alt_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] color
#   @return [Object, nil]
#
# @!attribute [rw] colorfulness
#   @return [Object, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] content_e_tag
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] fingerprint
#   @return [Object, nil]
#
# @!attribute [rw] height
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] iiif_url
#   @return [Object, nil]
#
# @!attribute [rw] is_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_teacher_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] lake_guid
#   @return [Object, nil]
#
# @!attribute [rw] lqip
#   @return [Object, nil]
#
# @!attribute [rw] phash
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] width
#   @return [Float, nil]
ImageListMatch = Struct.new(
  :ahash,
  :alt_text,
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :color,
  :colorfulness,
  :content,
  :content_e_tag,
  :credit_line,
  :fingerprint,
  :height,
  :id,
  :iiif_url,
  :is_educational_resource,
  :is_multimedia_resource,
  :is_teacher_resource,
  :lake_guid,
  :lqip,
  :phash,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :type,
  :updated_at,
  :width,
  keyword_init: true
)

# LandingPage entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] search_tag
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
LandingPage = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :search_tag,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for LandingPage#load.
#
# @!attribute [rw] id
#   @return [String]
LandingPageLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for LandingPage#list (any subset of LandingPage fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] search_tag
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
LandingPageListMatch = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :search_tag,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Place entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] tgn_id
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
Place = Struct.new(
  :api_link,
  :api_model,
  :id,
  :latitude,
  :longitude,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :tgn_id,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# Request payload for Place#load.
#
# @!attribute [rw] id
#   @return [String]
PlaceLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Place#list (any subset of Place fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] latitude
#   @return [Float, nil]
#
# @!attribute [rw] longitude
#   @return [Float, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] tgn_id
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
PlaceListMatch = Struct.new(
  :api_link,
  :api_model,
  :id,
  :latitude,
  :longitude,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :tgn_id,
  :timestamp,
  :title,
  :updated_at,
  keyword_init: true
)

# PressRelease entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
PressRelease = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for PressRelease#load.
#
# @!attribute [rw] id
#   @return [String]
PressReleaseLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for PressRelease#list (any subset of PressRelease fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
PressReleaseListMatch = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# PrintedPublication entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
PrintedPublication = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for PrintedPublication#load.
#
# @!attribute [rw] id
#   @return [String]
PrintedPublicationLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for PrintedPublication#list (any subset of PrintedPublication fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] copy
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
PrintedPublicationListMatch = Struct.new(
  :api_link,
  :api_model,
  :copy,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Product entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artist_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] exhibition_id
#   @return [String, nil]
#
# @!attribute [rw] external_sku
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [Object, nil]
#
# @!attribute [rw] max_compare_at_price
#   @return [Object, nil]
#
# @!attribute [rw] max_current_price
#   @return [Object, nil]
#
# @!attribute [rw] min_compare_at_price
#   @return [Object, nil]
#
# @!attribute [rw] min_current_price
#   @return [Object, nil]
#
# @!attribute [rw] price_display
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
Product = Struct.new(
  :api_link,
  :api_model,
  :artist_id,
  :artwork_id,
  :description,
  :exhibition_id,
  :external_sku,
  :id,
  :image_url,
  :max_compare_at_price,
  :max_current_price,
  :min_compare_at_price,
  :min_current_price,
  :price_display,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for Product#load.
#
# @!attribute [rw] id
#   @return [String]
ProductLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Product#list (any subset of Product fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artist_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] exhibition_id
#   @return [String, nil]
#
# @!attribute [rw] external_sku
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [Object, nil]
#
# @!attribute [rw] max_compare_at_price
#   @return [Object, nil]
#
# @!attribute [rw] max_current_price
#   @return [Object, nil]
#
# @!attribute [rw] min_compare_at_price
#   @return [Object, nil]
#
# @!attribute [rw] min_current_price
#   @return [Object, nil]
#
# @!attribute [rw] price_display
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
ProductListMatch = Struct.new(
  :api_link,
  :api_model,
  :artist_id,
  :artwork_id,
  :description,
  :exhibition_id,
  :external_sku,
  :id,
  :image_url,
  :max_compare_at_price,
  :max_current_price,
  :min_compare_at_price,
  :min_current_price,
  :price_display,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Publication entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] section_id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
Publication = Struct.new(
  :api_link,
  :api_model,
  :id,
  :section_id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for Publication#load.
#
# @!attribute [rw] id
#   @return [String]
PublicationLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Publication#list (any subset of Publication fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] section_id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
PublicationListMatch = Struct.new(
  :api_link,
  :api_model,
  :id,
  :section_id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Search entity data model.
#
# @!attribute [rw] api_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_boosted
#   @return [Boolean, nil]
#
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] thumbnail
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
Search = Struct.new(
  :api_id,
  :api_link,
  :api_model,
  :id,
  :is_boosted,
  :score,
  :thumbnail,
  :timestamp,
  :title,
  keyword_init: true
)

# Match filter for Search#list (any subset of Search fields).
#
# @!attribute [rw] api_id
#   @return [String, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_boosted
#   @return [Boolean, nil]
#
# @!attribute [rw] score
#   @return [Float, nil]
#
# @!attribute [rw] thumbnail
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
SearchListMatch = Struct.new(
  :api_id,
  :api_link,
  :api_model,
  :id,
  :is_boosted,
  :score,
  :thumbnail,
  :timestamp,
  :title,
  keyword_init: true
)

# Section entity data model.
#
# @!attribute [rw] accession
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] generic_page_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] publication_id
#   @return [String, nil]
#
# @!attribute [rw] publication_title
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
Section = Struct.new(
  :accession,
  :api_link,
  :api_model,
  :artwork_id,
  :content,
  :generic_page_id,
  :id,
  :publication_id,
  :publication_title,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for Section#load.
#
# @!attribute [rw] id
#   @return [String]
SectionLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Section#list (any subset of Section fields).
#
# @!attribute [rw] accession
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] generic_page_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] publication_id
#   @return [String, nil]
#
# @!attribute [rw] publication_title
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
SectionListMatch = Struct.new(
  :accession,
  :api_link,
  :api_model,
  :artwork_id,
  :content,
  :generic_page_id,
  :id,
  :publication_id,
  :publication_title,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Site entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] exhibition_id
#   @return [String, nil]
#
# @!attribute [rw] exhibition_title
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
Site = Struct.new(
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :description,
  :exhibition_id,
  :exhibition_title,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for Site#load.
#
# @!attribute [rw] id
#   @return [String]
SiteLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Site#list (any subset of Site fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] exhibition_id
#   @return [String, nil]
#
# @!attribute [rw] exhibition_title
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
SiteListMatch = Struct.new(
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :description,
  :exhibition_id,
  :exhibition_title,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Sound entity data model.
#
# @!attribute [rw] alt_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] content_e_tag
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_teacher_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] lake_guid
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] transcript
#   @return [Object, nil]
#
# @!attribute [rw] type
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
Sound = Struct.new(
  :alt_text,
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :content,
  :content_e_tag,
  :credit_line,
  :id,
  :is_educational_resource,
  :is_multimedia_resource,
  :is_teacher_resource,
  :lake_guid,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :transcript,
  :type,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for Sound#load.
#
# @!attribute [rw] id
#   @return [String]
SoundLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Sound#list (any subset of Sound fields).
#
# @!attribute [rw] alt_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] content_e_tag
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_teacher_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] lake_guid
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] transcript
#   @return [Object, nil]
#
# @!attribute [rw] type
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
SoundListMatch = Struct.new(
  :alt_text,
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :content,
  :content_e_tag,
  :credit_line,
  :id,
  :is_educational_resource,
  :is_multimedia_resource,
  :is_teacher_resource,
  :lake_guid,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :transcript,
  :type,
  :updated_at,
  :web_url,
  keyword_init: true
)

# StaticPage entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
StaticPage = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Request payload for StaticPage#load.
#
# @!attribute [rw] id
#   @return [String]
StaticPageLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for StaticPage#list (any subset of StaticPage fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] web_url
#   @return [Object, nil]
StaticPageListMatch = Struct.new(
  :api_link,
  :api_model,
  :id,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :web_url,
  keyword_init: true
)

# Text entity data model.
#
# @!attribute [rw] alt_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] content_e_tag
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_teacher_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] lake_guid
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
Text = Struct.new(
  :alt_text,
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :content,
  :content_e_tag,
  :credit_line,
  :id,
  :is_educational_resource,
  :is_multimedia_resource,
  :is_teacher_resource,
  :lake_guid,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :type,
  :updated_at,
  keyword_init: true
)

# Request payload for Text#load.
#
# @!attribute [rw] id
#   @return [String]
TextLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Text#list (any subset of Text fields).
#
# @!attribute [rw] alt_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] content_e_tag
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_teacher_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] lake_guid
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
TextListMatch = Struct.new(
  :alt_text,
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :content,
  :content_e_tag,
  :credit_line,
  :id,
  :is_educational_resource,
  :is_multimedia_resource,
  :is_teacher_resource,
  :lake_guid,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :type,
  :updated_at,
  keyword_init: true
)

# Tour entity data model.
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artist_title
#   @return [Object, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image
#   @return [Object, nil]
#
# @!attribute [rw] intro
#   @return [Object, nil]
#
# @!attribute [rw] intro_link
#   @return [Object, nil]
#
# @!attribute [rw] intro_transcript
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] weight
#   @return [Float, nil]
Tour = Struct.new(
  :api_link,
  :api_model,
  :artist_title,
  :artwork_title,
  :description,
  :id,
  :image,
  :intro,
  :intro_link,
  :intro_transcript,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :weight,
  keyword_init: true
)

# Request payload for Tour#load.
#
# @!attribute [rw] id
#   @return [String]
TourLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Tour#list (any subset of Tour fields).
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artist_title
#   @return [Object, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] image
#   @return [Object, nil]
#
# @!attribute [rw] intro
#   @return [Object, nil]
#
# @!attribute [rw] intro_link
#   @return [Object, nil]
#
# @!attribute [rw] intro_transcript
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
#
# @!attribute [rw] weight
#   @return [Float, nil]
TourListMatch = Struct.new(
  :api_link,
  :api_model,
  :artist_title,
  :artwork_title,
  :description,
  :id,
  :image,
  :intro,
  :intro_link,
  :intro_transcript,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :updated_at,
  :weight,
  keyword_init: true
)

# Video entity data model.
#
# @!attribute [rw] alt_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] content_e_tag
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_teacher_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] lake_guid
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
Video = Struct.new(
  :alt_text,
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :content,
  :content_e_tag,
  :credit_line,
  :id,
  :is_educational_resource,
  :is_multimedia_resource,
  :is_teacher_resource,
  :lake_guid,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :type,
  :updated_at,
  keyword_init: true
)

# Request payload for Video#load.
#
# @!attribute [rw] id
#   @return [String]
VideoLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Match filter for Video#list (any subset of Video fields).
#
# @!attribute [rw] alt_text
#   @return [Object, nil]
#
# @!attribute [rw] api_link
#   @return [Object, nil]
#
# @!attribute [rw] api_model
#   @return [Object, nil]
#
# @!attribute [rw] artwork_id
#   @return [String, nil]
#
# @!attribute [rw] artwork_title
#   @return [Object, nil]
#
# @!attribute [rw] content
#   @return [Object, nil]
#
# @!attribute [rw] content_e_tag
#   @return [Object, nil]
#
# @!attribute [rw] credit_line
#   @return [Object, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_educational_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_multimedia_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] is_teacher_resource
#   @return [Boolean, nil]
#
# @!attribute [rw] lake_guid
#   @return [Object, nil]
#
# @!attribute [rw] source_updated_at
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_all
#   @return [Object, nil]
#
# @!attribute [rw] suggest_autocomplete_boosted
#   @return [Object, nil]
#
# @!attribute [rw] timestamp
#   @return [Object, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [Object, nil]
#
# @!attribute [rw] updated_at
#   @return [Object, nil]
VideoListMatch = Struct.new(
  :alt_text,
  :api_link,
  :api_model,
  :artwork_id,
  :artwork_title,
  :content,
  :content_e_tag,
  :credit_line,
  :id,
  :is_educational_resource,
  :is_multimedia_resource,
  :is_teacher_resource,
  :lake_guid,
  :source_updated_at,
  :suggest_autocomplete_all,
  :suggest_autocomplete_boosted,
  :timestamp,
  :title,
  :type,
  :updated_at,
  keyword_init: true
)

