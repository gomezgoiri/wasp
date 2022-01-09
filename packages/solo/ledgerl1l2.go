package solo

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/iotaledger/wasp/packages/util"

	"github.com/iotaledger/hive.go/serializer/v2"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/kv/kvdecoder"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/stretchr/testify/require"
)

// L2Accounts returns all accounts on the chain with non-zero balances
func (ch *Chain) L2Accounts() []*iscp.AgentID {
	d, err := ch.CallView(accounts.Contract.Name, accounts.FuncViewAccounts.Name)
	require.NoError(ch.Env.T, err)
	keys := d.KeysSorted()
	ret := make([]*iscp.AgentID, 0, len(keys)-1)
	for _, key := range keys {
		aid, err := codec.DecodeAgentID([]byte(key))
		require.NoError(ch.Env.T, err)
		ret = append(ret, aid)
	}
	return ret
}

func (ch *Chain) parseAccountBalance(d dict.Dict, err error) *iscp.Assets {
	require.NoError(ch.Env.T, err)
	if d.IsEmpty() {
		return iscp.NewEmptyAssets()
	}
	ret, err := iscp.AssetsFromDict(d)
	require.NoError(ch.Env.T, err)
	return ret
}

func (ch *Chain) L2Ledger() map[string]*iscp.Assets {
	accs := ch.L2Accounts()
	ret := make(map[string]*iscp.Assets)
	for i := range accs {
		ret[accs[i].String()] = ch.L2AccountAssets(accs[i])
	}
	return ret
}

func (ch *Chain) L2LedgerString() string {
	l := ch.L2Ledger()
	keys := make([]string, 0, len(l))
	for aid := range l {
		keys = append(keys, aid)
	}
	sort.Strings(keys)
	ret := ""
	for _, aid := range keys {
		ret += aid + "\n"
		ret += "        " + l[aid].String() + "\n"
	}
	return ret
}

// L2AccountAssets return all assets contained in the on-chain account controlled by the 'agentID'
func (ch *Chain) L2AccountAssets(agentID *iscp.AgentID) *iscp.Assets {
	return ch.parseAccountBalance(
		ch.CallView(accounts.Contract.Name, accounts.FuncViewBalance.Name, accounts.ParamAgentID, agentID),
	)
}

func (ch *Chain) L2AccountIotas(agentID *iscp.AgentID) uint64 {
	return ch.L2AccountAssets(agentID).Iotas
}

func (ch *Chain) L2AccountNativeTokens(agentID *iscp.AgentID, tokenID *iotago.NativeTokenID) *big.Int {
	return ch.L2AccountAssets(agentID).AmountNativeToken(tokenID)
}

func (ch *Chain) L2CommonAccountAssets() *iscp.Assets {
	return ch.L2AccountAssets(ch.CommonAccount())
}

func (ch *Chain) L2CommonAccountIotas() uint64 {
	return ch.L2AccountAssets(ch.CommonAccount()).Iotas
}

func (ch *Chain) L2CommonAccountNativeTokens(tokenID *iotago.NativeTokenID) *big.Int {
	return ch.L2AccountAssets(ch.CommonAccount()).AmountNativeToken(tokenID)
}

// L2TotalAssetsInAccounts return total sum of assets contained in the on-chain accounts
func (ch *Chain) L2TotalAssetsInAccounts() *iscp.Assets {
	return ch.parseAccountBalance(
		ch.CallView(accounts.Contract.Name, accounts.FuncViewTotalAssets.Name),
	)
}

// L2TotalIotasInAccounts return total sum of iotas
func (ch *Chain) L2TotalIotasInAccounts() uint64 {
	return ch.L2TotalAssetsInAccounts().Iotas
}

func mustNativeTokenIDFromBytes(data []byte) *iotago.NativeTokenID {
	if len(data) != iotago.NativeTokenIDLength {
		panic("len(data) != iotago.NativeTokenIDLength")
	}
	ret := new(iotago.NativeTokenID)
	copy(ret[:], data)
	return ret
}

func (ch *Chain) GetOnChainTokenIDs() []*iotago.NativeTokenID {
	res, err := ch.CallView(accounts.Contract.Name, accounts.FuncGetNativeTokenIDRegistry.Name)
	require.NoError(ch.Env.T, err)
	ret := make([]*iotago.NativeTokenID, 0, len(res))
	for k := range res {
		ret = append(ret, mustNativeTokenIDFromBytes([]byte(k)))
	}
	return ret
}

func (ch *Chain) GetFoundryOutput(sn uint32) (*iotago.FoundryOutput, error) {
	res, err := ch.CallView(accounts.Contract.Name, accounts.FuncFoundryOutput.Name,
		accounts.ParamFoundrySN, sn,
	)
	if err != nil {
		return nil, err
	}
	outBin := res.MustGet(accounts.ParamFoundryOutputBin)
	out := &iotago.FoundryOutput{}
	_, err = out.Deserialize(outBin, serializer.DeSeriModeNoValidation, nil)
	require.NoError(ch.Env.T, err)
	return out, nil
}

func (ch *Chain) GetNativeTokenIDByFoundrySN(sn uint32) (iotago.NativeTokenID, error) {
	o, err := ch.GetFoundryOutput(sn)
	if err != nil {
		return iotago.NativeTokenID{}, err
	}
	return o.MustNativeTokenID(), nil
}

type foundryParams struct {
	ch        *Chain
	user      *cryptolib.KeyPair
	sch       iotago.TokenScheme
	tag       *iotago.TokenTag
	maxSupply *big.Int
}

// CreateFoundryGasBudgetIotas always takes 1000 iotas as gas budget and assets for the call
const (
	CreateFoundryGasBudgetIotas   = 1000
	MintTokensGasBudgetIotas      = 1000
	DestroyTokensGasBudgetIotas   = 1000
	SendToL2AccountGasBudgetIotas = 1000
)

func (ch *Chain) NewFoundryParams(maxSupply interface{}) *foundryParams {
	ret := &foundryParams{
		ch:        ch,
		maxSupply: util.ToBigInt(maxSupply),
	}
	return ret
}

func (fp *foundryParams) WithUser(user *cryptolib.KeyPair) *foundryParams {
	fp.user = user
	return fp
}

func (fp *foundryParams) WithTokenScheme(sch iotago.TokenScheme) *foundryParams {
	fp.sch = sch
	return fp
}

func (fp *foundryParams) WithTag(tag *iotago.TokenTag) *foundryParams {
	fp.tag = tag
	return fp
}

func (fp *foundryParams) CreateFoundry() (uint32, iotago.NativeTokenID, error) {
	par := dict.New()
	if fp.sch != nil {
		par.Set(accounts.ParamTokenScheme, codec.EncodeTokenScheme(fp.sch))
	}
	if fp.tag != nil {
		par.Set(accounts.ParamTokenTag, codec.EncodeTokenTag(*fp.tag))
	}
	if fp.maxSupply != nil {
		par.Set(accounts.ParamMaxSupply, codec.EncodeBigIntAbs(fp.maxSupply))
	}
	user := &fp.ch.OriginatorPrivateKey
	if fp.user != nil {
		user = fp.user
	}
	req := NewCallParamsFromDic(accounts.Contract.Name, accounts.FuncFoundryCreateNew.Name, par).
		WithGasBudget(CreateFoundryGasBudgetIotas).
		AddAssetsIotas(CreateFoundryGasBudgetIotas)
	res, err := fp.ch.PostRequestSync(req, user)

	retSN := uint32(0)
	var tokenID iotago.NativeTokenID
	if err == nil {
		resDeco := kvdecoder.New(res)
		retSN = resDeco.MustGetUint32(accounts.ParamFoundrySN)
		tokenID, err = fp.ch.GetNativeTokenIDByFoundrySN(retSN)
	}
	return retSN, tokenID, err
}

func toFoundrySN(foundry interface{}) uint32 {
	switch f := foundry.(type) {
	case uint32:
		return f
	case *iotago.NativeTokenID:
		return accounts.FoundrySNFromNativeTokenID(f)
	case iotago.NativeTokenID:
		return accounts.FoundrySNFromNativeTokenID(&f)
	}
	panic(fmt.Sprintf("toFoundrySN: type %T not supported", foundry))
}

func (ch *Chain) MintTokens(foundry, amount interface{}, user *cryptolib.KeyPair) error {
	req := NewCallParams(accounts.Contract.Name, accounts.FuncFoundryModifySupply.Name,
		accounts.ParamFoundrySN, toFoundrySN(foundry),
		accounts.ParamSupplyDeltaAbs, util.ToBigInt(amount),
	).
		WithGasBudget(MintTokensGasBudgetIotas).
		AddAssetsIotas(MintTokensGasBudgetIotas)
	if user == nil {
		user = &ch.OriginatorPrivateKey
	}
	_, err := ch.PostRequestSync(req, user)
	return err
}

// DestroyTokensOnL2 destroys tokens (identified by foundry SN) on user's on-chain account
func (ch *Chain) DestroyTokensOnL2(foundryOrTokenID, amount interface{}, user *cryptolib.KeyPair) error {
	req := NewCallParams(accounts.Contract.Name, accounts.FuncFoundryModifySupply.Name,
		accounts.ParamFoundrySN, toFoundrySN(foundryOrTokenID),
		accounts.ParamSupplyDeltaAbs, util.ToBigInt(amount),
		accounts.ParamDestroyTokens, true,
	).WithGasBudget(DestroyTokensGasBudgetIotas)

	if user == nil {
		user = &ch.OriginatorPrivateKey
	}
	_, err := ch.PostRequestSync(req, user)
	return err
}

// DestroyTokensOnL1 sends tokens as assets and destroys in the same transaction
func (ch *Chain) DestroyTokensOnL1(tokenID *iotago.NativeTokenID, amount interface{}, user *cryptolib.KeyPair) error {
	req := NewCallParams(accounts.Contract.Name, accounts.FuncFoundryModifySupply.Name,
		accounts.ParamFoundrySN, toFoundrySN(tokenID),
		accounts.ParamSupplyDeltaAbs, util.ToBigInt(amount),
		accounts.ParamDestroyTokens, true,
	).WithGasBudget(DestroyTokensGasBudgetIotas).AddAssetsIotas(1000)
	req.AddAssetsNativeTokens(tokenID, amount)
	req.AddNativeTokensAllowance(tokenID, amount)
	if user == nil {
		user = &ch.OriginatorPrivateKey
	}
	_, err := ch.PostRequestSync(req, user)
	return err
}

// DepositAssets deposits assets on user's on-chain account
func (ch *Chain) DepositAssets(assets *iscp.Assets, user *cryptolib.KeyPair) error {
	req := NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).AddAssets(assets)
	_, err := ch.PostRequestSync(req, user)
	return err
}

// DepositIotas deposits assets on user's on-chain account
func (ch *Chain) DepositIotas(amount uint64, user *cryptolib.KeyPair) error {
	return ch.DepositAssets(iscp.NewAssets(amount, nil), user)
}

// SendFromL1ToL2Account sends assets from L1 address to the target account on L2
func (ch *Chain) SendFromL1ToL2Account(assets *iscp.Assets, target *iscp.AgentID, user *cryptolib.KeyPair) error {
	req := NewCallParams(accounts.Contract.Name, accounts.FuncTransferAllowanceTo.Name,
		accounts.ParamAgentID, target)

	req.AddAssets(assets).
		AddAssetsIotas(SendToL2AccountGasBudgetIotas).
		AddAllowance(assets).
		WithGasBudget(SendToL2AccountGasBudgetIotas)
	_, err := ch.PostRequestSync(req, user)
	return err
}

func (ch *Chain) SendFromL1ToL2AccountIotas(iotas uint64, target *iscp.AgentID, user *cryptolib.KeyPair) error {
	return ch.SendFromL1ToL2Account(iscp.NewAssets(iotas, nil), target, user)
}

func (ch *Chain) SendFromL1ToL2AccountNativeTokens(id iotago.NativeTokenID, target *iscp.AgentID, amount interface{}, user *cryptolib.KeyPair) error {
	return ch.SendFromL1ToL2Account(iscp.NewEmptyAssets().AddNativeTokens(id, amount), target, user)
}

// SendFromL2ToL2Account moves assets on L2 from user's account to the target
func (ch *Chain) SendFromL2ToL2Account(transfer *iscp.Assets, target *iscp.AgentID, user *cryptolib.KeyPair) error {
	req := NewCallParams(accounts.Contract.Name, accounts.FuncTransferAllowanceTo.Name,
		accounts.ParamAgentID, target)

	req.AddAssetsIotas(SendToL2AccountGasBudgetIotas).
		AddAllowance(transfer).
		WithGasBudget(SendToL2AccountGasBudgetIotas)
	_, err := ch.PostRequestSync(req, user)
	return err
}

func (ch *Chain) SendFromL2ToL2AccountIotas(iotas uint64, target *iscp.AgentID, user *cryptolib.KeyPair) error {
	return ch.SendFromL2ToL2Account(iscp.NewAssets(iotas, nil), target, user)
}

func (ch *Chain) SendFromL2ToL2AccountNativeTokens(id iotago.NativeTokenID, target *iscp.AgentID, amount interface{}, user *cryptolib.KeyPair) error {
	return ch.SendFromL2ToL2Account(iscp.NewEmptyAssets().AddNativeTokens(id, amount), target, user)
}
