# **INFISICAL GO CLIENT**
An api client for go applications to interact with infisical's api.

## Example api calls

- ### An api call with the cli as `agent`
```go
// Example req - Get a User's Encrypted WorkSpace Key

// Create resty client
httpClient := resty.New()

// Set JWT as authToken if required
httpClient.SetAuthToken(JTWToken).
		SetHeader("Accept", "application/json")

// Configure request
request := models.GetEncryptedWorkspaceKeyRequest{
	WorkspaceId: workspaceId,
}

// configure Api Call for cli
config := models.ApiConfig{
    UserAgent: "cli"
    HostApiUrl: "https://app.infisical.com/api" // default, use self-hosted instance if any
}

workspaceKeyResponse, err := api.CallGetEncryptedWorkspaceKey(httpClient, request, config)
if err != nil {
	return fmt.Errorf("unable to get your encrypted workspace key. [err=%v]", err)
}
```

- ### An api call with `k8-operator` as `agent`
```go
// Example req - Get a User's service token details

// Create resty client
httpClient := resty.New()

// Set serviceToken as authToken if required
httpClient.SetAuthToken(serviceToken).
	SetHeader("Accept", "application/json")

// configure Api Call for k8-operator
config := models.ApiConfig{
    UserAgent: "k8-operator"
    HostApiUrl: "https://app.infisical.com/api" // default, use self-hosted instance if any
}

serviceTokenDetails, err := api.CallGetServiceTokenDetailsV2(httpClient, config)
if err != nil {
	return fmt.Errorf("unable to get service token details. [err=%v]", err)
}
```

- ### An api call with `terraform` as the `agent`
```go
// Example 1 -  Get a User's Accessible Environment for a Workspace
// Create resty client
httpClient := resty.New()

// Set Auth Token if needed
httpClient.SetAuthToken(userLoggedInDetails.JTWToken).
	SetHeader("Accept", "application/json")

// Configure request
request := models.GetAccessibleEnvironmentsRequest{WorkspaceId: workspaceId}

// configure Api Call for terraform
config := models.ApiConfig{
    UserAgent: "terraform"
    HostApiUrl: "https://app.infisical.com/api" // default, use self-hosted instance if any
}

response, err := api.CallGetAccessibleEnvironments(httpClient, request,config)
if err != nil {
	return err
}

// Example 2 -  Get a User's Encrypted secrets given a workspace and an environment

// Create resty client
httpClient := resty.New()

// Set Auth Token if needed
httpClient.SetAuthToken(userLoggedInDetails.JTWToken).
	SetHeader("Accept", "application/json")

// Configure request
request := models.GetEncryptedSecretsV3Request{
		WorkspaceId: serviceTokenDetails.Workspace,
		Environment: envSlug,
	}

// configure Api Call for terraform
config := models.ApiConfig{
    UserAgent: "terraform"
    HostApiUrl: "https://app.infisical.com/api" // default, use self-hosted instance if any
}

encryptedSecrets, err := client.CallGetSecretsV3(httpClient, request, config)
if err != nil {
	return nil, nil, err
}
```