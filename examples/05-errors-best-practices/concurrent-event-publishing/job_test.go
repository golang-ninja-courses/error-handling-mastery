package job

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestJob_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventStreamMock := NewMockeventStream(ctrl)
	job := &Job{eventStream: eventStreamMock}

	eventStreamMock.EXPECT().Publish(gomock.Any(), "user-id-1", "some event")
	eventStreamMock.EXPECT().Publish(gomock.Any(), "user-id-2", "another yet event")

	err := job.Handle(context.Background(), `{}`)
	require.NoError(t, err)
}
