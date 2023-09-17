package pairEndpoint

import (
	"errors"
	"time"

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
			"sessionNo": body.SessionNo,
			"createdAt": bson.M{
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
			SessionNo: body.SessionNo,
			Action:    value.Ptr("duplicate"),
		}
		if err := mng.PairLog.Create(newPairLog); err != nil {
			return response.Error(true, "Unable to create pair log", err)
		}

		// * Response
		return response.Error(false, "Duplicate commit within 5 seconds", nil)
	}

	// * Check for existing pair
	pairThreshold := time.Now().Add(-1 * time.Second)
	pairCommit := new(model.PairCommit)
	if err := mng.PairCommit.First(
		bson.M{
			"itemNo": body.ItemNo,
			"pairedWith": bson.M{
				"$exists": true,
				"$ne":     nil,
			},
			"createdAt": bson.M{
				"$gte": pairThreshold,
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

	if pairCommit == nil {
		// # Case of not found matching pair
		// * Construct adding pair commit
		newPairCommit := &model.PairCommit{
			SessionNo:  body.SessionNo,
			ItemNo:     body.ItemNo,
			PairedWith: nil,
		}
		if err := mng.PairCommit.Create(newPairCommit); err != nil {
			return response.Error(true, "Unable to create pair commit", err)
		}

		// * Sleep
		time.Sleep(1 * time.Second)

		// * Check for paired with
		pairCommit = new(model.PairCommit)
		if err := mng.PairCommit.First(
			bson.M{
				"pairedWith": newPairCommit.ID,
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
		forwardLink, pairedWith, err := procedures.Paired(*body.SessionNo, *pairCommit.SessionNo)
		if err != nil {
			return err
		}

		return c.JSON(response.Info(&payload.ParingCommitResponse{
			Matched:     value.TruePtr,
			ForwardLink: forwardLink,
			PairedWith:  pairedWith,
		}))
	}

	// * Pair commit
	forwardLink, pairedWith, err := procedures.Paired(*body.SessionNo, *pairCommit.SessionNo)
	if err != nil {
		return err
	}

	return c.JSON(response.Info(&payload.ParingCommitResponse{
		Matched:     value.TruePtr,
		ForwardLink: forwardLink,
		PairedWith:  pairedWith,
	}))
}
