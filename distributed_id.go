package distributedid

import (
	"github.com/bwmarrin/snowflake"
)

// ID ...
type ID interface {
	String() string
	Int64() int64
	Bytes() []byte
	Base64() string
}

// IDGenerator ID生成器
type IDGenerator interface {
	Generate() ID
}

// Snowflake ...
type Snowflake struct {
	options *Options
	node    *snowflake.Node
}

// Generate 生成唯一ID
func (sf *Snowflake) Generate() ID {
	if sf.node != nil {
		return sf.node.Generate()
	}
	return nil
}

// NewIDGenerator 创建ID生成器
func NewIDGenerator(ip string, options ...Option) (generator IDGenerator, err error) {
	var defaultOptions = &Options{
		Epoch:    1288834974657,
		NodeBits: 16,
		StepBits: 6,
	}

	for _, opt := range options {
		opt(defaultOptions)
	}

	if defaultOptions.Node == 0 {
		var node int64
		node, err = convertNode(ip)
		if err != nil {
			return
		}
		defaultOptions.Node = node
	}

	snowflake.Epoch = defaultOptions.Epoch
	snowflake.NodeBits = defaultOptions.NodeBits
	snowflake.StepBits = defaultOptions.StepBits

	sNode, err := snowflake.NewNode(defaultOptions.Node)
	if err != nil {
		return
	}
	generator = &Snowflake{
		options: defaultOptions,
		node:    sNode,
	}
	return
}
