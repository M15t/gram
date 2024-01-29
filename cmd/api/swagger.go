//go:build swagger

// Gram - The sword is only as powerful as the person wielding it.
//
// `Inspired by the legendary Norse sword Gram, this Golang project wields the power
// of simplicity and precision. A versatile tool for crafting robust and efficient
// applications with mythical coding prowess.`
//
// ## Authentication
// All API endpoints with the lock icon require authentication token.
// Firstly, grab the **access_token** from the response of `/auth/login`. Then include this header in all API calls:
//
// ```
// Authorization: Bearer ${access_token}
// ```
//
// Terms Of Service: N/A
//
// Version: %{VERSION}
// Contact: m15t <khanhnguyen1411@gmail.com>
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Security:
// - bearer: []
//
// SecurityDefinitions:
// bearer:
//	 type: apiKey
//	 name: Authorization
//	 in: header
//
// swagger:meta
package main

import (
	"embed"
)

//go:embed swagger-ui
var embedSwaggerui embed.FS

func init() {
	enableSwagger = true
	swaggerui = embedSwaggerui
}

// OK - No Content
// swagger:response ok
type swaggOKResp struct{}
