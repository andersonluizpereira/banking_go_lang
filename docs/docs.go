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
        "/v1/clients": {
            "get": {
                "description": "Retorna uma lista de todos os clientes cadastrados",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "Lista todos os clientes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Client"
                            }
                        }
                    },
                    "500": {
                        "description": "Mensagem de erro",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "Cria um novo cliente com as informações fornecidas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "Cria um novo cliente",
                "parameters": [
                    {
                        "description": "Cliente",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    },
                    "400": {
                        "description": "Mensagem de erro",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/clients/{accountNum}": {
            "get": {
                "description": "Busca um cliente pelo número da conta fornecido",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "Busca cliente por número da conta",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Número da conta",
                        "name": "accountNum",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Client"
                        }
                    },
                    "404": {
                        "description": "client not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Client": {
            "type": "object",
            "properties": {
                "account_num": {
                    "type": "string"
                },
                "balance": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
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
