package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ancalabrese/beaver-client-libraries/go/client"
	"github.com/ancalabrese/beaver-client-libraries/go/client/content"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

}

const (
	scheme = `{
"type": "object",
"properties": {
"expenseReport":{
"type":"object",
"properties":{
"isExpensable":{
"type":"boolean",
"description": "Set to True if expense is less than $150."
},
"reason":{
"type":"string",
"description": "provide the reason why 'isExpansable' is true or false"
}
}
},
"metadata":{
"type": "object",
"properties":{
"merchant": {
"type": "string",
"description": "Merchant name in the receipt."
},
"date": {
"type": "string",
"format": "date",
"description": "Date of the transaction."
},
"time": {
"type": "string",
"description": "Time of the transaction"
},
"sale": {
"type": "object",
"properties": {
"amount": {
"type": "number",
"description": "The total amount paid by the customer"
}
}
}
}
}
}
}`
)

func main() {
	apikey := os.Getenv("BEAVER_API_KEY")
	if apikey == "" {
		log.Fatal("invalid API key")
	}

	fmt.Println(apikey)
	apiKeyAuth := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("X-API-Key", apikey)
	})

	params := content.NewGenerateContentParams()
	params.SetSchema(scheme)

	f, err := os.Open("cv.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	params.SetAttachment(runtime.NamedReader("attachment", f))

	resp, err := client.Default.Content.GenerateContent(params, apiKeyAuth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("resp: %v\n", resp)

}
