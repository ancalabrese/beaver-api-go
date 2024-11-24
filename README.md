# Beaver API client for Go

Unlock the power of your data. Effortlessly extract actionable insights with AI from any source.

Integrate seamlessly with SDKs for all popular programming languages and get up and running in minutes





## Usage/Examples
Extracts key information from CVs (PDFs), including personal details, experience, education, and skills.   
Infers implicit data and supports complex queries for insights, returning results in JSON format.  Useful for candidate screening, skill gap analysis, and data-driven recruitment.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ancalabrese/beaver-api-go/client"
	"github.com/ancalabrese/beaver-api-go/client/content"
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
    "name": {
      "type": "string",
      "description": "The name of the candidate in the CV"
    },
    "summary": {
      "type": "string",
      "description": "A short summary of the candidate skills and experience"
    },
    "isMatch": {
      "type": "boolean",
      "description": "Set to true is candidate has experience with AI"
    },
    "reason": {
      "type": "string",
      "description": "Give a reason why isMatch was set either to true or false"
    }
  }
}`
)

type Result struct {
	Properties struct {
		Name    string `json:"name"`
		Summary string `json:"summary"`
		IsMatch bool   `json:"isMatch"`
		Reason  string `json:"reason"`
	} `json:"properties"`
}

func main() {
	apikey := os.Getenv("BEAVER_API_KEY")
	if apikey == "" {
		log.Fatal("invalid API key")
	}

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
		log.Fatal("Generate content failed: ", err)
	}

	var res Result
	err = json.Unmarshal([]byte(resp.GetPayload()), &res)
	if err != nil {
		log.Fatal("Resp unmarshalling failed: ", err)
	}

	fmt.Printf("resp: %+v\n", res)
}

```


## Authors

For support please reach out to:

- [@ancalabrese](https://antoniocalabrese.dev)


## FAQ

#### What is Beaver?

Beaver is a powerful SDK that allows developers to effortlessly extract structured data from any unstructured source. Simply provide the source file and a JSON schema, and our AI-powered engine will return the parsed data in your desired format.

#### How do I get an API Key

Sign up and get yourself a key at [Beaver]()


#### How does pricing work?
We offer flexible pricing plans based on your usage needs. Choose a plan that aligns with your data volume and request frequency. Visit [website]() for details.
