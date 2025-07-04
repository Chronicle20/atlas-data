package reactor

import (
	"atlas-data/database"
	"atlas-data/document"
	"atlas-data/xml"
	"context"
	"fmt"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"path/filepath"
	"strconv"
	"strings"
)

func NewStorage(l logrus.FieldLogger, db *gorm.DB) *document.Storage[string, RestModel] {
	return document.NewStorage(l, db, GetModelRegistry(), "REACTOR")
}

func Register(s *document.Storage[string, RestModel]) func(ctx context.Context) func(r model.Provider[RestModel]) error {
	return func(ctx context.Context) func(r model.Provider[RestModel]) error {
		return func(r model.Provider[RestModel]) error {
			m, err := r()
			if err != nil {
				return err
			}
			_, err = s.Add(ctx)(m)()
			if err != nil {
				return err
			}
			return nil
		}
	}
}

func extractPathAndID(path string) (string, uint32, error) {
	// Extract the base filename
	base := filepath.Base(path)

	// Trim the ".img.xml" extension
	if !strings.HasSuffix(base, ".img.xml") {
		return "", 0, fmt.Errorf("invalid file format: %s", base)
	}
	idStr := strings.TrimSuffix(base, ".img.xml")

	// Convert to uint32
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return "", 0, fmt.Errorf("failed to convert ID to uint32: %w", err)
	}

	// Extract the directory
	dir := filepath.Dir(path) + "/"

	return dir, uint32(id), nil
}

func RegisterReactor(db *gorm.DB) func(l logrus.FieldLogger) func(ctx context.Context) func(path string) error {
	return func(l logrus.FieldLogger) func(ctx context.Context) func(path string) error {
		return func(ctx context.Context) func(path string) error {
			return func(path string) error {
				parentPath, reactorId, err := extractPathAndID(path)
				if err != nil {
					return err
				}
				return database.ExecuteTransaction(db, func(tx *gorm.DB) error {
					return Register(NewStorage(l, tx))(ctx)(Read(l)(parentPath, reactorId, xml.FromParentPathProvider(7)))
				})
			}
		}
	}
}
