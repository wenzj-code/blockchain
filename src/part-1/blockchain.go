/*
 * @Author: your name
 * @Date: 2020-07-17 22:44:43
 * @LastEditTime: 2020-07-17 22:47:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \blockchain\src\part-1\blockchain.go
 */
package main

// BlockChain 是一个 Block 指针数组
type BlockChain struct {
	blocks []*Block
}

// NewBlockChain 创建一个有创世块的链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// AddBlock 向链中加入一个新块
// data 在实际中就是交易
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
