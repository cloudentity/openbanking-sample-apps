package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	"github.com/cloudentity/acp-client-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type CreateDomesticPaymentConsentRequest struct {
	Amount               string `json:"amount" binding:"required"`
	BankID               BankID `json:"bank_id" binding:"required"`
	AccountID            string `json:"account_id" binding:"required"`
	PayeeAccountName     string `json:"payee_account_name" binding:"required"`
	PayeeAccountNumber   string `json:"payee_account_number" binding:"required"`
	PayeeAccountSortCode string `json:"payee_account_sort_code" binding:"required"`
	PaymentReference     string `json:"payment_reference" binding:"required"`
}

func (s *Server) CreateDomesticPaymentConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			clients               Clients
			ok                    bool
			registerResponse      *openbanking.CreateDomesticPaymentConsentCreated
			paymentConsentRequest = CreateDomesticPaymentConsentRequest{}
			user                  User
			err                   error
		)

		if user, _, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if err = c.BindJSON(&paymentConsentRequest); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %+v", err))
			return
		}

		if clients, ok = s.Clients[paymentConsentRequest.BankID]; !ok {
			c.String(http.StatusBadRequest, fmt.Sprintf("client not configured for bank: %s", paymentConsentRequest.BankID))
		}

		schema := "UK.OBIE.SortCodeAccountNumber"
		authorisationType := "Single"
		account := models.DomesticPaymentConsentCreditorAccount{
			Identification: &paymentConsentRequest.PayeeAccountNumber,
			Name:           &paymentConsentRequest.PayeeAccountName,
			SchemeName:     &schema,
		}
		debtorAccount := models.DomesticPaymentConsentDebtorAccount{
			Identification: &paymentConsentRequest.AccountID,
			Name:           "myAccount", // todo
			SchemeName:     &schema,
		}
		id := uuid.New().String()[:10]
		currency := "USD"
		if registerResponse, err = clients.AcpPaymentsClient.Openbanking.CreateDomesticPaymentConsent(
			openbanking.NewCreateDomesticPaymentConsentParams().
				WithTid(clients.AcpPaymentsClient.TenantID).
				WithAid(clients.AcpPaymentsClient.ServerID).
				WithRequest(&models.DomesticPaymentConsentRequest{
					Data: &models.DomesticPaymentConsentRequestData{
						Authorisation: &models.DomesticPaymentConsentAuthorisation{
							AuthorisationType:  &authorisationType,
							CompletionDateTime: strfmt.DateTime(time.Now().Add(time.Hour)),
						},
						Initiation: &models.DomesticPaymentConsentDataInitiation{
							CreditorAccount:        &account,
							DebtorAccount:          &debtorAccount,
							EndToEndIdentification: &id,
							InstructedAmount: &models.DomesticPaymentConsentInstructedAmount{
								Amount:   &paymentConsentRequest.Amount,
								Currency: &currency,
							},
							InstructionIdentification: &id,
						},
						ReadRefundAccount: "No",
					},
					Risk: &models.PaymentRisk{},
				}),
			nil,
		); err != nil {
			errJson, _ := json.Marshal(err) // todo error handling for error mapping
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register domestic payment consent: %s", errJson))
			return
		}

		s.CreateConsentResponse(c, paymentConsentRequest.BankID, registerResponse.Payload.Data.ConsentID, user, clients.AcpPaymentsClient)
	}
}

func (s *Server) DomesticPaymentCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app        string
			appStorage = AppStorage{}
			err        error
		)

		if c.Query("error") != "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("acp returned an error: %s: %s", c.Query("error"), c.Query("error_description")))
			return
		}

		if app, err = c.Cookie("app"); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get app cookie: %+v", err))
			return
		}

		if err = s.SecureCookie.Decode("app", app, &appStorage); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to decode app storage: %+v", err))
			return
		}

		// todo get consent
		// todo create payment
		// todo redirect to confetti page

		c.SetCookie("app", "", -1, "/", "", false, true)

		c.Redirect(http.StatusFound, s.Config.UIURL)
	}
}
