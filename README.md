# **INFISICAL GO CLIENT**
An api client for go applications to interact with infisical's api.

## Example api calls

- ### An api call with the cli as `agent`
```go
// Example req - Get a User's Encrypted WorkSpace Key
import (
	"github.com/infisical/go-client/api/models"
	"github.com/infisical/go-client/api"
	)

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
client := api.ClientConfig{
    UserAgent: "cli"
    BaseUrl: "https://app.infisical.com/api" // default, use self-hosted instance if any
}

workspaceKeyResponse, err := client.GetEncryptedWorkspaceKey(httpClient, request)
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
client := api.ClientConfig{
    UserAgent: "k8-operator"
    BaseUrl: "https://app.infisical.com/api" // default, use self-hosted instance if any
}

serviceTokenDetails, err := client.GetServiceTokenDetailsV2(httpClient)
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
client := api.ClientConfig{
    UserAgent: "terraform"
    HostApiUrl: "https://app.infisical.com/api" // default, use self-hosted instance if any
}

response, err := client.GetAccessibleEnvironments(httpClient, request)
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
client := api.ClientConfig{
    UserAgent: "terraform"
    HostApiUrl: "https://app.infisical.com/api" // default, use self-hosted instance if any
}

encryptedSecrets, err := client.GetSecretsV3(httpClient, request)
if err != nil {
	return nil, nil, err
}
```