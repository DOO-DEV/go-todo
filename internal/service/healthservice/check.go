package healthservice

import (
	"context"
	"fmt"
)

func (s Service) HealthCheck(ctx context.Context, token string, service string) error {
	if s.healthToken != token {
		return fmt.Errorf("wrong token")
	}

	switch service {
	case "db":
		return s.checkDb(ctx)
	}

	return nil
}

func (s Service) checkDb(ctx context.Context) error {
	if err := s.db.PingContext(ctx); err != nil {
		return fmt.Errorf("can't connect to mysql db")
	}

	return nil
}
