package pairEndpoint

import (
	"errors"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend/modules/mng"
	"backend/procedures"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
	"backend/utils/value"
)

func CommitPostHandler(c *fiber.Ctx) error {
	// * Parse body
	body := new(payload.ParingCommit)
	if err := c.BodyParser(body); err != nil {
		return response.Error(true, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	// * Query duplicate commit
	duplicateThreshold := time.Now().Add(-5 * time.Second)
	duplicateCommit := new(model.PairCommit)
	if err := mng.PairCommit.First(
		bson.M{
			model.PairCommitSessionId: body.SessionId,
			model.FieldCreatedAt: bson.M{
				"$gte": duplicateThreshold,
			},
		},
		duplicateCommit,
		&options.FindOneOptions{
			Sort: bson.M{
				"createdAt": -1,
			},
		}); errors.Is(err, mongo.ErrNoDocuments) {
		duplicateCommit = nil
	} else if err != nil {
		return response.Error(true, "Unable to find pair commit", err)
	}

	// * Check for duplicate commit
	if duplicateCommit != nil {
		// * Create new pair log
		newPairLog := &model.PairLog{
			SessionId: body.SessionId,
			Action:    value.Ptr("duplicate"),
		}
		if err := mng.PairLog.Create(newPairLog); err != nil {
			return response.Error(true, "Unable to create pair log", err)
		}

		// * Response
		difference := 5*time.Second - time.Now().Sub(*duplicateCommit.CreatedAt)

		// * Sprintf Duplicate commit within 4.12 seconds
		message := spew.Sprintf("Cool-down %.0f seconds %d ms remaining", difference.Seconds(), difference.Milliseconds()%1000)
		return response.Error(false, message, nil)
	}

	// * Check for parable commit
	pairThreshold := time.Now().Add(-1 * time.Second)
	pairCommit := new(model.PairCommit)
	if err := mng.PairCommit.First(
		bson.M{
			"itemNo": body.ItemId,
			"createdAt": bson.M{
				"$gte": pairThreshold,
			},
			"$or": bson.A{
				bson.M{
					"pairedWith": nil,
				},
				bson.M{
					"pairedWith": bson.M{
						"$exists": false,
					},
				},
			},
		},
		pairCommit,
		&options.FindOneOptions{
			Sort: bson.M{
				"createdAt": -1,
			},
		}); errors.Is(err, mongo.ErrNoDocuments) {
		pairCommit = nil
	} else if err != nil {
		return response.Error(true, "Unable to find pair commit", err)
	}

	// * Construct new pair commit
	newCommit := &model.PairCommit{
		SessionId:  body.SessionId,
		ItemNo:     body.ItemId,
		PairedWith: nil,
	}

	if pairCommit == nil {
		// # Case of not found matching pair
		if err := mng.PairCommit.Create(newCommit); err != nil {
			return response.Error(true, "Unable to create new pair commit", err)
		}

		// * Sleep
		time.Sleep(1 * time.Second)

		// * Check for paired with
		pairCommit = new(model.PairCommit)
		if err := mng.PairCommit.First(
			bson.M{
				"pairedWith": newCommit.ID,
			},
			pairCommit,
		); errors.Is(err, mongo.ErrNoDocuments) {
			pairCommit = nil
		} else if err != nil {
			return response.Error(true, "Unable to find pair commit", err)
		}

		if pairCommit == nil {
			return c.JSON(response.Info(&payload.ParingCommitResponse{
				Matched:     value.FalsePtr,
				ForwardLink: nil,
				PairedWith:  nil,
			}))
		}

		// * Pair commit
		pairedWith, forwardLink, err := procedures.Paired(*body.SessionId, *pairCommit.SessionId, *newCommit.ID, *pairCommit.ID)
		if err != nil {
			return err
		}

		return c.JSON(response.Info(&payload.ParingCommitResponse{
			Matched:     value.TruePtr,
			ForwardLink: forwardLink,
			PairedWith:  pairedWith,
		}))
	}

	// * Create new pair commit
	newCommit.PairedWith = pairCommit.ID
	if err := mng.PairCommit.Create(newCommit); err != nil {
		return response.Error(true, "Unable to create new pair commit", err)
	}

	// * Update pair commit
	pairCommit.PairedWith = newCommit.ID
	if err := mng.PairCommit.Update(pairCommit); err != nil {
		return response.Error(true, "Unable to update pair commit", err)
	}

	// * Pair commit
	pairedWith, forwardLink, err := procedures.Paired(*body.SessionId, *pairCommit.SessionId, *newCommit.ID, *pairCommit.ID)
	if err != nil {
		return err
	}

	return c.JSON(response.Info(&payload.ParingCommitResponse{
		Matched:     value.TruePtr,
		ForwardLink: forwardLink,
		PairedWith:  pairedWith,
	}))
}
