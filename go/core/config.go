package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "ArtInstituteOfChicago",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://api.artic.edu/api/v1",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"agent": map[string]any{},
				"agent_role": map[string]any{},
				"agent_type": map[string]any{},
				"article": map[string]any{},
				"artwork": map[string]any{},
				"artwork_date_qualifier": map[string]any{},
				"artwork_place_qualifier": map[string]any{},
				"artwork_type": map[string]any{},
				"category_term": map[string]any{},
				"digital_publication": map[string]any{},
				"digital_publication_article": map[string]any{},
				"educator_resource": map[string]any{},
				"event": map[string]any{},
				"event_occurrence": map[string]any{},
				"event_program": map[string]any{},
				"exhibition": map[string]any{},
				"gallery": map[string]any{},
				"generic_page": map[string]any{},
				"highlight": map[string]any{},
				"hour": map[string]any{},
				"image": map[string]any{},
				"landing_page": map[string]any{},
				"place": map[string]any{},
				"press_release": map[string]any{},
				"printed_publication": map[string]any{},
				"product": map[string]any{},
				"publication": map[string]any{},
				"search": map[string]any{},
				"section": map[string]any{},
				"site": map[string]any{},
				"sound": map[string]any{},
				"static_page": map[string]any{},
				"text": map[string]any{},
				"tour": map[string]any{},
				"video": map[string]any{},
			},
		},
		"entity": map[string]any{
			"agent": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "birth_date",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "death_date",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_artist",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sort_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ulan_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "agent",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/agents",
								"parts": []any{
									"agents",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artists",
								"parts": []any{
									"artists",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/agents/{id}",
								"parts": []any{
									"agents",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artists/{id}",
								"parts": []any{
									"artists",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"agent_role": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "agent_role",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/agent-roles",
								"parts": []any{
									"agent-roles",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/agent-roles/{id}",
								"parts": []any{
									"agent-roles",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"agent_type": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "agent_type",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/agent-types",
								"parts": []any{
									"agent-types",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/agent-types/{id}",
								"parts": []any{
									"agent-types",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"article": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "article",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/articles",
								"parts": []any{
									"articles",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/articles/{id}",
								"parts": []any{
									"articles",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"artwork": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_artist_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_classification_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_image_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_material_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_style_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_subject_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_technique_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "artist_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_type_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "artwork_type_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "boost_rank",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "catalog_based_search_keyword_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "catalogue_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "category_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "category_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "classification_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "classification_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "classification_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "classification_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "color",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "colorfulness",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copyright_notice",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_end",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_qualifier_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "date_qualifier_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_start",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "department_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "department_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dimensions",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "dimensions_detail",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "document_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "edition",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "exhibition_history",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "fiscal_year",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "fiscal_year_deaccession",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gallery_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gallery_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "has_advanced_imaging",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "has_educational_resources",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "has_multimedia_resources",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "has_not_been_viewed_much",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_embedding",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "image_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inscriptions",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "internal_department_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_boosted",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_on_view",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_public_domain",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_zoomable",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latitude",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "latlon",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "longitude",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "main_reference_number",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "material_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "material_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "material_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "max_zoom_window_size",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "medium_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "nomisma_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "on_loan_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "pageviews",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "pageviews_recent",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "place_of_origin",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance_text",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "publication_history",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "publishing_verification_level",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "section_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "section_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "short_description",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "site_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sound_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "style_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "style_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "style_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "style_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "subject_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "subject_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "subject_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "technique_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "technique_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "technique_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "term_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "text_embedding",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "text_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "theme_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thumbnail",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "video_ids",
						"type": "`$ANY`",
					},
				},
				"name": "artwork",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artworks",
								"parts": []any{
									"artworks",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artworks/{id}",
								"parts": []any{
									"artworks",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"artwork_date_qualifier": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "artwork_date_qualifier",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-date-qualifiers",
								"parts": []any{
									"artwork-date-qualifiers",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-date-qualifiers/{id}",
								"parts": []any{
									"artwork-date-qualifiers",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"artwork_place_qualifier": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "artwork_place_qualifier",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-place-qualifiers",
								"parts": []any{
									"artwork-place-qualifiers",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-place-qualifiers/{id}",
								"parts": []any{
									"artwork-place-qualifiers",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"artwork_type": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "aat_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "artwork_type",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-types",
								"parts": []any{
									"artwork-types",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artwork-types/{id}",
								"parts": []any{
									"artwork-types",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"category_term": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "aat_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "parent_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "subtype",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "category_term",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/category-terms",
								"parts": []any{
									"category-terms",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/category-terms/{id}",
								"parts": []any{
									"category-terms",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"digital_publication": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "digital_publication",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publications",
								"parts": []any{
									"digital-publications",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publications/{id}",
								"parts": []any{
									"digital-publications",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"digital_publication_article": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "author_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "digital_publication_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "digital_publication_article",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publication-articles",
								"parts": []any{
									"digital-publication-articles",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publication-articles/{id}",
								"parts": []any{
									"digital-publication-articles",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"educator_resource": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "educator_resource",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/educator-resources",
								"parts": []any{
									"educator-resources",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/educator-resources/{id}",
								"parts": []any{
									"educator-resources",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_audience_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_event_type_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "audience_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "buy_button_caption",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "buy_button_text",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "date_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "door_time",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "end_date",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "end_time",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "entrance",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "event_host_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "event_host_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "event_type_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "header_description",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "hero_caption",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "is_admission_required",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_after_hours",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_free",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_member_exclusive",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_private",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_registration_required",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_sales_button_hidden",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_sold_out",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_ticketed",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_virtual_event",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "join_url",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "layout_type",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "list_description",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "location",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "program_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "program_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "rsvp_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "search_tags",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "short_description",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "slug",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "start_date",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "start_time",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "survey_url",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ticketed_event_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "title_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "virtual_event_passcode",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "virtual_event_url",
						"type": "`$ANY`",
					},
				},
				"name": "event",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/events",
								"parts": []any{
									"events",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/events/{id}",
								"parts": []any{
									"events",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event_occurrence": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "button_caption",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "button_text",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "button_url",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "end_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "event_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "is_private",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_sales_button_hidden",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_ticketed",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "location",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "off_sale_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "on_sale_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "short_description",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "start_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "title_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "event_occurrence",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/event-occurrences",
								"parts": []any{
									"event-occurrences",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/event-occurrences/{id}",
								"parts": []any{
									"event-occurrences",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event_program": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_affiliate_group",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_event_host",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "event_program",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/event-programs",
								"parts": []any{
									"event-programs",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/event-programs/{id}",
								"parts": []any{
									"event-programs",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"exhibition": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "aic_end_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "aic_start_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_image_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "document_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gallery_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gallery_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "is_featured",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_published",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "position",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "short_description",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "site_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "status",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "exhibition",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/exhibitions",
								"parts": []any{
									"exhibitions",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/exhibitions/{id}",
								"parts": []any{
									"exhibitions",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"gallery": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "floor",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_closed",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latitude",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "latlon",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "longitude",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "number",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tgn_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "gallery",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/galleries",
								"parts": []any{
									"galleries",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/galleries/{id}",
								"parts": []any{
									"galleries",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"generic_page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "search_tags",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "generic_page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/generic-pages",
								"parts": []any{
									"generic-pages",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/generic-pages/{id}",
								"parts": []any{
									"generic-pages",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"highlight": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "highlight",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/highlights",
								"parts": []any{
									"highlights",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/highlights/{id}",
								"parts": []any{
									"highlights",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"hour": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additional_text",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_is_closed",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_member_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_member_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_public_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "friday_public_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "monday_is_closed",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "monday_member_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "monday_member_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "monday_public_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "monday_public_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_is_closed",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_member_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_member_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_public_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "saturday_public_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "summary",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_is_closed",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_member_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_member_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_public_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "sunday_public_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_is_closed",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_member_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_member_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_public_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "thursday_public_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tuesday_is_closed",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tuesday_member_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tuesday_member_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tuesday_public_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tuesday_public_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_is_closed",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_member_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_member_open",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_public_close",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "wednesday_public_open",
						"type": "`$ANY`",
					},
				},
				"name": "hour",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/hours",
								"parts": []any{
									"hours",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/hours/{id}",
								"parts": []any{
									"hours",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"image": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ahash",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "alt_text",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "color",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "colorfulness",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content_e_tag",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "fingerprint",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "height",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "iiif_url",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "is_educational_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lake_guid",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "lqip",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "phash",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "width",
						"type": "`$NUMBER`",
					},
				},
				"name": "image",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/images",
								"parts": []any{
									"images",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/images/{id}",
								"parts": []any{
									"images",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"landing_page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "search_tags",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "landing_page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/landing-pages",
								"parts": []any{
									"landing-pages",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/landing-pages/{id}",
								"parts": []any{
									"landing-pages",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"place": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "latitude",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "longitude",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tgn_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "place",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/places",
								"parts": []any{
									"places",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/places/{id}",
								"parts": []any{
									"places",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"press_release": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "press_release",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/press-releases",
								"parts": []any{
									"press-releases",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/press-releases/{id}",
								"parts": []any{
									"press-releases",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"printed_publication": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "copy",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "printed_publication",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/printed-publications",
								"parts": []any{
									"printed-publications",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/printed-publications/{id}",
								"parts": []any{
									"printed-publications",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"product": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "exhibition_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "external_sku",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "max_compare_at_price",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "max_current_price",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "min_compare_at_price",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "min_current_price",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "price_display",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "product",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/products",
								"parts": []any{
									"products",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/products/{id}",
								"parts": []any{
									"products",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"publication": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "section_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "publication",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/publications",
								"parts": []any{
									"publications",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/publications/{id}",
								"parts": []any{
									"publications",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_boosted",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "score",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "thumbnail",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
				},
				"name": "search",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/agents/search",
								"parts": []any{
									"agents",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/articles/search",
								"parts": []any{
									"articles",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/artworks/search",
								"parts": []any{
									"artworks",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/category-terms/search",
								"parts": []any{
									"category-terms",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publication-articles/search",
								"parts": []any{
									"digital-publication-articles",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/digital-publications/search",
								"parts": []any{
									"digital-publications",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/educator-resources/search",
								"parts": []any{
									"educator-resources",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/event-occurrences/search",
								"parts": []any{
									"event-occurrences",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/event-programs/search",
								"parts": []any{
									"event-programs",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/events/search",
								"parts": []any{
									"events",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/exhibitions/search",
								"parts": []any{
									"exhibitions",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/galleries/search",
								"parts": []any{
									"galleries",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/generic-pages/search",
								"parts": []any{
									"generic-pages",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/highlights/search",
								"parts": []any{
									"highlights",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/hours/search",
								"parts": []any{
									"hours",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/images/search",
								"parts": []any{
									"images",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/landing-pages/search",
								"parts": []any{
									"landing-pages",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/mobile-sounds/search",
								"parts": []any{
									"mobile-sounds",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/places/search",
								"parts": []any{
									"places",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/press-releases/search",
								"parts": []any{
									"press-releases",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/printed-publications/search",
								"parts": []any{
									"printed-publications",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/products/search",
								"parts": []any{
									"products",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/publications/search",
								"parts": []any{
									"publications",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/search",
								"parts": []any{
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sections/search",
								"parts": []any{
									"sections",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sites/search",
								"parts": []any{
									"sites",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sounds/search",
								"parts": []any{
									"sounds",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/static-pages/search",
								"parts": []any{
									"static-pages",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/texts/search",
								"parts": []any{
									"texts",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/tours/search",
								"parts": []any{
									"tours",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "facet",
											"orig": "facet",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from",
											"orig": "from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/videos/search",
								"parts": []any{
									"videos",
									"search",
								},
								"select": map[string]any{
									"exist": []any{
										"facet",
										"from",
										"q",
										"query",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"section": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accession",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "content",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "generic_page_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "publication_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "publication_title",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "section",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/sections",
								"parts": []any{
									"sections",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sections/{id}",
								"parts": []any{
									"sections",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"site": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "exhibition_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "exhibition_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "site",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/sites",
								"parts": []any{
									"sites",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sites/{id}",
								"parts": []any{
									"sites",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sound": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_text",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content_e_tag",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_educational_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lake_guid",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transcript",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "type",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "sound",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/mobile-sounds",
								"parts": []any{
									"mobile-sounds",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/sounds",
								"parts": []any{
									"sounds",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/mobile-sounds/{id}",
								"parts": []any{
									"mobile-sounds",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/sounds/{id}",
								"parts": []any{
									"sounds",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"static_page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "web_url",
						"type": "`$ANY`",
					},
				},
				"name": "static_page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/static-pages",
								"parts": []any{
									"static-pages",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/static-pages/{id}",
								"parts": []any{
									"static-pages",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"text": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_text",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content_e_tag",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_educational_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lake_guid",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "text",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/texts",
								"parts": []any{
									"texts",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/texts/{id}",
								"parts": []any{
									"texts",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"tour": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artist_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "intro",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "intro_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "intro_transcript",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "weight",
						"type": "`$NUMBER`",
					},
				},
				"name": "tour",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/tours",
								"parts": []any{
									"tours",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/tours/{id}",
								"parts": []any{
									"tours",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"video": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alt_text",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_link",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "api_model",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_ids",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "artwork_titles",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "content_e_tag",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "credit_line",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_educational_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_multimedia_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "is_teacher_resource",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lake_guid",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "source_updated_at",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_all",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "suggest_autocomplete_boosted",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "timestamp",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$ANY`",
					},
				},
				"name": "video",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/videos",
								"parts": []any{
									"videos",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/videos/{id}",
								"parts": []any{
									"videos",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
