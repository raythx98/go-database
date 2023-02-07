package command

import (
	"context"
	"fmt"
	"github.com/raythx98/go-database/service/receiver"
)

type GetFullLinkCommand struct {
	Database      receiver.Database
	ShortenedLink string
}

func (c *GetFullLinkCommand) Execute() (string, error) {
	// business logic
	fullLink, err := c.Database.GetFullLink(context.Background(), c.ShortenedLink)
	if err != nil {
		return "", fmt.Errorf("error retrieving from database, %w", err)
	}
	return fullLink, nil
}
