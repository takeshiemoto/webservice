package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// 共通のテストコードはTestMainにまとめることが可能
var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	// テスト実行用のマルチプレクサ
	mux = http.NewServeMux()
	// テスト対象のハンドラ
	mux.HandleFunc("/post/", handleRequest)
	// 返されたHTTPレスポンスを取得
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	// テストしたいハンドラ宛のリクエストを作成
	request, _ := http.NewRequest("GET", "/post/1", nil)
	// テスト対象のハンドラにリクエストを送信
	mux.ServeHTTP(writer, request)
	// 結果をチェック
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	fmt.Printf("%+v", post)
	if post.ID != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	mockData := strings.NewReader(`{"content: "Updated post", "author": "David Gilmour"}`)
	request, _ := http.NewRequest("PUT", "/post/1", mockData)

	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
