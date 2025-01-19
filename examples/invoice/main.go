package main

import (
	"context"
	"fmt"
	"time"

	"github.com/djcopley/quickbooks-sdk-go"
	"github.com/djcopley/quickbooks-sdk-go/invoice"
	"golang.org/x/oauth2"
)

func main() {
	qbOAuthConfig := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{"com.intuit.quickbooks.accounting"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer",
			AuthURL:  "https://appcenter.intuit.com/connect/oauth2",
		},
		RedirectURL: "http://localhost:8080/oauth/callback",
	}

	// // use PKCE to protect against CSRF attacks
	// verifier := oauth2.GenerateVerifier()

	// url := qbOAuthConfig.AuthCodeURL(
	// 	"state",
	// 	oauth2.AccessTypeOffline,
	// 	oauth2.S256ChallengeOption(verifier),
	// )

	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer cancel()

	// eg, ctx := errgroup.WithContext(ctx)

	// listener, err := net.Listen("tcp4", "localhost:8080")
	// if err != nil {
	// 	panic(err)
	// }
	// eg.Go(func() error {
	// 	defer listener.Close()

	// 	<-ctx.Done()
	// 	return ctx.Err()
	// })

	// eg.Go(func() error {
	// 	mux := http.NewServeMux()

	// 	mux.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
	// 		queryParams := r.URL.Query()
	// 		code := queryParams.Get("code")
	// 		realmId := queryParams.Get("realmId")
	// 		// requestId := queryParams.Get("state")

	// 		tok, err := qbOAuthConfig.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	// 		if err != nil {
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 		}

	// 		fmt.Println(realmId)
	// 		fmt.Println(tok.AccessToken)
	// 		fmt.Println(tok.RefreshToken)

	// 	})

	// 	return http.Serve(listener, mux)
	// })

	// fmt.Println(url)

	// eg.Wait()

	t := &oauth2.Token{
		AccessToken:  "",
		RefreshToken: "",
		Expiry:       time.Now(),
	}

	realmId := ""

	client := qbOAuthConfig.Client(context.Background(), t)

	qbClient := quickbooks.NewService(
		quickbooks.Sandbox,
		client,
		realmId,
	)

	// _, err := quickbooks.NewQuery[account.Account](qbClient).All()
	// if err != nil {
	// 	panic(err)
	// }

	invoices, err := quickbooks.NewQuery[invoice.Invoice](qbClient).All()
	if err != nil {
		panic(err)
	}

	fmt.Println(invoices)
}
