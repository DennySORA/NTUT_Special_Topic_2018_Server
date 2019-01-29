package GraphQLModel

import (
	"SORA/Project/Go_Back_End_SEGA_Project/Base"
	"SORA/Project/Go_Back_End_SEGA_Project/Controller"
	"context"
)

// ============================================[Finish]

func (r *mutationResolver) CreateAccount(ctx context.Context, AccountIDPW Base.NewAccountIDPW, User Base.NewAccountUser) (Base.CreateReturn, error) {
	return Controller.ExaminationCreateAccount(
		AccountIDPW,
		User,
	)
}

func (r *queryResolver) LogIn(ctx context.Context, ID string, Password string) (Base.LogInToken, error) {
	return Controller.ExaminationLogIn(
		ID,
		Password,
	)
}

func (r *queryResolver) LogOut(ctx context.Context, Certification Base.InputCertification) (Base.StatusData, error) {
	return Controller.ExaminationLogOut(Certification)
}

func (r *queryResolver) CheckAccountHas(ctx context.Context, ID string) (Base.AccountHas, error) {
	return Controller.ExaminationCheckAccountHas(
		ID,
	)
}

// ============================================[Not Doing]

func (r *mutationResolver) ChangePassword(ctx context.Context, Certification Base.InputCertification, OldPW Base.AccountPW, NewPW Base.AccountPW, ConfirmationPW Base.AccountPW) (Base.CreateReturn, error) {
	panic("not implemented")
}
