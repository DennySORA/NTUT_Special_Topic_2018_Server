package GraphQLModel

import (
	"SORA/Project/Go_Back_End_SEGA_Project/Base"
	"SORA/Project/Go_Back_End_SEGA_Project/Controller"
	"context"
)

// ============================================[Finish]

func (r *queryResolver) GetUser(ctx context.Context, Certification Base.InputCertification) (Base.Users, error) {
	return Controller.ExaminationGetUser(
		Certification,
	)
}

// ============================================[Not Doing]

func (r *mutationResolver) UpdateUser(ctx context.Context, Certification Base.InputCertification, User Base.NewAccountUser) (Base.CreateReturn, error) {
	return Controller.ExaminationUpdateUser(
		Certification,
		User,
	)
}
