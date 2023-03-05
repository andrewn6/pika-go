package main

type PrefixRecord struct {
	prefix      string
	description *string
	secure      bool
}

type DecodePika struct {
	prefix        string
	tail          string
	snowflake     uint64
	node_id       uint32
	timestamp     uint32
	epoch         uint64
	seq           uint64
	version       int8
	prefix_record PrefixRecord
}
