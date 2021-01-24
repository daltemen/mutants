package repositories

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"log"
	"mutants/app/mutants"
	"strings"
)

type mySqlStrongWriteRepository struct {
	conn *gorm.DB
}

func NewMySqlStrongRepository(conn *gorm.DB) StrongWriteRepository {
	return &mySqlStrongWriteRepository{conn: conn}
}

func (m *mySqlStrongWriteRepository) SaveDna(ctx context.Context, human mutants.Human) error {
	uid := uuid.NewV4()
	humanType := human.GetType().ToString()

	segments := mapMatrixSegmentsToSlice(human.Dna)
	segmentsBytes, err := json.Marshal(segments)
	if err != nil {
		log.Println(err)
		return err // TODO: wrap error ?
	}

	return m.conn.Create(&DnaDB{ID: uid, HumanType: humanType, Segments: segmentsBytes}).Error // TODO: wrap error
}

type DnaDB struct {
	ID        uuid.UUID `gorm:"primary_key" sql:"type:CHAR(36)"`
	HumanType string    `sql:"type:CHAR(36)"`
	Segments  datatypes.JSON
}

type SegmentsSlice struct {
	Dna []string `json:"dna"`
}

func mapMatrixSegmentsToSlice(dna [][]string) SegmentsSlice {
	result := make([]string, len(dna))

	for i, row := range dna {
		result[i] = strings.Join(row, "")
	}
	return SegmentsSlice{
		Dna: result,
	}
}
