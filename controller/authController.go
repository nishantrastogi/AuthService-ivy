package controller

import (
	"authservice/constants"
	"authservice/model"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var credentials model.UserDetails

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Printf("Unable to parse user credentials. %s \n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validatePayload := model.NewValidateUserPasswordRequest(credentials.Username, credentials.Password)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(validatePayload)

	// Calling user REST service to validate if the user and password combination is correct.
	resp, err := http.Post("http://localhost:4000/api/user/validate", "application/json;charset=UTF-8", reqBodyBytes) // Call to auth server for token validation.
	if err != nil {
		fmt.Printf("An error occured during REST call: %s", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "User unauthorized.", resp.StatusCode)
		return
	}

	var vpResp model.ValidatePasswordResponse
	json.NewDecoder(resp.Body).Decode(&vpResp)

	// now we fill more metadata for the token, eg expiration time = 5 mins
	expirationTime := time.Now().Add(5 * time.Minute)
	// create the JWT claims, which is basically the data we want to send in the JWT token.
	claims := &model.Claims{
		Username: vpResp.Uname,
		Role:     vpResp.User_role,
		StandardClaims: jwt.StandardClaims{
			// in JWT expired time is expressed as unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 	Now we will declare the token with the algo used for signing and the claims.
	token := jwt.NewWithClaims(constants.SignMethod, claims)
	//Create jwt string
	tokenString, err := token.SignedString(constants.JwtKey)
	if err != nil {
		log.Printf("Error generating token string. %s.\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Created token: \n %s.", tokenString)
	signInresponse := model.NewSignInResponse(vpResp.Uname, vpResp.User_role, tokenString)
	json.NewEncoder(w).Encode(signInresponse)

}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var req model.ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Unable to parse enduser payload %s \n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := &model.Claims{}
	tknStr := req.TokenString
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return constants.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Printf("Creating response with %s and %s", claims.Username, claims.Role)
	validateResponse := model.NewValidateResponse(claims.Username, claims.Role)
	json.NewEncoder(w).Encode(validateResponse)

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var req model.ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Unable to parse enduser payload %s \n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := &model.Claims{}
	tknStr := req.TokenString
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return constants.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// Below code is set to 1 second as an example.
	if time.Since(time.Unix(claims.IssuedAt, 0)) < 1*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Too early to refresh the token")
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(constants.JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	refreshResponse := model.NewRefreshResponse(tokenString)
	json.NewEncoder(w).Encode(refreshResponse)
}
