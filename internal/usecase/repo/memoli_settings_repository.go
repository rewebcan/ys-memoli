package repo

import (
	"errors"
	"fmt"
	"github.com/rewebcan/ys-memoli/pkg/memoli"
)

var NotExistsError = errors.New("not exists")
var InvalidArgumentError = errors.New("invalid argument")

type MemoliSettingsRepository struct {
	mdb *memoli.Bucket
}

func NewMemoliDBSettingsRepository(mdb *memoli.Bucket) *MemoliSettingsRepository {
	return &MemoliSettingsRepository{mdb}
}

func (r *MemoliSettingsRepository) Get(key string) (string, error) {
	val, ok := r.mdb.Get("settings").(string)
	if !ok {
		return "", fmt.Errorf("%w: %s", NotExistsError, key)
	}

	return val, nil
}

func (r *MemoliSettingsRepository) Set(key string, value string) error {
	if value == "" {
		return fmt.Errorf("%w: %s", InvalidArgumentError, "value can not be empty")
	}
	r.mdb.Set(key, value)

	return nil
}
