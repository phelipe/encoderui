package encoders

import "encoding/base64"

type Base64 struct {
}

func (c Base64) Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (c Base64) Decode(data string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", DecodeError
	}
	return string(decoded), nil
}
