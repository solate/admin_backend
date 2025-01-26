package casbin_sdk

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entadapter "github.com/casbin/ent-adapter"
)

type CasbinSdk struct {
	DataSource string // mysql 数据库
}

func NewCasbin(dataSource string) (*casbin.Enforcer, error) {

	a, err := entadapter.NewAdapter("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	m, err := model.NewModelFromString(`

[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act || r.sub == "root"

	`)
	if err != nil {
		return nil, err
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	return e, nil
}
