# Miele Checker
Simple script. GETs a specific dishwasher model, searches for specific words, sends specific email when a price drop is detected.

Add it to a crontab for best results!

# Spinup

```
glide install
touch keys.go
```

Add relevant global variables to keys.go. Here is a template:

```go
package main

var emailSubject = "Miele dishwasher price drop detected (g-4920-sci-clst-cleansteel-integrated-60cm-wide)"
var emailSender = "user@example.com"
var emailReceiver = "user@example.com"
var mailgunDomain = "mail.example.com"
var mailgunAPIKey = "key-EXAMPLEAPIKEY"
var mailgunPublicAPIKey = "pubkey-EXAMPLEAPIKEY"
```

Run the app!

```
go run *.go
```