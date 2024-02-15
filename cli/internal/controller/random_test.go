package controller

import (
	"driving-journal-estimate/public/calendar"
	"driving-journal-estimate/public/day"
	"driving-journal-estimate/public/logger"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRandomController_Random(t *testing.T) {
	type fields struct {
		Month *calendar.Month
	}
	type args struct {
		c *gin.Context
	}
	type result struct {
		message string
		code    int
	}
	req, _ := http.NewRequest("GET", "/random", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Params = gin.Params{
		gin.Param{"total", "30000"},
		gin.Param{"days", "30"},
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result result
	}{
		{
			name: "success",
			fields: fields{
				Month: &calendar.Month{
					Days:   make([]*day.Config, 0),
					Logger: logger.New(logger.DEFAULT),
					Total:  30000,
				},
			},
			args: args{
				c: c,
			},
			result: result{
				message: `{"error":"missing form body"}`,
				code:    http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RandomController{
				Month: tt.fields.Month,
			}
			r.Random(tt.args.c)

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
			if string(b) != tt.result.message {
				t.Errorf("Random() message = '%s', want '%s'", string(b), tt.result.message)
			}
		})
	}
}
