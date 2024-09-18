package services

type Service struct {
	Authorizer Authorizer
	JWTer      JWTer
	Timer      Timer
}
