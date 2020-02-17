package images

import (
	"io"
	"net/http"
	"os"
	"path"
)

func ImportImage(w http.ResponseWriter, r *http.Request) {
	imgpth := path.Join("assets", "image10.jpg")
	img, err := os.Open(imgpth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, img)

}
