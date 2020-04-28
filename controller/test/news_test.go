package controller_test

import (
	"casbindemo/config"
	"casbindemo/logger"
	"casbindemo/model"
	"casbindemo/route"
	"flag"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(req *http.Request) *httptest.ResponseRecorder {
	flag.Parse()
	config.Setup()
	logger.Setup()
	model.Setup()
	r := route.Setup()
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var auth = "admin"

// 列表
func TestGetNewsList(t *testing.T) {
	req, err := http.NewRequest("GET", "/news", nil)
	assert.Nil(t, err)
	req.Header.Set("Authority", auth)
	w := performRequest(req)
	assert.Equal(t, 200, w.Code)

}

// 详情
func TestGetNewsDetail(t *testing.T) {
	req, err := http.NewRequest("GET", "/news/1", nil)
	assert.Nil(t, err)
	req.Header.Set("Authority", auth)
	w := performRequest(req)
	assert.Equal(t, 200, w.Code)
}

// 添加
func TestAddNews(t *testing.T) {
	req, err := http.NewRequest("POST", "/news", nil)
	assert.Nil(t, err)
	req.Header.Set("Authority", auth)
	w := performRequest(req)
	assert.Equal(t, 200, w.Code)
}

// 编辑
func TestEditNews(t *testing.T) {
	req, err := http.NewRequest("PUT", "/news/1", nil)
	assert.Nil(t, err)
	req.Header.Set("Authority", auth)
	w := performRequest(req)
	assert.Equal(t, 200, w.Code)
}

// 删除
func TestDelNews(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/news/1", nil)
	assert.Nil(t, err)
	req.Header.Set("Authority", auth)
	w := performRequest(req)
	assert.Equal(t, 200, w.Code)
}
