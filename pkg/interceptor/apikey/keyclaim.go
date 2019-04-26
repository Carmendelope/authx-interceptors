package apikey

const KeyAccessPrimitive = "APIKEY"

// KeyClaim with a set of primitives.
type KeyClaim struct {
	Primitives     [] string `json:"access,omitempty"`
}

func NewDefaultKeyClaim() *KeyClaim{
	return &KeyClaim{
		Primitives: []string{KeyAccessPrimitive},
	}
}