// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/EyeOfCode",
        "contact": {
            "name": "API Support",
            "email": "champuplove@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/user/{id}": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the API's update user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User update details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the API's delete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Delete endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/admin/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get paginated list of users with optional filtering",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "List users",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Post the API's login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login endpoint",
                "parameters": [
                    {
                        "description": "User login",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/logout": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Post the API's logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Refresh token",
                        "name": "X-Refresh-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/auth/refresh": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Post the API's refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh endpoint",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/register": {
            "post": {
                "description": "Post the API's register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register endpoint",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/category": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Post the API's create category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Create Category endpoint",
                "parameters": [
                    {
                        "description": "Category details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/category/list": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the API's get all categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get All Categories endpoint",
                "responses": {}
            }
        },
        "/category/{id}": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the API's delete category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Delete Category endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/file/shop/{shop_id}/download/{file_id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the API's download file store",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "file-store"
                ],
                "summary": "Download File Store endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shop ID",
                        "name": "shop_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "File Store ID",
                        "name": "file_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/other/example/gallery": {
            "get": {
                "description": "Get the API's get Gallery",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "other"
                ],
                "summary": "Get Gallery endpoint",
                "responses": {}
            }
        },
        "/shop": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Post the API's create shop",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "Create Shop endpoint",
                "parameters": [
                    {
                        "maxLength": 30,
                        "minLength": 3,
                        "type": "string",
                        "description": "Shop name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Shop budget",
                        "name": "budget",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "csv",
                        "description": "Multiple files to upload",
                        "name": "files",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/shop/list": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get paginated list of shops with optional filtering",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "List shops",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/shop/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the API's get shop",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "Get Shop endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shop ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the API's update shop",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "Update Shop endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shop ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 30,
                        "minLength": 3,
                        "type": "string",
                        "description": "Shop name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Shop budget",
                        "name": "budget",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "csv",
                        "description": "Multiple files to upload",
                        "name": "files",
                        "in": "formData"
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the API's delete shop",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "Delete Shop endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shop ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/profile": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the API's get profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Profile endpoint",
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.CategoryRequest": {
            "type": "object",
            "required": [
                "name",
                "shop_id"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                },
                "shop_id": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "name",
                "password"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Enter the token with the ` + "`" + `Bearer: ` + "`" + ` prefix, e.g. \"Bearer abcde12345\".",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "Example Go Fiber Project API",
	Description:      "A RESTful API server with user authentication and MongoDB integration",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
