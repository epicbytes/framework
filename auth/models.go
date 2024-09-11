package auth

import (
	"encoding/base64"
	"github.com/goccy/go-json"
)

type PolicyKey string

func (key PolicyKey) String() string {
	return string(key)
}

type ResourceAccess struct {
	Cascade struct {
		Roles []string `json:"roles"`
	} `json:"cascade"`
}

type BaseUserModel struct {
	ID                string                 `json:"id"`
	Username          string                 `json:"username"`
	Name              string                 `json:"name"`
	FirstName         string                 `json:"first_name"`
	LastName          string                 `json:"last_name"`
	Email             string                 `json:"email"`
	ResourceAccessRaw map[string]interface{} `json:"resource_access"`
}

func (u *BaseUserModel) Encode() string {
	b, _ := json.Marshal(u)
	sEnc := base64.StdEncoding.EncodeToString(b)
	return sEnc
}

func (u *BaseUserModel) Decode(data string) error {
	sDec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(sDec, &u)
	if err != nil {
		return err
	}
	return nil
}
