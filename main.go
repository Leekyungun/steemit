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

func (steemit *Steemit) getAccount(accountName []string) (*GetAccountResponse, error) {
	var params []interface{}
	var account [1][]string
	account[0] = accountName

	params = append(params, "database_api")
	params = append(params, "get_accounts")
	params = append(params, account)

	rpcRequest := RpcRequest{1, "2.0", "call", params}
	requestBody, err := json.Marshal(rpcRequest)
	if err != nil {
		fmt.Println("getAccount", "json.Marshal", err)
		return nil, err
	}

	response, err := steemit.BaseRequest(requestBody)
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
	requestBody, err := json.Marshal(rpcRequest)
	if err != nil {
		fmt.Println("GetConfig", "err", err)
		return nil, err
	}

	response, err := steemit.BaseRequest(requestBody)
	if err != nil {
		return nil, err
	}

	globalPropertiesResponse := new(GlobalPropertiesResponse)
	err = json.Unmarshal(response, globalPropertiesResponse)
	if err != nil {
		fmt.Println("GetConfig", "err", err)
		return nil, err
	}

	return globalPropertiesResponse, nil
}

func main() {
	baseUrl := "https://api.steemit.com"
	steemit := NewSteemit(baseUrl)
	account, err := steemit.getAccount([]string{"lku"})
	if err != nil {
		return
	}

	accountAsset := account.Asset()

	fmt.Printf("%+v\n", accountAsset)

	globalProperty, err := steemit.GetConfig()
	if err != nil {
		return
	}

	globalProperties := globalProperty.Convert()
	fmt.Printf("%+v\n", globalProperties)

	// 내 스팀 = 토탈스팀 * 내 주식 / 토탈 주식
	steemPower := globalProperties.TotalVestingFundSteem.Mul(*accountAsset.VestShares).Div(*globalProperties.TotalVestingShares)
	fmt.Println("steemPower", steemPower)
}
