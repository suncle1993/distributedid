package distributedid

// Options snowflake 配置option
type Options struct {
	Node     int64
	Epoch    int64
	NodeBits uint8
	StepBits uint8
}

// Option ...
type Option func(options *Options)

// Node 参考snowflake.Node
func Node(node int64) Option {
	return func(options *Options) {
		options.Node = node
	}
}

// Epoch 参考snowflake.Epoch
func Epoch(epoch int64) Option {
	return func(options *Options) {
		options.Epoch = epoch
	}
}

// NodeBits 参考snowflake.NodeBits
func NodeBits(nodeBits uint8) Option {
	return func(options *Options) {
		options.NodeBits = nodeBits
	}
}

// StepBits 参考snowflake.StepBits
func StepBits(stepBits uint8) Option {
	return func(options *Options) {
		options.StepBits = stepBits
	}
}
