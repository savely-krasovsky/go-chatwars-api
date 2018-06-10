package cwapi

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strings"
)

func (res *Response) UnmarshalJSON(b []byte) error {
	type alias Response
	temp := struct {
		Action  string          `json:"action"`
		Payload json.RawMessage `json:"payload"`
		*alias
	}{
		alias: (*alias)(res),
	}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	var payload resPayload
	if err := json.Unmarshal(temp.Payload, &payload); err != nil {
		return err
	}
	res.Payload.RequiredOperation = payload.RequiredOperation
	res.Action = temp.Action

	switch temp.Action {
	case "createAuthCode":
		var payload ResCreateAuthCode
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResCreateAuthCode = &payload
	case "grantToken":
		var payload ResGrantToken
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResGrantToken = &payload
	case "authAdditionalOperation":
		var payload ResAuthAdditionalOperation
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResAuthAdditionalOperation = &payload
	case "grantAdditionalOperation":
		var payload ResGrantAdditionalOperation
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResGrantAdditionalOperation = &payload
	case "authorizePayment":
		var payload ResAuthorizePayment
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResAuthorizePayment = &payload
	case "pay":
		var payload ResPay
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResPay = &payload
	case "payout":
		var payload ResPayout
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResPayout = &payload
	case "getInfo":
		var payload ResGetInfo
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResGetInfo = &payload
	case "requestProfile":
		var payload ResRequestProfile
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResRequestProfile = &payload
	case "requestStock":
		var payload ResRequestStock
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResRequestStock = &payload
	case "wantToBuy":
		var payload ResWantToBuy
		if err := json.Unmarshal(temp.Payload, &payload); err != nil {
			return err
		}
		res.Payload.ResWantToBuy = &payload
	default:
		res.Action = "unknownMethod"
	}

	return nil
}

func (payload *reqPayload) MarshalJSON() ([]byte, error) {
	if payload.reqCreateAuthCode != nil {
		return json.Marshal(payload.reqCreateAuthCode)
	}
	if payload.reqGrantToken != nil {
		return json.Marshal(payload.reqGrantToken)
	}
	if payload.reqAuthAdditionalOperation != nil {
		return json.Marshal(payload.reqAuthAdditionalOperation)
	}
	if payload.reqGrantAdditionalOperation != nil {
		return json.Marshal(payload.reqGrantAdditionalOperation)
	}
	if payload.reqAuthorizePayment != nil {
		return json.Marshal(payload.reqAuthorizePayment)
	}
	if payload.reqPay != nil {
		return json.Marshal(payload.reqPay)
	}
	if payload.reqPayout != nil {
		return json.Marshal(payload.reqPayout)
	}
	if payload.reqWantToBuy != nil {
		return json.Marshal(payload.reqWantToBuy)
	}

	return json.Marshal(nil)
}

// Create new client, you can set server optional param, defaults to Chat Wars 2 server (or EU), accepts those variants:
// cw2, eu, cw3, ru
func NewClient(user string, password string, server ...string) (*Client, error) {
	rabbitUrl := fmt.Sprintf(CW2, user, password)

	if len(server) > 0 {
		if strings.ToLower(server[0]) == "cw2" || strings.ToLower(server[0]) == "eu" {
			rabbitUrl = fmt.Sprintf(CW2, user, password)
		} else if strings.ToLower(server[0]) == "cw3" || strings.ToLower(server[0]) == "ru" {
			rabbitUrl = fmt.Sprintf(CW3, user, password)
		}
	}

	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		return nil, err
	}

	chForUpdates, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	chForPublish, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	client := Client{
		User:              user,
		Password:          password,
		connection:        conn,
		channelForUpdates: chForUpdates,
		channelForPublish: chForPublish,
		RabbitUrl:         rabbitUrl,
	}

	client.Updates = make(chan Response, 100)
	err = client.startUpdateConsumer()
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *Client) reStartConsumers() error {
	var err error

	if c.Updates != nil {
		err = c.startUpdateConsumer()
		if err != nil {
			return err
		}
	}

	if c.Deals != nil {
		err = c.startDealsConsumer()
		if err != nil {
			return err
		}

	}

	if c.Offers != nil {
		err = c.startOffersConsumer()
		if err != nil {
			return err
		}

	}

	if c.SexDigest != nil {
		err = c.startSexDigestConsumer()
		if err != nil {
			return err
		}
	}

	if c.YellowPages != nil {
		err = c.startYellowPages()
		if err != nil {
			return err
		}
	}
	return nil
}

// Start consumer for base events
func (c *Client) startUpdateConsumer() error {
	updates, err := c.channelForUpdates.Consume(
		fmt.Sprintf("%s_i", c.User),
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for update := range updates {
			if update.RoutingKey == fmt.Sprintf("%s_i", c.User) {
				var res Response
				err := json.Unmarshal(update.Body, &res)
				if err != nil {
					log.Println(err)
				}

				var userID int

				switch res.Action {
				case "createAuthCode":
					userID = res.Payload.ResCreateAuthCode.UserID
				case "grantToken":
					userID = res.Payload.ResGrantToken.UserID
				case "authAdditionalOperation":
					userID = res.Payload.ResAuthAdditionalOperation.UserID
				case "grantAdditionalOperation":
					userID = res.Payload.ResGrantAdditionalOperation.UserID
				case "authorizePayment":
					userID = res.Payload.ResAuthorizePayment.UserID
				case "pay":
					userID = res.Payload.ResPay.UserID
				case "payout":
					userID = res.Payload.ResPayout.UserID
				case "requestProfile":
					userID = res.Payload.ResRequestProfile.UserID
				case "requestStock":
					userID = res.Payload.ResRequestStock.UserID
				case "wantToBuy":
					userID = res.Payload.ResWantToBuy.UserID
				}

				// trying to load update with this salt
				if waiter, found := c.waiters.Load(userID); found {
					// found? send it to waiter channel
					waiter.(chan Response) <- res

					// trying to prevent memory leak
					close(waiter.(chan Response))
				}

				c.Updates <- res
			}
		}
	}()
	return nil
}

// Close connection and active channel
func (c *Client) CloseConnection() error {
	close(c.Updates)
	close(c.Deals)
	close(c.Offers)
	close(c.SexDigest)
	close(c.YellowPages)

	if err := c.channelForUpdates.Close(); err != nil {
		return err
	}
	if err := c.channelForPublish.Close(); err != nil {
		return err
	}
	if err := c.connection.Close(); err != nil {
		return err
	}

	return nil
}

func (c *Client) makeRequest(req []byte) (err error) {
	err = c.channelForPublish.Publish(
		fmt.Sprintf("%s_ex", c.User),
		fmt.Sprintf("%s_o", c.User),
		true,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        req,
		},
	)
	// If channel closed
	if err != nil && err.(*amqp.Error).Code == 504 {
		// Open new
		conn, err := amqp.Dial(c.RabbitUrl)
		if err != nil {
			return err
		}

		chForPublish, err := conn.Channel()
		if err != nil {
			return err
		}

		chForUpdates, err := conn.Channel()
		if err != nil {
			return err
		}

		// force close old channels and connection to close old consumers
		c.channelForUpdates.Close()
		c.channelForPublish.Close()
		c.connection.Close()
		// Reassign it and restart consumers
		c.connection = conn
		c.channelForPublish = chForPublish
		c.channelForUpdates = chForUpdates

		err = c.reStartConsumers()
		if err != nil {
			return err
		}

		// And try again
		if err := c.makeRequest(req); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
