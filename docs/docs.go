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
        "/v1/users": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "{ \"code\": \"INVALID_ENTITY\", \"message\": \"Invalid\tentity\" },{ \"code\": \"USER_ALREADY_EXISTS\", \"message\": \"Already\texists\tuser\" }",
                        "schema": {
                            "$ref": "#/definitions/exception.ApplicationException"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}": {
            "get": {
                "description": "Get information about a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    },
                    "400": {
                        "description": "{ \"code\": \"USER_NOT_FOUND\", \"message\": \"User\tnot\tfound\" }",
                        "schema": {
                            "$ref": "#/definitions/exception.ApplicationException"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "{ \"code\": \"USER_NOT_FOUND\", \"message\": \"User\tnot\tfound\" }",
                        "schema": {
                            "$ref": "#/definitions/exception.ApplicationException"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.RegisterUser": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 0
                },
                "name": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 8
                }
            }
        },
        "exception.ApplicationException": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
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
