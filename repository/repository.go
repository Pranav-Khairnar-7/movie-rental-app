package repository

import (
	"sync"

	"github.com/jinzhu/gorm"
)

type QueryProcessor func(db *gorm.DB) *gorm.DB

type Repository struct {
	db *gorm.DB
}

var (
	instance *Repository
	once     sync.Once
)

// Initialize the singleton repository once with a DB instance
func GetRepository(db *gorm.DB) *Repository {
	once.Do(func() {
		instance = &Repository{db: db}
	})
	return instance
}

// Create inserts a new record; model must be a pointer to struct
func (r *Repository) Create(model interface{}) error {
	return r.db.Create(model).Error
}

// GetByID fetches a record by primary key; model must be a pointer to struct instance where result is stored
func (r *Repository) GetByID(model interface{}, id interface{}) error {
	return r.db.First(model, "id = ?", id).Error
}

// Get
func (r *Repository) GetWithFilters(model interface{}, processors []QueryProcessor) error {
	query := r.db.Model(model)

	for _, proc := range processors {
		if proc != nil {
			query = proc(query)
		}
	}
	return query.Find(model).Limit(1).Error
}

// GetAll fetches records into a slice pointed by destSlicePtr.
// destSlicePtr must be a pointer to a slice of structs, e.g. *[]Movie
// modelInstance is a zero value instance of the struct (e.g. Movie{}) used to tell GORM which table to query.
// limit and offset are optional (pass nil to ignore).
// processor is optional query processor for joins/filters.
// preloads is a variadic list of association names to preload.
func (r *Repository) GetAll(
	destSlicePtr interface{},
	modelInstance interface{},
	limit, offset *int,
	processor []QueryProcessor,
	preloads ...string,
) error {
	query := r.db.Model(modelInstance)

	for _, proc := range processor {
		if proc != nil {
			query = proc(query)
		}
	}

	for _, assoc := range preloads {
		query = query.Preload(assoc)
	}

	if limit != nil && *limit > 0 {
		query = query.Limit(*limit)
	}
	if offset != nil && *offset >= 0 {
		query = query.Offset(*offset)
	}

	return query.Find(destSlicePtr).Error
}

// Update updates fields in the model using the provided map of field names to values.
// model must be a pointer to a struct instance that includes primary key for where clause.
func (r *Repository) Update(model interface{}, fields map[string]interface{}) error {
	return r.db.Model(model).Updates(fields).Error
}

// Delete performs a soft delete (sets deleted_at) on the model.
// model must be a pointer to struct instance.
func (r *Repository) Delete(model interface{}) error {
	return r.db.Delete(model).Error
}

// UnscopedDelete performs a hard delete bypassing soft delete on the model.
// model must be a pointer to struct instance.
func (r *Repository) UnscopedDelete(model interface{}) error {
	return r.db.Unscoped().Delete(model).Error
}
