// Code generated by goctl. DO NOT EDIT.

package friend

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	friendFieldNames          = builder.RawFieldNames(&Friend{})
	friendRows                = strings.Join(friendFieldNames, ",")
	friendRowsExpectAutoSet   = strings.Join(stringx.Remove(friendFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	friendRowsWithPlaceHolder = strings.Join(stringx.Remove(friendFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	friendModel interface {
		Insert(ctx context.Context, data *Friend) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Friend, error)
		FindByUser(ctx context.Context, userId int64) ([]*Friend, error)
		Update(ctx context.Context, data *Friend) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFriendModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Friend struct {
		Id       int64         `db:"id"`
		User1Id  sql.NullInt64 `db:"user1_id"`
		User2Id  sql.NullInt64 `db:"user2_id"`
		CreateAt time.Time     `db:"create_at"`
		UpdateAt time.Time     `db:"update_at"`
		DeleteAt sql.NullTime  `db:"delete_at"`
	}
)

func newFriendModel(conn sqlx.SqlConn) *defaultFriendModel {
	return &defaultFriendModel{
		conn:  conn,
		table: "`friend`",
	}
}

func (m *defaultFriendModel) withSession(session sqlx.Session) *defaultFriendModel {
	return &defaultFriendModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`friend`",
	}
}

func (m *defaultFriendModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultFriendModel) FindOne(ctx context.Context, id int64) (*Friend, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendRows, m.table)
	var resp Friend
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFriendModel) FindByUser(ctx context.Context, userId int64) ([]*Friend, error) {
	query := fmt.Sprintf("select %s from %s where `user1_id` = ? or `user2_id` = ? limit 1", friendRows, m.table)
	var resp []*Friend
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFriendModel) Insert(ctx context.Context, data *Friend) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, friendRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.User1Id, data.User2Id, data.DeleteAt)
	return ret, err
}

func (m *defaultFriendModel) Update(ctx context.Context, data *Friend) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, friendRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.User1Id, data.User2Id, data.DeleteAt, data.Id)
	return err
}

func (m *defaultFriendModel) tableName() string {
	return m.table
}
