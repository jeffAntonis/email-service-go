package endpoints

import "email-service-go/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
