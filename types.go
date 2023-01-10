package main

import (
	"github.com/shopspring/decimal"
	"strings"
)

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
			CurrentMana    string `json:"current_mana"`
			LastUpdateTime int    `json:"last_update_time"`
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

type AccountAsset struct {
	SteemBalance            *decimal.Decimal
	SBDBalance              *decimal.Decimal
	RewardSteemBalance      *decimal.Decimal
	RewardSBDBalance        *decimal.Decimal
	RewardSteemPowerBalance *decimal.Decimal
	RewardVestsShares       *decimal.Decimal
	VestShares              *decimal.Decimal
	DelegationVestShares    *decimal.Decimal
}

func ConvertElem(origin string) (*decimal.Decimal, error) {
	numberString := strings.Split(origin, " ")
	number, err := decimal.NewFromString(numberString[0])
	if err != nil {
		return nil, err
	}

	return &number, nil
}

func (account *GetAccountResponse) Asset() *AccountAsset {
	accountAsset := new(AccountAsset)
	accountAsset.SteemBalance, _ = ConvertElem(account.Result[0].Balance)
	accountAsset.SBDBalance, _ = ConvertElem(account.Result[0].SbdBalance)
	accountAsset.RewardSteemBalance, _ = ConvertElem(account.Result[0].RewardSteemBalance)
	accountAsset.RewardSBDBalance, _ = ConvertElem(account.Result[0].RewardSbdBalance)
	accountAsset.RewardSteemPowerBalance, _ = ConvertElem(account.Result[0].RewardVestingSteem)
	accountAsset.RewardVestsShares, _ = ConvertElem(account.Result[0].RewardVestingBalance)
	accountAsset.VestShares, _ = ConvertElem(account.Result[0].VestingShares)
	accountAsset.DelegationVestShares, _ = ConvertElem(account.Result[0].DelegatedVestingShares)

	return accountAsset
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

func (properties *GlobalPropertiesResponse) Convert() *GlobalProperties {
	globalProperties := new(GlobalProperties)
	globalProperties.TotalVestingShares, _ = ConvertElem(properties.Result.TotalVestingShares)
	globalProperties.TotalVestingFundSteem, _ = ConvertElem(properties.Result.TotalVestingFundSteem)
	return globalProperties
}

type GlobalProperties struct {
	TotalVestingShares    *decimal.Decimal
	TotalVestingFundSteem *decimal.Decimal
}
