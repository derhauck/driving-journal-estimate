package controller

import (
	"driving-journal-estimate/public/calendar"
	"driving-journal-estimate/public/day"
	"driving-journal-estimate/public/logger"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
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
