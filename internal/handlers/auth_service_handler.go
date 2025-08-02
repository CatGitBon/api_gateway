package handlers

import (
	authService "github.com/CatGitBon/auth_service/pkg"
	"log"
	"net/http"
)

var authClient authService.AuthServiceClient

func SetAuthClient(client authService.AuthServiceClient) {
	authClient = client
}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	req := &authService.AuthRequest{
		UserId:   "user124",
		Password: "password123",
	}
	log.Printf(req.UserId, req.Password)

	res, err := authClient.Authenticate(r.Context(), req)
	log.Printf(res.Message, err)
	if err != nil {
		http.Error(w, "Failed to authenticate", http.StatusInternalServerError)
		return
	}

	if res.Success {
		w.Write([]byte("Authentication successful"))
	} else {
		w.Write([]byte("Authentication failed"))
	}

}
