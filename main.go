package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Person -
type Person struct {
	Name string
	Age  int
}

type Steemit struct {
	BaseUrl string
}

type RpcRequest struct {
	Id      int           `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type GetAccountResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  []struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Owner struct {
			WeightThreshold int             `json:"weight_threshold"`
			AccountAuths    []interface{}   `json:"account_auths"`
			KeyAuths        [][]interface{} `json:"key_auths"`
		} `json:"owner"`
		Active struct {
			WeightThreshold int             `json:"weight_threshold"`
			AccountAuths    []interface{}   `json:"account_auths"`
			KeyAuths        [][]interface{} `json:"key_auths"`
		} `json:"active"`
		Posting struct {
			WeightThreshold int             `json:"weight_threshold"`
			AccountAuths    []interface{}   `json:"account_auths"`
			KeyAuths        [][]interface{} `json:"key_auths"`
		} `json:"posting"`
		MemoKey             string `json:"memo_key"`
		JsonMetadata        string `json:"json_metadata"`
		PostingJsonMetadata string `json:"posting_json_metadata"`
		Proxy               string `json:"proxy"`
		LastOwnerUpdate     string `json:"last_owner_update"`
		LastAccountUpdate   string `json:"last_account_update"`
		Created             string `json:"created"`
		Mined               bool   `json:"mined"`
		RecoveryAccount     string `json:"recovery_account"`
		LastAccountRecovery string `json:"last_account_recovery"`
		ResetAccount        string `json:"reset_account"`
		CommentCount        int    `json:"comment_count"`
		LifetimeVoteCount   int    `json:"lifetime_vote_count"`
		PostCount           int    `json:"post_count"`
		CanVote             bool   `json:"can_vote"`
		VotingManabar       struct {
			CurrentMana    string `json:"current_mana"`
			LastUpdateTime int    `json:"last_update_time"`
		} `json:"voting_manabar"`
		DownvoteManabar struct {
			CurrentMana    int64 `json:"current_mana"`
			LastUpdateTime int   `json:"last_update_time"`
		} `json:"downvote_manabar"`
		VotingPower                   int           `json:"voting_power"`
		Balance                       string        `json:"balance"`
		SavingsBalance                string        `json:"savings_balance"`
		SbdBalance                    string        `json:"sbd_balance"`
		SbdSeconds                    string        `json:"sbd_seconds"`
		SbdSecondsLastUpdate          string        `json:"sbd_seconds_last_update"`
		SbdLastInterestPayment        string        `json:"sbd_last_interest_payment"`
		SavingsSbdBalance             string        `json:"savings_sbd_balance"`
		SavingsSbdSeconds             string        `json:"savings_sbd_seconds"`
		SavingsSbdSecondsLastUpdate   string        `json:"savings_sbd_seconds_last_update"`
		SavingsSbdLastInterestPayment string        `json:"savings_sbd_last_interest_payment"`
		SavingsWithdrawRequests       int           `json:"savings_withdraw_requests"`
		RewardSbdBalance              string        `json:"reward_sbd_balance"`
		RewardSteemBalance            string        `json:"reward_steem_balance"`
		RewardVestingBalance          string        `json:"reward_vesting_balance"`
		RewardVestingSteem            string        `json:"reward_vesting_steem"`
		VestingShares                 string        `json:"vesting_shares"`
		DelegatedVestingShares        string        `json:"delegated_vesting_shares"`
		ReceivedVestingShares         string        `json:"received_vesting_shares"`
		VestingWithdrawRate           string        `json:"vesting_withdraw_rate"`
		NextVestingWithdrawal         string        `json:"next_vesting_withdrawal"`
		Withdrawn                     int           `json:"withdrawn"`
		ToWithdraw                    int           `json:"to_withdraw"`
		WithdrawRoutes                int           `json:"withdraw_routes"`
		CurationRewards               int           `json:"curation_rewards"`
		PostingRewards                int           `json:"posting_rewards"`
		ProxiedVsfVotes               []int         `json:"proxied_vsf_votes"`
		WitnessesVotedFor             int           `json:"witnesses_voted_for"`
		LastPost                      string        `json:"last_post"`
		LastRootPost                  string        `json:"last_root_post"`
		LastVoteTime                  string        `json:"last_vote_time"`
		PostBandwidth                 int           `json:"post_bandwidth"`
		PendingClaimedAccounts        int           `json:"pending_claimed_accounts"`
		VestingBalance                string        `json:"vesting_balance"`
		Reputation                    string        `json:"reputation"`
		TransferHistory               []interface{} `json:"transfer_history"`
		MarketHistory                 []interface{} `json:"market_history"`
		PostHistory                   []interface{} `json:"post_history"`
		VoteHistory                   []interface{} `json:"vote_history"`
		OtherHistory                  []interface{} `json:"other_history"`
		WitnessVotes                  []interface{} `json:"witness_votes"`
		TagsUsage                     []interface{} `json:"tags_usage"`
		GuestBloggers                 []interface{} `json:"guest_bloggers"`
	} `json:"result"`
	Id int `json:"id"`
}

func NewSteemit(baseUrl string) *Steemit {
	return &Steemit{baseUrl}
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
	getAccountResponse := new(GetAccountResponse)

	respBody, err := io.ReadAll(resp.Body)
	if err == nil {
		err := json.Unmarshal(respBody, getAccountResponse)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
	}

	return getAccountResponse, nil
}

type GlobalPropertiesResponse struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		HeadBlockNumber                 int    `json:"head_block_number"`
		HeadBlockId                     string `json:"head_block_id"`
		Time                            string `json:"time"`
		CurrentWitness                  string `json:"current_witness"`
		TotalPow                        int    `json:"total_pow"`
		NumPowWitnesses                 int    `json:"num_pow_witnesses"`
		VirtualSupply                   string `json:"virtual_supply"`
		CurrentSupply                   string `json:"current_supply"`
		ConfidentialSupply              string `json:"confidential_supply"`
		InitSbdSupply                   string `json:"init_sbd_supply"`
		CurrentSbdSupply                string `json:"current_sbd_supply"`
		ConfidentialSbdSupply           string `json:"confidential_sbd_supply"`
		TotalVestingFundSteem           string `json:"total_vesting_fund_steem"`
		TotalVestingShares              string `json:"total_vesting_shares"`
		TotalRewardFundSteem            string `json:"total_reward_fund_steem"`
		TotalRewardShares2              string `json:"total_reward_shares2"`
		PendingRewardedVestingShares    string `json:"pending_rewarded_vesting_shares"`
		PendingRewardedVestingSteem     string `json:"pending_rewarded_vesting_steem"`
		SbdInterestRate                 int    `json:"sbd_interest_rate"`
		SbdPrintRate                    int    `json:"sbd_print_rate"`
		MaximumBlockSize                int    `json:"maximum_block_size"`
		RequiredActionsPartitionPercent int    `json:"required_actions_partition_percent"`
		CurrentAslot                    int    `json:"current_aslot"`
		RecentSlotsFilled               string `json:"recent_slots_filled"`
		ParticipationCount              int    `json:"participation_count"`
		LastIrreversibleBlockNum        int    `json:"last_irreversible_block_num"`
		VotePowerReserveRate            int    `json:"vote_power_reserve_rate"`
		DelegationReturnPeriod          int    `json:"delegation_return_period"`
		ReverseAuctionSeconds           int    `json:"reverse_auction_seconds"`
		AvailableAccountSubsidies       int    `json:"available_account_subsidies"`
		SbdStopPercent                  int    `json:"sbd_stop_percent"`
		SbdStartPercent                 int    `json:"sbd_start_percent"`
		NextMaintenanceTime             string `json:"next_maintenance_time"`
		LastBudgetTime                  string `json:"last_budget_time"`
		ContentRewardPercent            int    `json:"content_reward_percent"`
		VestingRewardPercent            int    `json:"vesting_reward_percent"`
		SpsFundPercent                  int    `json:"sps_fund_percent"`
		SpsIntervalLedger               string `json:"sps_interval_ledger"`
		DownvotePoolPercent             int    `json:"downvote_pool_percent"`
	} `json:"result"`
}

// curl -s --data '{"jsonrpc":"2.0", "method":"condenser_api.get_dynamic_global_properties", "params":[], "id":1}' https://api.steemit.com
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
