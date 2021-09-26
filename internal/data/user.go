package data

import (
	"context"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/data/ent/linuseridentiy"
)

var _ biz.LinUserRepo = (*linUserRepo)(nil)

type linUserRepo struct {
	data *Data
}

func NewLinUserRepo(data *Data) biz.LinUserRepo {
	return &linUserRepo{
		data: data,
	}
}

func (r *linUserRepo) GetLinUserIdentityByIdentifier(ctx context.Context, identifier string) (*ent.LinUserIdentiy, error) {

	po, err := r.data.db.LinUserIdentiy.Query().Where(linuseridentiy.Identifier(identifier)).First(ctx)
	if err != nil {
		return nil, err
	}
	return po, nil
}
