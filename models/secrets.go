package models

import "time"

type GetEncryptedSecretsV3Request struct {
	Environment   string `json:"environment"`
	WorkspaceId   string `json:"workspaceId"`
	SecretPath    string `json:"secretPath"`
	IncludeImport bool   `json:"include_imports"`
	ETag          string `json:"etag,omitempty"`
}

type EncryptedSecretV3 struct {
	ID        string `json:"_id"`
	Version   int    `json:"version"`
	Workspace string `json:"workspace"`
	Type      string `json:"type"`
	Tags      []struct {
		ID        string `json:"_id"`
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		Workspace string `json:"workspace"`
	} `json:"tags"`
	Environment             string    `json:"environment"`
	SecretKeyCiphertext     string    `json:"secretKeyCiphertext"`
	SecretKeyIV             string    `json:"secretKeyIV"`
	SecretKeyTag            string    `json:"secretKeyTag"`
	SecretValueCiphertext   string    `json:"secretValueCiphertext"`
	SecretValueIV           string    `json:"secretValueIV"`
	SecretValueTag          string    `json:"secretValueTag"`
	SecretCommentCiphertext string    `json:"secretCommentCiphertext"`
	SecretCommentIV         string    `json:"secretCommentIV"`
	SecretCommentTag        string    `json:"secretCommentTag"`
	Algorithm               string    `json:"algorithm"`
	KeyEncoding             string    `json:"keyEncoding"`
	Folder                  string    `json:"folder"`
	V                       int       `json:"__v"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type ImportedSecretV3 struct {
	Environment string              `json:"environment"`
	FolderId    string              `json:"folderId"`
	SecretPath  string              `json:"secretPath"`
	Secrets     []EncryptedSecretV3 `json:"secrets"`
}

type GetEncryptedSecretsV3Response struct {
	Secrets         []EncryptedSecretV3 `json:"secrets"`
	ImportedSecrets []ImportedSecretV3  `json:"imports,omitempty"`
}

type CreateSecretV3Request struct {
	SecretName              string `json:"secretName"`
	WorkspaceID             string `json:"workspaceId"`
	Type                    string `json:"type"`
	Environment             string `json:"environment"`
	SecretKeyCiphertext     string `json:"secretKeyCiphertext"`
	SecretKeyIV             string `json:"secretKeyIV"`
	SecretKeyTag            string `json:"secretKeyTag"`
	SecretValueCiphertext   string `json:"secretValueCiphertext"`
	SecretValueIV           string `json:"secretValueIV"`
	SecretValueTag          string `json:"secretValueTag"`
	SecretCommentCiphertext string `json:"secretCommentCiphertext"`
	SecretCommentIV         string `json:"secretCommentIV"`
	SecretCommentTag        string `json:"secretCommentTag"`
	SecretPath              string `json:"secretPath"`
}

type DeleteSecretV3Request struct {
	SecretName  string `json:"secretName"`
	WorkspaceId string `json:"workspaceId"`
	Environment string `json:"environment"`
	Type        string `json:"type"`
	SecretPath  string `json:"secretPath"`
}

type UpdateSecretByNameV3Request struct {
	SecretName            string `json:"secretName"`
	WorkspaceID           string `json:"workspaceId"`
	Environment           string `json:"environment"`
	Type                  string `json:"type"`
	SecretPath            string `json:"secretPath"`
	SecretValueCiphertext string `json:"secretValueCiphertext"`
	SecretValueIV         string `json:"secretValueIV"`
	SecretValueTag        string `json:"secretValueTag"`
}

type GetSingleSecretByNameV3Request struct {
	SecretName  string `json:"secretName"`
	WorkspaceId string `json:"workspaceId"`
	Environment string `json:"environment"`
	Type        string `json:"type"`
	SecretPath  string `json:"secretPath"`
}

type GetSingleSecretByNameSecretResponse struct {
	Secrets []struct {
		ID                      string    `json:"_id"`
		Version                 int       `json:"version"`
		Workspace               string    `json:"workspace"`
		Type                    string    `json:"type"`
		Environment             string    `json:"environment"`
		SecretKeyCiphertext     string    `json:"secretKeyCiphertext"`
		SecretKeyIV             string    `json:"secretKeyIV"`
		SecretKeyTag            string    `json:"secretKeyTag"`
		SecretValueCiphertext   string    `json:"secretValueCiphertext"`
		SecretValueIV           string    `json:"secretValueIV"`
		SecretValueTag          string    `json:"secretValueTag"`
		SecretCommentCiphertext string    `json:"secretCommentCiphertext"`
		SecretCommentIV         string    `json:"secretCommentIV"`
		SecretCommentTag        string    `json:"secretCommentTag"`
		Algorithm               string    `json:"algorithm"`
		KeyEncoding             string    `json:"keyEncoding"`
		Folder                  string    `json:"folder"`
		V                       int       `json:"__v"`
		CreatedAt               time.Time `json:"createdAt"`
		UpdatedAt               time.Time `json:"updatedAt"`
	} `json:"secrets"`
}
