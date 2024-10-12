package repository

import "context"

func (r *repository) RemoveProductById(ctx context.Context, id uint32) error {
	tx := r.db.MustBegin()
	tx.MustExec("DELETE FROM product WHERE Id = $1", id)
	return tx.Commit()
}
