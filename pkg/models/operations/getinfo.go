// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/smartcar/go-sdk-v2/pkg/models/shared"
	"net/http"
)

type GetInfoResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// return User's information
	UserInfo *shared.UserInfo
}
