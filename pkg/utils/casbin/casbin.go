package casbin_sdk

import (
	"sync"

	"entgo.io/ent/dialect"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entadapter "github.com/casbin/ent-adapter"
)

var (
	enforcer *casbin.Enforcer
	once     sync.Once
	initErr  error
)

func NewCasbin(dataSource string) (*casbin.Enforcer, error) {
	once.Do(func() {
		initCasbin(dataSource)
	})

	return enforcer, initErr
}
func initCasbin(dataSource string) {
	a, err := entadapter.NewAdapter(dialect.Postgres, dataSource)
	if err != nil {
		initErr = err
		return
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
		initErr = err
		return
	}

	enforcer, err = casbin.NewEnforcer(m, a)
	if err != nil {
		initErr = err
	}
}
