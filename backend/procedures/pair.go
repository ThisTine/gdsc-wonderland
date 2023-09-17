package procedures

import (
	"backend/modules/mng"
	"backend/types/model"
	"backend/types/response"
	"backend/utils/value"
)

func Paired(sessionNo string, pairedSessionNo string) (forwardUrl *string, pairedWith *string, e error) {
	// * Generate session hash
	sessionHash, decodedSessionNo, err := GenerateSessionHash(sessionNo)
	if err != nil {
		return nil, nil, err
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

	return forwardLink, decodedSessionNo, nil
}
