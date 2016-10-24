package app

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleCreateSuccess(t *testing.T) {
	d := InitDB()
	r := InitRouter(d.Begin())

	firstName := "asdf"

	jsonStr := fmt.Sprintf(`{
        "firstName": "%s",
        "lastName": "asdf",
        "email": "a@b.com"
    }`, firstName)

	fmt.Println("YOYOYO", jsonStr)
	req, _ := http.NewRequest("POST", "/person", strings.NewReader(jsonStr))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var jsonMap map[string]string

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&jsonMap)

	defer req.Body.Close()

	fmt.Println("YOYOYO", jsonMap, resp.Body)
	assert.Equal(t, jsonMap["firstName"], firstName)
}
