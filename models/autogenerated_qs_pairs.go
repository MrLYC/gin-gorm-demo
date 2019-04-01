// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set PairQuerySet

// PairQuerySet is an queryset type for Pair
type PairQuerySet struct {
	db *gorm.DB
}

// NewPairQuerySet constructs new PairQuerySet
func NewPairQuerySet(db *gorm.DB) PairQuerySet {
	return PairQuerySet{
		db: db.Model(&Pair{}),
	}
}

func (qs PairQuerySet) w(db *gorm.DB) PairQuerySet {
	return NewPairQuerySet(db)
}

// Create is an autogenerated method
// nolint: dupl
func (o *Pair) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Pair) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) All(ret *[]Pair) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) CreatedAtEq(createdAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("create_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) CreatedAtGt(createdAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("create_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) CreatedAtGte(createdAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("create_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) CreatedAtLt(createdAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("create_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) CreatedAtLte(createdAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("create_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) CreatedAtNe(createdAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("create_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) Delete() error {
	return qs.db.Delete(Pair{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Pair{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Pair{})
	return db.RowsAffected, db.Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeletedAtEq(deletedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeletedAtGt(deletedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeletedAtGte(deletedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeletedAtIsNotNull() PairQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeletedAtIsNull() PairQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeletedAtLt(deletedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeletedAtLte(deletedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) DeletedAtNe(deletedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) GetUpdater() PairUpdater {
	return NewPairUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) IDEq(ID int) PairQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) IDGt(ID int) PairQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) IDGte(ID int) PairQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) IDIn(ID ...int) PairQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) IDLt(ID int) PairQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) IDLte(ID int) PairQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) IDNe(ID int) PairQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) IDNotIn(ID ...int) PairQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// KeyEq is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) KeyEq(key string) PairQuerySet {
	return qs.w(qs.db.Where("key = ?", key))
}

// KeyIn is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) KeyIn(key ...string) PairQuerySet {
	if len(key) == 0 {
		qs.db.AddError(errors.New("must at least pass one key in KeyIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("key IN (?)", key))
}

// KeyNe is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) KeyNe(key string) PairQuerySet {
	return qs.w(qs.db.Where("key != ?", key))
}

// KeyNotIn is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) KeyNotIn(key ...string) PairQuerySet {
	if len(key) == 0 {
		qs.db.AddError(errors.New("must at least pass one key in KeyNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("key NOT IN (?)", key))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) Limit(limit int) PairQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) Offset(offset int) PairQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs PairQuerySet) One(ret *Pair) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) OrderAscByCreatedAt() PairQuerySet {
	return qs.w(qs.db.Order("create_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) OrderAscByDeletedAt() PairQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) OrderAscByID() PairQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) OrderAscByUpdatedAt() PairQuerySet {
	return qs.w(qs.db.Order("update_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) OrderDescByCreatedAt() PairQuerySet {
	return qs.w(qs.db.Order("create_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) OrderDescByDeletedAt() PairQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) OrderDescByID() PairQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) OrderDescByUpdatedAt() PairQuerySet {
	return qs.w(qs.db.Order("update_at DESC"))
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) UpdatedAtEq(updatedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("update_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) UpdatedAtGt(updatedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("update_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) UpdatedAtGte(updatedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("update_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) UpdatedAtLt(updatedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("update_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) UpdatedAtLte(updatedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("update_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) UpdatedAtNe(updatedAt time.Time) PairQuerySet {
	return qs.w(qs.db.Where("update_at != ?", updatedAt))
}

// ValueEq is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) ValueEq(value string) PairQuerySet {
	return qs.w(qs.db.Where("value = ?", value))
}

// ValueIn is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) ValueIn(value ...string) PairQuerySet {
	if len(value) == 0 {
		qs.db.AddError(errors.New("must at least pass one value in ValueIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("value IN (?)", value))
}

// ValueNe is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) ValueNe(value string) PairQuerySet {
	return qs.w(qs.db.Where("value != ?", value))
}

// ValueNotIn is an autogenerated method
// nolint: dupl
func (qs PairQuerySet) ValueNotIn(value ...string) PairQuerySet {
	if len(value) == 0 {
		qs.db.AddError(errors.New("must at least pass one value in ValueNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("value NOT IN (?)", value))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u PairUpdater) SetCreatedAt(createdAt time.Time) PairUpdater {
	u.fields[string(PairDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u PairUpdater) SetDeletedAt(deletedAt *time.Time) PairUpdater {
	u.fields[string(PairDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u PairUpdater) SetID(ID int) PairUpdater {
	u.fields[string(PairDBSchema.ID)] = ID
	return u
}

// SetKey is an autogenerated method
// nolint: dupl
func (u PairUpdater) SetKey(key string) PairUpdater {
	u.fields[string(PairDBSchema.Key)] = key
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u PairUpdater) SetUpdatedAt(updatedAt time.Time) PairUpdater {
	u.fields[string(PairDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetValue is an autogenerated method
// nolint: dupl
func (u PairUpdater) SetValue(value string) PairUpdater {
	u.fields[string(PairDBSchema.Value)] = value
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u PairUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u PairUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set PairQuerySet

// ===== BEGIN of Pair modifiers

// PairDBSchemaField describes database schema field. It requires for method 'Update'
type PairDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f PairDBSchemaField) String() string {
	return string(f)
}

// PairDBSchema stores db field names of Pair
var PairDBSchema = struct {
	ID        PairDBSchemaField
	CreatedAt PairDBSchemaField
	UpdatedAt PairDBSchemaField
	DeletedAt PairDBSchemaField
	Key       PairDBSchemaField
	Value     PairDBSchemaField
}{

	ID:        PairDBSchemaField("id"),
	CreatedAt: PairDBSchemaField("create_at"),
	UpdatedAt: PairDBSchemaField("update_at"),
	DeletedAt: PairDBSchemaField("deleted_at"),
	Key:       PairDBSchemaField("key"),
	Value:     PairDBSchemaField("value"),
}

// Update updates Pair fields by primary key
// nolint: dupl
func (o *Pair) Update(db *gorm.DB, fields ...PairDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"create_at":  o.CreatedAt,
		"update_at":  o.UpdatedAt,
		"deleted_at": o.DeletedAt,
		"key":        o.Key,
		"value":      o.Value,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Pair %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// PairUpdater is an Pair updates manager
type PairUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewPairUpdater creates new Pair updater
// nolint: dupl
func NewPairUpdater(db *gorm.DB) PairUpdater {
	return PairUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Pair{}),
	}
}

// ===== END of Pair modifiers

// ===== END of all query sets