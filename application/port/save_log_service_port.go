package port

import "github.com/frederico-neres/my-own-log-management--litelogs/application/domain"

type SaveLogServicePort interface {
	Exec(logDomain *domain.LogDomain) error
}
