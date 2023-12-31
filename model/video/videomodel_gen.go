// Code generated by goctl. DO NOT EDIT.

package video

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
	videoFieldNames          = builder.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	videoModel interface {
		Insert(ctx context.Context, data *Video) (sql.Result, error)
		TransactionInsert(ctx context.Context, session sqlx.Session, data *Video) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Video, error)
		FindAll(ctx context.Context) ([]*Video, error)
		FindAllByUser(ctx context.Context, userId int64) ([]*Video, error)
		Update(ctx context.Context, data *Video) error
		Delete(ctx context.Context, id int64) error
	}

	defaultVideoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Video struct {
		Id            int64        `db:"id"`
		Author        int64       `db:"author"`         // The author
		PlayUrl       string       `db:"play_url"`       // The play_url
		CoverUrl      string       `db:"cover_url"`      // The cover_url
		Title         string       `db:"title"`          // The title
		IsFavorite    string       `db:"is_favorite"`    // The is_favorite
		FavoriteCount int64        `db:"favorite_count"` // The favorite_count
		CommentCount  int64        `db:"comment_count"`  // The comment_count
		CreateAt      time.Time    `db:"create_at"`
		UpdateAt      time.Time    `db:"update_at"`
		DeleteAt      sql.NullTime `db:"delete_at"`
	}
)

func newVideoModel(conn sqlx.SqlConn) *defaultVideoModel {
	return &defaultVideoModel{
		conn:  conn,
		table: "`video`",
	}
}

func (m *defaultVideoModel) withSession(session sqlx.Session) *defaultVideoModel {
	return &defaultVideoModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`video`",
	}
}

func (m *defaultVideoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultVideoModel) FindOne(ctx context.Context, id int64) (*Video, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	var resp Video
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

func (m *defaultVideoModel) FindAll(ctx context.Context) ([]*Video, error) {
	query := fmt.Sprintf("select %s from %s", videoRows, m.table)
	var resp []*Video
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) FindAllByUser(ctx context.Context, userId int64) ([]*Video, error) {
	query := fmt.Sprintf("select %s from %s where `author` = ?", videoRows, m.table)
	var resp []*Video
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) Insert(ctx context.Context, data *Video) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Author, data.PlayUrl, data.CoverUrl, data.Title, data.IsFavorite, data.FavoriteCount, data.CommentCount, data.DeleteAt)
	return ret, err
}


func (m *defaultVideoModel) TransactionInsert(ctx context.Context, session sqlx.Session, data *Video) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.Author, data.PlayUrl, data.CoverUrl, data.Title, data.IsFavorite, data.FavoriteCount, data.CommentCount, data.DeleteAt)
	return ret, err
}

func (m *defaultVideoModel) Update(ctx context.Context, data *Video) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Author, data.PlayUrl, data.CoverUrl, data.Title, data.IsFavorite, data.FavoriteCount, data.CommentCount, data.DeleteAt, data.Id)
	return err
}

func (m *defaultVideoModel) tableName() string {
	return m.table
}
