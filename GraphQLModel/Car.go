package GraphQLModel

import (
	"SORA/Project/Go_Back_End_SEGA_Project/Base"
	"SORA/Project/Go_Back_End_SEGA_Project/Controller"
	"context"
)

// ============================================[Finish]

func (r *queryResolver) GetCarID(ctx context.Context, Certification Base.InputCertification) ([]Base.CarData, error) {
	return Controller.ExaminationGetCarID(
		Certification,
	)
}

func (r *mutationResolver) AddCarID(ctx context.Context, InputCarNews Base.CarNews) (Base.CarIDReturn, error) {
	return Controller.ExaminationAddCarID(
		InputCarNews,
	)
}

func (r *queryResolver) GetTemporarilyToken(ctx context.Context, Certification Base.InputCertification) (Base.TemporarilyTokenData, error) {
	return Controller.ExaminationGetTemporarilyToken(
		Certification,
	)
}

// ============================================[Not Doing]

func (r *queryResolver) DeleteCarID(ctx context.Context, Certification Base.InputCertification, CarID string) (Base.StatusData, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateCarName(ctx context.Context, Certification Base.InputCertification, CarNameData Base.NewCarName) (Base.CreateReturn, error) {
	panic("not implemented")
}
