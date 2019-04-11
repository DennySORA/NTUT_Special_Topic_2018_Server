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

func (r *mutationResolver) UpdateUser(ctx context.Context, certification Base.InputCertification, user Base.NewAccountUser) (*Base.CreateReturn, error) {
	return Controller.ExaminationUpdateUser(certification, user)
}
func (r *mutationResolver) CreateAccount(ctx context.Context, accountIDPw Base.NewAccountIDPw, user Base.NewAccountUser) (*Base.CreateReturn, error) {
	return Controller.ExaminationCreateAccount(accountIDPw, user)
}
func (r *mutationResolver) ChangePassword(ctx context.Context, certification Base.InputCertification, oldPw Base.AccountPw, newPw Base.AccountPw, confirmationPw Base.AccountPw) (*Base.CreateReturn, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddCarID(ctx context.Context, inputCarNews Base.CarNews) (*Base.CarIDReturn, error) {
	return Controller.ExaminationAddCarID(inputCarNews)
}
func (r *mutationResolver) UpdateCarName(ctx context.Context, certification Base.InputCertification, carNameData Base.NewCarName) (*Base.CreateReturn, error) {
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

func (r *queryResolver) GetUser(ctx context.Context, certification Base.InputCertification, GetHistorysNumber int) (*Base.Users, error) {
	return Controller.ExaminationGetUser(certification, GetHistorysNumber)
}
func (r *queryResolver) LogIn(ctx context.Context, id string, password string, information Base.Logformation) (*Base.LogInToken, error) {
	return Controller.ExaminationLogIn(id, password, information)
}
func (r *queryResolver) LogOut(ctx context.Context, certification Base.InputCertification, information Base.Logformation) (*Base.StatusData, error) {
	return Controller.ExaminationLogOut(certification, information)
}
func (r *queryResolver) CheckAccountHas(ctx context.Context, id string) (*Base.AccountHas, error) {
	return Controller.ExaminationCheckAccountHas(id)
}
func (r *queryResolver) GetCarID(ctx context.Context, certification Base.InputCertification) ([]Base.CarData, error) {
	return Controller.ExaminationGetCarID(certification)
}
func (r *queryResolver) DeleteCarID(ctx context.Context, certification Base.InputCertification, carID string) (*Base.StatusData, error) {
	panic("not implemented")
}
func (r *queryResolver) GetMonitorStatus(ctx context.Context, certification Base.InputCertification, selectObject string) (*Base.MonitorData, error) {
	panic("not implemented")
}
func (r *queryResolver) GetSecurityStatus(ctx context.Context, certification Base.InputCertification, selectObject string) (*Base.SecurityData, error) {
	panic("not implemented")
}
func (r *queryResolver) GetTemporarilyToken(ctx context.Context, certification Base.InputCertification) (*Base.TemporarilyTokenData, error) {
	return Controller.ExaminationGetTemporarilyToken(certification)
}
