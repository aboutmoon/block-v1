package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Version int64 // 区块版本号,表示本区块遵循的验证规则
	PrevBlockHash []byte //前区块的hash值,使用SHA256(SHA256(父区块头))计算
	MerKelRoot []byte //该区块中交易的Merkle树根的哈希值，同样采用SHA256(SHA256())计算
	TimeStamp int64 //该区块产生的近似时间，精确到秒，必须大于前11个区块的时间的中值，同时全节点也会拒绝哪些超出自己两个小时时间戳的区块
	Bits int64 // 该区块工作量证明算法的难度目标，已经使用特定算法编码
	Nonce int64 // 为了找到满足难度目标所设定的随机数，为了解决32为随机数在算力飞升的情况下不够用的问题，规定时间戳和coinbase交易信息均可修改，以此扩展nonce的位数

	Hash []byte // 当前区块的hash值，为了简化代码
	Data []byte // 交易信息
}

func (block *Block) SetHash() {
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(block.Nonce),
		block.Data}
	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block  {
	var block Block
	block = Block{
		Version: 1,
		PrevBlockHash: prevBlockHash,
		// Hash todo:
		MerKelRoot: []byte{},
		TimeStamp: time.Now().Unix(),
		Bits: 1,
		Nonce: 1,
		Data: []byte(data),
	}
	block.SetHash()
	return &block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}
