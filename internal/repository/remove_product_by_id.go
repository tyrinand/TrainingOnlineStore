package repository

import "context"

func (r *repository) RemoveProductById(ctx context.Context, id uint32) error {
	_, err := r.db.Exec("DELETE FROM product WHERE Id = $1", id)
	return err
}
