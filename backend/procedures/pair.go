package procedures

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"backend/modules/mng"
	"backend/types/model"
	"backend/types/response"
)

func Paired(sessionId primitive.ObjectID, pairedSessionId primitive.ObjectID, commitId primitive.ObjectID, pairedCommitId primitive.ObjectID) (pairedEmail *string, forwardUrl *string, e error) {
	// * Query session info
	session := new(model.Session)
	if err := mng.Session.FindByID(sessionId, session); err != nil {
		return nil, nil, response.Error(true, "Unable to fetch session", err)
	}

	// * Query paired session info
	pairedSession := new(model.Session)
	if err := mng.Session.FindByID(pairedSessionId, pairedSession); err != nil {
		return nil, nil, response.Error(true, "Unable to fetch paired session", err)
	}

	// * Generate forward link
	forwardLink := GenerateForwardLink(*session.Hash)

	// * Create match
	match := &model.PairMatch{
		CommitA: &commitId,
		CommitB: &pairedCommitId,
		Diff:    nil,
	}
	if err := mng.PairMatch.Create(match); err != nil {
		return nil, nil, response.Error(true, "Unable to create pair match", err)
	}

	return pairedSession.Email, forwardLink, nil
}
