package domain

type Media struct {
	MediaItems []MediaItems `json:"mediaItems"`
}

type MediaItems struct {
	Id            string        `json:"id"`
	ProductUrl    string        `json:"productUrl"`
	BaseUrl       string        `json:"baseUrl"`
	MimeType      string        `json:"mimeType"`
	MediaMetadata MediaMetaData `json:"mediaMetadata"`
	Filename      string        `json:"filename"`
}
type MediaMetaData struct {
	CreationTime string `json:"creationTime"`
	Width        string `json:"width"`
	Height       string `json:"height"`
}

func (mediaMetaData MediaMetaData) getImageSizeForDownload() string {
	return "=w" + mediaMetaData.Width + "-h" + mediaMetaData.Height

}

func (mediaItems MediaItems) GetUrlForDownload() string {
	return mediaItems.BaseUrl + mediaItems.MediaMetadata.getImageSizeForDownload()

}
