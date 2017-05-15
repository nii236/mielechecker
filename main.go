package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	mailgun "github.com/mailgun/mailgun-go"
)

func main() {
	dishwasherURL := "http://www.mieleshop.com.au/g-4920-sci-clst-cleansteel-integrated-60cm-wide/p_10074510_Dishwasherappliances_Integrateddishwashers/miele-au"
	doc, err := goquery.NewDocument(dishwasherURL)
	if err != nil {
		log.Fatal(err)
	}
	priceText := doc.Find("#productSelection").Find(".priceProduct").First()
	if strings.Contains(priceText.Text(), "Now") || strings.Contains(priceText.Text(), "Was") {
		log.Println("Price drop detected.")
		log.Println(priceText.Text())
		sendEmail(priceText.Text())
		return
	}
	log.Println("Price drop not detected.")
}

func sendEmail(emailText string) {
	mg := mailgun.NewMailgun(mailgunDomain, mailgunAPIKey, mailgunPublicAPIKey)
	message := mailgun.NewMessage(
		emailSender,
		emailSubject,
		emailText,
		emailReceiver)
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID: %s Resp: %s\n", id, resp)
}
