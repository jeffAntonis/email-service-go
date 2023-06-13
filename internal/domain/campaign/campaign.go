package campaign

import (
	internalerros "email-service-go/internal/internal-erros"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	Id        string `validate:"required"`
	Name      string `validate:"min=5,max=24"`
	CreatedOn time.Time
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	campaign := &Campaign{
		Id:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}

	err := internalerros.ValidadeStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}
