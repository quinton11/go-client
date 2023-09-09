package client

import (
	"github.com/quinton11/go-infisical/models"

	"fmt"

	"github.com/go-resty/resty/v2"
)

func GetUserWorkSpaces(httpClient *resty.Client, config models.ApiConfig) (models.GetWorkSpacesResponse, error) {
	var workSpacesResponse models.GetWorkSpacesResponse
	response, err := httpClient.
		R().
		SetResult(&workSpacesResponse).
		SetHeader("User-Agent", config.UserAgent).
		Get(fmt.Sprintf("%v/v1/workspace", config.HostApiUrl))

	if err != nil {
		return models.GetWorkSpacesResponse{}, err
	}

	if response.IsError() {
		return models.GetWorkSpacesResponse{}, fmt.Errorf("CallGetAllWorkSpacesUserBelongsTo: Unsuccessful response:  [response=%v]", response)
	}

	return workSpacesResponse, nil
}

func GetEncryptedWorkSpaceKey(httpClient *resty.Client, request models.GetEncryptedWorkspaceKeyRequest, config models.ApiConfig) (models.GetEncryptedWorkspaceKeyResponse, error) {
	endpoint := fmt.Sprintf("%v/v2/workspace/%v/encrypted-key", config.HostApiUrl, request.WorkspaceId)
	var result models.GetEncryptedWorkspaceKeyResponse
	response, err := httpClient.
		R().
		SetResult(&result).
		SetHeader("User-Agent", config.UserAgent).
		Get(endpoint)

	if err != nil {
		return models.GetEncryptedWorkspaceKeyResponse{}, fmt.Errorf("CallGetEncryptedWorkspaceKey: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.GetEncryptedWorkspaceKeyResponse{}, fmt.Errorf("CallGetEncryptedWorkspaceKey: Unsuccessful response: [response=%s]", response)
	}

	return result, nil
}

func GetServiceAccountWorkSpacePermissionsV2(httpClient *resty.Client, config models.ApiConfig) (models.ServiceAccountWorkspacePermissions, error) {
	var serviceAccountWorkspacePermissionsResponse models.ServiceAccountWorkspacePermissions
	response, err := httpClient.
		R().
		SetResult(&serviceAccountWorkspacePermissionsResponse).
		SetHeader("User-Agent", config.UserAgent).
		Get(fmt.Sprintf("%v/v2/service-accounts/<service-account-id>/permissions/workspace", config.HostApiUrl))

	if err != nil {
		return models.ServiceAccountWorkspacePermissions{}, fmt.Errorf("CallGetServiceAccountWorkspacePermissionsV2: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return models.ServiceAccountWorkspacePermissions{}, fmt.Errorf("CallGetServiceAccountWorkspacePermissionsV2: Unsuccessful response: [response=%s]", response)
	}

	return serviceAccountWorkspacePermissionsResponse, nil
}

func GetAccessibleEnvironments(httpClient *resty.Client, request models.GetAccessibleEnvironmentsRequest, config models.ApiConfig) (models.GetAccessibleEnvironmentsResponse, error) {
	var accessibleEnvironmentsResponse models.GetAccessibleEnvironmentsResponse
	response, err := httpClient.
		R().
		SetResult(&accessibleEnvironmentsResponse).
		SetHeader("User-Agent", config.UserAgent).
		Get(fmt.Sprintf("%v/v2/workspace/%s/environments", config.HostApiUrl, request.WorkspaceId))

	if err != nil {
		return models.GetAccessibleEnvironmentsResponse{}, err
	}

	if response.IsError() {
		return models.GetAccessibleEnvironmentsResponse{}, fmt.Errorf("CallGetAccessibleEnvironments: Unsuccessful response:  [response=%v] [response-code=%v] [url=%s]", response, response.StatusCode(), response.Request.URL)
	}

	return accessibleEnvironmentsResponse, nil
}
