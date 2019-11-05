package lcd

import (
	"encoding/json"
	"fmt"

	"github.com/NetCloth/go-sdk/client/types"
)

type (
	AccountValue struct {
		Address       string       `json:"address"`
		Coins         []types.Coin `json:"coins"`
		AccountNumber string       `json:"account_number"`
		Sequence      string       `json:"sequence"`
	}
	AccountResult struct {
		Type  string       `json:"type"`
		Value AccountValue `json:"value"`
	}

	AccountBody struct {
		Height string        `json:"height"`
		Result AccountResult `json:"result"`
	}
	AccountInfo struct {
		Type  string           `json:"type"`
		Value AccountInfoValue `json:"value"`
	}

	IPALObj struct {
		Address string `json:"user_address"`
		IP      string `json:"ip"`
	}

	AccountInfoValue struct {
		AccountNumber string       `json:"account_number"`
		Address       string       `json:"address"`
		Sequence      string       `json:"sequence"`
		Coins         []types.Coin `json:"coins"`
	}
)

func (c *client) QueryAccount(address string) (accountInfo AccountBody, err error) {
	path := fmt.Sprintf(UriQueryAccount, address)

	if _, body, err := c.httpClient.Get(path, nil); err != nil {
		return accountInfo, err
	} else {
		if err := json.Unmarshal(body, &accountInfo); err != nil {
			return accountInfo, err
		} else {
			return accountInfo, nil
		}
	}
}

func (c *client) QueryIPAL(address string) (IPALObj, error) {
	var (
		ipalObj IPALObj
	)
	path := fmt.Sprintf(UriQueryIPAL, address)
	fmt.Println(path)

	if _, body, err := c.httpClient.Get(path, nil); err != nil {
		return ipalObj, err
	} else {
		if err := json.Unmarshal(body, &ipalObj); err != nil {
			return ipalObj, err
		} else {
			return ipalObj, nil
		}
	}
}

func (c *client) QueryAIPALList() string {
	return ""
}
