package middleware

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/gogo/status"
	"google.golang.org/grpc"
)

func UnaryClientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	err := invoker(ctx, method, req, reply, cc, opts...)

	st := status.Convert(err)
	var reconstituted error
	for _, det := range st.Details() {
		if t, ok := det.(*errors.EncodedError); ok {
			reconstituted = errors.DecodeError(ctx, *t)
		}
	}

	if reconstituted != nil {
		err = reconstituted
	}

	return err
}
