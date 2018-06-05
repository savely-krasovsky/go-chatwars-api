package cwapi

import (
	"encoding/json"
	"errors"
	"time"
)

// Access request from your application to the user.
func (c *Client) CreateAuthCode(userID int) error {
	payload, err := json.Marshal(&reqPayload{
		reqCreateAuthCode: &reqCreateAuthCode{
			userID,
		},
	})
	if err != nil {
		return err
	}

	req := &Request{
		Action:  "createAuthCode",
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Sync-version of CreateAuthCode method.
func (c *Client) CreateAuthCodeSync(userID int) (*Response, error) {
	payload, err := json.Marshal(&reqPayload{
		reqCreateAuthCode: &reqCreateAuthCode{
			userID,
		},
	})
	if err != nil {
		return nil, err
	}

	req := &Request{
		Action:  "createAuthCode",
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	// wait response from main loop in NewClient()
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
		// or timeout
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}

// Exchange auth code for access token.
func (c *Client) GrantToken(userID int, authCode string) error {
	payload, err := json.Marshal(&reqPayload{
		reqGrantToken: &reqGrantToken{
			userID,
			authCode,
		},
	})
	if err != nil {
		return err
	}

	req := &Request{
		Action:  "grantToken",
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Sync-version of GrantToken method.
func (c *Client) GrantTokenSync(userID int, authCode string) (*Response, error) {
	payload, err := json.Marshal(&reqPayload{
		reqGrantToken: &reqGrantToken{
			userID,
			authCode,
		},
	})
	if err != nil {
		return nil, err
	}

	req := &Request{
		Action:  "grantToken",
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}

// Sends request to broaden tokens operations set to user.
func (c *Client) AuthAdditionalOperation(token string, operation string) error {
	payload, err := json.Marshal(&reqPayload{
		reqAuthAdditionalOperation: &reqAuthAdditionalOperation{
			operation,
		},
	})
	if err != nil {
		return err
	}

	req := &Request{
		Action:  "authAdditionalOperation",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Sync-version of AuthAdditionalOperation method.
func (c *Client) AuthAdditionalOperationSync(token string, operation string, userID int) (*Response, error) {
	payload, err := json.Marshal(&reqPayload{
		reqAuthAdditionalOperation: &reqAuthAdditionalOperation{
			operation,
		},
	})
	if err != nil {
		return nil, err
	}

	req := &Request{
		Action:  "authAdditionalOperation",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}

// Completes the authAdditionalOperation action.
func (c *Client) GrantAdditionalOperation(token string, requestedID string, authCode string) error {
	payload, err := json.Marshal(&reqPayload{
		reqGrantAdditionalOperation: &reqGrantAdditionalOperation{
			requestedID,
			authCode,
		},
	})
	if err != nil {
		return err
	}

	req := &Request{
		Action:  "grantAdditionalOperation",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Sync-version of GrantAdditionalOperation method.
func (c *Client) GrantAdditionalOperationSync(token string, requestedID string, authCode string, userID int) (*Response, error) {
	payload, err := json.Marshal(&reqPayload{
		reqGrantAdditionalOperation: &reqGrantAdditionalOperation{
			requestedID,
			authCode,
		},
	})
	if err != nil {
		return nil, err
	}

	req := &Request{
		Action:  "grantAdditionalOperation",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}

// Sends authorization request to user with confirmation code in it.
func (c *Client) AuthorizePayment(token string, transactionID string, pouchesAmount int) error {
	payload, err := json.Marshal(&reqPayload{
		reqAuthorizePayment: &reqAuthorizePayment{
			transactionID,
			map[string]int{
				"pouches": pouchesAmount,
			},
		},
	})
	if err != nil {
		return err
	}

	req := &Request{
		Action:  "authorizePayment",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Sync-version of AuthorizePayment method.
func (c *Client) AuthorizePaymentSync(token string, transactionID string, pouchesAmount int, userID int) (*Response, error) {
	payload, err := json.Marshal(&reqPayload{
		reqAuthorizePayment: &reqAuthorizePayment{
			transactionID,
			map[string]int{
				"pouches": pouchesAmount,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	req := &Request{
		Action:  "authorizePayment",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}

// Previously, transfers held an amount of gold from users account to application’s balance.
func (c *Client) Pay(token string, transactionID string, pouchesAmount int, confirmCode string) error {
	payload, err := json.Marshal(&reqPayload{
		reqPay: &reqPay{
			transactionID,
			map[string]int{
				"pouches": pouchesAmount,
			},
			confirmCode,
		},
	})
	if err != nil {
		return err
	}

	req := &Request{
		Action:  "pay",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Sync-version of Pay method.
func (c *Client) PaySync(token string, transactionID string, pouchesAmount int, confirmCode string, userID int) (*Response, error) {
	payload, err := json.Marshal(&reqPayload{
		reqPay: &reqPay{
			transactionID,
			map[string]int{
				"pouches": pouchesAmount,
			},
			confirmCode,
		},
	})
	if err != nil {
		return nil, err
	}

	req := &Request{
		Action:  "pay",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}

// Transfers of a given amount of gold (or pouches) from the application’s balance to users account.
func (c *Client) Payout(token string, transactionID string, pouchesAmount int, message string) error {
	payload, err := json.Marshal(&reqPayload{
		reqPayout: &reqPayout{
			transactionID,
			map[string]int{
				"pouches": pouchesAmount,
			},
			message,
		},
	})
	if err != nil {
		return err
	}

	req := &Request{
		Action:  "payout",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Sync-version of Payout method.
func (c *Client) PayoutSync(token string, transactionID string, pouchesAmount int, message string, userID int) (*Response, error) {
	payload, err := json.Marshal(&reqPayload{
		reqPayout: &reqPayout{
			transactionID,
			map[string]int{
				"pouches": pouchesAmount,
			},
			message,
		},
	})
	if err != nil {
		return nil, err
	}

	req := &Request{
		Action:  "payout",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}

// Request current info about your application. E.g. balance, limits, status.
func (c *Client) GetInfo() error {
	req := &Request{
		Action: "getInfo",
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Request brief user profile information.
func (c *Client) RequestProfile(token string) error {
	req := &Request{
		Action: "requestProfile",
		Token:  token,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Sync-version of RequstProfile method.
func (c *Client) RequestProfileSync(token string, userID int) (*Response, error) {
	req := &Request{
		Action: "requestProfile",
		Token:  token,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}

// Request users stock information.
func (c *Client) RequestStock(token string) error {
	req := &Request{
		Action: "requestStock",
		Token:  token,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Sync-version of RequestStock method.
func (c *Client) RequestStockSync(token string, userID int) (*Response, error) {
	req := &Request{
		Action: "requestStock",
		Token:  token,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}

// Buys something on exchange.
func (c *Client) WantToBuy(token string, itemCode string, quantity int, price int, exactPrice bool) error {
	payload, err := json.Marshal(&reqPayload{
		reqWantToBuy: &reqWantToBuy{
			ItemCode:   itemCode,
			Quantity:   quantity,
			Price:      price,
			ExactPrice: exactPrice,
		},
	})
	if err != nil {
		return err
	}

	req := &Request{
		Action:  "wantToBuy",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.makeRequest(body)
	if err != nil {
		return err
	}

	return nil
}

// Buys something on exchange.
func (c *Client) WantToBuySync(token string, itemCode string, quantity int, price int, exactPrice bool, userID int) (*Response, error) {
	payload, err := json.Marshal(&reqPayload{
		reqWantToBuy: &reqWantToBuy{
			ItemCode:   itemCode,
			Quantity:   quantity,
			Price:      price,
			ExactPrice: exactPrice,
		},
	})
	if err != nil {
		return nil, err
	}

	req := &Request{
		Action:  "wantToBuy",
		Token:   token,
		Payload: payload,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = c.makeRequest(body)
	if err != nil {
		return nil, err
	}

	waiter := make(chan Response, 1)
	c.waiters.Store(userID, waiter)

	select {
	case response := <-waiter:
		if response.GetResultEnum() != Ok {
			return &response, errors.New(string(response.GetResultEnum()))
		}
		return &response, nil
	case <-time.After(10 * time.Second):
		c.waiters.Delete(userID)
		return nil, errors.New("timeout")
	}
}
