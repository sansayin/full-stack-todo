
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}) (*{{.upperStartCamelObject}}, error) {
	data.DeleteTime = time.Unix(0,0)
	data.DelState = globalkey.DelStateNo
 	var resp {{.upperStartCamelObject}}

	{{if .withCache}}{{.keys}}
	_, err :=m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("insert into %s (%s) values ({{.expression}}) returning *", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	if session != nil{
		return nil, session.QueryRowCtx(ctx, &resp, query,{{.expressionValues}})
	}
	return nil, conn.QueryRowCtx(ctx, &resp, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}
	query := fmt.Sprintf("insert into %s (%s) values ({{.expression}}) returning", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	if session != nil{
		return session.ExecCtx(ctx,query,{{.expressionValues}})
	}
	return nil,m.conn.QueryRowCts(ctx, &resp, query, {{.expressionValues}}){{end}}

  return &resp, err
}
