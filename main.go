package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Steemit struct {
	BaseUrl string
}

func NewSteemit(baseUrl string) *Steemit {
	return &Steemit{baseUrl}
}

func (steemit *Steemit) BaseRequest(body []byte) ([]byte, error) {
	title := "Post"

	BodyBuffer := bytes.NewBuffer(body)
	resp, err := http.Post(steemit.BaseUrl, "application/json", BodyBuffer)
	if err != nil {
		fmt.Println(title, "http.Post", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(title, "Body.Close", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println(title, "resp.StatusCode not 200", resp)
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(title, "io.ReadAll", err)
		return nil, err
	}

	return respBody, nil
}

func (steemit *Steemit) getAccount() (*GetAccountResponse, error) {
	var Params []interface{}
	var Account [1][1]string
	Account[0][0] = "lku"

	Params = append(Params, "database_api")
	Params = append(Params, "get_accounts")
	Params = append(Params, Account)

	rpcRequest := RpcRequest{1, "2.0", "call", Params}
	pbytes, err := json.Marshal(rpcRequest)
	if err != nil {
		fmt.Println("getAccount", "json.Marshal", err)
		return nil, err
	}

	response, err := steemit.BaseRequest(pbytes)
	if err != nil {
		return nil, err
	}

	// Response 체크.
	getAccountResponse := new(GetAccountResponse)

	err = json.Unmarshal(response, getAccountResponse)
	if err != nil {
		fmt.Println("getAccount", "json.Unmarshal", err)
		return nil, err
	}

	return getAccountResponse, nil
}

func (steemit *Steemit) GetConfig() (*GlobalPropertiesResponse, error) {
	Params := make([]interface{}, 0)
	rpcRequest := RpcRequest{2, "2.0", "condenser_api.get_dynamic_global_properties", Params}
	pbytes, err := json.Marshal(rpcRequest)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(steemit.BaseUrl, "application/json", buff)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	defer resp.Body.Close()

	// Response 체크.
	globalPropertiesResponse := new(GlobalPropertiesResponse)

	respBody, err := io.ReadAll(resp.Body)
	if err == nil {
		err := json.Unmarshal(respBody, globalPropertiesResponse)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
	}

	return globalPropertiesResponse, nil
}

func main() {
	baseUrl := "https://api.steemit.com"
	steemit := NewSteemit(baseUrl)
	account, err := steemit.getAccount()
	if err != nil {
		return
	}

	fmt.Println("account", account)
	fmt.Println("steem balance", account.Result[0].Balance)
	fmt.Println("SBD balance", account.Result[0].SbdBalance)
	fmt.Println("reward steem balance", account.Result[0].RewardSteemBalance)
	fmt.Println("reward SBD balance", account.Result[0].RewardSbdBalance)
	fmt.Println("reward steem power balance", account.Result[0].RewardVestingSteem)
	fmt.Println("reward steem power balance", account.Result[0].RewardVestingBalance)

	fmt.Println("vest shares", account.Result[0].VestingShares)
	fmt.Println("vest balance", account.Result[0].VestingBalance)
	fmt.Println("deleted vest shares", account.Result[0].DelegatedVestingShares)

	globalProperty, err := steemit.GetConfig()
	if err != nil {
		return
	}

	fmt.Println("============")
	fmt.Println("globalProperty")
	fmt.Println("globalProperty.Result.TotalVestingShares", globalProperty.Result.TotalVestingShares)
	fmt.Println("globalProperty.Result.TotalVestingFundSteem", globalProperty.Result.TotalVestingFundSteem)

	/*
			내 스팀 = 토탈스팀 * 내 주식 / 토탈 주식
			내 스팀 = 167050752.555 * 10675732.181640 / 297947788335.180300
		167037797.286 * 10679990.460432 /297902892385.261137 = 5,988.4013453584
		167037797.286 * 5187.305547 /297902892385.261137 = 5,988.4013453584
	*/

}
