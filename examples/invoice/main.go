package main

import (
	"context"
	"github.com/djcopley/quickbooks-sdk-go"
	"github.com/djcopley/quickbooks-sdk-go/invoice"
	"golang.org/x/oauth2"
	"log"
)

func main() {
	qbOAuthConfig := &oauth2.Config{
		ClientID:     "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ClientSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Scopes:       []string{"com.intuit.quickbooks.accounting"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer",
			AuthURL:  "https://appcenter.intuit.com/connect/oauth2",
		},
		RedirectURL: "localhost:8080/oauth/callback",
	}
	t := &oauth2.Token{
		AccessToken:  "",
		RefreshToken: "",
	}

	realmId := "xxxxxxxxxxxxxxxx"

	client := qbOAuthConfig.Client(context.Background(), t)

	qbClient := quickbooks.NewService(
		quickbooks.Sandbox,
		client,
		realmId,
	)

	obj := &invoice.Invoice{
		Line: []invoice.Line{
			{
				DetailType: "SalesItemLineDetail",
				SalesItemLineDetail: invoice.SalesItemLineDetail{
					ItemRef: invoice.ItemRef{
						Name:  "Services",
						Value: "1",
					},
				},
				Amount: 100.0,
			},
		},
		CustomerRef: invoice.CustomerRef{
			Value: "1",
		},
	}

	var response invoice.Response
	err := qbClient.CreateEntity(obj, &response)
	if err != nil {
		log.Fatal(err)
	}
}
