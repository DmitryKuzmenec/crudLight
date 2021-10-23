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
  "swagger": "2.0",
  "info": {
    "title": "API crudLight",
    "version": "1.0"
  },
  "host": "localhost",
  "paths": {
    "/user": {
      "post": {
        "description": "Creates a user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "operationId": "create",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Server error"
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "description": "Returns a user by ID",
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "operationId": "get",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "500": {
            "description": "Server error"
          }
        }
      },
      "delete": {
        "description": "Deletes user by ID",
        "tags": [
          "User"
        ],
        "operationId": "delete",
        "responses": {
          "200": {
            "description": "OK"
          },
          "500": {
            "description": "Server error"
          }
        }
      },
      "patch": {
        "description": "Updates a user by ID",
        "consumes": [
          "aplication/json"
        ],
        "tags": [
          "User"
        ],
        "operationId": "update",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "500": {
            "description": "Server error"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "birth_date": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "API crudLight",
    "version": "1.0"
  },
  "host": "localhost",
  "paths": {
    "/user": {
      "post": {
        "description": "Creates a user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "operationId": "create",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Server error"
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "description": "Returns a user by ID",
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "operationId": "get",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "500": {
            "description": "Server error"
          }
        }
      },
      "delete": {
        "description": "Deletes user by ID",
        "tags": [
          "User"
        ],
        "operationId": "delete",
        "responses": {
          "200": {
            "description": "OK"
          },
          "500": {
            "description": "Server error"
          }
        }
      },
      "patch": {
        "description": "Updates a user by ID",
        "consumes": [
          "aplication/json"
        ],
        "tags": [
          "User"
        ],
        "operationId": "update",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "500": {
            "description": "Server error"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "birth_date": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
}