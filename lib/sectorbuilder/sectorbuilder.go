package sectorbuilder

import (
	"unsafe"

	"github.com/filecoin-project/go-lotus/chain/address"
	sectorbuilder "github.com/filecoin-project/go-sectorbuilder"
)

type SectorSealingStatus = sectorbuilder.SectorSealingStatus

type StagedSectorMetadata = sectorbuilder.StagedSectorMetadata

const CommLen = sectorbuilder.CommitmentBytesLen

type SectorBuilder struct {
	handle unsafe.Pointer
}

type SectorBuilderConfig struct {
	SectorSize  uint64
	Miner       address.Address
	SealedDir   string
	StagedDir   string
	MetadataDir string
}

func New(cfg *SectorBuilderConfig) (*SectorBuilder, error) {
	var proverId [31]byte
	copy(proverId[:], cfg.Miner.Payload())

	sbp, err := sectorbuilder.InitSectorBuilder(cfg.SectorSize, 2, 2, 1, cfg.MetadataDir, [31]byte{}, cfg.SealedDir, cfg.StagedDir, 16)
	if err != nil {
		return nil, err
	}

	return &SectorBuilder{
		handle: sbp,
	}, nil
}

func (sb *SectorBuilder) Destroy() {
	sectorbuilder.DestroySectorBuilder(sb.handle)
}

func (sb *SectorBuilder) AddPiece(pieceKey string, pieceSize uint64, piecePath string) (uint64, error) {
	return sectorbuilder.AddPiece(sb.handle, pieceKey, pieceSize, piecePath)
}

// TODO: should *really really* return an io.ReadCloser
func (sb *SectorBuilder) ReadPieceFromSealedSector(pieceKey string) ([]byte, error) {
	return sectorbuilder.ReadPieceFromSealedSector(sb.handle, pieceKey)
}

func (sb *SectorBuilder) SealAllStagedSectors() error {
	return sectorbuilder.SealAllStagedSectors(sb.handle)
}

func (sb *SectorBuilder) SealStatus(sector uint64) (SectorSealingStatus, error) {
	return sectorbuilder.GetSectorSealingStatusByID(sb.handle, sector)
}

func (sb *SectorBuilder) GetAllStagedSectors() ([]StagedSectorMetadata, error) {
	return sectorbuilder.GetAllStagedSectors(sb.handle)
}

func (sb *SectorBuilder) GeneratePoSt(sortedCommRs [][CommLen]byte, challengeSeed [CommLen]byte) ([][]byte, []uint64, error) {
	// Wait, this is a blocking method with no way of interrupting it?
	// does it checkpoint itself?
	return sectorbuilder.GeneratePoSt(sb.handle, sortedCommRs, challengeSeed)
}

var UserBytesForSectorSize = sectorbuilder.GetMaxUserBytesPerStagedSector

func VerifySeal(sectorSize uint64, commR, commD, commRStar [CommLen]byte, proverID address.Address, sectorID uint64, proof []byte) (bool, error) {
	panic("TODO")
	// return sectorbuilder.VerifySeal(sectorSize, commR, commD, commRStar, providerID, sectorID, proof)
}

func VerifyPost(sectorSize uint64, sortedCommRs [][CommLen]byte, challengeSeed [CommLen]byte, proofs [][]byte, faults []uint64) (bool, error) {
	// sectorbuilder.VerifyPost()
	panic("no")
}