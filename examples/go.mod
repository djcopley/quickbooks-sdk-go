module github.com/djcopley/quickbooks-sdk-go/examples

go 1.23.4

replace github.com/djcopley/quickbooks-sdk-go => ../

require (
	github.com/djcopley/quickbooks-sdk-go v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	golang.org/x/oauth2 v0.26.0
)

require github.com/Jeffail/gabs/v2 v2.7.0 // indirect
