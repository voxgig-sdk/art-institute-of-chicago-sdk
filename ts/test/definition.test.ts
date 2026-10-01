import { describe, test } from 'node:test'
import { SDK } from '..'
import { runDefinitionPoint } from './definition-runner'
import { isControlSkipped } from './utility'


// Generated from the API definition, not from the model this SDK was built
// from: the route, the declared query parameters, the credential the security
// scheme names, and the definition's own response example.
const PLAN: any[] = [
  {
    "entity": "agent_role",
    "accessor": "AgentRole",
    "op": "list",
    "method": "GET",
    "path": "/agent-roles",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "agent_role",
    "accessor": "AgentRole",
    "op": "load",
    "method": "GET",
    "path": "/agent-roles/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "agent_type",
    "accessor": "AgentType",
    "op": "list",
    "method": "GET",
    "path": "/agent-types",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "agent_type",
    "accessor": "AgentType",
    "op": "load",
    "method": "GET",
    "path": "/agent-types/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "article",
    "accessor": "Article",
    "op": "list",
    "method": "GET",
    "path": "/articles",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "article",
    "accessor": "Article",
    "op": "load",
    "method": "GET",
    "path": "/articles/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "artwork",
    "accessor": "Artwork",
    "op": "list",
    "method": "GET",
    "path": "/artworks",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "artwork",
    "accessor": "Artwork",
    "op": "load",
    "method": "GET",
    "path": "/artworks/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "artwork_date_qualifier",
    "accessor": "ArtworkDateQualifier",
    "op": "list",
    "method": "GET",
    "path": "/artwork-date-qualifiers",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "artwork_date_qualifier",
    "accessor": "ArtworkDateQualifier",
    "op": "load",
    "method": "GET",
    "path": "/artwork-date-qualifiers/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "artwork_place_qualifier",
    "accessor": "ArtworkPlaceQualifier",
    "op": "list",
    "method": "GET",
    "path": "/artwork-place-qualifiers",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "artwork_place_qualifier",
    "accessor": "ArtworkPlaceQualifier",
    "op": "load",
    "method": "GET",
    "path": "/artwork-place-qualifiers/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "artwork_type",
    "accessor": "ArtworkType",
    "op": "list",
    "method": "GET",
    "path": "/artwork-types",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "artwork_type",
    "accessor": "ArtworkType",
    "op": "load",
    "method": "GET",
    "path": "/artwork-types/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "category_term",
    "accessor": "CategoryTerm",
    "op": "list",
    "method": "GET",
    "path": "/category-terms",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "category_term",
    "accessor": "CategoryTerm",
    "op": "load",
    "method": "GET",
    "path": "/category-terms/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "digital_publication",
    "accessor": "DigitalPublication",
    "op": "list",
    "method": "GET",
    "path": "/digital-publications",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "digital_publication",
    "accessor": "DigitalPublication",
    "op": "load",
    "method": "GET",
    "path": "/digital-publications/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "digital_publication_article",
    "accessor": "DigitalPublicationArticle",
    "op": "list",
    "method": "GET",
    "path": "/digital-publication-articles",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "digital_publication_article",
    "accessor": "DigitalPublicationArticle",
    "op": "load",
    "method": "GET",
    "path": "/digital-publication-articles/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "educator_resource",
    "accessor": "EducatorResource",
    "op": "list",
    "method": "GET",
    "path": "/educator-resources",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "educator_resource",
    "accessor": "EducatorResource",
    "op": "load",
    "method": "GET",
    "path": "/educator-resources/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "event",
    "accessor": "Event",
    "op": "list",
    "method": "GET",
    "path": "/events",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "event",
    "accessor": "Event",
    "op": "load",
    "method": "GET",
    "path": "/events/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "event_occurrence",
    "accessor": "EventOccurrence",
    "op": "list",
    "method": "GET",
    "path": "/event-occurrences",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "event_occurrence",
    "accessor": "EventOccurrence",
    "op": "load",
    "method": "GET",
    "path": "/event-occurrences/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "event_program",
    "accessor": "EventProgram",
    "op": "list",
    "method": "GET",
    "path": "/event-programs",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "event_program",
    "accessor": "EventProgram",
    "op": "load",
    "method": "GET",
    "path": "/event-programs/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "exhibition",
    "accessor": "Exhibition",
    "op": "list",
    "method": "GET",
    "path": "/exhibitions",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "exhibition",
    "accessor": "Exhibition",
    "op": "load",
    "method": "GET",
    "path": "/exhibitions/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "gallery",
    "accessor": "Gallery",
    "op": "list",
    "method": "GET",
    "path": "/galleries",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "gallery",
    "accessor": "Gallery",
    "op": "load",
    "method": "GET",
    "path": "/galleries/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "generic_page",
    "accessor": "GenericPage",
    "op": "list",
    "method": "GET",
    "path": "/generic-pages",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "generic_page",
    "accessor": "GenericPage",
    "op": "load",
    "method": "GET",
    "path": "/generic-pages/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "highlight",
    "accessor": "Highlight",
    "op": "list",
    "method": "GET",
    "path": "/highlights",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "highlight",
    "accessor": "Highlight",
    "op": "load",
    "method": "GET",
    "path": "/highlights/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "hour",
    "accessor": "Hour",
    "op": "list",
    "method": "GET",
    "path": "/hours",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "hour",
    "accessor": "Hour",
    "op": "load",
    "method": "GET",
    "path": "/hours/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "image",
    "accessor": "Image",
    "op": "list",
    "method": "GET",
    "path": "/images",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "image",
    "accessor": "Image",
    "op": "load",
    "method": "GET",
    "path": "/images/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "landing_page",
    "accessor": "LandingPage",
    "op": "list",
    "method": "GET",
    "path": "/landing-pages",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "landing_page",
    "accessor": "LandingPage",
    "op": "load",
    "method": "GET",
    "path": "/landing-pages/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "place",
    "accessor": "Place",
    "op": "list",
    "method": "GET",
    "path": "/places",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "place",
    "accessor": "Place",
    "op": "load",
    "method": "GET",
    "path": "/places/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "press_release",
    "accessor": "PressRelease",
    "op": "list",
    "method": "GET",
    "path": "/press-releases",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "press_release",
    "accessor": "PressRelease",
    "op": "load",
    "method": "GET",
    "path": "/press-releases/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "printed_publication",
    "accessor": "PrintedPublication",
    "op": "list",
    "method": "GET",
    "path": "/printed-publications",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "printed_publication",
    "accessor": "PrintedPublication",
    "op": "load",
    "method": "GET",
    "path": "/printed-publications/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "product",
    "accessor": "Product",
    "op": "list",
    "method": "GET",
    "path": "/products",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "product",
    "accessor": "Product",
    "op": "load",
    "method": "GET",
    "path": "/products/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "publication",
    "accessor": "Publication",
    "op": "list",
    "method": "GET",
    "path": "/publications",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "publication",
    "accessor": "Publication",
    "op": "load",
    "method": "GET",
    "path": "/publications/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "section",
    "accessor": "Section",
    "op": "list",
    "method": "GET",
    "path": "/sections",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "section",
    "accessor": "Section",
    "op": "load",
    "method": "GET",
    "path": "/sections/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "site",
    "accessor": "Site",
    "op": "list",
    "method": "GET",
    "path": "/sites",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "site",
    "accessor": "Site",
    "op": "load",
    "method": "GET",
    "path": "/sites/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "static_page",
    "accessor": "StaticPage",
    "op": "list",
    "method": "GET",
    "path": "/static-pages",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "static_page",
    "accessor": "StaticPage",
    "op": "load",
    "method": "GET",
    "path": "/static-pages/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "text",
    "accessor": "Text",
    "op": "list",
    "method": "GET",
    "path": "/texts",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "text",
    "accessor": "Text",
    "op": "load",
    "method": "GET",
    "path": "/texts/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "tour",
    "accessor": "Tour",
    "op": "list",
    "method": "GET",
    "path": "/tours",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "tour",
    "accessor": "Tour",
    "op": "load",
    "method": "GET",
    "path": "/tours/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "video",
    "accessor": "Video",
    "op": "list",
    "method": "GET",
    "path": "/videos",
    "args": [],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  },
  {
    "entity": "video",
    "accessor": "Video",
    "op": "load",
    "method": "GET",
    "path": "/videos/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "query": [],
    "auth": null,
    "status": 200,
    "sample": [
      {}
    ],
    "idField": "id"
  }
]


describe('definition', () => {
  for (const point of PLAN) {
    test(point.entity + '.' + point.op + ' ' + point.method + ' ' + point.path, async (t) => {
      const control = isControlSkipped('entityOp', point.entity + '.' + point.op, 'definition')
      if (control.skip) {
        t.skip(control.reason || 'skipped via sdk-test-control.json')
        return
      }
      await runDefinitionPoint(SDK, point)
    })
  }
})
