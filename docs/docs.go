// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Frens Repo",
            "url": "http://www.github.com/bwoff11/frens"
        },
        "license": {
            "name": "MIT License",
            "url": "http://www.github.com/bwoff11/frens/docs/LICENSE.md"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bookmarks/{bookmarkId}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the details of a specific bookmark based on the provided ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookmarks"
                ],
                "summary": "Retrieve a bookmark by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bookmark ID",
                        "name": "bookmarkId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "frens.moe",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Frens API",
	Description:      "ActivityPub social network API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
