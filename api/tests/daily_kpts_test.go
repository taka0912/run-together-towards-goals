package tests

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/api"
	"testing"
)

func TestGetHello(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name string
		args args
		ans  string
		res  bool
	}{
		{
			name: "normal-test",
			args: args{userName: "test"},
			ans:  "Hello, test!!",
			res:  true,
		},
		{
			name: "wrong-test",
			args: args{userName: "test"},
			ans:  "Hello, taro!!",
			res:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.GetHello(tt.args.userName); (got != tt.ans) == tt.res {
				t.Errorf("GetHello() = %v, want %v", got, tt.ans)
			}
		})
	}
}

func TestHandler_PostDailyKpt(t *testing.T) {
	tests := []struct {
		name    string
		field   string
		request Request
		urlBody api.DailyKpt
		want    responseJson
	}{
		{
			name: "normal test",
			request: Request{
				Method: "POST",
				Url:    "http://localhost:8080/api/daily_kpt/add",
				Body: api.DailyKpt{
					UserID:  "1",
					Keep:    "test",
					Problem: "test",
					Try:     "test",
				},
			},
			want: responseJson{Code: 200, Msg: "Created"},
		},
		{
			name: "UserID not found",
			request: Request{
				Method: "POST",
				Url:    "http://localhost:8080/api/daily_kpt/add",
				Body: api.DailyKpt{
					UserID: "0",
				},
			},
			want: responseJson{Code: 400},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tryTestRequest(tt.request)
			resCode := res.(map[string]interface{})["code"]
			resMsg := res.(map[string]interface{})["msg"]
			if resCode != tt.want.Code {
				t.Errorf("Get Code = %v, want %v, Get Msg %v", resCode, tt.want.Code, resMsg)
			}
		})
	}
}
