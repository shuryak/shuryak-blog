// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/articles/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates an article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Creates an article",
                "parameters": [
                    {
                        "description": "article to create",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        },
        "/articles/delete": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes article by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Deletes article by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID to delete",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.DeleteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        },
        "/articles/getById": {
            "get": {
                "description": "Gets article by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Gets article by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID to get",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.GetByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        },
        "/articles/getMany": {
            "get": {
                "description": "Gets collection of articles",
                "produces": [
                    "application/json"
                ],
                "summary": "Gets collection of articles",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset to get",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "count to get",
                        "name": "count",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.GetManyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        },
        "/articles/update": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updates article by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Updates article by ID",
                "parameters": [
                    {
                        "description": "article to update",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.UpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "get": {
                "description": "Gets access and refresh tokens",
                "produces": [
                    "application/json"
                ],
                "summary": "Method to login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "get": {
                "description": "User registration",
                "produces": [
                    "application/json"
                ],
                "summary": "Method to register",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/errors.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.Response": {
            "type": "object",
            "properties": {
                "error": {}
            }
        },
        "pb.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "pb.RegisterResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "routes.CreateRequest": {
            "type": "object",
            "required": [
                "content",
                "custom_id",
                "thumbnail",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "object",
                    "additionalProperties": true
                },
                "custom_id": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3,
                    "example": "article-url"
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://smth.com/thumbnail.png"
                },
                "title": {
                    "type": "string",
                    "maxLength": 150,
                    "minLength": 5,
                    "example": "How to ..."
                }
            }
        },
        "routes.CreateResponse": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "integer",
                    "example": 42
                },
                "content": {
                    "type": "object",
                    "additionalProperties": true
                },
                "created_at": {
                    "type": "string",
                    "example": "2022-10-07T14:26:06.510465Z"
                },
                "custom_id": {
                    "type": "string",
                    "example": "article-url"
                },
                "id": {
                    "type": "integer",
                    "example": 1000
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://smth.com/thumbnail.png"
                },
                "title": {
                    "type": "string",
                    "example": "How to ..."
                }
            }
        },
        "routes.DeleteResponse": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "integer",
                    "example": 42
                },
                "content": {
                    "type": "object",
                    "additionalProperties": true
                },
                "created_at": {
                    "type": "string",
                    "example": "2022-10-07T14:26:06.510465Z"
                },
                "custom_id": {
                    "type": "string",
                    "example": "article-url"
                },
                "id": {
                    "type": "integer",
                    "example": 1000
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://smth.com/thumbnail.png"
                },
                "title": {
                    "type": "string",
                    "example": "How to ..."
                }
            }
        },
        "routes.GetByIdResponse": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "integer",
                    "example": 42
                },
                "content": {
                    "type": "object",
                    "additionalProperties": true
                },
                "created_at": {
                    "type": "string",
                    "example": "2022-10-07T14:26:06.510465Z"
                },
                "custom_id": {
                    "type": "string",
                    "example": "article-url"
                },
                "id": {
                    "type": "integer",
                    "example": 1000
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://smth.com/thumbnail.png"
                },
                "title": {
                    "type": "string",
                    "example": "How to ..."
                }
            }
        },
        "routes.GetManyResponse": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/routes.GetByIdResponse"
                    }
                }
            }
        },
        "routes.UpdateRequest": {
            "type": "object",
            "required": [
                "content",
                "custom_id",
                "id",
                "thumbnail",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "object",
                    "additionalProperties": true
                },
                "custom_id": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3,
                    "example": "article-url"
                },
                "id": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1000
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://smth.com/thumbnail.png"
                },
                "title": {
                    "type": "string",
                    "maxLength": 150,
                    "minLength": 5,
                    "example": "How to ..."
                }
            }
        },
        "routes.UpdateResponse": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "integer",
                    "example": 42
                },
                "content": {
                    "type": "object",
                    "additionalProperties": true
                },
                "created_at": {
                    "type": "string",
                    "example": "2022-10-07T14:26:06.510465Z"
                },
                "custom_id": {
                    "type": "string",
                    "example": "article-url"
                },
                "id": {
                    "type": "integer",
                    "example": 1000
                },
                "thumbnail": {
                    "type": "string",
                    "example": "https://smth.com/thumbnail.png"
                },
                "title": {
                    "type": "string",
                    "example": "How to ..."
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
