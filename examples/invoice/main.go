package main

import (
	"context"
	"encoding/json"
	"github.com/djcopley/quickbooks-sdk-go"
	"github.com/djcopley/quickbooks-sdk-go/model"
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

	qbClient := quickbooks.NewClient(
		quickbooks.Sandbox,
		client,
		realmId,
	)

	obj := &model.Invoice{
		Line: []model.Line{
			{
				DetailType: "SalesItemLineDetail",
				SalesItemLineDetail: model.SalesItemLineDetail{
					ItemRef: model.ItemRef{
						Name:  "services",
						Value: "1",
					},
				},
				Amount: 100,
			},
		},
		CustomerRef: model.CustomerRef{
			Value: "1",
		},
	}

	res, err := qbClient.CreateObject(obj)
	if err != nil {
		log.Println("Unable to create object:", err)
		return
	}

	var newObj model.InvoiceResponse
	err = json.Unmarshal(res, &newObj)
	if err != nil {
		log.Println("Unable to unmarshal json:", err)
		return
	}

	log.Println(newObj)
}
