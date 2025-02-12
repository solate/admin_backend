package userrepo

import (
	"context"
	"time"

	"admin_backend/pkg/common"
	"admin_backend/pkg/common/contextutil"
	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/ent/generated/predicate"
	"admin_backend/pkg/ent/generated/user"
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

func (r *UserRepo) Update(ctx context.Context, update *generated.User) (int, error) {
	now := time.Now().UnixMilli()
	update.UpdatedAt = now
	return r.db.User.Update().
		SetUpdatedAt(now).
		SetPhone(update.Phone).
		SetUserName(update.UserName).
		SetNickName(update.NickName).
		SetEmail(update.Email).
		SetSex(update.Sex).
		SetStatus(update.Status).
		Where(user.UserID(update.UserID)).Save(ctx)
}

// func (r *UserRepo) GetByID(ctx context.Context, id int) (*generated.User, error) {
// 	return r.db.User.Query().Where(user.ID(id)).Only(ctx)
// }

func (r *UserRepo) GetByUserID(ctx context.Context, userID uint64) (*generated.User, error) {
	return r.Get(ctx, []predicate.User{user.UserID(userID)})
}

func (r *UserRepo) GetByPhone(ctx context.Context, phone string) (*generated.User, error) {
	return r.Get(ctx, []predicate.User{user.Phone(phone)})
}

func (r *UserRepo) GetByUserName(ctx context.Context, userName string) (*generated.User, error) {
	return r.Get(ctx, []predicate.User{user.UserName(userName)})
}

// defaultQuery 默认查询条件
func (r *UserRepo) defaultQuery(ctx context.Context, where []predicate.User) []predicate.User {
	where = append(where, user.DeletedAtIsNil())
	where = append(where, user.TenantCode(contextutil.GetTenantCodeFromCtx(ctx)))
	return where
}

func (r *UserRepo) Get(ctx context.Context, where []predicate.User) (*generated.User, error) {
	where = r.defaultQuery(ctx, where)
	return r.db.User.Query().Where(where...).Only(ctx)
}

func (r *UserRepo) PageList(ctx context.Context, current, limit int, where []predicate.User) ([]*generated.User, int, error) {

	where = r.defaultQuery(ctx, where)

	offset := common.Offset(current, limit)
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
func (r *UserRepo) DeleteByUserID(ctx context.Context, delete *generated.User) (int, error) {
	now := time.Now().UnixMilli()
	delete.DeletedAt = &now
	return r.db.User.Update().
		SetDeletedAt(now).
		Where(user.UserID(delete.UserID)).Save(ctx)
}

// 用户登出
func (r *UserRepo) Logout(ctx context.Context, userID uint64) (int, error) {
	return r.db.User.Update().
		SetToken("").
		Where(user.UserID(userID)).Save(ctx)
}

// 用户登录
func (r *UserRepo) UpdateToken(ctx context.Context, userID uint64, token string) (int, error) {
	return r.db.User.Update().
		SetToken(token).
		SetUpdatedAt(time.Now().UnixMilli()).
		Where(user.UserID(userID)).Save(ctx)
}
