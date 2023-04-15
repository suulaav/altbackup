package backup

import (
	"encoding/json"
	"github.com/suulaav/altbackup/pkg/altutils"
	"github.com/suulaav/altbackup/pkg/service/domain"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func ExecuteBackupService(token *oauth2.Token) {
	url := "https://photoslibrary.googleapis.com/v1/mediaItems"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	altutils.CheckError(err)
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	resp, err := client.Do(req)
	altutils.CheckError(err)
	responseBody, _ := io.ReadAll(resp.Body)
	var media *domain.Media
	json.Unmarshal(responseBody, &media)
	resp.Body.Close()
	downloadAllFiles(media)
}

func downloadAllFiles(medias *domain.Media) {
	for _, mediaItem := range medias.MediaItems {
		wg.Add(1)
		go downloadFile(mediaItem)
	}
	wg.Wait()
}

func downloadFile(items domain.MediaItems) {
	defer wg.Done()
	os.Mkdir("./downloads", os.ModePerm)
	output, err := os.Create("./downloads/" + items.Filename)
	altutils.CheckError(err)
	response, err := http.Get(items.GetUrlForDownload())
	altutils.CheckError(err)
	defer response.Body.Close()
	_, err = io.Copy(output, response.Body)
	altutils.CheckError(err)
	log.Println(items.Filename, "downloaded successfully")
}
