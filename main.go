package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	openai "github.com/sashabaranov/go-openai"
)

// Read the contents of a file and store as a string
func readFileToString(filePath string) (string, error) {
	log.Printf("Pulling content from %v to use as blog topic.\n", filePath)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	fileContentString := string(fileContent)

	return fileContentString, nil
}

// Write the results from ChatGPT to a file
func writeStringToFile(response string, filePath string) error {
	fileName := filepath.Base(filePath)
	outputFile := "article-on-" + fileName
	outputFilePath := "articles/" + outputFile
	outputDirPath := "articles"
	switch _, err := os.Stat(outputDirPath); {
	case os.IsNotExist(err):
		err := os.MkdirAll(outputDirPath, 0755)
		if err != nil {
			log.Fatalf("Error creating directory %v\n", outputDirPath)
		} else {
			log.Printf("Directory '%v' created successfully.\n", outputDirPath)
		}
	case err != nil:
		log.Printf("Directory '%v' created successfully.\n", outputDirPath)
	default:
		log.Printf("Directory '%v' already exists.", outputDirPath)
	}
	log.Printf("Writing blog post to %v.\n", outputFilePath)
	err := os.WriteFile(outputFilePath, []byte(response), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Generate a blog post on a topic from a summary of a file
func chatGptQuery(fileContentString string, apikey string) (string, error) {
	log.Println("Sending request to ChatGPT...")
	client := openai.NewClient(apikey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `You are an Cyber Threat Intelligence Analyst.
                    This month is Cybersecurity Awareness month. This month you will be writing blog posts
                    for the entire company on any topic I provide for you. 
                    I will give you a topic, and you will write a blog post 800 words or less in length
                    that is targeted toward a general audience in a technology company.
                    string.
                    You you should use the following tone of voice:
                    1. Positivity – Scare tactics don’t work. Instead of using scary “hackers in hoodies” imagery, talk
                    about the benefits of reducing cyber risks and how strengthening cybersecurity can protect
                    what matters most in our lives.
                    2. Approachability – Cybersecurity seems like a complex subject to many, but it’s really all about
                    people. Make cybersecurity relatable and share how practicing good cyber hygiene is
                    something that anyone can do.
                    3. Simplicity – Avoid jargon and be sure to define acronyms
                    4. Back to basics – Even just practicing cyber hygiene basics provides a solid baseline of security. 
                    `,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fileContentString,
				},
			},
		},
	)

	if err != nil {
		//log.Fatalf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, err
}

func main() {
	filePath := os.Args[1]
	apikey := os.Getenv("OPENAI_API_KEY")
	if apikey == "" {
		log.Fatalf("apikey is not present")
	}

	fileContentString, err := readFileToString(filePath)
	if err != nil {
		log.Fatalf("File read error: %v\n", err)
	}

	response, err := chatGptQuery(fileContentString, apikey)
	if err != nil {
		log.Fatalf("ChatGPT query error: %v\n", err)
	}

	err = writeStringToFile(response, filePath)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Println(response)
}
