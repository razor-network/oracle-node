package cmd

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"os"
	"path/filepath"
	"razor/cmd/mocks"
	"razor/path"
	pathPkgMocks "razor/path/mocks"
	"razor/rpc"
	"razor/utils"
	utilsPkgMocks "razor/utils/mocks"
	"strings"
	"testing"
	"time"
)

var (
	utilsMock             *utilsPkgMocks.Utils
	clientUtilsMock       *utilsPkgMocks.ClientUtils
	timeUtilsMock         *utilsPkgMocks.TimeUtils
	osUtilsMock           *utilsPkgMocks.OSUtils
	coinUtilsMock         *utilsPkgMocks.CoinUtils
	merkleUtilsMock       *utilsPkgMocks.MerkleTreeInterface
	ioUtilsMock           *utilsPkgMocks.IOUtils
	abiUtilsMock          *utilsPkgMocks.ABIUtils
	bindUtilsMock         *utilsPkgMocks.BindUtils
	blockManagerUtilsMock *utilsPkgMocks.BlockManagerUtils
	stakeManagerUtilsMock *utilsPkgMocks.StakeManagerUtils
	assetManagerUtilsMock *utilsPkgMocks.AssetManagerUtils
	voteManagerUtilsMock  *utilsPkgMocks.VoteManagerUtils
	bindingUtilsMock      *utilsPkgMocks.BindingsUtils
	jsonUtilsMock         *utilsPkgMocks.JsonUtils
	stakedTokenUtilsMock  *utilsPkgMocks.StakedTokenUtils
	retryUtilsMock        *utilsPkgMocks.RetryUtils
	fileUtilsMock         *utilsPkgMocks.FileUtils
	gasUtilsMock          *utilsPkgMocks.GasUtils
	cmdUtilsMock          *mocks.UtilsCmdInterface
	flagSetMock           *mocks.FlagSetInterface
	transactionMock       *mocks.TransactionInterface
	stakeManagerMock      *mocks.StakeManagerInterface
	blockManagerMock      *mocks.BlockManagerInterface
	voteManagerMock       *mocks.VoteManagerInterface
	keystoreMock          *mocks.KeystoreInterface
	tokenManagerMock      *mocks.TokenManagerInterface
	assetManagerMock      *mocks.AssetManagerInterface
	cryptoMock            *mocks.CryptoInterface
	viperMock             *mocks.ViperInterface
	timeMock              *mocks.TimeInterface
	stringMock            *mocks.StringInterface
	abiMock               *mocks.AbiInterface
	osMock                *mocks.OSInterface
	pathMock              *pathPkgMocks.PathInterface
	osPathMock            *pathPkgMocks.OSInterface
)

func SetUpMockInterfaces() {
	utilsMock = new(utilsPkgMocks.Utils)
	razorUtils = utilsMock

	clientUtilsMock = new(utilsPkgMocks.ClientUtils)
	clientUtils = clientUtilsMock

	cmdUtilsMock = new(mocks.UtilsCmdInterface)
	cmdUtils = cmdUtilsMock

	fileUtilsMock = new(utilsPkgMocks.FileUtils)
	fileUtils = fileUtilsMock

	timeUtilsMock = new(utilsPkgMocks.TimeUtils)
	utils.Time = timeUtilsMock

	osUtilsMock = new(utilsPkgMocks.OSUtils)
	utils.OS = osUtilsMock

	coinUtilsMock = new(utilsPkgMocks.CoinUtils)
	utils.CoinInterface = coinUtilsMock

	merkleUtilsMock = new(utilsPkgMocks.MerkleTreeInterface)
	merkleUtils = merkleUtilsMock

	ioUtilsMock = new(utilsPkgMocks.IOUtils)
	utils.IOInterface = ioUtilsMock

	abiUtilsMock = new(utilsPkgMocks.ABIUtils)
	utils.ABIInterface = abiUtilsMock

	bindUtilsMock = new(utilsPkgMocks.BindUtils)
	utils.BindingsInterface = bindingUtilsMock

	blockManagerUtilsMock = new(utilsPkgMocks.BlockManagerUtils)
	utils.BlockManagerInterface = blockManagerUtilsMock

	stakeManagerUtilsMock = new(utilsPkgMocks.StakeManagerUtils)
	utils.StakeManagerInterface = stakeManagerUtilsMock

	assetManagerUtilsMock = new(utilsPkgMocks.AssetManagerUtils)
	utils.AssetManagerInterface = assetManagerUtilsMock

	voteManagerUtilsMock = new(utilsPkgMocks.VoteManagerUtils)
	utils.VoteManagerInterface = voteManagerUtilsMock

	bindingUtilsMock = new(utilsPkgMocks.BindingsUtils)
	utils.BindInterface = bindUtilsMock

	jsonUtilsMock = new(utilsPkgMocks.JsonUtils)
	utils.JsonInterface = jsonUtilsMock

	stakedTokenUtilsMock = new(utilsPkgMocks.StakedTokenUtils)
	utils.StakedTokenInterface = stakedTokenUtilsMock

	retryUtilsMock = new(utilsPkgMocks.RetryUtils)
	utils.RetryInterface = retryUtilsMock

	flagSetMock = new(mocks.FlagSetInterface)
	flagSetUtils = flagSetMock

	transactionMock = new(mocks.TransactionInterface)
	transactionUtils = transactionMock

	fileUtilsMock = new(utilsPkgMocks.FileUtils)
	fileUtils = fileUtilsMock

	gasUtilsMock = new(utilsPkgMocks.GasUtils)
	gasUtils = gasUtilsMock

	stakeManagerMock = new(mocks.StakeManagerInterface)
	stakeManagerUtils = stakeManagerMock

	blockManagerMock = new(mocks.BlockManagerInterface)
	blockManagerUtils = blockManagerMock

	voteManagerMock = new(mocks.VoteManagerInterface)
	voteManagerUtils = voteManagerMock

	keystoreMock = new(mocks.KeystoreInterface)
	keystoreUtils = keystoreMock

	tokenManagerMock = new(mocks.TokenManagerInterface)
	tokenManagerUtils = tokenManagerMock

	assetManagerMock = new(mocks.AssetManagerInterface)
	assetManagerUtils = assetManagerMock

	cryptoMock = new(mocks.CryptoInterface)
	cryptoUtils = cryptoMock

	viperMock = new(mocks.ViperInterface)
	viperUtils = viperMock

	timeMock = new(mocks.TimeInterface)
	timeUtils = timeMock

	stringMock = new(mocks.StringInterface)
	stringUtils = stringMock

	abiMock = new(mocks.AbiInterface)
	abiUtils = abiMock

	osMock = new(mocks.OSInterface)
	osUtils = osMock

	pathMock = new(pathPkgMocks.PathInterface)
	pathUtils = pathMock
	path.PathUtilsInterface = pathMock

	osPathMock = new(pathPkgMocks.OSInterface)
	path.OSUtilsInterface = osPathMock
}

var privateKey, _ = ecdsa.GenerateKey(crypto.S256(), rand.Reader)
var TxnOpts, _ = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31000)) // Used any random big int for chain ID

var rpcManager = rpc.RPCManager{
	BestEndpoint: &rpc.RPCEndpoint{
		Client: &ethclient.Client{},
	},
}
var rpcParameters = rpc.RPCParameters{
	Ctx:        context.Background(),
	RPCManager: &rpcManager,
}

func TestInvokeFunctionWithRetryAttempts(t *testing.T) {
	tests := []struct {
		name         string
		methodName   string
		timeout      time.Duration
		expectError  bool
		expectedErr  string
		expectedVals bool
	}{
		{
			name:         "Normal Case - Fast Method",
			methodName:   "FastMethod",
			timeout:      5 * time.Second,
			expectError:  false,
			expectedErr:  "",
			expectedVals: true,
		},
		{
			name:         "Timeout Case - Slow Method",
			methodName:   "SlowMethod",
			timeout:      0 * time.Second,
			expectError:  true,
			expectedErr:  "context deadline exceeded",
			expectedVals: false,
		},
	}

	// Dummy RPC struct
	dummyRPC := &DummyRPC{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up context with timeout for each test case
			ctx, cancel := context.WithTimeout(context.Background(), tt.timeout)
			defer cancel()

			SetUpMockInterfaces()
			retryUtilsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(4))

			localRPCParameters := rpc.RPCParameters{
				Ctx:        ctx,
				RPCManager: &rpcManager,
			}
			returnedValues, err := utils.InvokeFunctionWithRetryAttempts(localRPCParameters, dummyRPC, tt.methodName)
			fmt.Println("Error: ", err)

			if tt.expectError {
				assert.Error(t, err)
				assert.True(t, strings.Contains(err.Error(), tt.expectedErr), "Expected error to contain: %v, but got: %v", tt.expectedErr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			if tt.expectedVals {
				assert.NotNil(t, returnedValues)
			} else {
				assert.Nil(t, returnedValues)
			}
		})
	}
}

// Dummy interface with methods
type DummyRPC struct{}

// A fast method that simulates successful execution, added client as the parameter because generic retry functions expects client as the first parameter
func (d *DummyRPC) FastMethod(client *ethclient.Client) error {
	return nil
}

// A slow method that simulates a long-running process, added client as the parameter because generic retry functions expects client as the first parameter
func (d *DummyRPC) SlowMethod(client *ethclient.Client) error {
	fmt.Println("Sleeping...")
	time.Sleep(3 * time.Second) // Simulate delay to trigger timeout
	return nil
}

func setupTestEndpointsEnvironment() {
	var testDir = "/tmp/test_rzr"
	pathMock.On("GetDefaultPath").Return(testDir, nil)
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		log.Fatalf("failed to create test directory: %s", err.Error())
	}

	mockEndpoints := `["https://testnet.skalenodes.com/v1/juicy-low-small-testnet"]`
	mockFilePath := filepath.Join(testDir, "endpoints.json")
	err = os.WriteFile(mockFilePath, []byte(mockEndpoints), 0644)
	if err != nil {
		log.Fatalf("failed to write mock endpoints.json: %s", err.Error())
	}
}
