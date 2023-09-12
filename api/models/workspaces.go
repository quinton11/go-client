package models

import "time"

type GetWorkSpacesResponse struct {
	Workspaces []struct {
		ID           string `json:"_id"`
		Name         string `json:"name"`
		Plan         string `json:"plan,omitempty"`
		V            int    `json:"__v"`
		Organization string `json:"organization,omitempty"`
	} `json:"workspaces"`
}

type GetEncryptedWorkspaceKeyRequest struct {
	WorkspaceId string `json:"workspaceId"`
}

type GetEncryptedWorkspaceKeyResponse struct {
	ID           string `json:"_id"`
	EncryptedKey string `json:"encryptedKey"`
	Nonce        string `json:"nonce"`
	Sender       struct {
		ID             string    `json:"_id"`
		Email          string    `json:"email"`
		RefreshVersion int       `json:"refreshVersion"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		V              int       `json:"__v"`
		FirstName      string    `json:"firstName"`
		LastName       string    `json:"lastName"`
		PublicKey      string    `json:"publicKey"`
	} `json:"sender"`
	Receiver  string    `json:"receiver"`
	Workspace string    `json:"workspace"`
	V         int       `json:"__v"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ServiceAccountWorkspacePermission struct {
	ID             string `json:"_id"`
	ServiceAccount string `json:"serviceAccount"`
	Workspace      struct {
		ID                 string `json:"_id"`
		Name               string `json:"name"`
		AutoCapitalization bool   `json:"autoCapitalization"`
		Organization       string `json:"organization"`
		Environments       []struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
			ID   string `json:"_id"`
		} `json:"environments"`
	} `json:"workspace"`
	Environment string `json:"environment"`
	Read        bool   `json:"read"`
	Write       bool   `json:"write"`
}

type ServiceAccountWorkspacePermissions struct {
	ServiceAccountWorkspacePermission []ServiceAccountWorkspacePermissions `json:"serviceAccountWorkspacePermissions"`
}

type GetAccessibleEnvironmentsRequest struct {
	WorkspaceId string `json:"workspaceId"`
}

type GetAccessibleEnvironmentsResponse struct {
	AccessibleEnvironments []struct {
		Name          string `json:"name"`
		Slug          string `json:"slug"`
		IsWriteDenied bool   `json:"isWriteDenied"`
	} `json:"accessibleEnvironments"`
}
