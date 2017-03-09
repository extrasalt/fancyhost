package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	h, _ := os.Hostname()
	rdr := strings.NewReader(h)
	hash := sha256.New()
	io.Copy(hash, rdr)
	color := hex.EncodeToString(hash.Sum(nil))[:6]

	htmltempl := `
		<head>
			<title>%[2]s</title>
			<style type="text/css">
			body{
				background: #%[1]s;
				color: white;
				font: 3vw Arial;
				display: flex;
				justify-content: center;
				align-items:center;
			}
			</style>
		</head>
		<body>
			%[2]s
		</body>
		`
	html := fmt.Sprintf(htmltempl, color, h)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(html))
	})

	http.ListenAndServe(":8000", nil)
}
