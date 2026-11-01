package handlers

import "net/http"

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	page := `<html><head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
	<p> Only <span style="color:pink;">Cool</span> people can read this</p>
</body></html>
`
	w.Write([]byte(page))
}
