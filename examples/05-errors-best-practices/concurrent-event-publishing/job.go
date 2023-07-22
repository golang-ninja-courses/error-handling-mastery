package job

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

//go:generate mockgen -source=$GOFILE -destination=job_mocks_test.go -package=job

type eventStream interface {
	Publish(ctx context.Context, userID string, event string) error
}

type Job struct {
	eventStream eventStream
}

func (j *Job) Handle(ctx context.Context, payload string) error {
	_, err := parsePayload(payload)
	if err != nil {
		return fmt.Errorf("parse payload: %v", err)
	}

	wg, ctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		err = j.eventStream.Publish(ctx, "user-id-1", "some event")
		if err != nil {
			return fmt.Errorf("publish first event: %v", err)
		}
		return nil
	})

	wg.Go(func() error {
		err = j.eventStream.Publish(ctx, "user-id-2", "another yet event")
		if err != nil {
			return fmt.Errorf("publish second event: %v", err)
		}
		return nil
	})

	return wg.Wait()
}

func parsePayload(_ string) (any, error) {
	return nil, nil
}
