package registry
// ----------------------------------------------------------------------------
// DO NOT EDIT THIS FILE
// This file was generated by `swagger generate operation`
//
// See hack/generate-swagger-api.sh
// ----------------------------------------------------------------------------

// AuthenticateOKBody authenticate o k body
// swagger:model AuthenticateOKBody
type AuthenticateOKBody struct {

	// An opaque token used to authenticate a user after a successful login
	// Required: true
	IdentityToken string `json:"IdentityToken"`

	// The status of the authentication
	// Required: true
	Status string `json:"Status"`
}
