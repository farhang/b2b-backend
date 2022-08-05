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
        "/assets/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "asset"
                ],
                "summary": "Get user information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login with email",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login response model including access token",
                        "schema": {
                            "$ref": "#/definitions/domain.LoginResponseDTO"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "Registration data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RegisterRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/auth/reset-password": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ResetPasswordRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login response model including access token",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/emails/{email}/send-verification-code": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "summary": "Send verification code to user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Email verification code",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.ResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/emails/{email}/verify": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "email"
                ],
                "summary": "Send verification code to user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Email verification data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.VerifyRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "true",
                        "schema": {
                            "type": "bool"
                        }
                    }
                }
            }
        },
        "/plans/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plan"
                ],
                "summary": "Get plans",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plan"
                ],
                "summary": "Add new plan",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.PlanStoreRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login response model including access token",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/profiles/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "get profiles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/profiles/me/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "get my profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/profiles/{id}/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "get profile by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Profile id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "update profile",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Profile id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Profile",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateProfileRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Email verification code",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/transactions/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Get user information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/transactions/deposit": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "create deposit transaction",
                "parameters": [
                    {
                        "description": "Registration data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.DepositRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/transactions/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Get user information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/transactions/profit": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "create withdraw transaction",
                "parameters": [
                    {
                        "description": "Withdraw data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ProfitRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/transactions/withdraw": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "create withdraw transaction",
                "parameters": [
                    {
                        "description": "Withdraw data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.WithDrawRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/users/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Add new user",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.ResponseDTO"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.UserResponseDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users/{id}/transactions/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.ResponseDTO": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.DepositRequestDTO": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.LoginRequestDTO": {
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
        "domain.LoginResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/domain.LoginResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.LoginResponseData": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "domain.PlanStoreRequestDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "profitPercent": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "domain.ProfitRequestDTO": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.RegisterRequestDTO": {
            "type": "object",
            "required": [
                "confirm_password",
                "email"
            ],
            "properties": {
                "company_name": {
                    "type": "string"
                },
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "mobile_number": {
                    "type": "string"
                },
                "mobile_number_company": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                }
            }
        },
        "domain.ResetPasswordRequestDTO": {
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
        "domain.UpdateProfileRequestDTO": {
            "type": "object",
            "properties": {
                "company_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "mobile_number": {
                    "type": "string"
                },
                "mobile_number_company": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "plan_id": {
                    "type": "integer"
                },
                "position": {
                    "type": "string"
                }
            }
        },
        "domain.UserResponseDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "domain.VerifyRequestDTO": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                }
            }
        },
        "domain.WithDrawRequestDTO": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Turkey Exchange API Documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
