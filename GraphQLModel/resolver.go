package GraphQLModel

import (
	"SORA/Base"
	"SORA/Controller"
	"context"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) UpdateUser(ctx context.Context, token string, user Base.NewAccountUser) (*Base.CreateReturn, error) {
	return Controller.ExaminationUpdateUser(token, user)
}
func (r *mutationResolver) CreateAccount(ctx context.Context, accountIDPw Base.NewAccountIDPw, user Base.NewAccountUser) (*Base.CreateReturn, error) {
	return Controller.ExaminationCreateAccount(accountIDPw, user)
}
func (r *mutationResolver) ChangePassword(ctx context.Context, token string, oldPw string, newPw string) (*Base.CreateReturn, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddCarID(ctx context.Context, accountID string, carName string, temporarilyToken string) (*Base.CarIDReturn, error) {
	return Controller.ExaminationAddCarID(accountID, carName, temporarilyToken)
}
func (r *mutationResolver) UpdateCarName(ctx context.Context, newCarName string, carToken string) (*Base.CreateReturn, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateMonitor(ctx context.Context, inputMonitorData Base.SecurityStatus) (*Base.CreateReturn, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateSecurity(ctx context.Context, inputSecurityData Base.SecurityStatus) (*Base.CreateReturn, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddSecurity(ctx context.Context, inputSecurityData Base.SecurityStatus) (*Base.CreateReturn, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetUser(ctx context.Context, token string, getHistorysNumber int) (*Base.Users, error) {
	return Controller.ExaminationGetUser(token, getHistorysNumber)
}
func (r *queryResolver) LogIn(ctx context.Context, accountID string, password string, information Base.Logformation) (*Base.LogInToken, error) {
	return Controller.ExaminationLogIn(accountID, password, information)
}
func (r *queryResolver) LogOut(ctx context.Context, token string, information Base.Logformation) (*Base.StatusData, error) {
	return Controller.ExaminationLogOut(token, information)
}
func (r *queryResolver) CheckAccountHas(ctx context.Context, accountID string) (*Base.AccountHas, error) {
	return Controller.ExaminationCheckAccountHas(accountID)
}
func (r *queryResolver) GetCarID(ctx context.Context, token string) ([]Base.CarData, error) {
	return Controller.ExaminationGetCarID(token)
}
func (r *queryResolver) DeleteCarID(ctx context.Context, token string, carToken string) (*Base.StatusData, error) {
	panic("not implemented")
}
func (r *queryResolver) GetMonitorStatus(ctx context.Context, token string, selectObject string) (*Base.MonitorData, error) {
	panic("not implemented")
}
func (r *queryResolver) GetSecurityStatus(ctx context.Context, token string, selectObject string) (*Base.SecurityData, error) {
	panic("not implemented")
}
func (r *queryResolver) GetTemporarilyToken(ctx context.Context, token string) (*Base.TemporarilyTokenData, error) {
	return Controller.ExaminationGetTemporarilyToken(token)
}
