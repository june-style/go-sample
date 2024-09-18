package entities

import (
	"strconv"
	"time"

	"github.com/june-style/go-sample/domain/derrors"
)

type RegisteredUser struct {
	accessKey string
	userID    string
	createdAt time.Time
}

type RegisteredUserSet []*RegisteredUser

func NewRegisteredUser(accessKey, userID string, createdAt time.Time) *RegisteredUser {
	return &RegisteredUser{
		accessKey: accessKey,
		userID:    userID,
		createdAt: createdAt,
	}
}

func (d *RegisteredUser) AccessKey() string {
	if d == nil {
		return ""
	}
	return d.accessKey
}

func (d *RegisteredUser) UserID() string {
	if d == nil {
		return ""
	}
	return d.userID
}

func (d *RegisteredUser) CreatedAt() time.Time {
	if d == nil {
		return time.Time{}
	}
	return d.createdAt
}

func CreateRegisteredUser() *RegisteredUser {
	return &RegisteredUser{
		accessKey: GenUUID(),
		userID:    GenXID(),
		createdAt: Now(),
	}
}

type UserProfile struct {
	userID    string
	name      string
	createdAt time.Time
}

type UserProfileSet []*UserProfile

func NewUserProfile(userID, name string, createdAt time.Time) *UserProfile {
	return &UserProfile{
		userID:    userID,
		name:      name,
		createdAt: createdAt,
	}
}

func (u *UserProfile) UserID() string {
	if u == nil {
		return ""
	}
	return u.userID
}

func (u *UserProfile) Name() string {
	if u == nil {
		return ""
	}
	return u.name
}

func (u *UserProfile) CreatedAt() time.Time {
	if u == nil {
		return time.Time{}
	}
	return u.createdAt
}

type UserSession struct {
	userID    string
	sessionID string
}

type UserSessionSet []*UserSession

func NewUserSession(userID, sessionID string) *UserSession {
	return &UserSession{
		userID:    userID,
		sessionID: sessionID,
	}
}

func (u *UserSession) UserID() string {
	if u == nil {
		return ""
	}
	return u.userID
}

func (u *UserSession) SessionID() string {
	if u == nil {
		return ""
	}
	return u.sessionID
}

func CreateUserSession(userID, secretSalt string) (*UserSession, error) {
	if userID == "" {
		return nil, derrors.Wrapf(ErrParameterDoesNotExists, "user-id is empty")
	}
	if secretSalt == "" {
		return nil, derrors.Wrapf(ErrParameterDoesNotExists, "secret-salt is empty")
	}
	return &UserSession{
		userID: userID,
		sessionID: GenSHA1(
			userID,
			derrors.Caller(1).Func(),
			strconv.Itoa(int(Now().UnixNano())),
			secretSalt,
		),
	}, nil
}
