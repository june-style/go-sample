package entities

type Repository struct {
	// tx
	DynamoDB DynamoDB
	// users
	RegisteredUser RegisteredUserRepository
	UserProfile    UserProfileRepository
	UserSession    UserSessionRepository
}
