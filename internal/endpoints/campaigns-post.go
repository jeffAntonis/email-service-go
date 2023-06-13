package endpoints

import (
	"email-service-go/internal/contract"
	internalerros "email-service-go/internal/internal-erros"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaignService.Create(request)

	if err != nil {
		if errors.Is(err, internalerros.ErrInternal) {
			return nil, 500, err
		} else {
			return nil, 400, err
		}
	}
	return map[string]string{"id": id}, 201, nil
}
