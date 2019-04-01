package models

// WithoutSoftDelete :
func (qs PairQuerySet) WithoutSoftDelete() PairQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}
