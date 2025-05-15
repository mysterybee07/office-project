package user

import "go.uber.org/fx"

var Module = fx.Module("user",
	fx.Options(fx.Options(
		fx.Provide(
			NewUserRepository,
			NewUserService,
			NewUserController,
			NewUserRoute,
		),
	)),
	fx.Invoke(RegisterRoute),
)
