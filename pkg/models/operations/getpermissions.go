// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"smartcar/pkg/models/shared"
)

type GetPermissionsRequest struct {
	// Number of vehicles to return
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Index to start vehicle list at
	Offset    *int64 `queryParam:"style=form,explode=true,name=offset"`
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetPermissionsResponse struct {
	ContentType string
	// The applications permissions
	Permission  *shared.Permission
	StatusCode  int
	RawResponse *http.Response
}
