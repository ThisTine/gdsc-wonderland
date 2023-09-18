package procedures

import (
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"backend/modules/mng"
	"backend/types/model"
	"backend/types/response"
)

func Paired(sessionId primitive.ObjectID, pairedSessionId primitive.ObjectID, commit *model.PairCommit, pairedCommit *model.PairCommit) (pairedEmail *string, forwardUrl *string, duration *time.Duration, e error) {
	// * Query session info
	session := new(model.Session)
	if err := mng.Session.FindByID(sessionId, session); err != nil {
		return nil, nil, nil, response.Error(true, "Unable to fetch session", err)
	}

	// * Query paired session info
	pairedSession := new(model.Session)
	if err := mng.Session.FindByID(pairedSessionId, pairedSession); err != nil {
		return nil, nil, nil, response.Error(true, "Unable to fetch paired session", err)
	}

	// * Generate forward link
	forwardLink := GenerateForwardLink(*session.Hash)

	// * Calculate diff
	diff := time.Duration(math.Abs(float64(commit.CreatedAt.Sub(*pairedCommit.CreatedAt))))

	// * Create match
	match := &model.PairMatch{
		CommitA: commit.ID,
		CommitB: pairedCommit.ID,
		Diff:    &diff,
	}
	if err := mng.PairMatch.Create(match); err != nil {
		return nil, nil, nil, response.Error(true, "Unable to create pair match", err)
	}

	return pairedSession.Email, forwardLink, &diff, nil
}
