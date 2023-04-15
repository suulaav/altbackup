package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/suulaav/altbackup/pkg/service/backup"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var config *oauth2.Config

func StartBackupService(appConfig AppConfigs) {
	oAuthCredentials, _ := json.Marshal(map[string]interface{}{
		"web": appConfig.OAuth})
	var err error
	config, err = google.ConfigFromJSON(oAuthCredentials, "https://www.googleapis.com/auth/photoslibrary.readonly")
	if err != nil {
		log.Fatalf("Error reading OAuth2 config: %v", err)
	}
	startServer(config)
}

func startServer(config *oauth2.Config) {
	var err error
	authURL := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	err = openBrowser(authURL)
	if err != nil {
		log.Fatalf("Error opening browser: %v", err)
	}
	http.HandleFunc("/oauth2callback", waitForCallback)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}

func waitForCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Error exchanging code for token", http.StatusInternalServerError)
		log.Fatalf("Error exchanging code for token: %v", err)
	}
	fmt.Fprintf(w, "Your Media Is Being Downlaoded You Can Close this Window Now ")
	backup.ExecuteBackupService(token)

}

func openBrowser(url string) error {
	var err error
	switch {
	case os.Getenv("OS") == "Windows_NT":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case os.Getenv("OS") == "darwin":
		err = exec.Command("open", url).Start()
	case isCommandAvailable("xdg-open"):
		err = exec.Command("xdg-open", url).Start()
	case isCommandAvailable("gnome-open"):
		err = exec.Command("gnome-open", url).Start()
	case isCommandAvailable("kde-open"):
		err = exec.Command("kde-open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
