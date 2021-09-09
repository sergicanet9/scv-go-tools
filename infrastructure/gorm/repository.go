package gorm

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type GormRepository struct {
	db           *gorm.DB
	constructor  func() interface{}
	defaultJoins []string
}

// NewGormRepository returns a new base repository that implements TransactionRepository
func NewGormRepository(db *gorm.DB, constructor func() interface{}, defaultJoins ...string) *GormRepository {
	return &GormRepository{
		db,
		constructor, // constructor of the repository target entity
		defaultJoins,
	}
}

func (r *GormRepository) Create(entity interface{}, include ...string) (interface{}, error) {
	res := r.dbWithPreloads(include).Create(entity)
	return entity, r.handleError(res)
}

func (r *GormRepository) Get(filter interface{}, include ...string) (result []interface{}, _ error) {
	res := r.dbWithPreloads(include).Where(filter).Find(result)
	return result, r.handleError(res)
}

func (r *GormRepository) GetByID(ID int, include ...string) (result interface{}, _ error) {
	res := r.dbWithPreloads(include).First(result, ID)
	return result, r.handleError(res)
}

func (r *GormRepository) GetRaw(query string, params map[string]interface{}) (result interface{}, _ error) {
	res := r.db.Raw(query, params)
	return result, r.handleError(res)
}

func (r *GormRepository) Update(ID int, entity interface{}, include ...string) error {
	var updates map[string]interface{}
	bytes, err := json.Marshal(entity)
	if err != nil {
		return fmt.Errorf("error: %w", err)

	}
	json.Unmarshal(bytes, &updates)

	res := r.dbWithPreloads(include).First(r.constructor(), ID).Updates(updates)
	return r.handleError(res)
}

func (r *GormRepository) Delete(ID int, include ...string) error {
	res := r.dbWithPreloads(include).Delete(r.constructor(), ID)
	return r.handleError(res)
}

func (r *GormRepository) handleError(res *gorm.DB) error {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("error: %w", res.Error)
		return err
	}
	return nil
}

func (r *GormRepository) dbWithPreloads(preloads []string) *gorm.DB {
	dbConn := r.db

	for _, join := range r.defaultJoins {
		dbConn = dbConn.Joins(join)
	}

	for _, preload := range preloads {
		dbConn = dbConn.Preload(preload)
	}

	return dbConn
}
