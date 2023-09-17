package procedures

import (
	"encoding/base64"

	mod "backend/modules"
	"backend/types/response"
)

func GenerateForwardLink(sessionNo *string) (*string, *response.ErrorInstance) {
	// * Base64 decode sessionNo
	decoded, err := base64.StdEncoding.DecodeString(*sessionNo)
	if err != nil {
		return nil, response.Error(true, "Unable to decode session no", err)
	}

	// * Generate forward link
	forwardLink := mod.Conf.ForwardLink + string(decoded)
	return &forwardLink, nil
}
