package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/qndaa/backend/src/db"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestLogin(t *testing.T) {

	users := []LoginParams{
		{
			Username: "marko",
			Password: "marko",
		}, {
			Username: "igor",
			Password: "igor",
		},
	}

	tt := []struct {
		name   string
		code   int
		method string
		want   uint
		body   *LoginParams
	}{
		{
			name:   "Login success!",
			code:   200,
			method: http.MethodPost,
			want:   0,
			body:   &users[0],
		}, {
			name:   "Login failed!",
			code:   400,
			method: http.MethodPost,
			want:   0,
			body:   nil,
		}, {
			name:   "Login failed!",
			code:   400,
			method: http.MethodPost,
			want:   0,
			body:   &users[1],
		},
	}

	db.Connect()
	db.AutoMigration()
	a := App{}
	a.Init()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			b, err := json.Marshal(tc.body)
			if err != nil {
				fmt.Errorf("Problems: %v", err)
			}
			r := httptest.NewRequest(tc.method, fmt.Sprintf("/api/login"), bytes.NewBuffer(b))
			rr := httptest.NewRecorder()
			a.GetRouter().ServeHTTP(rr, r)
			if rr.Code != tc.code {
				t.Errorf("Want status '%d', got '%d'", tc.code, rr.Code)
			}
		})
	}
}

func TestLoginUnauthorized(t *testing.T) {

	tt := struct {
		name   string
		code   int
		method string
		want   uint
		body   *struct {
			Username int `json:"username"`
		}
	}{
		name:   "Unauthorized",
		code:   401,
		method: http.MethodPost,
		want:   0,
		body: &struct {
			Username int `json:"username"`
		}{Username: 1},
	}

	db.Connect()
	db.AutoMigration()
	a := App{}
	a.Init()

	t.Run(tt.name, func(t *testing.T) {
		b, _ := json.Marshal(tt.body)
		b = append(b, '"')
		r := httptest.NewRequest(tt.method, fmt.Sprintf("/api/login"), bytes.NewBuffer(b))
		rr := httptest.NewRecorder()
		a.GetRouter().ServeHTTP(rr, r)
		if rr.Code != tt.code {
			t.Errorf("Want status '%d', got '%d'", tt.code, rr.Code)
		}

	})
}
