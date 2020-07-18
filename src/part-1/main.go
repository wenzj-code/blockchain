/*
 * @Author: your name
 * @Date: 2020-07-17 22:44:43
 * @LastEditTime: 2020-07-17 23:09:58
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \blockchain\src\part-1\main.go
 */
package main

import (
	"fmt"
)

/**
 * @description: main funtion
 * @param {type}
 * @return:
 */
func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Ivan dd")
	bc.AddBlock("Send 2 more BTC ddto Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hashx: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println()
	}
}
