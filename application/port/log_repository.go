package port

import "github.com/frederico-neres/my-own-log-management--litelogs/application/domain"

type LogRepository interface {
	Save(logDomain *domain.LogDomain) error
}
