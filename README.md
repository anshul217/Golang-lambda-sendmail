# Golang lambda to send mail


[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

send_mail is a golang lamda function to send mail.
Just add your sender email and password and compile it.

### Installation

It requires Golang to compile

Here are the steps to compile and make zip for aws lambda
```sh
$ GOOS=linux GOARCH=amd64 go build -o send_mail send_mail.go
$ zip send_mail.zip send_mail
```
Above steps will create a zip that can be uploaded to aws lambda



### Test
Make lambda test like this to test
{
    "email":"youremail",
    "body": "test mail"
}
 - This will send mail to you with given body.

License
----

MIT

