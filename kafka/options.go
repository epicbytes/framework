package kafka

import "time"

type consumerOption struct {
	Partition int32
	Offset    int64
	timeout   time.Duration
}

type ConsumerOption func(option *consumerOption)

func WithPartition(partition int32) ConsumerOption {
	return func(b *consumerOption) {
		b.Partition = partition
	}
}

func WithOffset(offset int64) ConsumerOption {
	return func(b *consumerOption) {
		b.Offset = offset
	}
}

func WithTimeout(timeout time.Duration) ConsumerOption {
	return func(b *consumerOption) {
		b.timeout = timeout
	}
}
