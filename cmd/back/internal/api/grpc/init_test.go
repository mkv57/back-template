package grpc_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "github.com/ZergsLaw/back-template/api/user/v1"
	grpcapi "github.com/ZergsLaw/back-template/cmd/back/internal/api/grpc"
	"github.com/ZergsLaw/back-template/cmd/back/internal/app"
	dom "github.com/ZergsLaw/back-template/internal/dom"
	"github.com/ZergsLaw/back-template/internal/metrics"
	"github.com/ZergsLaw/back-template/internal/testhelper"
)

var (
	username = "username"
	fullName = "Elon Musk"
	email    = "email@mail.com"
	password = "password"
	userID   = uuid.Must(uuid.NewV4())

	token   = app.Token{Value: "token"}
	session = app.Session{
		ID:     uuid.Must(uuid.NewV4()),
		UserID: uuid.Must(uuid.NewV4()),
	}
	adminSession = app.Session{
		ID:     uuid.Must(uuid.NewV4()),
		UserID: uuid.Must(uuid.NewV4()),
	}
	jediSession = app.Session{
		ID:     uuid.Must(uuid.NewV4()),
		UserID: uuid.Must(uuid.NewV4()),
	}
	user = app.User{
		ID:        uuid.Must(uuid.NewV4()),
		Email:     email,
		Name:      username,
		PassHash:  []byte(password),
		AvatarID:  uuid.Must(uuid.NewV4()),
		Status:    dom.UserStatusDefault,
		FullName:  "Edgar Sipki",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	errAny = errors.New("any err")
	origin = app.Origin{
		IP:        net.ParseIP("127.0.0.1"),
		UserAgent: "grpc-go/1.68.1",
	}
)

type clients struct {
	pb.UserExternalAPIClient
}

func start(t *testing.T, userStatus dom.UserStatus) (context.Context, *clients, *Mockapplication, *require.Assertions) {
	t.Helper()
	assert := require.New(t)
	ctx := testhelper.Context(t)

	ctrl := gomock.NewController(t)
	mockApp := NewMockapplication(ctrl)

	namespace := testhelper.Namespace(t)

	reg := prometheus.NewPedanticRegistry()
	m := metrics.New(reg, namespace)

	server := grpcapi.New(ctx, m, mockApp, reg, namespace)
	addr := testhelper.Addr(t, assert)
	ln, err := net.Listen("tcp", addr.String())
	assert.NoError(err)

	go func() {
		err := server.Serve(ln)
		assert.NoError(err)
	}()

	switch userStatus { //nolint:exhaustive
	case dom.UserStatusAdmin:
		mockApp.EXPECT().Auth(gomock.Any(), token.Value).Return(&adminSession, nil).AnyTimes()
	case dom.UserStatusJedi:
		mockApp.EXPECT().Auth(gomock.Any(), token.Value).Return(&jediSession, nil).AnyTimes()
	default:
		mockApp.EXPECT().Auth(gomock.Any(), token.Value).Return(&session, nil).AnyTimes()
	}

	conn, err := grpc.NewClient(addr.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(err)

	t.Cleanup(func() {
		err := conn.Close()
		assert.NoError(err)
		assert.NoError(err)
		server.GracefulStop()
	})

	return ctx, &clients{UserExternalAPIClient: pb.NewUserExternalAPIClient(conn)}, mockApp, assert
}

func auth(ctx context.Context) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.MD{
		"authorization": {fmt.Sprintf("Bearer %s", token.Value)},
	})
}
