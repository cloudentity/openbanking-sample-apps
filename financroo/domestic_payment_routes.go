package main

import (
	"fmt"
	"net/http"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type CreateDomesticPaymentRequest struct {
 // todo require minimal amount of data to create domestic payment
}

func (s *Server) CreateDomesticPayment() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			bankID             = BankID(c.Param("bankId"))
			clients            Clients
			ok                 bool
			registerResponse   *openbanking.CreateDomesticPaymentConsentCreated
			encodedCookieValue string
			loginURL           string
			data               = gin.H{}
			paymentRequest     = CreateDomesticPaymentRequest{}
			user               User
			err                error
		)

		if user, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if err = c.BindJSON(&paymentRequest); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %+v", err))
			return
		}

		if clients, ok = s.Clients[bankID]; !ok {
			c.String(http.StatusBadRequest, fmt.Sprintf("client not configured for bank: %s", bankID))
		}

		// todo create consent request
		if registerResponse, err = clients.AcpClient.Openbanking.CreateDomesticPaymentConsent(
			openbanking.NewCreateDomesticPaymentConsentParams().
				WithTid(clients.AcpClient.TenantID).
				WithAid(clients.AcpClient.ServerID).
				WithRequest(&models.DomesticPaymentConsentRequest{
					Data: &models.DomesticPaymentConsentRequestData{
						Authorisation:     &models.DomesticPaymentConsentAuthorisation{
							AuthorisationType:  nil,
							CompletionDateTime: strfmt.DateTime{},
						},
						Initiation:        &models.DomesticPaymentConsentDataInitiation{
							CreditorAccount:           nil,
							CreditorPostalAddress:     nil,
							DebtorAccount:             nil,
							EndToEndIdentification:    nil,
							InstructedAmount:          nil,
							InstructionIdentification: nil,
							LocalInstrument:           "",
							RemittanceInformation:     nil,
							SupplementaryData:         nil,
						},
						ReadRefundAccount: "No",
						SCASupportData:    nil,
					},
				}),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register domestic payment consent: %+v", err))
			return
		}

		s.ConsentCreatedResponse(c, bankID, registerResponse.Payload.Data.ConsentID, user, loginURL, err, clients, encodedCookieValue, data)
	}
}
