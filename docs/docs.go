// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/user/register": {
            "post": {
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Custom UserRequest request",
                        "name": "UserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/api/v1/user/status": {
            "put": {
                "description": "Get User informations from oauth server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User informations from oauth server",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.UserStatusUpdate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/api/v1/user/{userid}": {
            "delete": {
                "description": "Get User informations from oauth server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User informations from oauth server",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.UserDetailResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/api/v1/user/{userid}/detail": {
            "get": {
                "description": "Get User informations from oauth server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User informations from oauth server",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.UserDetailResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.User": {
            "type": "object",
            "properties": {
                "createdOn": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "isDeleted": {
                    "type": "boolean"
                },
                "modifiedOn": {
                    "type": "string"
                },
                "profilePicture": {
                    "type": "string"
                },
                "reOC": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                },
                "userRegisterId": {
                    "type": "string"
                },
                "userType": {
                    "type": "string"
                }
            }
        },
        "user.UserDetailResponse": {
            "type": "object",
            "properties": {
                "ansp_id": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "created_date": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "organization_id": {
                    "type": "string"
                },
                "organization_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "profile_picture": {
                    "type": "string"
                },
                "reoc": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "user_roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.UserRole"
                    }
                },
                "user_sora_role": {
                    "type": "string"
                },
                "user_type": {
                    "type": "string"
                }
            }
        },
        "user.UserRequest": {
            "type": "object",
            "required": [
                "email",
                "user_role",
                "user_type"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "reoc": {
                    "type": "string"
                },
                "user_role": {
                    "type": "string"
                },
                "user_type": {
                    "type": "string"
                }
            }
        },
        "user.UserRole": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.UserStatusUpdate": {
            "type": "object",
            "properties": {
                "active": {
                    "description": "true for activation, false for deactivation",
                    "type": "boolean"
                },
                "country": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
