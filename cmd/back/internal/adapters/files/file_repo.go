package files

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ZergsLaw/back-template/cmd/back/internal/app"
	"github.com/ZergsLaw/back-template/internal/logger"
)

var _ app.FileStore = &Client{}

const (
	headerSrcName = `src_name`
	headerUserID  = `user_id`
	bucketFile    = `user.files`
)

type (
	// Config provide connection info for message broker.
	Config struct {
		Secure       bool
		Endpoint     string
		AccessKey    string
		SecretKey    string
		SessionToken string
		Region       string
	}
	// Client provided data from and to message broker.
	Client struct {
		store *minio.Client
		m     Metrics
	}
)

// New build and returns new file store instance.
func New(ctx context.Context, reg *prometheus.Registry, namespace string, cfg Config) (*Client, error) {
	const subsystem = "file_store"
	m := NewMetrics(reg, namespace, subsystem, []string{})

	transport, err := minio.DefaultTransport(cfg.Secure)
	if err != nil {
		return nil, fmt.Errorf("minio.DefaultTransport: %w", err)
	}

	opts := &minio.Options{
		Creds:        credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, cfg.SessionToken),
		Secure:       cfg.Secure,
		Transport:    transport,
		Region:       cfg.Region,
		BucketLookup: minio.BucketLookupAuto,
	}
	client, err := minio.New(cfg.Endpoint, opts)
	if err != nil {
		return nil, fmt.Errorf("minio.New: %w, opts: %+v", err, cfg)
	}

	var lastErr error
	exist, err := client.BucketExists(ctx, bucketFile)
	for err != nil {
		logger.FromContext(ctx).Error("couldn't check bucket", slog.String(logger.Error.String(), err.Error()))
		exist, err = client.BucketExists(ctx, bucketFile)
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("client.BucketExists: %w", lastErr)
		}

		lastErr = err
	}

	err = client.MakeBucket(ctx, bucketFile, minio.MakeBucketOptions{
		Region: cfg.Region,
	})
	for !exist && err != nil {
		logger.FromContext(ctx).Error("couldn't make bucket", slog.String(logger.Error.String(), err.Error()))
		err = client.MakeBucket(ctx, bucketFile, minio.MakeBucketOptions{
			Region: cfg.Region,
		})
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("client.MakeBucket: %w", lastErr)
		}

		lastErr = err
	}

	return &Client{
		store: client,
		m:     m,
	}, nil
}

// Close implements io.Closer.
func (*Client) Close() error {
	return nil
}

var errBucketNotFound = errors.New("bucket not found")

// HealthCheck checks the connection to the file store.
func (c *Client) HealthCheck(ctx context.Context) error {
	v, err := c.store.BucketExists(ctx, bucketFile)
	if err != nil {
		return fmt.Errorf("store.BucketExists: %w", err)
	}

	if !v {
		return errBucketNotFound
	}

	return nil
}
