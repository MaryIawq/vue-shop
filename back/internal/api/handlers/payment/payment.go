package paymentAPI

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestParams struct {
	Price          int64  `json:"price"`
	IdempotenceKey string `json:"idempotence_key"`
}

func GetToken(ctx *gin.Context) {
	var params requestParams

	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Printf("proccessOrder #1\nError: %s\n", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &params)
	if err != nil {
		fmt.Printf("proccessOrder #2\nError: %s\n", err.Error())
		return
	}

	shopID := 229691
	secretKey := "test_fuSfgXoUycU9Prz4129EA2C0LsHJeSqyUe5XMayNim8"
	idempotenceKey := params.IdempotenceKey

	url := "https://api.yookassa.ru/v3/payments"

	payload := []byte(`{
		"amount": {
			"value": ` + fmt.Sprint(params.Price) + `,
			"currency": "RUB"
		},
		"confirmation": {
			"type": "embedded"
		},
		"capture": true,
		"description": "Заказ из Vue Candies Shop"
	}`)

	// Create an HTTP client
	client := &http.Client{}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotence-Key", idempotenceKey)

	// Encode credentials for basic authentication
	auth := fmt.Sprint(shopID) + ":" + secretKey
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("proccessOrder #1\nError: %s\n", err.Error())
		return
	}

	type yokassaResponse struct {
		Confirmation struct {
			Token string `json:"confirmation_token"`
		} `json:"confirmation"`
	}

	var response yokassaResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("proccessOrder #2\nError: %s\n", err.Error())
		return
	}

	log.Printf("\n%s\n", response)
	ctx.JSON(200, response.Confirmation.Token)
}
