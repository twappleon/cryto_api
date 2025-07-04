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
        "/eth/balance": {
            "post": {
                "description": "Get the balance of native tokens (ETH/TRX) for an address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Get native token balance",
                "parameters": [
                    {
                        "description": "Balance query details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.BalanceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.BalanceResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/eth/connect": {
            "post": {
                "description": "Connect to a blockchain node using the provided URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Connect to blockchain node",
                "parameters": [
                    {
                        "description": "Connection details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ConnectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response"
                        }
                    }
                }
            }
        },
        "/eth/contract/deploy": {
            "post": {
                "description": "Deploy a new smart contract to the blockchain",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Deploy smart contract",
                "parameters": [
                    {
                        "description": "Contract deployment details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ContractDeployRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.ContractResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/eth/transfer/native": {
            "post": {
                "description": "Send native tokens (ETH/TRX) to an address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Send native tokens",
                "parameters": [
                    {
                        "description": "Transfer details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.TransferRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.TransactionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/eth/wallet/generate": {
            "post": {
                "description": "Generate a new blockchain wallet",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Generate new wallet",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.WalletResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/tron/balance": {
            "post": {
                "description": "Get the balance of native tokens (ETH/TRX) for an address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Get native token balance",
                "parameters": [
                    {
                        "description": "Balance query details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.BalanceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.BalanceResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/tron/connect": {
            "post": {
                "description": "Connect to a blockchain node using the provided URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Connect to blockchain node",
                "parameters": [
                    {
                        "description": "Connection details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ConnectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response"
                        }
                    }
                }
            }
        },
        "/tron/contract/deploy": {
            "post": {
                "description": "Deploy a new smart contract to the blockchain",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Deploy smart contract",
                "parameters": [
                    {
                        "description": "Contract deployment details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ContractDeployRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.ContractResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/tron/transfer/native": {
            "post": {
                "description": "Send native tokens (ETH/TRX) to an address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Send native tokens",
                "parameters": [
                    {
                        "description": "Transfer details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.TransferRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.TransactionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/tron/wallet/generate": {
            "post": {
                "description": "Generate a new blockchain wallet",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum",
                    "tron"
                ],
                "summary": "Generate new wallet",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.WalletResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.BalanceRequest": {
            "type": "object",
            "required": [
                "address"
            ],
            "properties": {
                "address": {
                    "description": "查詢地址",
                    "type": "string"
                }
            }
        },
        "types.BalanceResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "查詢地址",
                    "type": "string"
                },
                "balance": {
                    "description": "餘額",
                    "type": "number"
                }
            }
        },
        "types.ConnectRequest": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "description": "節點 URL",
                    "type": "string"
                }
            }
        },
        "types.ContractDeployRequest": {
            "type": "object",
            "required": [
                "abi",
                "bytecode",
                "private_key"
            ],
            "properties": {
                "abi": {
                    "description": "合約 ABI",
                    "type": "string"
                },
                "bytecode": {
                    "description": "合約 bytecode",
                    "type": "string"
                },
                "constructor_args": {
                    "description": "建構子參數",
                    "type": "array",
                    "items": {}
                },
                "private_key": {
                    "description": "部署者私鑰",
                    "type": "string"
                }
            }
        },
        "types.ContractResponse": {
            "type": "object",
            "properties": {
                "contract_address": {
                    "description": "合約地址",
                    "type": "string"
                },
                "result": {
                    "description": "執行結果"
                },
                "tx_hash": {
                    "description": "交易雜湊",
                    "type": "string"
                }
            }
        },
        "types.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "狀態碼",
                    "type": "integer"
                },
                "data": {
                    "description": "回傳資料"
                },
                "message": {
                    "description": "訊息",
                    "type": "string"
                }
            }
        },
        "types.TransactionResponse": {
            "type": "object",
            "properties": {
                "tx_hash": {
                    "description": "交易雜湊",
                    "type": "string"
                }
            }
        },
        "types.TransferRequest": {
            "type": "object",
            "required": [
                "amount",
                "from_private_key",
                "to_address"
            ],
            "properties": {
                "amount": {
                    "description": "轉帳金額",
                    "type": "number"
                },
                "from_private_key": {
                    "description": "發送方私鑰",
                    "type": "string"
                },
                "to_address": {
                    "description": "接收方地址",
                    "type": "string"
                }
            }
        },
        "types.WalletResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "錢包地址",
                    "type": "string"
                },
                "private_key": {
                    "description": "私鑰",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Blockchain SDK API",
	Description:      "API for interacting with Ethereum and Tron blockchains",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
