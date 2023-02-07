package command

import (
	"context"
	"fmt"
	"github.com/aidarkhanov/nanoid"
	"github.com/raythx98/go-database/service/receiver"
	database "github.com/raythx98/go-database/sqlc/output"
)

type AddLinkCommand struct {
	Database     receiver.Database
	FullLink     string `json:"full_link"`
	CustomLink   string `json:"custom_link,omitempty"`
	InvalidateAt string `json:"invalidate_at,omitempty"` // TODO: add datetime unmarshal
	NumRedirects int    `json:"number_redirects,omitempty"`
}

func (c *AddLinkCommand) Execute() (string, error) {
	// business logic
	shortenedLink, _ := nanoid.Generate(nanoid.DefaultAlphabet, 8)

	err := c.Database.AddLink(context.Background(), database.AddLinkParams{Link: c.FullLink, ShortenedLink: shortenedLink})
	if err != nil {
		return "", fmt.Errorf("error inserting to database, %w", err)
	}

	return shortenedLink, nil
}
