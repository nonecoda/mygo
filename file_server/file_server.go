package fileserver

import (
	"fmt"
	"log"
	"net/http"
)

const LOCAL_DIR_PATH string = ``

const PORT uint16 = 5000

func Do() {
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", PORT),
		http.FileServer(
			http.Dir(LOCAL_DIR_PATH),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
