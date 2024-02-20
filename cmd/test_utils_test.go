package cmd

import (
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	accountsPkgMocks "razor/accounts/mocks"
	"razor/cmd/mocks"
	"razor/path"
	pathPkgMocks "razor/path/mocks"
	"razor/utils"
	utilsPkgMocks "razor/utils/mocks"
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
	accountUtilsMock      *utilsPkgMocks.AccountsUtils
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
	accountsMock          *accountsPkgMocks.AccountInterface
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

	accountUtilsMock = new(utilsPkgMocks.AccountsUtils)
	utils.AccountsInterface = accountUtilsMock

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

	osPathMock = new(pathPkgMocks.OSInterface)
	path.OSUtilsInterface = osPathMock

	accountsMock = new(accountsPkgMocks.AccountInterface)
	accountUtils = accountsMock
}

var privateKey, _ = ecdsa.GenerateKey(crypto.S256(), rand.Reader)
var TxnOpts, _ = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31000))
