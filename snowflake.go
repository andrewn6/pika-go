package main\

import (
  "strconv"
  "time"
)

type Snowflake struct {
  epoch uint64
  nodeID, seq uint32
  lastSequenceExhaustion uint64
}

type DecodedSnowflake {
  id uint64
  timestamp uint64
  nodeID uint32
  seq uint32
  epoch uint64
}
