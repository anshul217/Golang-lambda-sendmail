package main

import (
	"fmt"
	"net/smtp"
	"context"
    "github.com/aws/aws-lambda-go/lambda"
)

type MailEvent struct {
    Email string `json:"email"`
    Body string `json:"body"`
}

func HandleRequest(ctx context.Context, mail_event MailEvent) (string, error) {

	from := ""
	
	to :=  mail_event.Email
	
	pass := ""
	
	message := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		string(mail_event.Body)
	
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
	
		from, []string{to}, []byte(message))
	
	if err != nil {
		fmt.Sprintf("smtp error :%s", err)
	
		return fmt.Sprintf("smtp error :%s", err), nil
	}
	
	fmt.Sprintf("mail sent.")
	
	return fmt.Sprintf("We can reach you back shortly"), nil
}

func main(){
	
	 lambda.Start(HandleRequest)
}

