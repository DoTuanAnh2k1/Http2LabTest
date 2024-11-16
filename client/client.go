package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

const ServerUrl = "http://localhost:3654/test"

const (
	Mess1 = `The quick brown fox jumps over the lazy dog, a phrase often used to test typewriters and keyboards due to its inclusion of every letter in the English alphabet. However, beyond its linguistic utility, this sentence can serve as a reminder of the simplicity and elegance of concise communication. In our increasingly digital world, where applications strive to process vast amounts of text, ensuring proper handling of diverse characters, punctuation marks, and even whitespace is crucial. Developers often design systems that analyze, transform, and display text with precision, but they must also account for edge cases, such as excessive line breaks, unexpected symbols, or mixed encodings. Whether you are testing a text editor, a translation tool, or an AI chatbot, challenges like these provide valuable insights into the robustness and versatility of your software.`
	Mess2 = `The quick brown fox jumps over the lazy dog, a phrase often used to test typewriters and keyboards due to its inclusion of every letter in the English alphabet. However, beyond its linguistic utility, this sentence can serve as a reminder of the simplicity and elegance of concise communication. In our increasingly digital world, where applications strive to process vast amounts of text, ensuring proper handling of diverse characters, punctuation marks, and even whitespace is crucial. Developers often design systems that analyze, transform, and display text with precision, but they must also account for edge cases, such as excessive line breaks, unexpected symbols, or mixed encodings. Whether you are testing a text editor, a translation tool, or an AI chatbot, challenges like these provide valuable insights into the robustness and versatility of your software.`
	Mess3 = `The quick brown fox jumps over the lazy dog, a phrase often used to test typewriters and keyboards due to its inclusion of every letter in the English alphabet. However, beyond its linguistic utility, this sentence can serve as a reminder of the simplicity and elegance of concise communication. In our increasingly digital world, where applications strive to process vast amounts of text, ensuring proper handling of diverse characters, punctuation marks, and even whitespace is crucial. Developers often design systems that analyze, transform, and display text with precision, but they must also account for edge cases, such as excessive line breaks, unexpected symbols, or mixed encodings. Whether you are testing a text editor, a translation tool, or an AI chatbot, challenges like these provide valuable insights into the robustness and versatility of your software.`
)

func dialTlsContext(ctx context.Context, network, address string, cfg *tls.Config) (net.Conn, error) {
	return net.Dial(network, address)
}

var client = &http.Client{
	Transport: &http2.Transport{
		AllowHTTP:                  true,
		StrictMaxConcurrentStreams: true,
		DialTLSContext:             dialTlsContext,
	},
	Timeout: 10 * time.Second,
}

func triggerHandle(w http.ResponseWriter, r *http.Request) {

	body, _ := json.Marshal(Mess1)
	req, _ := http.NewRequest(http.MethodPost, ServerUrl, bytes.NewBuffer(body))

	client.Do(req)
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/trigger", triggerHandle)

	return mux
}

func newServer() *http.Server {
	mux := newRouter()
	return &http.Server{
		Addr:    ":3317",
		Handler: mux,
	}
}

func main() {
	clientServer := newServer()
	err := clientServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
