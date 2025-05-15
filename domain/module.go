package domain

import (
	"github.com/mysterybee07/office-project-setup/domain/user"
	"go.uber.org/fx"
)

var Module = fx.Module("domain",
	fx.Options(
		user.Module,
	),
)
