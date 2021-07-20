package main

import (
	"fmt"
	"imageTransform/primitive"
	"io"
	"log"
	"net/http"
	"path/filepath"

)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
    <body>
        <form action="/upload" method="post" enctype="multipart/form-data">
            <div class="file-uploader__message-area">
                <p>Select a file to upload</p>
            </div>
            <div class="file-chooser">
                <input class="file-chooser__input" type="file" name="image" id="image">
            </div>
                <br><input class="file-uploader__submit-button" type="submit" value="Upload">
        </form>
    </body>
</html>`
	fmt.Fprint(w, html)

}

//uploadHandler to upload image that is persisted throughout one request
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)[1:]
	out, err := primitive.Transform(file, 50)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	switch ext {
	case "jpg":
		fallthrough
	case "png":
		w.Header().Set("Content_Type", "image/png")
	case "jpeg":
		w.Header().Set("Content_Type", "image/jpeg")	
	default:
		http.Error(w, fmt.Sprintf("invalid image type %s ,", ext), http.StatusBadRequest)	
		return	
	}
	io.Copy(w,out)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
