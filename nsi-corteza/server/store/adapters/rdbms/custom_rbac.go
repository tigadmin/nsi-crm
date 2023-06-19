package rdbms

import (
    "context"
)

func (s Store) TransferRbacRules(ctx context.Context, src, dst uint64) (err error) {
	var (
		transfer = s.Dialect.GOQU().Update(rbacRuleTable).
			Set(goqu.Record{"rel_role": dst}).
			Where(goqu.Ex{"rel_role": src})
	)

	return s.Exec(ctx, transfer)
}
