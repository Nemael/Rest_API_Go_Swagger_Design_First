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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A simple Book management API",
    "title": "Books API",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/book/{id}": {
      "get": {
        "summary": "Return a book from ID",
        "operationId": "getBook",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A book",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/books": {
      "get": {
        "summary": "Lists all books",
        "operationId": "getBooks",
        "responses": {
          "200": {
            "description": "An array of books",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Book"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "summary": "Adds a book",
        "operationId": "addBook",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Book"
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
    "/checkout": {
      "patch": {
        "operationId": "checkoutBook",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/return": {
      "patch": {
        "operationId": "returnBook",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Book": {
      "type": "object",
      "required": [
        "id",
        "title",
        "author",
        "quantity"
      ],
      "properties": {
        "author": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "quantity": {
          "type": "integer",
          "format": "int64"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A simple Book management API",
    "title": "Books API",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/book/{id}": {
      "get": {
        "summary": "Return a book from ID",
        "operationId": "getBook",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A book",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/books": {
      "get": {
        "summary": "Lists all books",
        "operationId": "getBooks",
        "responses": {
          "200": {
            "description": "An array of books",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Book"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "summary": "Adds a book",
        "operationId": "addBook",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Book"
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
    "/checkout": {
      "patch": {
        "operationId": "checkoutBook",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/return": {
      "patch": {
        "operationId": "returnBook",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Book"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Book": {
      "type": "object",
      "required": [
        "id",
        "title",
        "author",
        "quantity"
      ],
      "properties": {
        "author": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "quantity": {
          "type": "integer",
          "format": "int64"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}
