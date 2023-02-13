package main

import (
	"fmt"
	"github/mrqwer/slangTelBot/bot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func main() {
	// Basic configurations
	_, exists := os.LookupEnv("TELEGRAM_API_TOKEN")

	if !exists {
		fmt.Println("No token found")
		return
	}

	//	apiKey := os.Getenv("APIKEYOPENAI")
	//
	//	client := openaigo.NewClient(apiKey)
	//
	//	request := openaigo.CompletionRequestBody{
	//		Model:     "text-davinci-003",
	//		Prompt:    []string{"what is your name"},
	//		MaxTokens: 128,
	//	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//	response, err := client.Completion(nil, request)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//txt := ""
	//for _, v := range response.Choices {
	//	txt += v.Text + " "
	//}
	//fmt.Println(txt)
	//	fmt.Println(response.Choices[0].Text)

	// end
	bot.Bot()

}
