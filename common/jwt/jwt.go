package jwt

import (
	"github.com/gogf/gf/errors/gerror"

	"github.com/gogf/gf/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"github.com/micro-mesh/common/model"
)

type MemberClaims struct {
	*jwt.StandardClaims
	Member *model.ContextMember
}

type UserClaims struct {
	*jwt.StandardClaims
	User *model.ContextUser
}

var jwtString = "afqeoasdfg_PL345t2345sfmwprqwoergm2p5409gdfgwergwrifkmqpoeifmqp39ef"

func GeneMemberToken(contextMember *model.ContextMember) (string, error) {
	claims := &MemberClaims{
		Member: contextMember,
		StandardClaims: &jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			//ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
			Subject: "api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtString)
	if err != nil {
		return "", gerror.Wrap(err, "")
	}
	return tokenString, nil
}

func ParseMemberToken(tokenString string) (*model.ContextMember, error) {
	claims := &MemberClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return gconv.Bytes(jwtString), nil
	})

	if err == nil && token.Valid {
		return claims.Member, nil
	}

	return nil, gerror.Wrap(err, "")

}

func GeneUserToken(contextUser *model.ContextUser) (string, error) {
	claims := &UserClaims{
		User: contextUser,
		StandardClaims: &jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			// ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
			Subject: "backend",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(gconv.Bytes(jwtString))
	if err != nil {
		return "", gerror.Wrap(err, "产生token失败")
	}
	return tokenString, nil
}

func ParseUserToken(tokenString string) (*model.ContextUser, error) {
	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return gconv.Bytes(jwtString), nil
	})

	if err == nil && token.Valid {
		return claims.User, nil
	}

	return nil, gerror.Wrap(err, "")
}
