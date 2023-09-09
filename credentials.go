// Package client provides api functions as a client
// to interact with the infisical api service
package client

import (
	"fmt"

	"github.com/quinton11/go-infisical/models"

	"net/http"

	"github.com/go-resty/resty/v2"
)

func GetServiceTokenDetailsV2(httpClient *resty.Client, config models.ApiConfig) (models.GetServiceTokenDetailsResponse, error) {
	var tokenDetailsResponse models.GetServiceTokenDetailsResponse
	response, err := httpClient.
		R().
		SetResult(&tokenDetailsResponse).
		SetHeader("User-Agent", config.UserAgent).
		Get(fmt.Sprintf("%v/v2/service-token", config.HostApiUrl))

	if err != nil {
		return models.GetServiceTokenDetailsResponse{}, fmt.Errorf("CallGetServiceTokenDetails: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.GetServiceTokenDetailsResponse{}, fmt.Errorf("CallGetServiceTokenDetails: Unsuccessful response: [response=%s]", response)
	}

	return tokenDetailsResponse, nil
}

func GetServiceTokenAccountDetailsV2(httpClient *resty.Client, config models.ApiConfig) (models.ServiceAccountDetailsResponse, error) {
	var serviceAccountDetailsResponse models.ServiceAccountDetailsResponse
	response, err := httpClient.
		R().
		SetResult(&serviceAccountDetailsResponse).
		SetHeader("User-Agent", config.UserAgent).
		Get(fmt.Sprintf("%v/v2/service-accounts/me", config.HostApiUrl))

	if err != nil {
		return models.ServiceAccountDetailsResponse{}, fmt.Errorf("CallGetServiceTokenAccountDetailsV2: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.ServiceAccountDetailsResponse{}, fmt.Errorf("CallGetServiceTokenAccountDetailsV2: Unsuccessful response: [response=%s]", response)
	}

	return serviceAccountDetailsResponse, nil
}

func Login1V2(httpClient *resty.Client, request models.GetLoginOneV2Request, config models.ApiConfig) (models.GetLoginOneV2Response, error) {
	var loginOneV2Response models.GetLoginOneV2Response
	response, err := httpClient.
		R().
		SetResult(&loginOneV2Response).
		SetHeader("User-Agent", config.UserAgent).
		SetBody(request).
		Post(fmt.Sprintf("%v/v3/auth/login1", config.HostApiUrl))

	if err != nil {
		return models.GetLoginOneV2Response{}, fmt.Errorf("CallLogin1V3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.GetLoginOneV2Response{}, fmt.Errorf("CallLogin1V3: Unsuccessful response: [response=%s]", response)
	}

	return loginOneV2Response, nil
}

func VerifyMFAToken(httpClient *resty.Client, request models.VerifyMfaTokenRequest, config models.ApiConfig) (*models.VerifyMfaTokenResponse, *models.VerifyMfaTokenErrorResponse, error) {
	var verifyMfaTokenResponse models.VerifyMfaTokenResponse
	var responseError models.VerifyMfaTokenErrorResponse
	response, err := httpClient.
		R().
		SetResult(&verifyMfaTokenResponse).
		SetHeader("User-Agent", config.UserAgent).
		SetError(&responseError).
		SetBody(request).
		Post(fmt.Sprintf("%v/v2/auth/mfa/verify", config.HostApiUrl))

	cookies := response.Cookies()

	cookieName := "jid"
	var refreshToken *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			refreshToken = cookie
			break
		}
	}

	// When MFA is enabled
	if refreshToken != nil {
		verifyMfaTokenResponse.RefreshToken = refreshToken.Value
	}

	if err != nil {
		return nil, nil, fmt.Errorf("CallVerifyMfaToken: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return nil, &responseError, nil
	}

	return &verifyMfaTokenResponse, nil, nil
}

func Login2V2(httpClient *resty.Client, request models.GetLoginTwoV2Request, config models.ApiConfig) (models.GetLoginTwoV2Response, error) {
	var loginTwoV2Response models.GetLoginTwoV2Response
	response, err := httpClient.
		R().
		SetResult(&loginTwoV2Response).
		SetHeader("User-Agent", config.UserAgent).
		SetBody(request).
		Post(fmt.Sprintf("%v/v3/auth/login2", config.HostApiUrl))

	cookies := response.Cookies()

	cookieName := "jid"
	var refreshToken *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			refreshToken = cookie
			break
		}
	}

	// When MFA is enabled
	if refreshToken != nil {
		loginTwoV2Response.RefreshToken = refreshToken.Value
	}

	if err != nil {
		return models.GetLoginTwoV2Response{}, fmt.Errorf("CallLogin2V3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.GetLoginTwoV2Response{}, fmt.Errorf("CallLogin2V3: Unsuccessful response: [response=%s]", response)
	}

	return loginTwoV2Response, nil
}

// Checks if the attached JWToken is still valid.
// Set the http client's Auth Token with the user's JWT before passing
func IsAuthenticated(httpClient *resty.Client, config models.ApiConfig) bool {
	var workSpacesResponse models.GetWorkSpacesResponse
	response, err := httpClient.
		R().
		SetResult(&workSpacesResponse).
		SetHeader("User-Agent", config.UserAgent).
		Post(fmt.Sprintf("%v/v1/auth/checkAuth", config.HostApiUrl))

	if err != nil {
		return false
	}

	if response.IsError() {
		fmt.Printf("\nCallIsAuthenticated: Unsuccessful response:  [response=%v]", response)
		return false
	}
	return true
}

func GetNewAccessTokenWithRefreshToken(httpClient *resty.Client, refreshToken string, config models.ApiConfig) (models.GetNewAccessTokenWithRefreshTokenResponse, error) {
	var newAccessToken models.GetNewAccessTokenWithRefreshTokenResponse
	response, err := httpClient.
		R().
		SetResult(&newAccessToken).
		SetHeader("User-Agent", config.UserAgent).
		SetCookie(&http.Cookie{
			Name:  "jid",
			Value: refreshToken,
		}).
		Post(fmt.Sprintf("%v/v1/auth/token", config.HostApiUrl))

	if err != nil {
		return models.GetNewAccessTokenWithRefreshTokenResponse{}, err
	}

	if response.IsError() {
		return models.GetNewAccessTokenWithRefreshTokenResponse{}, fmt.Errorf("CallGetNewAccessTokenWithRefreshToken: Unsuccessful response:  [response=%v]", response)
	}

	return newAccessToken, nil
}

func GetServiceAccountKeysV2(httpClient *resty.Client, request models.GetServiceAccountKeysRequest, config models.ApiConfig) (models.GetServiceAccountKeysResponse, error) {
	var serviceAccountKeysResponse models.GetServiceAccountKeysResponse
	response, err := httpClient.
		R().
		SetResult(&serviceAccountKeysResponse).
		SetHeader("User-Agent", config.UserAgent).
		Get(fmt.Sprintf("%v/v2/service-accounts/%v/keys", config.HostApiUrl, request.ServiceAccountId))

	if err != nil {
		return models.GetServiceAccountKeysResponse{}, fmt.Errorf("CallGetServiceAccountKeysV2: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.GetServiceAccountKeysResponse{}, fmt.Errorf("CallGetServiceAccountKeysV2: Unsuccessful response: [response=%s]", response)
	}

	return serviceAccountKeysResponse, nil
}
