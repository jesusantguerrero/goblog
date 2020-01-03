package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var postJSON = `{
	"title": "The post2",
	"content": "this is a really long post"
}`

var postUpdateJSON = `{
	"title": "The post updated",
	"content": "this is a really long post"
}`

var createdPostId string

func TestGetPosts(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/posts")

	if assert.NoError(t, GetResources(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	t.Log(rec.Body.String())
}

func TestGetPost(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, GetResource(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreatePost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/posts")

	if assert.NoError(t, CreateResource(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		createdPostId = rec.Body.String()
	}
}

func TestUpdatePost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(postUpdateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/posts/:id")
	c.SetPath("id")
	c.SetParamValues("2")

	if assert.NoError(t, CreateResource(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
