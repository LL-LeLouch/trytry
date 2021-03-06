// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"trytry/common/globalkey"
	"trytry/common/xerr"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userAuthFieldNames          = builder.RawFieldNames(&UserAuth{})
	userAuthRows                = strings.Join(userAuthFieldNames, ",")
	userAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userAuthRowsWithPlaceHolder = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheTrytryUserAuthIdPrefix              = "cache:trytry:userAuth:id:"
	cacheTrytryUserAuthAuthTypeAuthKeyPrefix = "cache:trytry:userAuth:authType:authKey:"
	cacheTrytryUserAuthUserIdAuthTypePrefix  = "cache:trytry:userAuth:userId:authType:"
)

type (
	userAuthModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserAuth, error)
		FindOneByAuthTypeAuthKey(ctx context.Context, authType string, authKey string) (*UserAuth, error)
		FindOneByUserIdAuthType(ctx context.Context, userId int64, authType string) (*UserAuth, error)
		Update(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *UserAuth) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUserAuthModel struct {
		sqlc.CachedConn
		table string
	}

	UserAuth struct {
		Id         int64     `pb:"id"`
		UserId     int64     `pb:"user_id"`
		AuthType   string    `pb:"auth_type"` // 平台类型
		AuthKey    string    `pb:"auth_key"`  // 平台唯一id
		DelState   int64     `pb:"del_state"`
		DeleteTime time.Time `pb:"delete_time"`
		CreateTime time.Time `pb:"create_time"`
		UpdateTime time.Time `pb:"update_time"`
		Version    int64     `pb:"version"` // 版本号
	}
)

func newUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserAuthModel {
	return &defaultUserAuthModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_auth`",
	}
}

func (m *defaultUserAuthModel) Insert(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	trytryUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	trytryUserAuthIdKey := fmt.Sprintf("%s%v", cacheTrytryUserAuthIdPrefix, data.Id)
	trytryUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userAuthRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.UserId, data.AuthType, data.AuthKey, data.DelState, data.DeleteTime, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.UserId, data.AuthType, data.AuthKey, data.DelState, data.DeleteTime, data.Version)
	}, trytryUserAuthAuthTypeAuthKeyKey, trytryUserAuthIdKey, trytryUserAuthUserIdAuthTypeKey)
}

func (m *defaultUserAuthModel) FindOne(ctx context.Context, id int64) (*UserAuth, error) {
	trytryUserAuthIdKey := fmt.Sprintf("%s%v", cacheTrytryUserAuthIdPrefix, id)
	var resp UserAuth
	err := m.QueryRowCtx(ctx, &resp, trytryUserAuthIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userAuthRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindOneByAuthTypeAuthKey(ctx context.Context, authType string, authKey string) (*UserAuth, error) {
	trytryUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthAuthTypeAuthKeyPrefix, authType, authKey)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, trytryUserAuthAuthTypeAuthKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `auth_type` = ? and `auth_key` = ? and del_state = ? limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, authType, authKey, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindOneByUserIdAuthType(ctx context.Context, userId int64, authType string) (*UserAuth, error) {
	trytryUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthUserIdAuthTypePrefix, userId, authType)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, trytryUserAuthUserIdAuthTypeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `auth_type` = ? and del_state = ? limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, authType, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) Update(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error) {
	trytryUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	trytryUserAuthIdKey := fmt.Sprintf("%s%v", cacheTrytryUserAuthIdPrefix, data.Id)
	trytryUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAuthRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.UserId, data.AuthType, data.AuthKey, data.DelState, data.DeleteTime, data.Version, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.UserId, data.AuthType, data.AuthKey, data.DelState, data.DeleteTime, data.Version, data.Id)
	}, trytryUserAuthAuthTypeAuthKeyKey, trytryUserAuthIdKey, trytryUserAuthUserIdAuthTypeKey)
}

func (m *defaultUserAuthModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *UserAuth) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	trytryUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	trytryUserAuthIdKey := fmt.Sprintf("%s%v", cacheTrytryUserAuthIdPrefix, data.Id)
	trytryUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, userAuthRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.UserId, data.AuthType, data.AuthKey, data.DelState, data.DeleteTime, data.Version, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.UserId, data.AuthType, data.AuthKey, data.DelState, data.DeleteTime, data.Version, data.Id, oldVersion)
	}, trytryUserAuthAuthTypeAuthKeyKey, trytryUserAuthIdKey, trytryUserAuthUserIdAuthTypeKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil
}

func (m *defaultUserAuthModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	trytryUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	trytryUserAuthIdKey := fmt.Sprintf("%s%v", cacheTrytryUserAuthIdPrefix, id)
	trytryUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheTrytryUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, trytryUserAuthAuthTypeAuthKeyKey, trytryUserAuthIdKey, trytryUserAuthUserIdAuthTypeKey)
	return err
}

func (m *defaultUserAuthModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTrytryUserAuthIdPrefix, primary)
}
func (m *defaultUserAuthModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", userAuthRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultUserAuthModel) tableName() string {
	return m.table
}
