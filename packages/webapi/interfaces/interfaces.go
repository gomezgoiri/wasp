package interfaces

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/vm/core/blocklog"
	"github.com/iotaledger/wasp/packages/webapi/dto"
	"github.com/iotaledger/wasp/packages/webapi/models"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/chain"

	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/registry"
)

var (
	ErrUnableToGetLatestState = errors.New("unable to get latest state")
	ErrUnableToGetReceipt     = errors.New("unable to get request receipt from block state")
	ErrAlreadyProcessed       = errors.New("request already processed")
	ErrNoBalanceOnAccount     = errors.New("no balance on account")
	ErrInvalidNonce           = errors.New("invalid nonce")
	ErrChainNotFound          = errors.New("chain not found")
	ErrCantDeleteLastUser     = errors.New("you can't delete the last user")
)

type APIController interface {
	Name() string
	RegisterPublic(publicAPI echoswagger.ApiGroup, mocker Mocker)
	RegisterAdmin(adminAPI echoswagger.ApiGroup, mocker Mocker)
}

type ChainService interface {
	ActivateChain(chainID isc.ChainID) error
	SetChainRecord(chainRecord *registry.ChainRecord) error
	DeactivateChain(chainID isc.ChainID) error
	GetAllChainIDs() ([]isc.ChainID, error)
	HasChain(chainID isc.ChainID) bool
	GetChainByID(chainID isc.ChainID) chain.Chain
	GetChainInfoByChainID(chainID isc.ChainID) (*dto.ChainInfo, error)
	GetContracts(chainID isc.ChainID) (dto.ContractsMap, error)
	GetEVMChainID(chainID isc.ChainID) (uint16, error)
	GetState(chainID isc.ChainID, stateKey []byte) (state []byte, err error)
	WaitForRequestProcessed(ctx context.Context, chainID isc.ChainID, requestID isc.RequestID, timeout time.Duration) (*isc.Receipt, *isc.VMError, error)
}

type EVMService interface {
	HandleJSONRPC(chainID isc.ChainID, request *http.Request, response *echo.Response) error
	GetRequestID(chainID isc.ChainID, hash string) (isc.RequestID, error)
}

type MetricsService interface {
	GetAllChainsMetrics() *dto.ChainMetrics
	GetChainConsensusPipeMetrics(chainID isc.ChainID) *models.ConsensusPipeMetrics
	GetChainConsensusWorkflowMetrics(chainID isc.ChainID) *models.ConsensusWorkflowMetrics
	GetChainMetrics(chainID isc.ChainID) *dto.ChainMetrics
}

var ErrPeerNotFound = errors.New("couldn't find peer")

type NodeService interface {
	AddAccessNode(chainID isc.ChainID, peer string) error
	DeleteAccessNode(chainID isc.ChainID, peer string) error
	SetNodeOwnerCertificate(publicKey *cryptolib.PublicKey, ownerAddress iotago.Address) ([]byte, error)
	ShutdownNode()
}

type RegistryService interface {
	GetChainRecordByChainID(chainID isc.ChainID) (*registry.ChainRecord, error)
}

type CommitteeService interface {
	GetCommitteeInfo(chainID isc.ChainID) (*dto.ChainNodeInfo, error)
	GetPublicKey() *cryptolib.PublicKey
}

type PeeringService interface {
	DistrustPeer(name string) (*dto.PeeringNodeIdentity, error)
	GetIdentity() *dto.PeeringNodeIdentity
	GetRegisteredPeers() []*dto.PeeringNodeStatus
	GetTrustedPeers() ([]*dto.PeeringNodeIdentity, error)
	IsPeerTrusted(publicKey *cryptolib.PublicKey) error
	TrustPeer(name string, pubkey *cryptolib.PublicKey, peeringURL string) (*dto.PeeringNodeIdentity, error)
}

type OffLedgerService interface {
	EnqueueOffLedgerRequest(chainID isc.ChainID, request []byte) error
	ParseRequest(payload []byte) (isc.OffLedgerRequest, error)
}

type UserService interface {
	AddUser(username string, password string, permissions []string) error
	DeleteUser(username string) error
	GetUser(username string) (*models.User, error)
	GetUsers() []*models.User
	UpdateUserPassword(username string, password string) error
	UpdateUserPermissions(username string, permissions []string) error
}

type VMService interface {
	CallView(chainState state.State, chain chain.Chain, contractName isc.Hname, functionName isc.Hname, params dict.Dict) (dict.Dict, error)
	CallViewByChainID(chainID isc.ChainID, contractName isc.Hname, functionName isc.Hname, params dict.Dict) (dict.Dict, error)
	ParseReceipt(chain chain.Chain, receipt *blocklog.RequestReceipt) (*isc.Receipt, *isc.VMError, error)
	GetReceipt(chainID isc.ChainID, requestID isc.RequestID) (ret *isc.Receipt, vmError *isc.VMError, err error)
}

type Mocker interface {
	Get(i interface{}) interface{}
}
