package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type CreateDomesticPaymentRequest struct {
	Amount               string `json:"amount" binding:"required"`
	BankID               string `json:"bank_id" binding:"required"`
	AccountID            string `json:"account_id" binding:"required"`
	PayeeAccountName     string `json:"payee_account_name" binding:"required"`
	PayeeAccountNumber   string `json:"payee_account_number" binding:"required"`
	PayeeAccountSortCode string `json:"payee_account_sort_code" binding:"required"`
	PaymentReference     string `json:"payment_reference" binding:"required"`
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

		authorisationType := "Single"
		account := models.DomesticPaymentConsentCreditorAccount{
			Identification:          &paymentRequest.PayeeAccountNumber,
			Name:                    &paymentRequest.PayeeAccountName,
			SchemeName:              &authorisationType, // todo some schema name??
		}
		debtorAccount := models.DomesticPaymentConsentDebtorAccount{
			Identification: &paymentRequest.AccountID,
			SchemeName:     &authorisationType, // todo some schema name??
		}
		id := uuid.New().String()
		currency := "usd"
		if registerResponse, err = clients.AcpClient.Openbanking.CreateDomesticPaymentConsent(
			openbanking.NewCreateDomesticPaymentConsentParams().
				WithTid(clients.AcpClient.TenantID).
				WithAid(clients.AcpClient.ServerID).
				WithRequest(&models.DomesticPaymentConsentRequest{
					Data: &models.DomesticPaymentConsentRequestData{
						Authorisation:     &models.DomesticPaymentConsentAuthorisation{
							AuthorisationType:  &authorisationType,
							CompletionDateTime: strfmt.DateTime(time.Now().Add(time.Hour)),
						},
						Initiation:        &models.DomesticPaymentConsentDataInitiation{
							CreditorAccount:           &account,
							DebtorAccount:             &debtorAccount,
							EndToEndIdentification:    &id,
							InstructedAmount:          &models.DomesticPaymentConsentInstructedAmount{
								Amount:   &paymentRequest.Amount,
								Currency: &currency,
							},
							InstructionIdentification: &id,
						},
						ReadRefundAccount: "No",
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
