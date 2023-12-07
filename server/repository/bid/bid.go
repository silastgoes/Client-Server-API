package bid

import (
	"context"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silastgoes/client-server-api/models"
	"gorm.io/gorm"
)

func Save(db *gorm.DB, bid *models.Bid) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	if err := db.WithContext(ctx).Create(bid).Error; err != nil {
		return err
	}
	return nil
}
