package model

import "time"

// iss (issuer): Issuer of the JWT
// sub (subject): Subject of the JWT (the user)
// aud (audience): Recipient for which the JWT is intended
// exp (expiration time): Time after which the JWT expires
// nbf (not before time): Time before which the JWT must not be accepted for processing
// iat (issued at time): Time at which the JWT was issued; can be used to determine age of the JWT
// jti (JWT ID): Unique identifier; can be used to prevent the JWT from being replayed (allows a token to be used only once)

type StandardClaim struct {
	Jti string `json:"jti"`
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Exp uint64 `json:"exp"`
	Nbf uint64 `json:"nbf"`
	Iat uint64 `json:"iat"`
}

type AccessClaim struct {
	StandardClaim
	UserID   uint64    `json:"user_id"`
	Username string    `json:"username"`
	Dob      time.Time `json:"dob"`
}
