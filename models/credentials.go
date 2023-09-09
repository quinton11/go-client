// Package models contain custom data types for interacting
// with this client
package models

import "time"

type ApiConfig struct {
	UserAgent  string `json:"userAgent"`
	HostApiUrl string `json:"hostApiUrl"`
}

type GetLoginOneV2Request struct {
	Email           string `json:"email"`
	ClientPublicKey string `json:"clientPublicKey"`
}

type GetLoginOneV2Response struct {
	ServerPublicKey string `json:"serverPublicKey"`
	Salt            string `json:"salt"`
}

type GetLoginTwoV2Request struct {
	Email       string `json:"email"`
	ClientProof string `json:"clientProof"`
}

type GetLoginTwoV2Response struct {
	MfaEnabled          bool   `json:"mfaEnabled"`
	EncryptionVersion   int    `json:"encryptionVersion"`
	Token               string `json:"token"`
	PublicKey           string `json:"publicKey"`
	EncryptedPrivateKey string `json:"encryptedPrivateKey"`
	Iv                  string `json:"iv"`
	Tag                 string `json:"tag"`
	ProtectedKey        string `json:"protectedKey"`
	ProtectedKeyIV      string `json:"protectedKeyIV"`
	ProtectedKeyTag     string `json:"protectedKeyTag"`
	RefreshToken        string `json:"RefreshToken"`
}

type VerifyMfaTokenRequest struct {
	Email    string `json:"email"`
	MFAToken string `json:"mfaToken"`
}

type VerifyMfaTokenResponse struct {
	EncryptionVersion   int    `json:"encryptionVersion"`
	Token               string `json:"token"`
	PublicKey           string `json:"publicKey"`
	EncryptedPrivateKey string `json:"encryptedPrivateKey"`
	Iv                  string `json:"iv"`
	Tag                 string `json:"tag"`
	ProtectedKey        string `json:"protectedKey"`
	ProtectedKeyIV      string `json:"protectedKeyIV"`
	ProtectedKeyTag     string `json:"protectedKeyTag"`
	RefreshToken        string `json:"refreshToken"`
}

type VerifyMfaTokenErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Context struct {
		Code      string `json:"code"`
		TriesLeft int    `json:"triesLeft"`
	} `json:"context"`
	Level       int           `json:"level"`
	LevelName   string        `json:"level_name"`
	StatusCode  int           `json:"status_code"`
	DatetimeIso time.Time     `json:"datetime_iso"`
	Application string        `json:"application"`
	Extra       []interface{} `json:"extra"`
}

type GetNewAccessTokenWithRefreshTokenResponse struct {
	Token string `json:"token"`
}

type GetServiceAccountKeysRequest struct {
	ServiceAccountId string `json:"id"`
}

type ServiceAccountKey struct {
	ID             string    `json:"_id"`
	EncryptedKey   string    `json:"encryptedKey"`
	Nonce          string    `json:"nonce"`
	Sender         string    `json:"sender"`
	ServiceAccount string    `json:"serviceAccount"`
	Workspace      string    `json:"workspace"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type GetServiceAccountKeysResponse struct {
	ServiceAccountKeys []ServiceAccountKey `json:"serviceAccountKeys"`
}

type GetServiceTokenDetailsResponse struct {
	ID           string    `json:"_id"`
	Name         string    `json:"name"`
	Workspace    string    `json:"workspace"`
	ExpiresAt    time.Time `json:"expiresAt"`
	EncryptedKey string    `json:"encryptedKey"`
	Iv           string    `json:"iv"`
	Tag          string    `json:"tag"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Scopes       []struct {
		Environment string `json:"environment"`
		SecretPath  string `json:"secretPath"`
	} `json:"scopes"`
}

type ServiceAccountDetailsResponse struct {
	ServiceAccount struct {
		ID           string    `json:"_id"`
		Name         string    `json:"name"`
		Organization string    `json:"organization"`
		PublicKey    string    `json:"publicKey"`
		LastUsed     time.Time `json:"lastUsed"`
		ExpiresAt    time.Time `json:"expiresAt"`
	} `json:"serviceAccount"`
}
