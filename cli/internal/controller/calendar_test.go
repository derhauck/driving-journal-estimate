package controller

import (
	"derhauck/driving-journal-estimate/public/calendar"
	"derhauck/driving-journal-estimate/public/day"
	"derhauck/driving-journal-estimate/public/logger"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func NewMonth(total float32) *calendar.Month {
	return &calendar.Month{
		make([]*day.Config, 0),
		total,
		logger.New(logger.DEFAULT),
	}
}
func NewContext() (*gin.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	return c, rec
}
func TestRandomController_Random(t *testing.T) {
	type fields struct {
		Month *calendar.Month
	}
	type args struct {
		req func() *http.Request
	}
	type result struct {
		message string
		code    int
	}

	gin.SetMode(gin.TestMode)
	tests := []struct {
		name   string
		fields fields
		args   args
		result result
	}{
		{
			name: "missing form parameters",
			fields: fields{
				Month: NewMonth(30000),
			},
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("GET", "/random", nil)
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					return req
				},
			},
			result: result{
				message: `{"error":"missing form body"}`,
				code:    http.StatusInternalServerError,
			},
		}, {
			name: "success",
			fields: fields{
				Month: NewMonth(30000),
			},
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("GET", "/random", nil)
					req.Form = make(url.Values)
					req.Form.Set("total", "30000")
					req.Form.Set("days", "30")

					return req
				},
			},
			result: result{
				message: `{"message":{"Days":[{"`,
				code:    http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CalendarController{
				Month: tt.fields.Month,
			}
			c, rec := NewContext()
			c.Request = tt.args.req()
			r.Random(c)

			if rec.Code != tt.result.code {
				t.Errorf("Random() status = %v, want %v", rec.Code, tt.result.code)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Error(err)
				}
			}(rec.Result().Body)

			b, err := io.ReadAll(rec.Result().Body)
			if err != nil {
				t.Error(err)
			}
			if !strings.Contains(string(b), tt.result.message) {
				t.Errorf("Random() message = '%s', should contain '%s'", string(b), tt.result.message)
			}
		})
	}
}

func TestCalendarController_Configuration(t *testing.T) {
	type fields struct {
		Month  *calendar.Month
		Logger logger.Inf
	}
	type args struct {
		req func() *http.Request
	}
	type result struct {
		message string
		code    int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		result result
	}{
		{
			name: "success with default",
			fields: fields{
				Month: NewMonth(30000),
			},
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/config", nil)
					req.Header.Set("Content-Type", "application/json")
					req.Body = io.NopCloser(strings.NewReader(`{"days":[{"date":"first","count":1}]}`))

					return req
				},
			},
			result: result{
				message: `{"message":{"Days":[{"`,
				code:    http.StatusOK,
			},
		},
		{
			name: "success with total",
			fields: fields{
				Month: NewMonth(30000),
			},
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/config", nil)
					req.Header.Set("Content-Type", "application/json")
					req.Body = io.NopCloser(strings.NewReader(`{"total": 30, "days":[{"date":"first","count":1}]}`))

					return req
				},
			},
			result: result{
				message: `{"message":{"Days":[{"`,
				code:    http.StatusOK,
			},
		},
		{
			name: "success with baseline",
			fields: fields{
				Month: NewMonth(30000),
			},
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/config", nil)
					req.Header.Set("Content-Type", "application/json")
					req.Body = io.NopCloser(strings.NewReader(`{"baseline":0.3,"days":[{"date":"first","count":1}]}`))

					return req
				},
			},
			result: result{
				message: `{"message":{"Days":[{"`,
				code:    http.StatusOK,
			},
		},
		{
			name: "fail empty body",
			fields: fields{
				Month: NewMonth(30000),
			},
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/config", nil)
					req.Header.Set("Content-Type", "application/json")

					return req
				},
			},
			result: result{
				message: `{"error":"invalid request"}`,
				code:    http.StatusInternalServerError,
			},
		},
		{
			name: "fail empty body",
			fields: fields{
				Month: NewMonth(30000),
			},
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest("POST", "/config", nil)
					req.Header.Set("Content-Type", "application/json")

					return req
				},
			},
			result: result{
				message: `{"error":"invalid request"}`,
				code:    http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CalendarController{
				Month:  tt.fields.Month,
				Logger: tt.fields.Logger,
			}

			c, rec := NewContext()
			c.Request = tt.args.req()

			r.Configuration(c)
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Error(err)
				}
			}(rec.Result().Body)

			b, err := io.ReadAll(rec.Result().Body)

			if err != nil {
				t.Error(err)
			}
			if rec.Code != tt.result.code {
				t.Errorf("Configuration() status = %v, want %v", rec.Code, tt.result.code)
			}
			if !strings.Contains(string(b), tt.result.message) {
				t.Errorf("Random() message = '%s', should contain '%s'", string(b), tt.result.message)
			}
		})
	}
}
