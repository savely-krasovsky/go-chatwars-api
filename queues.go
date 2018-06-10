package cwapi

import (
	"encoding/json"
	"fmt"
	"log"
)

// Initializes deals public exchange.
func (c *Client) InitDeals() error {
	c.Deals = make(chan Deal, 100)
	err := c.startDealsConsumer()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) startDealsConsumer() error {
	updates, err := c.channelForUpdates.Consume(
		fmt.Sprintf("%s_deals", c.User),
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
			var res Deal
			err := json.Unmarshal(update.Body, &res)
			if err != nil {
				log.Println(err)
			}

			c.Deals <- res
		}
	}()
	return nil
}

// // Initializes offers public exchange.
func (c *Client) InitOffers() error {
	c.Offers = make(chan Offer, 100)
	err := c.startOffersConsumer()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) startOffersConsumer() error {
	updates, err := c.channelForUpdates.Consume(
		fmt.Sprintf("%s_offers", c.User),
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
			var res Offer
			err := json.Unmarshal(update.Body, &res)
			if err != nil {
				log.Println(err)
			}

			c.Offers <- res
		}
	}()
	return nil
}

// Initializes sex_digest public exchange.
func (c *Client) InitSexDigest() error {
	c.SexDigest = make(chan []SexDigestItem, 1)
	err := c.startSexDigestConsumer()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) startSexDigestConsumer() error {
	updates, err := c.channelForUpdates.Consume(
		fmt.Sprintf("%s_sex_digest", c.User),
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
			var res []SexDigestItem
			err := json.Unmarshal(update.Body, &res)
			if err != nil {
				log.Println(err)
			}

			c.SexDigest <- res
		}
	}()
	return nil
}

// Initializes yellow_pages public exchange.
func (c *Client) InitYellowPages() error {
	c.YellowPages = make(chan []YellowPage, 1)
	err := c.startYellowPages()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) startYellowPages() error {
	updates, err := c.channelForUpdates.Consume(
		fmt.Sprintf("%s_yellow_pages", c.User),
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
			var res []YellowPage
			err := json.Unmarshal(update.Body, &res)
			if err != nil {
				log.Println(err)
			}

			c.YellowPages <- res
		}
	}()
	return nil
}
