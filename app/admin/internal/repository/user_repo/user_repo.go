package user_repo

import (
	"context"
	"time"

	"github.com/solate/admin_backend/pkg/ent/generated"
	"github.com/solate/admin_backend/pkg/ent/generated/predicate"
	"github.com/solate/admin_backend/pkg/ent/generated/user"
)

type UserRepo struct {
	db *generated.Client
}

// NewUserRepo 创建用户仓储实例
func NewUserRepo(db *generated.Client) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *generated.User) (*generated.User, error) {
	now := time.Now().UnixMilli()
	return r.db.User.Create().
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetTenantCode(user.TenantCode).
		SetUserID(user.UserID).
		SetPhone(user.Phone).
		SetUserName(user.UserName).
		SetPwdHashed(user.PwdHashed).
		SetPwdSalt(user.PwdSalt).
		SetStatus(user.Status).
		SetNickName(user.NickName).
		SetEmail(user.Email).
		SetSex(user.Sex).
		Save(ctx)
}

func (r *UserRepo) Update(ctx context.Context, user *generated.User) (*generated.User, error) {
	now := time.Now().UnixMilli()
	user.UpdatedAt = now
	return r.db.User.UpdateOne(user).Save(ctx)
}

// func (r *UserRepo) GetByID(ctx context.Context, id int) (*generated.User, error) {
// 	return r.db.User.Query().Where(user.ID(id)).Only(ctx)
// }

func (r *UserRepo) GetByUserID(ctx context.Context, userID uint64) (*generated.User, error) {
	return r.db.User.Query().Where(user.UserID(userID)).Only(ctx)
}

func (r *UserRepo) GetByPhone(ctx context.Context, phone string) (*generated.User, error) {
	return r.db.User.Query().Where(user.Phone(phone)).Only(ctx)
}

func (r *UserRepo) GetByUserName(ctx context.Context, userName string) (*generated.User, error) {
	return r.db.User.Query().Where(user.UserName(userName)).Only(ctx)
}

func (r *UserRepo) PageList(ctx context.Context, offset, limit int, where []predicate.User) ([]*generated.User, int, error) {
	query := r.db.User.Query().Where(where...).Order(generated.Desc(user.FieldCreatedAt))

	// 查询总数
	total, err := query.Count(ctx)
	if err != nil || total == 0 {
		return nil, 0, err
	}

	// 分页查询
	list, err := query.Offset(offset).Limit(limit).All(ctx)
	return list, total, err
}

// DeleteByUserID 根据用户ID删除用户，软删除
func (r *UserRepo) DeleteByUserID(ctx context.Context, user *generated.User) error {
	now := time.Now().UnixMilli()
	user.DeletedAt = &now
	_, err := r.Update(ctx, user)
	return err
}
