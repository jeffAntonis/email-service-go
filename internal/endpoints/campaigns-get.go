package endpoints

import (
	internalerros "email-service-go/internal/internal-erros"
	"errors"
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaigns, err := h.CampaignService.Repository.Get()

	if err != nil {
		if errors.Is(err, internalerros.ErrInternal) {
			return nil, 500, err
		} else {
			return nil, 400, err
		}
	}

	return campaigns, 200, nil
}
