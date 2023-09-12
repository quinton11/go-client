package api

import (
	"github.com/quinton11/go-client/api/models"

	"fmt"

	"github.com/go-resty/resty/v2"
)

func (c *ClientConfig) GetUserWorkSpaces(httpClient *resty.Client) (models.GetWorkSpacesResponse, error) {
	var workSpacesResponse models.GetWorkSpacesResponse
	response, err := httpClient.
		R().
		SetResult(&workSpacesResponse).
		SetHeader("User-Agent", c.UserAgent).
		Get(fmt.Sprintf("%v/v1/workspace", c.BaseUrl))

	if err != nil {
		return models.GetWorkSpacesResponse{}, err
	}

	if response.IsError() {
		return models.GetWorkSpacesResponse{}, fmt.Errorf("CallGetAllWorkSpacesUserBelongsTo: Unsuccessful response:  [response=%v]", response)
	}

	return workSpacesResponse, nil
}

func (c *ClientConfig) GetEncryptedWorkSpaceKey(httpClient *resty.Client, request models.GetEncryptedWorkspaceKeyRequest) (models.GetEncryptedWorkspaceKeyResponse, error) {
	endpoint := fmt.Sprintf("%v/v2/workspace/%v/encrypted-key", c.BaseUrl, request.WorkspaceId)
	var result models.GetEncryptedWorkspaceKeyResponse
	response, err := httpClient.
		R().
		SetResult(&result).
		SetHeader("User-Agent", c.UserAgent).
		Get(endpoint)

	if err != nil {
		return models.GetEncryptedWorkspaceKeyResponse{}, fmt.Errorf("CallGetEncryptedWorkspaceKey: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.GetEncryptedWorkspaceKeyResponse{}, fmt.Errorf("CallGetEncryptedWorkspaceKey: Unsuccessful response: [response=%s]", response)
	}

	return result, nil
}

func (c *ClientConfig) GetServiceAccountWorkSpacePermissionsV2(httpClient *resty.Client) (models.ServiceAccountWorkspacePermissions, error) {
	var serviceAccountWorkspacePermissionsResponse models.ServiceAccountWorkspacePermissions
	response, err := httpClient.
		R().
		SetResult(&serviceAccountWorkspacePermissionsResponse).
		SetHeader("User-Agent", c.UserAgent).
		Get(fmt.Sprintf("%v/v2/service-accounts/<service-account-id>/permissions/workspace", c.BaseUrl))

	if err != nil {
		return models.ServiceAccountWorkspacePermissions{}, fmt.Errorf("CallGetServiceAccountWorkspacePermissionsV2: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.ServiceAccountWorkspacePermissions{}, fmt.Errorf("CallGetServiceAccountWorkspacePermissionsV2: Unsuccessful response: [response=%s]", response)
	}

	return serviceAccountWorkspacePermissionsResponse, nil
}

func (c *ClientConfig) GetAccessibleEnvironments(httpClient *resty.Client, request models.GetAccessibleEnvironmentsRequest) (models.GetAccessibleEnvironmentsResponse, error) {
	var accessibleEnvironmentsResponse models.GetAccessibleEnvironmentsResponse
	response, err := httpClient.
		R().
		SetResult(&accessibleEnvironmentsResponse).
		SetHeader("User-Agent", c.UserAgent).
		Get(fmt.Sprintf("%v/v2/workspace/%s/environments", c.BaseUrl, request.WorkspaceId))

	if err != nil {
		return models.GetAccessibleEnvironmentsResponse{}, err
	}

	if response.IsError() {
		return models.GetAccessibleEnvironmentsResponse{}, fmt.Errorf("CallGetAccessibleEnvironments: Unsuccessful response:  [response=%v] [response-code=%v] [url=%s]", response, response.StatusCode(), response.Request.URL)
	}

	return accessibleEnvironmentsResponse, nil
}
