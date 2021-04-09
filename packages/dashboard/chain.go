package dashboard

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"

	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/registry"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/vm/core/blob"
	"github.com/iotaledger/wasp/plugins/chains"
	"github.com/iotaledger/wasp/plugins/database"
	"github.com/labstack/echo/v4"
)

//go:embed templates/chain.tmpl
var tplChain string

func chainBreadcrumb(e *echo.Echo, chainID coretypes.ChainID) Tab {
	return Tab{
		Path:  e.Reverse("chain"),
		Title: fmt.Sprintf("Chain %.8s…", chainID),
		Href:  e.Reverse("chain", chainID.Base58()),
	}
}

func (d *Dashboard) initChain(e *echo.Echo, r renderer) {
	route := e.GET("/chain/:chainid", d.handleChain)
	route.Name = "chain"
	r[route.Path] = d.makeTemplate(e, tplChain, tplWs)
}

func (d *Dashboard) handleChain(c echo.Context) error {
	chainid, err := coretypes.ChainIDFromBase58(c.Param("chainid"))
	if err != nil {
		return err
	}

	tab := chainBreadcrumb(c.Echo(), *chainid)

	result := &ChainTemplateParams{
		BaseTemplateParams: d.BaseParams(c, tab),
		ChainID:            *chainid,
	}

	result.ChainRecord, err = registry.ChainRecordFromRegistry(chainid)
	if err != nil {
		return err
	}

	if result.ChainRecord != nil && result.ChainRecord.Active {
		result.VirtualState, result.Block, _, err = state.LoadSolidState(database.GetInstance(), chainid)
		if err != nil {
			return err
		}

		theChain := chains.AllChains().Get(chainid)

		result.Committee.Size = theChain.Committee().Size()
		result.Committee.Quorum = theChain.Committee().Quorum()
		// result.Committee.NumPeers = theChain.Committee().NumPeers()
		result.Committee.HasQuorum = theChain.Committee().QuorumIsAlive()
		result.Committee.PeerStatus = theChain.Committee().PeerStatus()
		result.RootInfo, err = fetchRootInfo(theChain)
		if err != nil {
			return err
		}

		result.Accounts, err = fetchAccounts(theChain)
		if err != nil {
			return err
		}

		result.TotalAssets, err = fetchTotalAssets(theChain)
		if err != nil {
			return err
		}

		result.Blobs, err = fetchBlobs(theChain)
		if err != nil {
			return err
		}
	}

	return c.Render(http.StatusOK, c.Path(), result)
}

func fetchAccounts(chain chain.Chain) ([]coretypes.AgentID, error) {
	accounts, err := callView(chain, accounts.Interface.Hname(), accounts.FuncAccounts, nil)
	if err != nil {
		return nil, fmt.Errorf("accountsc view call failed: %v", err)
	}

	ret := make([]coretypes.AgentID, 0)
	for k := range accounts {
		agentid, _, err := codec.DecodeAgentID([]byte(k))
		if err != nil {
			return nil, err
		}
		ret = append(ret, agentid)
	}
	return ret, nil
}

func fetchTotalAssets(chain chain.Chain) (map[ledgerstate.Color]uint64, error) {
	bal, err := callView(chain, accounts.Interface.Hname(), accounts.FuncTotalAssets, nil)
	if err != nil {
		return nil, err
	}
	return accounts.DecodeBalances(bal)
}

func fetchBlobs(chain chain.Chain) (map[hashing.HashValue]uint32, error) {
	ret, err := callView(chain, blob.Interface.Hname(), blob.FuncListBlobs, nil)
	if err != nil {
		return nil, err
	}
	return blob.DecodeDirectory(ret)
}

type ChainTemplateParams struct {
	BaseTemplateParams

	ChainID coretypes.ChainID

	ChainRecord  *registry.ChainRecord
	Block        state.Block
	VirtualState state.VirtualState
	RootInfo     RootInfo
	Accounts     []coretypes.AgentID
	TotalAssets  map[ledgerstate.Color]uint64
	Blobs        map[hashing.HashValue]uint32
	Committee    struct {
		Size       uint16
		Quorum     uint16
		NumPeers   uint16
		HasQuorum  bool
		PeerStatus []*chain.PeerStatus
	}
}
