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

type DecodedSnowflake struct {
  id uint64
  timestamp uint64
  nodeID uint32
  seq uint32
  epoch uint64
}

func NewSnowflakeWithNodeID(epoch uint64, nodeID uint32) *Snowflake {
  return &Snowflake {
    epoch: epoch,
    nodeID: nodeID,
    seq: 0,
    lastSequenceExhaustion: 0,
  }
}

func (sf *Snowflake) gen() string {
  return sf.genWithTs(uint64(time.Now().Unix()))
}
 
func (sf *Snowflake) genWithTs(timestamp uint64) string {
  if sf.seq >= 4095 && timestamp == sf.lastSequenceExhaustion {
    for time.Now().Unix()-int64(timestamp) < 1 {
      // 
    }
  }

  id := ((timestamp - sf.epoch) << 22) | (uint64(sf.nodeID) << 12) | uint64(sf.seq)
  sf.seq = (sf.seq + 1) & 4095

  if sf.seq == 0 {
    sf.lastSequenceExhaustion = timestamp
  }

  return strconv.FormatUint(id, 10)
}

func (sf *Snowflake) decode(id string) DecodedSnowflake {
  sfID, _ := strconv.ParseUint(id, 10, 64)
  timestamp := (sfID >> 22) + sf.epoch
  nodeID := (sfID >> 12) & 0b1111111111
  seq := sfID & 0b1111111111

  return DecodedSnowflake {
    ID: sfID,
    Timestamp: timestamp,
    NodeID: uint32(nodeID),
    Seq: seq,
    Epoch: sf.epoch,
  }
}
