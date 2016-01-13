package mobitool

import (
	"encoding/binary"
	"log"
	"os"
)

type PalmDatabase struct {
	Name               [32]byte
	Attributes         [2]byte
	Version            int16
	CreationDate       int32
	ModificationDate   int32
	LastBackupDate     int32
	ModificationNumber int32
	AppInfoID          int32
	SortInfoID         int32
	Type               int32
	Creator            int32
	UniqueIdSeed       int32
	NextRecordListID   int32
	NumberOfRecords    int16
}

type PDBRecordHeader struct {
	Offset     int32
	Attributes byte
	_          [3]byte
}

type PalmDocHeader struct {
	Compression     int16
	_               int16
	TextLength      int32
	RecordCount     int16
	RecordSize      int16
	CurrentPosition int32
}

type MobiHeader struct {
	Identifier                       [4]byte
	MobiType                         int32
	TextEncoding                     int32
	UniqueID                         int32
	FileVersion                      int32
	OrtographicIndex                 int32
	InflectionIndex                  int32
	IndexNames                       int32
	IndexKeys                        int32
	ExtraIndexes                     [6]int32
	FirstNonBookIndex                int32
	FullNameOffset                   int32
	FullNameLength                   int32
	Locale                           int32
	InputLanguage                    int32
	OutputLanguage                   int32
	MinVersion                       int32
	FirstImageIndex                  int32
	HuffmanRecordOffset              int32
	HuffmanRecordCount               int32
	HuffmanTableOffset               int32
	HuffmanTableLength               int32
	ExthFlags                        int32
	_                                [2]byte
	_                                [4]byte
	DRMOffset                        int32
	DRMCount                         int32
	DRMSize                          int32
	DRMFlags                         int32
	_                                [8]byte
	FirstContentRecordNumber         int32
	LastContentRecordNumber          int32
	_                                [4]byte
	FCISRecordNumber                 int32
	_                                [4]byte
	FLISRecordNumber                 int32
	_                                [4]byte
	_                                [8]byte
	_                                [4]byte
	FirstCompilationDataSectionCount int32
	NumberOfCompilationDataSections  int32
	_                                [4]byte
	ExtraRecordDataFlags             int32
	IndxRecordOffset                 int32
	_                                [4]byte
	_                                [4]byte
	_                                [4]byte
	_                                [4]byte
	_                                [4]byte
	_                                [4]byte
}

type MobiFileContent struct {
	PDHead PalmDocHeader
	Header MobiHeader
	Data   [4]byte
}

func Decode(filename string) (PalmDatabase, error) {
	rd, err := os.Open(filename)
	if err != nil {
		log.Print("Error while open file ", filename, err)
		return PalmDatabase{}, err
	}

	var content PalmDatabase

	if err = binary.Read(rd, binary.BigEndian, &content); err != nil {
		log.Print("Error while decode file ", filename, err)
		return PalmDatabase{}, err
	}

	var i int16
	for i = 0; i < content.NumberOfRecords; i++ {
		var header PDBRecordHeader
		if err = binary.Read(rd, binary.BigEndian, &header); err != nil {
			log.Print("Error while decode header ", filename, err)
			return PalmDatabase{}, err
		}
		log.Printf("Header %d : %v", i, header)
	}

	return content, nil
}
