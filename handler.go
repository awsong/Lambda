/*
package main

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	//qr, _ := qrcode.New(asin, qrcode.High)
	//qr.Write(512, w)
	log.Println(string(evt))
	return map[string]string{"key": string(evt)}, nil
}
*/
package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	qrcode "github.com/skip2/go-qrcode"
)

// Handle is the exported handler called by AWS Lambda.
var Handle apigatewayproxy.Handler

func init() {
	ln := net.Listen()

	// Amazon API Gateway binary media types are supported out of the box.
	// If you don't send or receive binary data, you can safely set it to nil.
	//Handle = apigatewayproxy.New(ln, []string{"image/png"}).Handle
	Handle = apigatewayproxy.New(ln, nil).Handle

	// Any Go framework complying with the Go http.Handler interface can be used.
	// This includes, but is not limited to, Vanilla Go, Gin, Echo, Gorrila, Goa, etc.
	go http.Serve(ln, http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/asin/":
		asin := r.URL.Query()["asin"][0]
		log.Println(r.URL.RawQuery)
		log.Println(asin)
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Transfer-Encoding", "base64")
		qr, _ := qrcode.New("http://www.google.com", qrcode.High)
		content, _ := qr.PNG(512)
		//qr.Write(512, w)
		//w.Header().Set("Content-Length", strconv.Itoa(len(content)))
		encoder := base64.NewEncoder(base64.StdEncoding, w)
		encoder.Write(content)
		encoder.Close()
	case "/page/":
		w.Header().Set("Content-Type", "text/html")
		log.Println(r.URL.RawQuery)
		asins := strings.Split(r.URL.Query()["asins"][0], "\n")
		for _, asin := range asins {
			w.Write([]byte(`<div><img src="/Prod/asin/?asin=`))
			w.Write([]byte(asin))
			w.Write([]byte(`">` + asin + `</img></div><br/>`))
		}
		w.Write([]byte("</body></html>"))
	case "/":
		w.Write([]byte(index))
	}
}
