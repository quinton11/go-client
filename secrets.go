package client

import (
	"github.com/quinton11/go-infisical/models"

	"fmt"

	"github.com/go-resty/resty/v2"
)

// Terraform, CLI, k8
func GetSecretsV3(httpClient *resty.Client, request models.GetEncryptedSecretsV3Request, config models.ApiConfig) (models.GetEncryptedSecretsV3Response, error) {
	var secretsResponse models.GetEncryptedSecretsV3Response

	httpRequest := httpClient.
		R().
		SetResult(&secretsResponse).
		SetHeader("User-Agent", config.UserAgent).
		SetQueryParam("environment", request.Environment).
		SetQueryParam("workspaceId", request.WorkspaceId)

	if request.ETag != "" {
		httpRequest.SetHeader("If-None-Match", request.ETag)
	}

	if request.IncludeImport {
		httpRequest.SetQueryParam("include_imports", "true")
	}

	if request.SecretPath != "" {
		httpRequest.SetQueryParam("secretPath", request.SecretPath)
	}

	response, err := httpRequest.Get(fmt.Sprintf("%v/v3/secrets", config.HostApiUrl))

	if err != nil {
		return models.GetEncryptedSecretsV3Response{}, fmt.Errorf("CallGetSecretsV3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		if response.StatusCode() == 401 {
			return models.GetEncryptedSecretsV3Response{}, fmt.Errorf("CallGetSecretsV3: Request to access secrets with [environment=%v] [path=%v] [workspaceId=%v] is denied. Please check if your authentication method has access to requested scope", request.Environment, request.SecretPath, request.WorkspaceId)
		} else {
			return models.GetEncryptedSecretsV3Response{}, fmt.Errorf("CallGetSecretsV3: Unsuccessful response. Please make sure your secret path, workspace and environment name are all correct [response=%v]", response.RawResponse)
		}
	}

	return secretsResponse, nil
}

func CreateSecretsV3(httpClient *resty.Client, request models.CreateSecretV3Request, config models.ApiConfig) error {
	var secretsResponse models.GetEncryptedSecretsV3Response
	response, err := httpClient.
		R().
		SetResult(&secretsResponse).
		SetHeader("User-Agent", config.UserAgent).
		SetBody(request).
		Post(fmt.Sprintf("%v/v3/secrets/%s", config.HostApiUrl, request.SecretName))

	if err != nil {
		return fmt.Errorf("CallCreateSecretsV3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return fmt.Errorf("CallCreateSecretsV3: Unsuccessful response. Please make sure your secret path, workspace and environment name are all correct [response=%s]", response)
	}

	return nil
}

func DeleteSecretsV3(httpClient *resty.Client, request models.DeleteSecretV3Request, config models.ApiConfig) error {
	var secretsResponse models.GetEncryptedSecretsV3Response
	response, err := httpClient.
		R().
		SetResult(&secretsResponse).
		SetHeader("User-Agent", config.UserAgent).
		SetBody(request).
		Delete(fmt.Sprintf("%v/v3/secrets/%s", config.HostApiUrl, request.SecretName))

	if err != nil {
		return fmt.Errorf("CallDeleteSecretsV3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return fmt.Errorf("CallDeleteSecretsV3: Unsuccessful response. Please make sure your secret path, workspace and environment name are all correct [response=%s]", response)
	}

	return nil
}

func UpdateSecretsV3(httpClient *resty.Client, request models.UpdateSecretByNameV3Request, config models.ApiConfig) error {
	var secretsResponse models.GetEncryptedSecretsV3Response
	response, err := httpClient.
		R().
		SetResult(&secretsResponse).
		SetHeader("User-Agent", config.UserAgent).
		SetBody(request).
		Patch(fmt.Sprintf("%v/v3/secrets/%s", config.HostApiUrl, request.SecretName))

	if err != nil {
		return fmt.Errorf("CallUpdateSecretsV3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return fmt.Errorf("CallUpdateSecretsV3: Unsuccessful response. Please make sure your secret path, workspace and environment name are all correct [response=%s]", response)
	}

	return nil
}

func GetSingleSecretByNameV3(httpClient *resty.Client, request models.CreateSecretV3Request, config models.ApiConfig) error {
	var secretsResponse models.GetEncryptedSecretsV3Response
	response, err := httpClient.
		R().
		SetResult(&secretsResponse).
		SetHeader("User-Agent", config.UserAgent).
		SetBody(request).
		Post(fmt.Sprintf("%v/v3/secrets/%s", config.HostApiUrl, request.SecretName))

	if err != nil {
		return fmt.Errorf("CallGetSingleSecretByNameV3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		return fmt.Errorf("CallGetSingleSecretByNameV3: Unsuccessful response. Please make sure your secret path, workspace and environment name are all correct [response=%s]", response)
	}

	return nil
}
