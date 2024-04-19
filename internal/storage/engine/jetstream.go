package engine

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

func NewJetStream(kvm jetstream.KeyValueManager) Engine {
	return &js{
		kvm: kvm,
		kvs: make(map[string]jetstream.KeyValue),
	}
}

type js struct {
	m sync.Mutex

	kvm jetstream.KeyValueManager
	kvs map[string]jetstream.KeyValue
}

func (js *js) Set(ctx context.Context, key string, data []byte) error {
	bucket, key := js.getBucketAndKey(key)
	kv, err := js.getBucket(ctx, bucket)
	if err != nil {
		return err
	}

	zap.L().Info("set", zap.String("bucket", bucket), zap.String("key", key))

	_, err = kv.Create(ctx, key, data)
	if errors.Is(err, jetstream.ErrKeyExists) {
		return ErrAlreadyExists
	}

	return err
}

func (js *js) Get(ctx context.Context, key string) ([]byte, error) {
	data, _, err := js.get(ctx, key)
	if errors.Is(err, jetstream.ErrKeyNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return data, nil
}

func (js *js) get(ctx context.Context, key string) ([]byte, uint64, error) {
	bucket, key := js.getBucketAndKey(key)
	kv, err := js.getBucket(ctx, bucket)
	if err != nil {
		return nil, 0, err
	}

	entry, err := kv.Get(ctx, key)
	if err != nil {
		return nil, 0, err
	}

	return entry.Value(), entry.Revision(), nil
}

func (js *js) Update(
	ctx context.Context,
	key string,
	f UpdateFunc,
) error {
	for {
		data, version, err := js.get(ctx, key)
		if err != nil {
			return err
		}

		newData, err := f(ctx, data)
		switch {
		case errors.Is(err, ErrNoAction):
			return nil
		case errors.Is(err, ErrRetry):
			continue
		case err != nil:
			return err
		}

		// Move out of loop.
		bucket, key := js.getBucketAndKey(key)
		kv, err := js.getBucket(ctx, bucket)
		if err != nil {
			return err
		}

		_, err = kv.Update(ctx, key, newData, version)
		// Check if is a conflict.
		if errors.Is(err, jetstream.ErrKeyExists) {
			continue
		} else if err != nil {
			zap.L().Error("failed to update", zap.Error(err))
			return err
		}

		return nil
	}
}

func (js *js) Delete(ctx context.Context, key string) (bool, error) {
	bucket, key := js.getBucketAndKey(key)
	kv, err := js.getBucket(ctx, bucket)
	if err != nil {
		return false, err
	}

	err = kv.Delete(ctx, key)
	if errors.Is(err, jetstream.ErrKeyNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (js *js) getBucketAndKey(key string) (string, string) {
	b, a, _ := strings.Cut(key, "/")
	b = strings.ReplaceAll(b, ":", "_")
	return b, a
}

func (js *js) getBucket(ctx context.Context, bucket string) (jetstream.KeyValue, error) {
	js.m.Lock()
	defer js.m.Unlock()
	if kv, ok := js.kvs[bucket]; ok {
		return kv, nil
	}

	kv, err := js.kvm.KeyValue(ctx, bucket)
	if err != nil {
		return nil, err
	}

	js.kvs[bucket] = kv
	return kv, nil
}
