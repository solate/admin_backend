// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/loginlog"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoginLogCreate is the builder for creating a LoginLog entity.
type LoginLogCreate struct {
	config
	mutation *LoginLogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (llc *LoginLogCreate) SetCreatedAt(i int64) *LoginLogCreate {
	llc.mutation.SetCreatedAt(i)
	return llc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (llc *LoginLogCreate) SetNillableCreatedAt(i *int64) *LoginLogCreate {
	if i != nil {
		llc.SetCreatedAt(*i)
	}
	return llc
}

// SetTenantCode sets the "tenant_code" field.
func (llc *LoginLogCreate) SetTenantCode(s string) *LoginLogCreate {
	llc.mutation.SetTenantCode(s)
	return llc
}

// SetLogID sets the "log_id" field.
func (llc *LoginLogCreate) SetLogID(s string) *LoginLogCreate {
	llc.mutation.SetLogID(s)
	return llc
}

// SetUserID sets the "user_id" field.
func (llc *LoginLogCreate) SetUserID(s string) *LoginLogCreate {
	llc.mutation.SetUserID(s)
	return llc
}

// SetUserName sets the "user_name" field.
func (llc *LoginLogCreate) SetUserName(s string) *LoginLogCreate {
	llc.mutation.SetUserName(s)
	return llc
}

// SetIP sets the "ip" field.
func (llc *LoginLogCreate) SetIP(s string) *LoginLogCreate {
	llc.mutation.SetIP(s)
	return llc
}

// SetMessage sets the "message" field.
func (llc *LoginLogCreate) SetMessage(s string) *LoginLogCreate {
	llc.mutation.SetMessage(s)
	return llc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (llc *LoginLogCreate) SetNillableMessage(s *string) *LoginLogCreate {
	if s != nil {
		llc.SetMessage(*s)
	}
	return llc
}

// SetUserAgent sets the "user_agent" field.
func (llc *LoginLogCreate) SetUserAgent(s string) *LoginLogCreate {
	llc.mutation.SetUserAgent(s)
	return llc
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (llc *LoginLogCreate) SetNillableUserAgent(s *string) *LoginLogCreate {
	if s != nil {
		llc.SetUserAgent(*s)
	}
	return llc
}

// SetBrowser sets the "browser" field.
func (llc *LoginLogCreate) SetBrowser(s string) *LoginLogCreate {
	llc.mutation.SetBrowser(s)
	return llc
}

// SetNillableBrowser sets the "browser" field if the given value is not nil.
func (llc *LoginLogCreate) SetNillableBrowser(s *string) *LoginLogCreate {
	if s != nil {
		llc.SetBrowser(*s)
	}
	return llc
}

// SetOs sets the "os" field.
func (llc *LoginLogCreate) SetOs(s string) *LoginLogCreate {
	llc.mutation.SetOs(s)
	return llc
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (llc *LoginLogCreate) SetNillableOs(s *string) *LoginLogCreate {
	if s != nil {
		llc.SetOs(*s)
	}
	return llc
}

// SetDevice sets the "device" field.
func (llc *LoginLogCreate) SetDevice(s string) *LoginLogCreate {
	llc.mutation.SetDevice(s)
	return llc
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (llc *LoginLogCreate) SetNillableDevice(s *string) *LoginLogCreate {
	if s != nil {
		llc.SetDevice(*s)
	}
	return llc
}

// SetLoginTime sets the "login_time" field.
func (llc *LoginLogCreate) SetLoginTime(i int64) *LoginLogCreate {
	llc.mutation.SetLoginTime(i)
	return llc
}

// SetNillableLoginTime sets the "login_time" field if the given value is not nil.
func (llc *LoginLogCreate) SetNillableLoginTime(i *int64) *LoginLogCreate {
	if i != nil {
		llc.SetLoginTime(*i)
	}
	return llc
}

// Mutation returns the LoginLogMutation object of the builder.
func (llc *LoginLogCreate) Mutation() *LoginLogMutation {
	return llc.mutation
}

// Save creates the LoginLog in the database.
func (llc *LoginLogCreate) Save(ctx context.Context) (*LoginLog, error) {
	llc.defaults()
	return withHooks(ctx, llc.sqlSave, llc.mutation, llc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (llc *LoginLogCreate) SaveX(ctx context.Context) *LoginLog {
	v, err := llc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (llc *LoginLogCreate) Exec(ctx context.Context) error {
	_, err := llc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (llc *LoginLogCreate) ExecX(ctx context.Context) {
	if err := llc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (llc *LoginLogCreate) defaults() {
	if _, ok := llc.mutation.CreatedAt(); !ok {
		v := loginlog.DefaultCreatedAt
		llc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (llc *LoginLogCreate) check() error {
	if _, ok := llc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "LoginLog.created_at"`)}
	}
	if _, ok := llc.mutation.TenantCode(); !ok {
		return &ValidationError{Name: "tenant_code", err: errors.New(`generated: missing required field "LoginLog.tenant_code"`)}
	}
	if v, ok := llc.mutation.TenantCode(); ok {
		if err := loginlog.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "LoginLog.tenant_code": %w`, err)}
		}
	}
	if _, ok := llc.mutation.LogID(); !ok {
		return &ValidationError{Name: "log_id", err: errors.New(`generated: missing required field "LoginLog.log_id"`)}
	}
	if _, ok := llc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`generated: missing required field "LoginLog.user_id"`)}
	}
	if _, ok := llc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New(`generated: missing required field "LoginLog.user_name"`)}
	}
	if v, ok := llc.mutation.UserName(); ok {
		if err := loginlog.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "user_name", err: fmt.Errorf(`generated: validator failed for field "LoginLog.user_name": %w`, err)}
		}
	}
	if _, ok := llc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`generated: missing required field "LoginLog.ip"`)}
	}
	if v, ok := llc.mutation.IP(); ok {
		if err := loginlog.IPValidator(v); err != nil {
			return &ValidationError{Name: "ip", err: fmt.Errorf(`generated: validator failed for field "LoginLog.ip": %w`, err)}
		}
	}
	return nil
}

func (llc *LoginLogCreate) sqlSave(ctx context.Context) (*LoginLog, error) {
	if err := llc.check(); err != nil {
		return nil, err
	}
	_node, _spec := llc.createSpec()
	if err := sqlgraph.CreateNode(ctx, llc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	llc.mutation.id = &_node.ID
	llc.mutation.done = true
	return _node, nil
}

func (llc *LoginLogCreate) createSpec() (*LoginLog, *sqlgraph.CreateSpec) {
	var (
		_node = &LoginLog{config: llc.config}
		_spec = sqlgraph.NewCreateSpec(loginlog.Table, sqlgraph.NewFieldSpec(loginlog.FieldID, field.TypeInt))
	)
	_spec.OnConflict = llc.conflict
	if value, ok := llc.mutation.CreatedAt(); ok {
		_spec.SetField(loginlog.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := llc.mutation.TenantCode(); ok {
		_spec.SetField(loginlog.FieldTenantCode, field.TypeString, value)
		_node.TenantCode = value
	}
	if value, ok := llc.mutation.LogID(); ok {
		_spec.SetField(loginlog.FieldLogID, field.TypeString, value)
		_node.LogID = value
	}
	if value, ok := llc.mutation.UserID(); ok {
		_spec.SetField(loginlog.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := llc.mutation.UserName(); ok {
		_spec.SetField(loginlog.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := llc.mutation.IP(); ok {
		_spec.SetField(loginlog.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := llc.mutation.Message(); ok {
		_spec.SetField(loginlog.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if value, ok := llc.mutation.UserAgent(); ok {
		_spec.SetField(loginlog.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = value
	}
	if value, ok := llc.mutation.Browser(); ok {
		_spec.SetField(loginlog.FieldBrowser, field.TypeString, value)
		_node.Browser = value
	}
	if value, ok := llc.mutation.Os(); ok {
		_spec.SetField(loginlog.FieldOs, field.TypeString, value)
		_node.Os = value
	}
	if value, ok := llc.mutation.Device(); ok {
		_spec.SetField(loginlog.FieldDevice, field.TypeString, value)
		_node.Device = value
	}
	if value, ok := llc.mutation.LoginTime(); ok {
		_spec.SetField(loginlog.FieldLoginTime, field.TypeInt64, value)
		_node.LoginTime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LoginLog.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LoginLogUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (llc *LoginLogCreate) OnConflict(opts ...sql.ConflictOption) *LoginLogUpsertOne {
	llc.conflict = opts
	return &LoginLogUpsertOne{
		create: llc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LoginLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (llc *LoginLogCreate) OnConflictColumns(columns ...string) *LoginLogUpsertOne {
	llc.conflict = append(llc.conflict, sql.ConflictColumns(columns...))
	return &LoginLogUpsertOne{
		create: llc,
	}
}

type (
	// LoginLogUpsertOne is the builder for "upsert"-ing
	//  one LoginLog node.
	LoginLogUpsertOne struct {
		create *LoginLogCreate
	}

	// LoginLogUpsert is the "OnConflict" setter.
	LoginLogUpsert struct {
		*sql.UpdateSet
	}
)

// SetTenantCode sets the "tenant_code" field.
func (u *LoginLogUpsert) SetTenantCode(v string) *LoginLogUpsert {
	u.Set(loginlog.FieldTenantCode, v)
	return u
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateTenantCode() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldTenantCode)
	return u
}

// SetUserID sets the "user_id" field.
func (u *LoginLogUpsert) SetUserID(v string) *LoginLogUpsert {
	u.Set(loginlog.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateUserID() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldUserID)
	return u
}

// SetUserName sets the "user_name" field.
func (u *LoginLogUpsert) SetUserName(v string) *LoginLogUpsert {
	u.Set(loginlog.FieldUserName, v)
	return u
}

// UpdateUserName sets the "user_name" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateUserName() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldUserName)
	return u
}

// SetIP sets the "ip" field.
func (u *LoginLogUpsert) SetIP(v string) *LoginLogUpsert {
	u.Set(loginlog.FieldIP, v)
	return u
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateIP() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldIP)
	return u
}

// SetMessage sets the "message" field.
func (u *LoginLogUpsert) SetMessage(v string) *LoginLogUpsert {
	u.Set(loginlog.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateMessage() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldMessage)
	return u
}

// ClearMessage clears the value of the "message" field.
func (u *LoginLogUpsert) ClearMessage() *LoginLogUpsert {
	u.SetNull(loginlog.FieldMessage)
	return u
}

// SetUserAgent sets the "user_agent" field.
func (u *LoginLogUpsert) SetUserAgent(v string) *LoginLogUpsert {
	u.Set(loginlog.FieldUserAgent, v)
	return u
}

// UpdateUserAgent sets the "user_agent" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateUserAgent() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldUserAgent)
	return u
}

// ClearUserAgent clears the value of the "user_agent" field.
func (u *LoginLogUpsert) ClearUserAgent() *LoginLogUpsert {
	u.SetNull(loginlog.FieldUserAgent)
	return u
}

// SetBrowser sets the "browser" field.
func (u *LoginLogUpsert) SetBrowser(v string) *LoginLogUpsert {
	u.Set(loginlog.FieldBrowser, v)
	return u
}

// UpdateBrowser sets the "browser" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateBrowser() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldBrowser)
	return u
}

// ClearBrowser clears the value of the "browser" field.
func (u *LoginLogUpsert) ClearBrowser() *LoginLogUpsert {
	u.SetNull(loginlog.FieldBrowser)
	return u
}

// SetOs sets the "os" field.
func (u *LoginLogUpsert) SetOs(v string) *LoginLogUpsert {
	u.Set(loginlog.FieldOs, v)
	return u
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateOs() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldOs)
	return u
}

// ClearOs clears the value of the "os" field.
func (u *LoginLogUpsert) ClearOs() *LoginLogUpsert {
	u.SetNull(loginlog.FieldOs)
	return u
}

// SetDevice sets the "device" field.
func (u *LoginLogUpsert) SetDevice(v string) *LoginLogUpsert {
	u.Set(loginlog.FieldDevice, v)
	return u
}

// UpdateDevice sets the "device" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateDevice() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldDevice)
	return u
}

// ClearDevice clears the value of the "device" field.
func (u *LoginLogUpsert) ClearDevice() *LoginLogUpsert {
	u.SetNull(loginlog.FieldDevice)
	return u
}

// SetLoginTime sets the "login_time" field.
func (u *LoginLogUpsert) SetLoginTime(v int64) *LoginLogUpsert {
	u.Set(loginlog.FieldLoginTime, v)
	return u
}

// UpdateLoginTime sets the "login_time" field to the value that was provided on create.
func (u *LoginLogUpsert) UpdateLoginTime() *LoginLogUpsert {
	u.SetExcluded(loginlog.FieldLoginTime)
	return u
}

// AddLoginTime adds v to the "login_time" field.
func (u *LoginLogUpsert) AddLoginTime(v int64) *LoginLogUpsert {
	u.Add(loginlog.FieldLoginTime, v)
	return u
}

// ClearLoginTime clears the value of the "login_time" field.
func (u *LoginLogUpsert) ClearLoginTime() *LoginLogUpsert {
	u.SetNull(loginlog.FieldLoginTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.LoginLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LoginLogUpsertOne) UpdateNewValues() *LoginLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(loginlog.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.LogID(); exists {
			s.SetIgnore(loginlog.FieldLogID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LoginLog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LoginLogUpsertOne) Ignore() *LoginLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LoginLogUpsertOne) DoNothing() *LoginLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LoginLogCreate.OnConflict
// documentation for more info.
func (u *LoginLogUpsertOne) Update(set func(*LoginLogUpsert)) *LoginLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LoginLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantCode sets the "tenant_code" field.
func (u *LoginLogUpsertOne) SetTenantCode(v string) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetTenantCode(v)
	})
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateTenantCode() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateTenantCode()
	})
}

// SetUserID sets the "user_id" field.
func (u *LoginLogUpsertOne) SetUserID(v string) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateUserID() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateUserID()
	})
}

// SetUserName sets the "user_name" field.
func (u *LoginLogUpsertOne) SetUserName(v string) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetUserName(v)
	})
}

// UpdateUserName sets the "user_name" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateUserName() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateUserName()
	})
}

// SetIP sets the "ip" field.
func (u *LoginLogUpsertOne) SetIP(v string) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateIP() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateIP()
	})
}

// SetMessage sets the "message" field.
func (u *LoginLogUpsertOne) SetMessage(v string) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateMessage() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *LoginLogUpsertOne) ClearMessage() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearMessage()
	})
}

// SetUserAgent sets the "user_agent" field.
func (u *LoginLogUpsertOne) SetUserAgent(v string) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetUserAgent(v)
	})
}

// UpdateUserAgent sets the "user_agent" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateUserAgent() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateUserAgent()
	})
}

// ClearUserAgent clears the value of the "user_agent" field.
func (u *LoginLogUpsertOne) ClearUserAgent() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearUserAgent()
	})
}

// SetBrowser sets the "browser" field.
func (u *LoginLogUpsertOne) SetBrowser(v string) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetBrowser(v)
	})
}

// UpdateBrowser sets the "browser" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateBrowser() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateBrowser()
	})
}

// ClearBrowser clears the value of the "browser" field.
func (u *LoginLogUpsertOne) ClearBrowser() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearBrowser()
	})
}

// SetOs sets the "os" field.
func (u *LoginLogUpsertOne) SetOs(v string) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetOs(v)
	})
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateOs() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateOs()
	})
}

// ClearOs clears the value of the "os" field.
func (u *LoginLogUpsertOne) ClearOs() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearOs()
	})
}

// SetDevice sets the "device" field.
func (u *LoginLogUpsertOne) SetDevice(v string) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetDevice(v)
	})
}

// UpdateDevice sets the "device" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateDevice() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateDevice()
	})
}

// ClearDevice clears the value of the "device" field.
func (u *LoginLogUpsertOne) ClearDevice() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearDevice()
	})
}

// SetLoginTime sets the "login_time" field.
func (u *LoginLogUpsertOne) SetLoginTime(v int64) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetLoginTime(v)
	})
}

// AddLoginTime adds v to the "login_time" field.
func (u *LoginLogUpsertOne) AddLoginTime(v int64) *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.AddLoginTime(v)
	})
}

// UpdateLoginTime sets the "login_time" field to the value that was provided on create.
func (u *LoginLogUpsertOne) UpdateLoginTime() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateLoginTime()
	})
}

// ClearLoginTime clears the value of the "login_time" field.
func (u *LoginLogUpsertOne) ClearLoginTime() *LoginLogUpsertOne {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearLoginTime()
	})
}

// Exec executes the query.
func (u *LoginLogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for LoginLogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LoginLogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LoginLogUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LoginLogUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LoginLogCreateBulk is the builder for creating many LoginLog entities in bulk.
type LoginLogCreateBulk struct {
	config
	err      error
	builders []*LoginLogCreate
	conflict []sql.ConflictOption
}

// Save creates the LoginLog entities in the database.
func (llcb *LoginLogCreateBulk) Save(ctx context.Context) ([]*LoginLog, error) {
	if llcb.err != nil {
		return nil, llcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(llcb.builders))
	nodes := make([]*LoginLog, len(llcb.builders))
	mutators := make([]Mutator, len(llcb.builders))
	for i := range llcb.builders {
		func(i int, root context.Context) {
			builder := llcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LoginLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, llcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = llcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, llcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, llcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (llcb *LoginLogCreateBulk) SaveX(ctx context.Context) []*LoginLog {
	v, err := llcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (llcb *LoginLogCreateBulk) Exec(ctx context.Context) error {
	_, err := llcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (llcb *LoginLogCreateBulk) ExecX(ctx context.Context) {
	if err := llcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LoginLog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LoginLogUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (llcb *LoginLogCreateBulk) OnConflict(opts ...sql.ConflictOption) *LoginLogUpsertBulk {
	llcb.conflict = opts
	return &LoginLogUpsertBulk{
		create: llcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LoginLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (llcb *LoginLogCreateBulk) OnConflictColumns(columns ...string) *LoginLogUpsertBulk {
	llcb.conflict = append(llcb.conflict, sql.ConflictColumns(columns...))
	return &LoginLogUpsertBulk{
		create: llcb,
	}
}

// LoginLogUpsertBulk is the builder for "upsert"-ing
// a bulk of LoginLog nodes.
type LoginLogUpsertBulk struct {
	create *LoginLogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LoginLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LoginLogUpsertBulk) UpdateNewValues() *LoginLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(loginlog.FieldCreatedAt)
			}
			if _, exists := b.mutation.LogID(); exists {
				s.SetIgnore(loginlog.FieldLogID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LoginLog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LoginLogUpsertBulk) Ignore() *LoginLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LoginLogUpsertBulk) DoNothing() *LoginLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LoginLogCreateBulk.OnConflict
// documentation for more info.
func (u *LoginLogUpsertBulk) Update(set func(*LoginLogUpsert)) *LoginLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LoginLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantCode sets the "tenant_code" field.
func (u *LoginLogUpsertBulk) SetTenantCode(v string) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetTenantCode(v)
	})
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateTenantCode() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateTenantCode()
	})
}

// SetUserID sets the "user_id" field.
func (u *LoginLogUpsertBulk) SetUserID(v string) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateUserID() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateUserID()
	})
}

// SetUserName sets the "user_name" field.
func (u *LoginLogUpsertBulk) SetUserName(v string) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetUserName(v)
	})
}

// UpdateUserName sets the "user_name" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateUserName() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateUserName()
	})
}

// SetIP sets the "ip" field.
func (u *LoginLogUpsertBulk) SetIP(v string) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetIP(v)
	})
}

// UpdateIP sets the "ip" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateIP() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateIP()
	})
}

// SetMessage sets the "message" field.
func (u *LoginLogUpsertBulk) SetMessage(v string) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateMessage() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *LoginLogUpsertBulk) ClearMessage() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearMessage()
	})
}

// SetUserAgent sets the "user_agent" field.
func (u *LoginLogUpsertBulk) SetUserAgent(v string) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetUserAgent(v)
	})
}

// UpdateUserAgent sets the "user_agent" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateUserAgent() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateUserAgent()
	})
}

// ClearUserAgent clears the value of the "user_agent" field.
func (u *LoginLogUpsertBulk) ClearUserAgent() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearUserAgent()
	})
}

// SetBrowser sets the "browser" field.
func (u *LoginLogUpsertBulk) SetBrowser(v string) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetBrowser(v)
	})
}

// UpdateBrowser sets the "browser" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateBrowser() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateBrowser()
	})
}

// ClearBrowser clears the value of the "browser" field.
func (u *LoginLogUpsertBulk) ClearBrowser() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearBrowser()
	})
}

// SetOs sets the "os" field.
func (u *LoginLogUpsertBulk) SetOs(v string) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetOs(v)
	})
}

// UpdateOs sets the "os" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateOs() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateOs()
	})
}

// ClearOs clears the value of the "os" field.
func (u *LoginLogUpsertBulk) ClearOs() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearOs()
	})
}

// SetDevice sets the "device" field.
func (u *LoginLogUpsertBulk) SetDevice(v string) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetDevice(v)
	})
}

// UpdateDevice sets the "device" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateDevice() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateDevice()
	})
}

// ClearDevice clears the value of the "device" field.
func (u *LoginLogUpsertBulk) ClearDevice() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearDevice()
	})
}

// SetLoginTime sets the "login_time" field.
func (u *LoginLogUpsertBulk) SetLoginTime(v int64) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.SetLoginTime(v)
	})
}

// AddLoginTime adds v to the "login_time" field.
func (u *LoginLogUpsertBulk) AddLoginTime(v int64) *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.AddLoginTime(v)
	})
}

// UpdateLoginTime sets the "login_time" field to the value that was provided on create.
func (u *LoginLogUpsertBulk) UpdateLoginTime() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.UpdateLoginTime()
	})
}

// ClearLoginTime clears the value of the "login_time" field.
func (u *LoginLogUpsertBulk) ClearLoginTime() *LoginLogUpsertBulk {
	return u.Update(func(s *LoginLogUpsert) {
		s.ClearLoginTime()
	})
}

// Exec executes the query.
func (u *LoginLogUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the LoginLogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for LoginLogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LoginLogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
