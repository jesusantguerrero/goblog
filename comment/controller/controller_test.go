package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var commentJSON = `{
	"user_id": 1,
	"post_id": 1,
	"edit_mode": false,
	"content": "this is a really long comment"
}`

var commentUpdateJSON = `{
	"user_id": 1,
	"post_id": 1,
	"edit_mode": false,
	"content": "this is a really long comment updated"
}`

var createdPostId string

func TestGetComments(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/comments")

	ctrl := Controller{}
	ctrl._Boot(e)

	if assert.NoError(t, ctrl.GetResources(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	t.Log(rec.Body.String())
}

func TestGetComment(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/comments/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	ctrl := Controller{}
	ctrl._Boot(e)

	if assert.NoError(t, ctrl.GetResource(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateComment(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(commentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/comments")
	ctrl := Controller{}
	ctrl._Boot(e)

	if assert.NoError(t, ctrl.CreateResource(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		createdPostId = rec.Body.String()
	}
}

func TestUpdateComment(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(commentUpdateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/comments/:id")
	c.SetPath("id")
	c.SetParamValues("1")
	ctrl := Controller{}
	ctrl._Boot(e)

	if assert.NoError(t, ctrl.UpdateResource(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
