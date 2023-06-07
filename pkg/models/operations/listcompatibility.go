// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Smartcar-API/pkg/models/shared"
	"net/http"
)

type ListCompatibilityRequest struct {
	Country *string `queryParam:"style=form,explode=true,name=country"`
	Scope   *string `queryParam:"style=form,explode=true,name=scope"`
	Vin     *string `queryParam:"style=form,explode=true,name=vin"`
}

type ListCompatibilityResponse struct {
	// return Compatibility
	CompatibilityResponse *shared.CompatibilityResponse
	ContentType           string
	StatusCode            int
	RawResponse           *http.Response
}