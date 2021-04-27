package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
)

type Job struct {
	Name        string
	Description string
	Id          int
	Ids         []int
}

func main() {
	// data := struct {
	// 	string
	// 	Job
	// }{"gilbert ",
	// 	*new(Job)}
	
	data := Job{
		Name:        "gilbert",
		Description: "tests",
		Id:          1,
		Ids:         make([]int, 10),
	}
	data1 := Job{
		Name:        "gilbert",
		Description: "test",
		Id:          1,
		Ids:         make([]int, 10),
	}
	
	// ########### int hash ############
	hash, err := fnv.New64a().Write([]byte(data.Name))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hash: %d\n", hash)
	
	hash, err = fnv.New32().Write([]byte(data.Name))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hash: %d\n", hash)
	
	// ########### string hash ############
	md5 := md5.Sum([]byte(data.Name))
	sha1String := sha1.Sum([]byte(data.Name))
	sha256String := sha256.Sum256([]byte(data.Name))
	
	// ########### struct hash ############
	byteArr, _ := json.Marshal(data)
	sha256struct := sha1.Sum(byteArr)
	
	byteArr1, _ := json.Marshal(data1)
	sha256struct1 := sha1.Sum(byteArr1) // Result not diff from sha256struct. when no fields are exported
	
	fmt.Printf("md5: %x\n", md5)
	fmt.Printf("sha1: %x\n", sha1String)
	fmt.Printf("sha256: %x\n", sha256String)
	fmt.Printf("sha256struct: %x\n", sha256struct)
	fmt.Printf("sha256struct1: %x\n", sha256struct1)
}
