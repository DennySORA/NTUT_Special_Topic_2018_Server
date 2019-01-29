package GraphQLModel

import (
	"SORA/Project/Go_Back_End_SEGA_Project/Base"
	"context"
)

// ============================================[Finish]

// ============================================[Not Doing]

func (r *mutationResolver) UpdateMonitor(ctx context.Context, InputMonitorData Base.SecurityStatus) (Base.CreateReturn, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateSecurity(ctx context.Context, InputSecurityData Base.SecurityStatus) (Base.CreateReturn, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddSecurity(ctx context.Context, InputSecurityData Base.SecurityStatus) (Base.CreateReturn, error) {
	panic("not implemented")
}

func (r *queryResolver) GetMonitorStatus(ctx context.Context, Certification Base.InputCertification, SelectObject string) (Base.MonitorData, error) {
	panic("not implemented")
}

func (r *queryResolver) GetSecurityStatus(ctx context.Context, Certification Base.InputCertification, SelectObject string) (Base.SecurityData, error) {
	panic("not implemented")
}

