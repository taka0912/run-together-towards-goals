package api

import "testing"

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
			if got := GetHello(tt.args.userName); (got != tt.ans) == tt.res {
				t.Errorf("GetHello() = %v, want %v", got, tt.ans)
			}
		})
	}
}
