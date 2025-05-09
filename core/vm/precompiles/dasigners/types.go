package dasigners

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type BN254G1Point = struct {
	X *big.Int "json:\"X\""
	Y *big.Int "json:\"Y\""
}

type BN254G2Point = struct {
	X [2]*big.Int "json:\"X\""
	Y [2]*big.Int "json:\"Y\""
}

type IDASignersSignerDetail = struct {
	Signer common.Address "json:\"signer\""
	Socket string         "json:\"socket\""
	PkG1   BN254G1Point   "json:\"pkG1\""
	PkG2   BN254G2Point   "json:\"pkG2\""
}

type IDASignersParams = struct {
	TokensPerVote     *big.Int "json:\"tokensPerVote\""
	MaxVotesPerSigner *big.Int "json:\"maxVotesPerSigner\""
	MaxQuorums        *big.Int "json:\"maxQuorums\""
	EpochBlocks       *big.Int "json:\"epochBlocks\""
	EncodedSlices     *big.Int "json:\"encodedSlices\""
}

func NewBN254G1Point(b []byte) BN254G1Point {
	return BN254G1Point{
		X: new(big.Int).SetBytes(b[:32]),
		Y: new(big.Int).SetBytes(b[32:64]),
	}
}

func SerializeG1(p BN254G1Point) []byte {
	b := make([]byte, 0)
	b = append(b, common.LeftPadBytes(p.X.Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(p.Y.Bytes(), 32)...)
	return b
}

func NewBN254G2Point(b []byte) BN254G2Point {
	return BN254G2Point{
		X: [2]*big.Int{
			new(big.Int).SetBytes(b[:32]),
			new(big.Int).SetBytes(b[32:64]),
		},
		Y: [2]*big.Int{
			new(big.Int).SetBytes(b[64:96]),
			new(big.Int).SetBytes(b[96:128]),
		},
	}
}

func SerializeG2(p BN254G2Point) []byte {
	b := make([]byte, 0)
	b = append(b, common.LeftPadBytes(p.X[0].Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(p.X[1].Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(p.Y[0].Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(p.Y[1].Bytes(), 32)...)
	return b
}

var (
	signerKey                = []byte{0x00} // signer => registered signer info
	quorumKey                = []byte{0x01} // epoch => quorumId => quorum
	registrationKey          = []byte{0x02} // epoch => signer => signature(vrf)
	votesKey                 = []byte{0x03} // epoch => signer => votes
	quorumCountKey           = []byte{0x04} // epoch => quorum count
	epochNumberKey           = []byte{0x05} // epoch number
	epochBlockKey            = []byte{0x06} // epoch number => block number
	epochRegistrationKey     = []byte{0x07} // epoch number => registration count
	epochRegisteredSignerKey = []byte{0x08} // epoch number => index => signer
)

func SignerKey(account common.Address) common.Hash {
	return crypto.Keccak256Hash(append(signerKey, account.Bytes()...))
}

func QuorumKey(epochNumber uint64, quorumId uint64) common.Hash {
	return crypto.Keccak256Hash(append(append(quorumKey, common.Uint64ToBytes(epochNumber)...), common.Uint64ToBytes(quorumId)...))
}

func RegistrationKey(epochNumber uint64, account common.Address) common.Hash {
	return crypto.Keccak256Hash(append(append(registrationKey, common.Uint64ToBytes(epochNumber)...), account.Bytes()...))
}

func VotesKey(epochNumber uint64, account common.Address) common.Hash {
	return crypto.Keccak256Hash(append(append(votesKey, common.Uint64ToBytes(epochNumber)...), account.Bytes()...))
}

func QuorumCountKey(epochNumber uint64) common.Hash {
	return crypto.Keccak256Hash(append(quorumCountKey, common.Uint64ToBytes(epochNumber)...))
}

func EpochNumberKey() common.Hash {
	return crypto.Keccak256Hash(epochNumberKey)
}

func EpochBlockKey(epochNumber uint64) common.Hash {
	return crypto.Keccak256Hash(append(epochBlockKey, common.Uint64ToBytes(epochNumber)...))
}

func EpochRegistrationKey(epochNumber uint64) common.Hash {
	return crypto.Keccak256Hash(append(epochRegistrationKey, common.Uint64ToBytes(epochNumber)...))
}

func EpochRegisteredSignerKey(epochNumber uint64, index uint64) common.Hash {
	return crypto.Keccak256Hash(append(append(epochRegisteredSignerKey, common.Uint64ToBytes(epochNumber)...), common.Uint64ToBytes(index)...))
}
