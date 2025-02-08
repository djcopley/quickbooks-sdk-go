package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/djcopley/quickbooks-sdk-go"
	"github.com/djcopley/quickbooks-sdk-go/examples/env"
	"github.com/djcopley/quickbooks-sdk-go/invoice"
	"golang.org/x/oauth2"
)

func main() {
	environment := env.GetEnvironment()

	qbOAuthConfig := &oauth2.Config{
		ClientID:     environment.ClientID,
		ClientSecret: environment.ClientSecret,
		Scopes:       []string{"com.intuit.quickbooks.accounting"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer",
			AuthURL:  "https://appcenter.intuit.com/connect/oauth2",
		},
		RedirectURL: environment.RedirectURL,
	}

	t := &oauth2.Token{
		AccessToken:  environment.AccessToken,
		RefreshToken: environment.RefreshToken,
	}

	client := qbOAuthConfig.Client(context.Background(), t)

	qbClient := quickbooks.NewService(
		quickbooks.Sandbox,
		client,
		environment.RealmID,
	)

	invoices, err := quickbooks.NewQuery[invoice.Invoice](qbClient).All()
	if err != nil {
		panic(err)
	}

	pretty, err := json.MarshalIndent(invoices, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pretty))
}
