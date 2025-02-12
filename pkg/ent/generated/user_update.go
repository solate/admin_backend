// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/predicate"
	"admin_backend/pkg/ent/generated/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(i int64) *UserUpdate {
	uu.mutation.ResetUpdatedAt()
	uu.mutation.SetUpdatedAt(i)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(i *int64) *UserUpdate {
	if i != nil {
		uu.SetUpdatedAt(*i)
	}
	return uu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (uu *UserUpdate) AddUpdatedAt(i int64) *UserUpdate {
	uu.mutation.AddUpdatedAt(i)
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(i int64) *UserUpdate {
	uu.mutation.ResetDeletedAt()
	uu.mutation.SetDeletedAt(i)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(i *int64) *UserUpdate {
	if i != nil {
		uu.SetDeletedAt(*i)
	}
	return uu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (uu *UserUpdate) AddDeletedAt(i int64) *UserUpdate {
	uu.mutation.AddDeletedAt(i)
	return uu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uu *UserUpdate) ClearDeletedAt() *UserUpdate {
	uu.mutation.ClearDeletedAt()
	return uu
}

// SetTenantCode sets the "tenant_code" field.
func (uu *UserUpdate) SetTenantCode(s string) *UserUpdate {
	uu.mutation.SetTenantCode(s)
	return uu
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTenantCode(s *string) *UserUpdate {
	if s != nil {
		uu.SetTenantCode(*s)
	}
	return uu
}

// SetUserID sets the "user_id" field.
func (uu *UserUpdate) SetUserID(s string) *UserUpdate {
	uu.mutation.SetUserID(s)
	return uu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserID(s *string) *UserUpdate {
	if s != nil {
		uu.SetUserID(*s)
	}
	return uu
}

// SetUserName sets the "user_name" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserName(s *string) *UserUpdate {
	if s != nil {
		uu.SetUserName(*s)
	}
	return uu
}

// SetPwdHashed sets the "pwd_hashed" field.
func (uu *UserUpdate) SetPwdHashed(s string) *UserUpdate {
	uu.mutation.SetPwdHashed(s)
	return uu
}

// SetNillablePwdHashed sets the "pwd_hashed" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePwdHashed(s *string) *UserUpdate {
	if s != nil {
		uu.SetPwdHashed(*s)
	}
	return uu
}

// SetPwdSalt sets the "pwd_salt" field.
func (uu *UserUpdate) SetPwdSalt(s string) *UserUpdate {
	uu.mutation.SetPwdSalt(s)
	return uu
}

// SetNillablePwdSalt sets the "pwd_salt" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePwdSalt(s *string) *UserUpdate {
	if s != nil {
		uu.SetPwdSalt(*s)
	}
	return uu
}

// SetToken sets the "token" field.
func (uu *UserUpdate) SetToken(s string) *UserUpdate {
	uu.mutation.SetToken(s)
	return uu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (uu *UserUpdate) SetNillableToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetToken(*s)
	}
	return uu
}

// SetNickName sets the "nick_name" field.
func (uu *UserUpdate) SetNickName(s string) *UserUpdate {
	uu.mutation.SetNickName(s)
	return uu
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickName(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickName(*s)
	}
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetSex sets the "sex" field.
func (uu *UserUpdate) SetSex(i int) *UserUpdate {
	uu.mutation.ResetSex()
	uu.mutation.SetSex(i)
	return uu
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSex(i *int) *UserUpdate {
	if i != nil {
		uu.SetSex(*i)
	}
	return uu
}

// AddSex adds i to the "sex" field.
func (uu *UserUpdate) AddSex(i int) *UserUpdate {
	uu.mutation.AddSex(i)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(i int) *UserUpdate {
	uu.mutation.ResetStatus()
	uu.mutation.SetStatus(i)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(i *int) *UserUpdate {
	if i != nil {
		uu.SetStatus(*i)
	}
	return uu
}

// AddStatus adds i to the "status" field.
func (uu *UserUpdate) AddStatus(i int) *UserUpdate {
	uu.mutation.AddStatus(i)
	return uu
}

// SetRoleID sets the "role_id" field.
func (uu *UserUpdate) SetRoleID(u uint64) *UserUpdate {
	uu.mutation.ResetRoleID()
	uu.mutation.SetRoleID(u)
	return uu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRoleID(u *uint64) *UserUpdate {
	if u != nil {
		uu.SetRoleID(*u)
	}
	return uu
}

// AddRoleID adds u to the "role_id" field.
func (uu *UserUpdate) AddRoleID(u int64) *UserUpdate {
	uu.mutation.AddRoleID(u)
	return uu
}

// SetDeptID sets the "dept_id" field.
func (uu *UserUpdate) SetDeptID(u uint64) *UserUpdate {
	uu.mutation.ResetDeptID()
	uu.mutation.SetDeptID(u)
	return uu
}

// SetNillableDeptID sets the "dept_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeptID(u *uint64) *UserUpdate {
	if u != nil {
		uu.SetDeptID(*u)
	}
	return uu
}

// AddDeptID adds u to the "dept_id" field.
func (uu *UserUpdate) AddDeptID(u int64) *UserUpdate {
	uu.mutation.AddDeptID(u)
	return uu
}

// SetPostID sets the "post_id" field.
func (uu *UserUpdate) SetPostID(u uint64) *UserUpdate {
	uu.mutation.ResetPostID()
	uu.mutation.SetPostID(u)
	return uu
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePostID(u *uint64) *UserUpdate {
	if u != nil {
		uu.SetPostID(*u)
	}
	return uu
}

// AddPostID adds u to the "post_id" field.
func (uu *UserUpdate) AddPostID(u int64) *UserUpdate {
	uu.mutation.AddPostID(u)
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`generated: validator failed for field "User.user_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PwdHashed(); ok {
		if err := user.PwdHashedValidator(v); err != nil {
			return &ValidationError{Name: "pwd_hashed", err: fmt.Errorf(`generated: validator failed for field "User.pwd_hashed": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PwdSalt(); ok {
		if err := user.PwdSaltValidator(v); err != nil {
			return &ValidationError{Name: "pwd_salt", err: fmt.Errorf(`generated: validator failed for field "User.pwd_salt": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`generated: validator failed for field "User.phone": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uu *UserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdate {
	uu.modifiers = append(uu.modifiers, modifiers...)
	return uu
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(user.FieldDeletedAt, field.TypeInt64, value)
	}
	if uu.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeInt64)
	}
	if value, ok := uu.mutation.TenantCode(); ok {
		_spec.SetField(user.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := uu.mutation.UserID(); ok {
		_spec.SetField(user.FieldUserID, field.TypeString, value)
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uu.mutation.PwdHashed(); ok {
		_spec.SetField(user.FieldPwdHashed, field.TypeString, value)
	}
	if value, ok := uu.mutation.PwdSalt(); ok {
		_spec.SetField(user.FieldPwdSalt, field.TypeString, value)
	}
	if value, ok := uu.mutation.Token(); ok {
		_spec.SetField(user.FieldToken, field.TypeString, value)
	}
	if value, ok := uu.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Sex(); ok {
		_spec.SetField(user.FieldSex, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedSex(); ok {
		_spec.AddField(user.FieldSex, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt, value)
	}
	if value, ok := uu.mutation.RoleID(); ok {
		_spec.SetField(user.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := uu.mutation.AddedRoleID(); ok {
		_spec.AddField(user.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := uu.mutation.DeptID(); ok {
		_spec.SetField(user.FieldDeptID, field.TypeUint64, value)
	}
	if value, ok := uu.mutation.AddedDeptID(); ok {
		_spec.AddField(user.FieldDeptID, field.TypeUint64, value)
	}
	if value, ok := uu.mutation.PostID(); ok {
		_spec.SetField(user.FieldPostID, field.TypeUint64, value)
	}
	if value, ok := uu.mutation.AddedPostID(); ok {
		_spec.AddField(user.FieldPostID, field.TypeUint64, value)
	}
	_spec.AddModifiers(uu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(i int64) *UserUpdateOne {
	uuo.mutation.ResetUpdatedAt()
	uuo.mutation.SetUpdatedAt(i)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetUpdatedAt(*i)
	}
	return uuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (uuo *UserUpdateOne) AddUpdatedAt(i int64) *UserUpdateOne {
	uuo.mutation.AddUpdatedAt(i)
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(i int64) *UserUpdateOne {
	uuo.mutation.ResetDeletedAt()
	uuo.mutation.SetDeletedAt(i)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetDeletedAt(*i)
	}
	return uuo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (uuo *UserUpdateOne) AddDeletedAt(i int64) *UserUpdateOne {
	uuo.mutation.AddDeletedAt(i)
	return uuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uuo *UserUpdateOne) ClearDeletedAt() *UserUpdateOne {
	uuo.mutation.ClearDeletedAt()
	return uuo
}

// SetTenantCode sets the "tenant_code" field.
func (uuo *UserUpdateOne) SetTenantCode(s string) *UserUpdateOne {
	uuo.mutation.SetTenantCode(s)
	return uuo
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTenantCode(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTenantCode(*s)
	}
	return uuo
}

// SetUserID sets the "user_id" field.
func (uuo *UserUpdateOne) SetUserID(s string) *UserUpdateOne {
	uuo.mutation.SetUserID(s)
	return uuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUserID(*s)
	}
	return uuo
}

// SetUserName sets the "user_name" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUserName(*s)
	}
	return uuo
}

// SetPwdHashed sets the "pwd_hashed" field.
func (uuo *UserUpdateOne) SetPwdHashed(s string) *UserUpdateOne {
	uuo.mutation.SetPwdHashed(s)
	return uuo
}

// SetNillablePwdHashed sets the "pwd_hashed" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePwdHashed(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPwdHashed(*s)
	}
	return uuo
}

// SetPwdSalt sets the "pwd_salt" field.
func (uuo *UserUpdateOne) SetPwdSalt(s string) *UserUpdateOne {
	uuo.mutation.SetPwdSalt(s)
	return uuo
}

// SetNillablePwdSalt sets the "pwd_salt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePwdSalt(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPwdSalt(*s)
	}
	return uuo
}

// SetToken sets the "token" field.
func (uuo *UserUpdateOne) SetToken(s string) *UserUpdateOne {
	uuo.mutation.SetToken(s)
	return uuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetToken(*s)
	}
	return uuo
}

// SetNickName sets the "nick_name" field.
func (uuo *UserUpdateOne) SetNickName(s string) *UserUpdateOne {
	uuo.mutation.SetNickName(s)
	return uuo
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickName(*s)
	}
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetSex sets the "sex" field.
func (uuo *UserUpdateOne) SetSex(i int) *UserUpdateOne {
	uuo.mutation.ResetSex()
	uuo.mutation.SetSex(i)
	return uuo
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSex(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetSex(*i)
	}
	return uuo
}

// AddSex adds i to the "sex" field.
func (uuo *UserUpdateOne) AddSex(i int) *UserUpdateOne {
	uuo.mutation.AddSex(i)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(i int) *UserUpdateOne {
	uuo.mutation.ResetStatus()
	uuo.mutation.SetStatus(i)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetStatus(*i)
	}
	return uuo
}

// AddStatus adds i to the "status" field.
func (uuo *UserUpdateOne) AddStatus(i int) *UserUpdateOne {
	uuo.mutation.AddStatus(i)
	return uuo
}

// SetRoleID sets the "role_id" field.
func (uuo *UserUpdateOne) SetRoleID(u uint64) *UserUpdateOne {
	uuo.mutation.ResetRoleID()
	uuo.mutation.SetRoleID(u)
	return uuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRoleID(u *uint64) *UserUpdateOne {
	if u != nil {
		uuo.SetRoleID(*u)
	}
	return uuo
}

// AddRoleID adds u to the "role_id" field.
func (uuo *UserUpdateOne) AddRoleID(u int64) *UserUpdateOne {
	uuo.mutation.AddRoleID(u)
	return uuo
}

// SetDeptID sets the "dept_id" field.
func (uuo *UserUpdateOne) SetDeptID(u uint64) *UserUpdateOne {
	uuo.mutation.ResetDeptID()
	uuo.mutation.SetDeptID(u)
	return uuo
}

// SetNillableDeptID sets the "dept_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeptID(u *uint64) *UserUpdateOne {
	if u != nil {
		uuo.SetDeptID(*u)
	}
	return uuo
}

// AddDeptID adds u to the "dept_id" field.
func (uuo *UserUpdateOne) AddDeptID(u int64) *UserUpdateOne {
	uuo.mutation.AddDeptID(u)
	return uuo
}

// SetPostID sets the "post_id" field.
func (uuo *UserUpdateOne) SetPostID(u uint64) *UserUpdateOne {
	uuo.mutation.ResetPostID()
	uuo.mutation.SetPostID(u)
	return uuo
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePostID(u *uint64) *UserUpdateOne {
	if u != nil {
		uuo.SetPostID(*u)
	}
	return uuo
}

// AddPostID adds u to the "post_id" field.
func (uuo *UserUpdateOne) AddPostID(u int64) *UserUpdateOne {
	uuo.mutation.AddPostID(u)
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`generated: validator failed for field "User.user_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PwdHashed(); ok {
		if err := user.PwdHashedValidator(v); err != nil {
			return &ValidationError{Name: "pwd_hashed", err: fmt.Errorf(`generated: validator failed for field "User.pwd_hashed": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PwdSalt(); ok {
		if err := user.PwdSaltValidator(v); err != nil {
			return &ValidationError{Name: "pwd_salt", err: fmt.Errorf(`generated: validator failed for field "User.pwd_salt": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`generated: validator failed for field "User.phone": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uuo *UserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdateOne {
	uuo.modifiers = append(uuo.modifiers, modifiers...)
	return uuo
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(user.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(user.FieldDeletedAt, field.TypeInt64, value)
	}
	if uuo.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeInt64)
	}
	if value, ok := uuo.mutation.TenantCode(); ok {
		_spec.SetField(user.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UserID(); ok {
		_spec.SetField(user.FieldUserID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PwdHashed(); ok {
		_spec.SetField(user.FieldPwdHashed, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PwdSalt(); ok {
		_spec.SetField(user.FieldPwdSalt, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Token(); ok {
		_spec.SetField(user.FieldToken, field.TypeString, value)
	}
	if value, ok := uuo.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Sex(); ok {
		_spec.SetField(user.FieldSex, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedSex(); ok {
		_spec.AddField(user.FieldSex, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.RoleID(); ok {
		_spec.SetField(user.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := uuo.mutation.AddedRoleID(); ok {
		_spec.AddField(user.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := uuo.mutation.DeptID(); ok {
		_spec.SetField(user.FieldDeptID, field.TypeUint64, value)
	}
	if value, ok := uuo.mutation.AddedDeptID(); ok {
		_spec.AddField(user.FieldDeptID, field.TypeUint64, value)
	}
	if value, ok := uuo.mutation.PostID(); ok {
		_spec.SetField(user.FieldPostID, field.TypeUint64, value)
	}
	if value, ok := uuo.mutation.AddedPostID(); ok {
		_spec.AddField(user.FieldPostID, field.TypeUint64, value)
	}
	_spec.AddModifiers(uuo.modifiers...)
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
