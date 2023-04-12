// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/comment/": {
            "get": {
                "description": "Get All Comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Get All Comment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Comment with the given information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Create a new Comment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Comment has been created",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateComments"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/comment/{id}": {
            "get": {
                "description": "Get One Comment by ID Comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Get One Comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Comment by ID Comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Update Comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Comment has been Updated",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateComments"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Comment by ID Comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Delete Comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/photo/": {
            "get": {
                "description": "Get All Photo by ID Photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Get All Photo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Photo with the given information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Create a new Photo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The social media to create",
                        "name": "socialmedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreatePhotos"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/photo/{id}": {
            "get": {
                "description": "Get One Photo by ID Photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Get One Photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Photo with the given information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Update Photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The social media to create",
                        "name": "socialmedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdatePhotos"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Photo by ID Photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photo"
                ],
                "summary": "Delete Photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/socialmedia": {
            "get": {
                "description": "Get All Social Media Which Has Been Made",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Get All Social Media",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SocialMedia"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/socialmedia/": {
            "post": {
                "description": "CreateSocialMedia a new socialmedia with the given information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Create a new Social Media",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token} Ex : 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The social media to create",
                        "name": "socialmedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateSocial"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/socialmedia/{id}": {
            "get": {
                "description": "Get One Social Media by ID Social Media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Get One Social Media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Social Media by ID Social Media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Update Social Media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The Social Media Has Been Updated",
                        "name": "socialmedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateSocial"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Social Media by ID Social Media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Delete Social Media",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Authenticates a user with their email and password and returns a token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": "Login for authentication",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register a new user with the given information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "The user to create",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CreateComments": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "photo_id": {
                    "type": "integer"
                }
            }
        },
        "controllers.CreatePhotos": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "controllers.CreateSocial": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.RegisterRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.SuccessResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "controllers.UpdateComments": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "photo_id": {
                    "type": "integer"
                }
            }
        },
        "controllers.UpdatePhotos": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "controllers.UpdateSocial": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                }
            }
        },
        "helpers.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Comment": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "photo_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Photo": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.SocialMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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