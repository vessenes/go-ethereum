package main

import (
  _"fmt"
)

type Contract struct {
  t uint32 // contract is always 1
  amount uint64 // ???
  state  *Trie
}

func NewContract(amount uint64, root []byte) *Contract {
  contract := &Contract{t: 1, amount: amount}
  contract.state = NewTrie(Db, string(root))

  return  contract
}

func (c *Contract) MarshalRlp() []byte {
  return Encode([]interface{}{c.t, c.amount, c.state.root})
}

func (c *Contract) UnmarshalRlp(data []byte) {
  decoder := NewRlpDecoder(data)

  c.t = uint32(decoder.Get(0).AsUint())
  c.amount = decoder.Get(1).AsUint()
  c.state = NewTrie(Db, decoder.Get(2).AsString())
}

type Ether struct {
  t uint32
  amount uint64
  nonce string
}

func NewEtherFromData(data []byte) *Ether {
  ether := &Ether{}
  ether.UnmarshalRlp(data)

  return ether
}

func (e *Ether) MarshalRlp() []byte {
  return Encode([]interface{}{e.t, e.amount, e.nonce})
}

func (e *Ether) UnmarshalRlp(data []byte) {
  t, _ := Decode(data, 0)

  if slice, ok := t.([]interface{}); ok {
    if t, ok := slice[0].(uint8); ok {
      e.t = uint32(t)
    }

    if amount, ok := slice[1].(uint8); ok {
      e.amount = uint64(amount)
    }

    if nonce, ok := slice[2].([]uint8); ok {
      e.nonce = string(nonce)
    }
  }
}
