package procedures

import (
	"encoding/base64"
	"strings"

	"backend/modules/mng"
	"backend/types/model"
	"backend/types/response"
	"backend/utils/value"
)

func Paired(sessionNo string, pairedSessionNo string) (forwardUrl *string, pairedWith *string, e error) {
	// * Generate session hash
	sessionHash, errr := GenerateSessionHash(sessionNo)
	if errr != nil {
		return nil, nil, errr
	}

	// * Decode pair base64 sessionNo
	decoded, err := base64.StdEncoding.DecodeString(pairedSessionNo)
	if err != nil {
		return nil, nil, response.Error(true, "Unable to decode session no", err)
	}

	// * Generate forward link
	forwardLink := GenerateForwardLink(*sessionHash)

	// * Log successful pair
	newPairLog := &model.PairLog{
		SessionNo: &sessionNo,
		Action:    value.Ptr("pair"),
		Attribute: map[string]any{
			"hash":       sessionHash,
			"pairedWith": pairedSessionNo,
		},
	}
	if err := mng.PairLog.Create(newPairLog); err != nil {
		return nil, nil, response.Error(true, "Unable to create pair log", err)
	}

	return forwardLink, value.Ptr(strings.TrimSpace(string(decoded))), nil
}
