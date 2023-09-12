// Package client provides api functions as a client
// to interact with the infisical api service
package api

import (
	"fmt"

	"github.com/quinton11/go-client/api/models"

	"net/http"

	"github.com/go-resty/resty/v2"
)

type ClientConfig struct {
	UserAgent string `json:"UserAgent"`
	BaseUrl   string `json:"BaseUrl"`
}

func NewClient(userAgent string, hostUrl string) *ClientConfig {
	return &ClientConfig{UserAgent: userAgent, BaseUrl: hostUrl}
}

func (c *ClientConfig) GetServiceTokenDetailsV2(httpClient *resty.Client) (models.GetServiceTokenDetailsResponse, error) {
	var tokenDetailsResponse models.GetServiceTokenDetailsResponse
	response, err := httpClient.
		R().
		SetResult(&tokenDetailsResponse).
		SetHeader("User-Agent", c.UserAgent).
		Get(fmt.Sprintf("%v/v2/service-token", c.BaseUrl))

	if err != nil {
		return models.GetServiceTokenDetailsResponse{}, fmt.Errorf("CallGetServiceTokenDetails: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.GetServiceTokenDetailsResponse{}, fmt.Errorf("CallGetServiceTokenDetails: Unsuccessful response: [response=%s]", response)
	}

	return tokenDetailsResponse, nil
}

func (c *ClientConfig) GetServiceTokenAccountDetailsV2(httpClient *resty.Client) (models.ServiceAccountDetailsResponse, error) {
	var serviceAccountDetailsResponse models.ServiceAccountDetailsResponse
	response, err := httpClient.
		R().
		SetResult(&serviceAccountDetailsResponse).
		SetHeader("User-Agent", c.UserAgent).
		Get(fmt.Sprintf("%v/v2/service-accounts/me", c.BaseUrl))

	if err != nil {
		return models.ServiceAccountDetailsResponse{}, fmt.Errorf("CallGetServiceTokenAccountDetailsV2: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.ServiceAccountDetailsResponse{}, fmt.Errorf("CallGetServiceTokenAccountDetailsV2: Unsuccessful response: [response=%s]", response)
	}

	return serviceAccountDetailsResponse, nil
}

func (c *ClientConfig) Login1V2(httpClient *resty.Client, request models.GetLoginOneV2Request) (models.GetLoginOneV2Response, error) {
	var loginOneV2Response models.GetLoginOneV2Response
	response, err := httpClient.
		R().
		SetResult(&loginOneV2Response).
		SetHeader("User-Agent", c.UserAgent).
		SetBody(request).
		Post(fmt.Sprintf("%v/v3/auth/login1", c.BaseUrl))

	if err != nil {
		return models.GetLoginOneV2Response{}, fmt.Errorf("CallLogin1V3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.GetLoginOneV2Response{}, fmt.Errorf("CallLogin1V3: Unsuccessful response: [response=%s]", response)
	}

	return loginOneV2Response, nil
}

func (c *ClientConfig) VerifyMFAToken(httpClient *resty.Client, request models.VerifyMfaTokenRequest) (*models.VerifyMfaTokenResponse, *models.VerifyMfaTokenErrorResponse, error) {
	var verifyMfaTokenResponse models.VerifyMfaTokenResponse
	var responseError models.VerifyMfaTokenErrorResponse
	response, err := httpClient.
		R().
		SetResult(&verifyMfaTokenResponse).
		SetHeader("User-Agent", c.UserAgent).
		SetError(&responseError).
		SetBody(request).
		Post(fmt.Sprintf("%v/v2/auth/mfa/verify", c.BaseUrl))

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

func (c *ClientConfig) Login2V2(httpClient *resty.Client, request models.GetLoginTwoV2Request) (models.GetLoginTwoV2Response, error) {
	var loginTwoV2Response models.GetLoginTwoV2Response
	response, err := httpClient.
		R().
		SetResult(&loginTwoV2Response).
		SetHeader("User-Agent", c.UserAgent).
		SetBody(request).
		Post(fmt.Sprintf("%v/v3/auth/login2", c.BaseUrl))

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
func (c *ClientConfig) IsAuthenticated(httpClient *resty.Client) bool {
	var workSpacesResponse models.GetWorkSpacesResponse
	response, err := httpClient.
		R().
		SetResult(&workSpacesResponse).
		SetHeader("User-Agent", c.UserAgent).
		Post(fmt.Sprintf("%v/v1/auth/checkAuth", c.BaseUrl))

	if err != nil {
		return false
	}

	if response.IsError() {
		fmt.Printf("\nCallIsAuthenticated: Unsuccessful response:  [response=%v]", response)
		return false
	}
	return true
}

func (c *ClientConfig) GetNewAccessTokenWithRefreshToken(httpClient *resty.Client, refreshToken string) (models.GetNewAccessTokenWithRefreshTokenResponse, error) {
	var newAccessToken models.GetNewAccessTokenWithRefreshTokenResponse
	response, err := httpClient.
		R().
		SetResult(&newAccessToken).
		SetHeader("User-Agent", c.UserAgent).
		SetCookie(&http.Cookie{
			Name:  "jid",
			Value: refreshToken,
		}).
		Post(fmt.Sprintf("%v/v1/auth/token", c.BaseUrl))

	if err != nil {
		return models.GetNewAccessTokenWithRefreshTokenResponse{}, err
	}

	if response.IsError() {
		return models.GetNewAccessTokenWithRefreshTokenResponse{}, fmt.Errorf("CallGetNewAccessTokenWithRefreshToken: Unsuccessful response:  [response=%v]", response)
	}

	return newAccessToken, nil
}

func (c *ClientConfig) GetServiceAccountKeysV2(httpClient *resty.Client, request models.GetServiceAccountKeysRequest) (models.GetServiceAccountKeysResponse, error) {
	var serviceAccountKeysResponse models.GetServiceAccountKeysResponse
	response, err := httpClient.
		R().
		SetResult(&serviceAccountKeysResponse).
		SetHeader("User-Agent", c.UserAgent).
		Get(fmt.Sprintf("%v/v2/service-accounts/%v/keys", c.BaseUrl, request.ServiceAccountId))

	if err != nil {
		return models.GetServiceAccountKeysResponse{}, fmt.Errorf("CallGetServiceAccountKeysV2: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.GetServiceAccountKeysResponse{}, fmt.Errorf("CallGetServiceAccountKeysV2: Unsuccessful response: [response=%s]", response)
	}

	return serviceAccountKeysResponse, nil
}
