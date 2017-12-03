package main

import(
	"fmt"
	"net/http"
	"gopkg.in/square/go-jose.v2"
  "github.com/auth0-community/go-auth0"
  "github.com/koding/multiconfig"
)
type Auth struct {
  Jwks_uri string
  Auth0_api_issuer string
  Auth0_api_audience string
}

func checkJwt(w http.ResponseWriter, r *http.Request) bool {
  validation := false
  
  // config 
	m := multiconfig.NewWithPath("./config/auth.json")
  auth := new(Auth)
  m.MustLoad(auth)
  var aud []string
  aud = append(aud, auth.Auth0_api_audience)
  
  client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth.Jwks_uri})
  audience := aud

  configuration := auth0.NewConfiguration(client, audience, auth.Auth0_api_issuer, jose.RS256)
  validator := auth0.NewValidator(configuration)

  _, err := validator.ValidateRequest(r)

  if err != nil {
    fmt.Println("Token is not valid or missing token")

    w.WriteHeader(http.StatusUnauthorized)

  } else {
  fmt.Println("Token is valid")
  validation = true
  }
  return validation
}