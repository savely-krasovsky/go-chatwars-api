// Copyright (c) 2018 Savely Krasovsky

// Golang API Wrapper for Chat Wars Telegram MMORPG game
//
// Examples
//
// Async approach:
//
//	package main
//
//	import (
//		"github.com/L11R/go-chatwars-api"
//		"encoding/json"
//		"log"
//	)
//
//	func main() {
//		client, err := cwapi.NewClient("login", "password")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		// Async method, it just sends request
//		err = client.CreateAuthCode(YourUserID)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		// Here you will get response
//		for u := range client.Updates {
//			if u.GetActionsEnum() == cwapi.CreateAuthCode {
//				log.Println("Got response!")
//
//				b, err := json.MarshalIndent(u, "", "\t")
//				if err != nil {
//					log.Fatal(err)
//				}
//
//				log.Println(string(b))
//			}
//		}
//	}
//
// Sync approach:
//
//	package main
//
//	import (
//		"github.com/L11R/go-chatwars-api"
//		"log"
//		"bufio"
//		"os"
//		"fmt"
//		"encoding/json"
//	)
//
//	func main() {
//		client, err := cwapi.NewClient("login", "password")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		userID := 123456
//
//		res, err := client.CreateAuthCodeSync(userID)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		reader := bufio.NewReader(os.Stdin)
//		fmt.Print("Enter text: ")
//		b, _, err := reader.ReadLine()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		authCode := string(b)
//
//		res, err = client.GrantTokenSync(userID, authCode)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		res, err = client.RequestProfileSync(res.Payload.Token, userID)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		b, err = json.MarshalIndent(res.Payload, "", "\t")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		log.Println(string(b))
//	}
//
// Fan-out Exchange Routing Keys
//
// If you want to deal with routing keys, there are a bunch of methods:
//
//	InitDeals()
//	InitOffers()
//	InitSexDigest()
//	InitYellowPages()
//
// After initializing you need just to handle updates from those routes:
//
//	err := client.InitYellowPages()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for page := range client.YellowPages {
//		log.Println("Got page from Yellow Pages!")
//
//		b, err := json.MarshalIndent(page, "", "\t")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		log.Println(string(b))
//	}
//
// Feedback
//
// If you have any questions, you can ask them in Chat Wars Development chat:
// https://t.me/cwapi
package cwapi
