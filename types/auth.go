package types

type CreateTokenRequest struct {
	Email         string `json:"email"`
	PlainPassword string `json:"plainPassword"`
	TntName       string `json:"tntName"`
}

type CreateTokenResponse struct {
	LoginResult  bool   `json:"loginResult"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
	TntName      string `json:"tntName"`
}

type RefreshTokenResponse struct {
	LoginResult  bool   `json:"loginResult"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type DeleteTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
	TntName      string `json:"tntName"`
}

type DeleteTokenResponse struct {
}
