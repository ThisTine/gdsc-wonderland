package pairEndpoint

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"backend/modules/mng"
	"backend/procedures"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/text"
	"backend/utils/value"
)

func InitialGetHandler(c *fiber.Ctx) error {
	// * Parse query
	query := new(payload.InitialQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse query", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return response.Error(false, "Unable to validate query", err)
	}

	// * Generate session hash
	email, hash, err := procedures.ExtractSessionInfo(*query.SessionNo)
	if err != nil {
		return err
	}

	// * Get session id
	session := new(model.Session)
	if err := mng.Session.First(
		bson.M{
			model.SessionEmail: email,
		},
		session,
	); errors.Is(err, mongo.ErrNoDocuments) {
		// * Create new session
		session = &model.Session{
			Email:        email,
			Hash:         hash,
			MatchedCount: value.Ptr[int64](0),
		}
		if err := mng.Session.Create(session); err != nil {
			return response.Error(true, "Unable to create session", err)
		}
	} else if err != nil {
		return response.Error(true, "Unable to find session", err)
	}

	// * Get pictures
	var pictures []*model.PairPicture
	if err := mng.PairPicture.SimpleFind(
		&pictures,
		bson.M{
			model.PairPictureVisible: true,
		},
	); err != nil {
		return response.Error(true, "Unable to find pictures", err)
	}

	// * Map pictures to payload
	picturePayloads, err := value.Iterate(pictures, func(p *model.PairPicture) (*payload.Picture, *response.ErrorInstance) {
		// * Read file from resource
		bytes, err := os.ReadFile(filepath.Join("resources", "picture", *p.Filename))
		if err != nil {
			return nil, response.Error(true, "Unable to read picture file", err)
		}

		// * Generate src base64 string
		src := "data:image/png;base64,"
		src += base64.RawStdEncoding.EncodeToString(bytes)

		return &payload.Picture{
			Id:          p.ID,
			Src:         &src,
			Title:       p.Title,
			Description: p.Description,
		}, nil
	})
	if err != nil {
		return err
	}

	return c.JSON(response.Info(&payload.InitialResponse{
		SessionId: session.ID,
		Email:     session.Email,
		Pictures:  picturePayloads,
	}))
}
