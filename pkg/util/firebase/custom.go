package firebase

// AuthenticationResponse represents the response from firebase authentication
type AuthenticationResponse struct {
	Context          string `json:"context"`
	DisplayName      string `json:"displayName"`
	Email            string `json:"email"`
	EmailVerified    bool   `json:"emailVerified"`
	ExpiresIn        string `json:"expiresIn"`
	FederatedID      string `json:"federatedId"`
	FirstName        string `json:"firstName"`
	FullName         string `json:"fullName"`
	IDToken          string `json:"idToken"`
	IsNewUser        bool   `json:"isNewUser"`
	Kind             string `json:"kind"`
	LocalID          string `json:"localId"`
	OauthAccessToken string `json:"oauthAccessToken"`
	OauthExpireIn    int64  `json:"oauthExpireIn"`
	OauthIDToken     string `json:"oauthIdToken"`
	PhotoURL         string `json:"photoUrl"`
	ProviderID       string `json:"providerId"`
	RawUserInfo      string `json:"rawUserInfo"`
	RefreshToken     string `json:"refreshToken"`
}
