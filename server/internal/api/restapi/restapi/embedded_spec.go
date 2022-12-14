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
    "description": "test grpc server",
    "title": "test-server",
    "version": "1.0.0"
  },
  "paths": {
    "/create-api-key": {
      "post": {
        "operationId": "createApiKey",
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "$ref": "#/definitions/uuid"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/send-message": {
      "post": {
        "operationId": "sendMessage",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/sendMessage"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ok"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "sendMessage": {
      "description": "send message object",
      "type": "object",
      "required": [
        "uuid"
      ],
      "properties": {
        "message": {
          "type": "string"
        },
        "uuid": {
          "type": "string"
        }
      }
    },
    "uuid": {
      "description": "uuid object",
      "type": "object",
      "required": [
        "uuid"
      ],
      "properties": {
        "uuid": {
          "type": "string"
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
    "description": "test grpc server",
    "title": "test-server",
    "version": "1.0.0"
  },
  "paths": {
    "/create-api-key": {
      "post": {
        "operationId": "createApiKey",
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "$ref": "#/definitions/uuid"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/send-message": {
      "post": {
        "operationId": "sendMessage",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/sendMessage"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ok"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "sendMessage": {
      "description": "send message object",
      "type": "object",
      "required": [
        "uuid"
      ],
      "properties": {
        "message": {
          "type": "string"
        },
        "uuid": {
          "type": "string"
        }
      }
    },
    "uuid": {
      "description": "uuid object",
      "type": "object",
      "required": [
        "uuid"
      ],
      "properties": {
        "uuid": {
          "type": "string"
        }
      }
    }
  }
}`))
}
