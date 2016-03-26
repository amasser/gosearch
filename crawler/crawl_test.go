package crawler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllLinks(t *testing.T) {
	assert := assert.New(t)
	var input string = "hoge<a href=\"http://example.com/\">link title</a>hoge\n<a href=\"http://hoge.example.com/\">link title</a>"
	urls := getAllLinks(input)
	assert.Equal("http://example.com/", urls[0])
	assert.Equal("http://hoge.example.com/", urls[1])
}

func DummyCrawledHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hogehoge<a href=\"http://example.com/\">fugafuga")
}

func TestCrawl(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(http.HandlerFunc(DummyCrawledHandler))
	defer ts.Close()

	msg := make(chan string)
	tocrawl := make(chan URL)
	go crawl(ts.URL, 1, msg, tocrawl)

	assert.Equal(ts.URL+" is crawled.", <-msg)
	url := <-tocrawl
	assert.Equal("http://example.com/", url.url)
	assert.Equal(0, url.depth)
}
