package routergin

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/api/handlers"
	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/db/dbmemory"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/repos/linkrepo"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/repos/userrepo"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/search/fullurlsearch"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var virtualServer *httptest.Server
var httpClient *http.Client

func TestMain(m *testing.M) {
	virtualServer = newVirtualServer()
	httpClient = virtualServer.Client()
	exitCode := m.Run()
	virtualServer.Close()
	os.Exit(exitCode)
}

func TestPostUserRegist(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	t.Run("Successful request", func(t *testing.T) {
		r, _ := http.NewRequest("POST", virtualServer.URL+"/regist", strings.NewReader(`{
			"name": "ivan",
			"login": "ivan",
			"password": "123456"
		}`))

		resp, err := httpClient.Do(r)
		if err != nil {
			t.Error(err)
		}

		if !generalCheck(t, resp, http.StatusOK) {
			return
		}
	})

	t.Run("Skipped required fields", func(t *testing.T) {
		var jsonData = []string{
			`{
				"name": "michael",
				"login": "michael"
			}`,
			`{
				"name": "emily",
				"password": "123456"
			}`,
			`{
				"login": "jacob",
				"password": "123456"
			}`,
		}

		for _, jsonItem := range jsonData {
			reader := strings.NewReader(jsonItem)
			r, _ := http.NewRequest("POST", virtualServer.URL+"/regist", reader)
			resp, err := httpClient.Do(r)
			if !assert.NoError(t, err) {
				return
			}
			if !generalCheck(t, resp, http.StatusBadRequest) {
				return
			}
		}
	})
}

func TestPostLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	// Creating new user request
	r, _ := http.NewRequest("POST", virtualServer.URL+"/regist", strings.NewReader(`{
		"name": "igor",
		"login": "igor",
		"password": "123456"
	}`))
	_, _ = httpClient.Do(r)

	t.Run("Successful request", func(t *testing.T) {
		// Login request
		r, _ = http.NewRequest("POST", virtualServer.URL+"/login", strings.NewReader(`{
			"login": "igor",
			"password": "123456"
		}`))
		resp, err := httpClient.Do(r)
		if !assert.NoError(t, err) {
			return
		}

		message, err := io.ReadAll(resp.Body)
		if !assert.NoError(t, err) {
			return
		}
		if len(message) != 0 {
			t.Log(string(message))
		}
		if !assert.Equal(t, http.StatusOK, resp.StatusCode, "status must be 200") {
			return
		}
		if !assert.NotEmpty(t, message, "message cannot be empty") {
			return
		}

		decoder := json.NewDecoder(bytes.NewReader(message))
		var respData struct {
			Message string `json:"message"`
			Token   string `json:"token"`
		}
		err = decoder.Decode(&respData)
		if !assert.NoError(t, err) {
			return
		}

		if !assert.NotEmpty(t, respData.Token, "token cannot be empty") {
			return
		}
	})

	t.Run("Skipped required fields", func(t *testing.T) {
		r, _ = http.NewRequest("POST", virtualServer.URL+"/login", strings.NewReader(`{
			"login": "igor"
		}`))
		resp, err := httpClient.Do(r)
		if !assert.NoError(t, err) {
			return
		}
		if !generalCheck(t, resp, http.StatusBadRequest) {
			return
		}

		r, _ = http.NewRequest("POST", virtualServer.URL+"/login", strings.NewReader(`{
			"password": "123456"
		}`))
		resp, err = httpClient.Do(r)
		if !assert.NoError(t, err) {
			return
		}
		if !generalCheck(t, resp, http.StatusBadRequest) {
			return
		}
	})

	t.Run("Wrong password", func(t *testing.T) {
		r, _ = http.NewRequest("POST", virtualServer.URL+"/login", strings.NewReader(`{
			"login": "igor",
			"password": "wrong password"
		}`))
		resp, err := httpClient.Do(r)
		if !assert.NoError(t, err) {
			return
		}
		if !generalCheck(t, resp, http.StatusBadRequest) {
			return
		}
	})
}

func TestGetAccount(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	t.Run("Successful get user data request", func(t *testing.T) {
		login := createNewAcc()
		token := loginNewAcc(login)

		// Get request
		url := virtualServer.URL + "/account"
		r, _ := http.NewRequest("GET", url, strings.NewReader(""))
		r.Header.Set("Authorization", "Bearer "+token)
		resp, _ := httpClient.Do(r)

		// Check response status
		if !assert.Equal(t, http.StatusOK, resp.StatusCode, "account request status must be 200") {
			return
		}

		// Check response body
		decoder := json.NewDecoder(resp.Body)
		var accountData struct {
			Id    string `json:"id"`
			Login string `json:"login"`
			Name  string `json:"name"`
		}
		if !assert.NoError(t, decoder.Decode(&accountData)) {
			return
		}
		t.Log("id", accountData.Id)
		t.Log("login", accountData.Login)
		t.Log("name", accountData.Name)
	})
}

func TestDeleteAccount(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	t.Run("Successful delete", func(t *testing.T) {
		login := createNewAcc()
		token := loginNewAcc(login)

		// Delete request
		r, _ := http.NewRequest("DELETE", virtualServer.URL+"/account", strings.NewReader(""))
		r.Header.Set("Authorization", "Bearer "+token)
		resp, err := httpClient.Do(r)
		if !assert.NoError(t, err) {
			return
		}

		if !assert.Equal(t, http.StatusOK, resp.StatusCode, "deleting request status must be 200") {
			return
		}
	})
}

func TestPutAccount(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	t.Run("Successful update", func(t *testing.T) {
		login := createNewAcc()
		token := loginNewAcc(login)

		// Update request
		r, _ := http.NewRequest("PUT", virtualServer.URL+"/account", strings.NewReader(`{
			"name": "some name"
		}`))
		r.Header.Set("Authorization", "Bearer "+token)
		resp, err := httpClient.Do(r)
		if !assert.NoError(t, err) {
			return
		}

		if !generalCheck(t, resp, http.StatusOK) {
			return
		}
	})
}

func TestUseUrl(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	t.Run("Successful url use", func(t *testing.T) {
		login := createNewAcc()
		token := loginNewAcc(login)

		// Request create url
		r, _ := http.NewRequest("POST", virtualServer.URL+"/links", strings.NewReader(`{
			"short_url": "example",
			"full_url": "http://example.com"
		}`))
		r.Header.Set("Authorization", "Bearer "+token)
		_, _ = httpClient.Do(r)

		// Use-url request
		r, _ = http.NewRequest("GET", virtualServer.URL+"/use-url?short_url=example", nil)
		r.Header.Set("Authorization", "Bearer "+token)
		resp, err := httpClient.Do(r)
		if !assert.NoError(t, err) {
			return
		}

		// Check response
		if !assert.Equal(t, http.StatusOK, resp.StatusCode, "status must be 200") {
			return
		}
		decoder := json.NewDecoder(resp.Body)
		var urlData struct {
			FullUrl string `json:"fullUrl"`
		}
		if !assert.NoError(t, decoder.Decode(&urlData)) {
			return
		}
		if !assert.Equal(t, "http://example.com", urlData.FullUrl) {
			return
		}
	})
}

func TestGetLinks(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	t.Run("Successful get users link", func(t *testing.T) {
		login := createNewAcc()
		token := loginNewAcc(login)

		// Request create url
		r, _ := http.NewRequest("POST", virtualServer.URL+"/links", strings.NewReader(`{
			"short_url": "example",
			"full_url": "http://example.com"
		}`))
		r.Header.Set("Authorization", "Bearer "+token)
		_, _ = httpClient.Do(r)

		// Get links request
		r, _ = http.NewRequest("GET", virtualServer.URL+"/links", nil)
		r.Header.Set("Authorization", "Bearer "+token)
		resp, err := httpClient.Do(r)
		if !assert.NoError(t, err) {
			return
		}

		// Check response
		if !assert.Equal(t, http.StatusOK, resp.StatusCode, "status must be 200") {
			return
		}
		decoder := json.NewDecoder(resp.Body)
		var linksData []struct {
			Id         string `json:"id"`
			UserId     string `json:"user_id"`
			ShortUrl   string `json:"short_url"`
			FullUrl    string `json:"full_url"`
			UsageCount int    `json:"usage_count"`
			CreatedAt  string `json:"created_at"`
		}
		if !assert.NoError(t, decoder.Decode(&linksData)) {
			return
		}
		if !assert.Equal(t, "http://example.com", linksData[0].FullUrl) {
			return
		}
	})
}

func newVirtualServer() *httptest.Server {
	gin.SetMode(gin.TestMode)

	userDB := dbmemory.NewUserRepo()
	user := userrepo.NewUserRepo(userDB)

	linkDB := dbmemory.NewLinkRepo()
	link := linkrepo.NewLinkRepo(linkDB)
	fullUrlSearch := fullurlsearch.NewFullUrl(linkDB)

	hds := &handlers.Handlers{
		Userrepo:      user,
		Linkrepo:      link,
		Fullurlsearch: fullUrlSearch,
	}

	router := NewRouter(hds)
	srv := httptest.NewServer(router)

	return srv
}

func createNewAcc() (login string) {
	login = uuid.New().String()
	r, _ := http.NewRequest("POST", virtualServer.URL+"/regist", strings.NewReader(`{
		"name": "emiel",
		"login": "`+login+`",
		"password": "123456"
	}`))
	resp, _ := httpClient.Do(r)
	decoder := json.NewDecoder(resp.Body)
	var regData struct {
		UserId string `json:"userId"`
	}
	_ = decoder.Decode(&regData)
	//return login, regData.UserId
	return login
}

func loginNewAcc(login string) (token string) {
	r, _ := http.NewRequest("POST", virtualServer.URL+"/login", strings.NewReader(`{
		"login": "`+login+`",
		"password": "123456"
	}`))
	resp, _ := httpClient.Do(r)
	decoder := json.NewDecoder(resp.Body)
	var logData struct {
		Token string `json:"token"`
	}
	_ = decoder.Decode(&logData)

	return logData.Token
}

func generalCheck(t *testing.T, resp *http.Response, status int) bool {
	message, err := io.ReadAll(resp.Body)
	if !assert.NoError(t, err) {
		return false
	}

	if len(message) != 0 {
		t.Log(string(message))
	}

	if !assert.Equal(t, status, resp.StatusCode, "status must be "+strconv.Itoa(status)) {
		return false
	}
	if !assert.NotEmpty(t, message, "message cannot be empty") {
		return false
	}

	return true
}
