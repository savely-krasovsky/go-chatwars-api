package cwapi

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

// Initializes deals public exchange.
func (c *Client) InitDeals() error {
	c.Deals = make(chan Deal, 1)
	topic := fmt.Sprintf("%sdeals", c.PublicPrefix)
	pc, err := newPublicConsumer(topic)
	if err != nil {
		return err
	}
	go startConsuming(pc, c.dealHandler)
	return nil
}

func (c *Client) dealHandler(val []byte) error {
	var res Deal
	err := json.Unmarshal(val, &res)
	if err != nil {
		return fmt.Errorf("unmarshal %s error: %v", val, err)
	}
	c.Deals <- res
	return nil
}

// Initializes offers public exchange.
func (c *Client) InitDuels() error {
	c.Duels = make(chan Duel, 1)
	topic := fmt.Sprintf("%sduels", c.PublicPrefix)
	pc, err := newPublicConsumer(topic)
	if err != nil {
		return err
	}
	go startConsuming(pc, c.duelHandler)
	return nil
}

func (c *Client) duelHandler(val []byte) error {
	var res Duel
	err := json.Unmarshal(val, &res)
	if err != nil {
		return fmt.Errorf("unmarshal %s error: %v", val, err)
	}
	c.Duels <- res
	return nil
}

// Initializes offers public exchange.
func (c *Client) InitOffers() error {
	c.Offers = make(chan Offer, 1)
	topic := fmt.Sprintf("%soffers", c.PublicPrefix)
	pc, err := newPublicConsumer(topic)
	if err != nil {
		return err
	}
	go startConsuming(pc, c.offerHandler)
	return nil
}

func (c *Client) offerHandler(val []byte) error {
	var res Offer
	err := json.Unmarshal(val, &res)
	if err != nil {
		return fmt.Errorf("unmarshal %s error: %v", val, err)
	}
	c.Offers <- res
	return nil
}

// Initializes sex_digest public exchange.
func (c *Client) InitSexDigest() error {
	c.SexDigest = make(chan []SexDigestItem, 1)
	topic := fmt.Sprintf("%ssex_digest", c.PublicPrefix)
	pc, err := newPublicConsumer(topic)
	if err != nil {
		return err
	}
	go startConsuming(pc, c.sexDigestHandler)
	return nil
}

func (c *Client) sexDigestHandler(val []byte) error {
	var res []SexDigestItem
	err := json.Unmarshal(val, &res)
	if err != nil {
		return fmt.Errorf("unmarshal %s error: %v", val, err)
	}
	c.SexDigest <- res
	return nil
}

// Initializes yellow_pages public exchange.
func (c *Client) InitYellowPages() error {
	c.YellowPages = make(chan []YellowPage, 1)
	topic := fmt.Sprintf("%syellow_pages", c.PublicPrefix)
	pc, err := newPublicConsumer(topic)
	if err != nil {
		return err
	}
	go startConsuming(pc, c.yellowPagesHandler)
	return nil
}

func (c *Client) yellowPagesHandler(val []byte) error {
	var res []YellowPage
	err := json.Unmarshal(val, &res)
	if err != nil {
		return fmt.Errorf("unmarshal %s error: %v", val, err)
	}
	c.YellowPages <- res
	return nil
}

// Initializes au_digest public exchange.
func (c *Client) InitAuctionDigest() error {
	c.AuctionDigest = make(chan []AuctionDigestItem, 1)
	topic := fmt.Sprintf("%sau_digest", c.PublicPrefix)
	pc, err := newPublicConsumer(topic)
	if err != nil {
		return err
	}
	go startConsuming(pc, c.auctionDigestHandler)
	return nil
}

func (c *Client) auctionDigestHandler(val []byte) error {
	var res []AuctionDigestItem
	err := json.Unmarshal(val, &res)
	if err != nil {
		return fmt.Errorf("unmarshal %s error: %v", val, err)
	}
	c.AuctionDigest <- res
	return nil
}

func newPublicConsumer(topic string) (sarama.PartitionConsumer, error) {
	cfg := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{kafkaServer}, cfg)
	// надо закрывать
	if err != nil {
		return nil, err
	}
	pc, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return nil, err
	}
	return pc, nil
}

func startConsuming(updates sarama.PartitionConsumer, handler func([]byte) error) {
	defer func() {
		if err := updates.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for msg := range updates.Messages() {
		if err := handler(msg.Value); err != nil {
			log.Printf("Handler error: %v", err)
		}
	}
}
