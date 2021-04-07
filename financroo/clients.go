package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudentity/acp-client-go/client/oauth2"
	"github.com/cloudentity/acp-client-go/models"
	obc "github.com/cloudentity/openbanking-sample-apps/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Clients struct {
	AcpAccountsClient acpclient.Client
	AcpPaymentsClient acpclient.Client
	BankClient        OpenbankingClient
}

func (c *Clients) RenewAccountsToken(bank ConnectedBank) (*models.TokenResponse, error) {
	var (
		resp *oauth2.TokenOK
		err  error
	)

	if resp, err = c.AcpAccountsClient.Acp.Oauth2.Token(
		oauth2.NewTokenParams().
			WithAid(c.AcpAccountsClient.ServerID).
			WithTid(c.AcpAccountsClient.TenantID).
			WithClientID(&c.AcpAccountsClient.Config.ClientID).
			WithGrantType("refresh_token").
			WithRefreshToken(&bank.RefreshToken)); err != nil {
		return nil, errors.Wrapf(err, "can't renew access token for a bank: %s", bank.BankID)
	}

	return resp.Payload, nil
}

func InitClients(config Config) (map[BankID]Clients, error) {
	var (
		clients              = map[BankID]Clients{}
		acpAccountsWebClient acpclient.Client
		acpPaymentsWebClient acpclient.Client
		bankClient           OpenbankingClient
		err                  error
	)

	for _, bank := range config.Banks {
		if acpAccountsWebClient, err = NewAcpClient(config, bank, "/api/callback"); err != nil {
			return clients, errors.Wrapf(err, "failed to init acp web client for bank: %s", bank.ID)
		}

		if acpPaymentsWebClient, err = NewAcpClient(config, bank, "/api/domestic/callback"); err != nil {
			return clients, errors.Wrapf(err, "failed to init acp web client for bank: %s", bank.ID)
		}

		if bankClient, err = NewOpenbankingClient(bank); err != nil {
			return clients, errors.Wrapf(err, "failed to init client for bank: %s", bank.ID)
		}

		clients[bank.ID] = Clients{
			AcpAccountsClient: acpAccountsWebClient,
			AcpPaymentsClient: acpPaymentsWebClient,
			BankClient:        bankClient,
		}
	}

	return clients, nil
}

func NewAcpClient(c Config, cfg BankConfig, redirect string) (acpclient.Client, error) {
	var (
		issuerURL, authorizeURL, redirectURL *url.URL
		client                               acpclient.Client
		err                                  error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if authorizeURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/authorize", c.ACPURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if redirectURL, err = url.Parse(fmt.Sprintf("%s%s", c.UIURL, redirect)); err != nil {
		return client, err
	}

	requestObjectExpiration := time.Minute * 10
	config := acpclient.Config{
		ClientID:                    cfg.AcpClient.ClientID,
		IssuerURL:                   issuerURL,
		AuthorizeURL:                authorizeURL,
		RedirectURL:                 redirectURL,
		RequestObjectSigningKeyFile: cfg.AcpClient.KeyFile,
		RequestObjectExpiration:     &requestObjectExpiration,
		Scopes:                      []string{"accounts", "payments", "openid", "offline_access"},
		Timeout:                     cfg.AcpClient.Timeout,
		CertFile:                    cfg.AcpClient.CertFile,
		KeyFile:                     cfg.AcpClient.KeyFile,
		RootCA:                      cfg.AcpClient.RootCA,
	}

	if client, err = acpclient.New(config); err != nil {
		return client, err
	}

	return client, nil
}

func NewLoginClient(c Config) (acpclient.Client, error) {
	var (
		issuerURL *url.URL
		client    acpclient.Client
		err       error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", c.ACPInternalURL, c.Login.TenantID, c.Login.ServerID)); err != nil {
		return client, err
	}

	config := acpclient.Config{
		ClientID:  c.Login.ClientID,
		IssuerURL: issuerURL,
		Timeout:   c.Login.Timeout,
		RootCA:    c.Login.RootCA,
	}

	if client, err = acpclient.New(config); err != nil {
		return client, err
	}

	return client, nil
}

type OpenbankingClient struct {
	*obc.Openbanking
}

func NewOpenbankingClient(config BankConfig) (OpenbankingClient, error) {
	var (
		c   = OpenbankingClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.URL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	c.Openbanking = obc.New(httptransport.NewWithClient(
		u.Host,
		"/",
		[]string{u.Scheme},
		hc,
	), nil)

	return c, nil
}
