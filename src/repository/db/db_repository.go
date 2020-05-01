package db

import (
	"github.com/aaronbickhaus/bookstore_oauth-api/src/clients/cassandra"
	"github.com/aaronbickhaus/bookstore_oauth-api/src/domain/access_token"
	"github.com/aaronbickhaus/bookstore_oauth-api/src/utils/errors"
)

const(
	queryGetAccessToken = `SELECT access_token, client_id, user_id, expires FROM "oauth"."access_tokens" WHERE access_token = ?`
	queryCreateAccessToken = `INSERT INTO "oauth"."access_tokens"(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?)`
	queryUpdateExpirationTime = `UPDATE "oauth"."access_tokens" SET expires = ? WHERE access_token = ?`
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpriationTime(access_token.AccessToken)  *errors.RestErr
}

type dbRepository struct {}

func (r *dbRepository) UpdateExpriationTime(at access_token.AccessToken) *errors.RestErr {

		if err := cassandra.GetSession().Query(queryUpdateExpirationTime, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
		return nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) GetById(id string)  (*access_token.AccessToken, *errors.RestErr) {

	var result access_token.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
			&result.AccessToken, &result.ClientId, &result.UserId, &result.Expires); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
	}
  return &result, nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}