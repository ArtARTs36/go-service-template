// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Example service",
    "title": "Example service",
    "version": "0.0.1"
  },
  "host": "localhost",
  "basePath": "/",
  "paths": {
    "/cars/{id}": {
      "get": {
        "description": "Get Car",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/car-get-response"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "car-get-response": {
      "description": "Car",
      "properties": {
        "id": {
          "type": "integer"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Example service",
    "title": "Example service",
    "version": "0.0.1"
  },
  "host": "localhost",
  "basePath": "/",
  "paths": {
    "/cars/{id}": {
      "get": {
        "description": "Get Car",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/car-get-response"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "car-get-response": {
      "description": "Car",
      "properties": {
        "id": {
          "type": "integer"
        }
      }
    }
  }
}`))
}
