package encoders

import (
	"testing"
)

func TestBase64_Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should encode a string in base64",
			args: args{
				data: "test_encode_string",
			},
			want: "dGVzdF9lbmNvZGVfc3RyaW5n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Base64{}
			if got := c.Encode(tt.args.data); got != tt.want {
				t.Errorf("Base64.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64_Decode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should decode a string in base64",
			args: args{
				data: "dGVzdF9lbmNvZGVfc3RyaW5n",
			},
			want:    "test_encode_string",
			wantErr: false,
		},
		{
			name: "Should return error",
			args: args{
				data: "this_is_not_base_64",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Base64{}
			got, err := c.Decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Base64.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
