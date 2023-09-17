package procedures

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"

	"backend/modules"
	"backend/modules/mng"
	"backend/types/model"
	"backend/types/response"
	"backend/utils/value"
)

func GenerateForwardLink(sessionHash string) *string {
	// * Generate forward link
	forwardLink := strings.Replace(mod.Conf.ForwardLink, "{{SESSION_HASH}}", sessionHash, 1)
	return &forwardLink
}

func GenerateSessionHash(sessionNo string) (*string, *string, *response.ErrorInstance) {
	// * Decode base64 sessionNo
	decoded, err := base64.StdEncoding.DecodeString(sessionNo)
	if err != nil {
		return nil, nil, response.Error(true, "Unable to decode session no", err)
	}

	// * Fetch session hash
	config := new(model.Config)
	if err := mng.Config.First(nil, config); err != nil {
		return nil, nil, response.Error(true, "Unable to fetch config", err)
	}

	// * Generate session hash
	h := hmac.New(sha256.New, []byte(*config.Secret))
	h.Write(decoded)
	sha := hex.EncodeToString(h.Sum(nil))

	// * Return session hash
	return &sha, value.Ptr(string(decoded)), nil
}
