package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func s1c1() error {
	log.Print("Set 1 Challenge 1")
	var (
		// input is hex-encoded
		input string = `49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d`
		// output is a base64 representation of the input
		expectedOutput string = `SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t`
	)

	bytes, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	output := base64.StdEncoding.EncodeToString(bytes)

	log.Print("Hex Input:              ", input)
	log.Print("Expected Base64 Output: ", expectedOutput)
	log.Print("Actual Base64 Output:   ", output)

	if output != expectedOutput {
		return fmt.Errorf("Computed output does not match the expected output")
	}

	return nil
}

func s1c2() error {
	log.Print("Set 1 Challenge 2")
	var (
		input          string = `1c0111001f010100061a024b53535009181c`
		inputBytes     []byte
		key            string = `686974207468652062756c6c277320657965`
		keyBytes       []byte
		expectedOutput string = `746865206b696420646f6e277420706c6179`
		output         string
		outputBytes    []byte
		err            error
	)

	inputBytes, err = hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	keyBytes, err = hex.DecodeString(key)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(inputBytes); i++ {
		outputBytes = append(outputBytes, inputBytes[i]^keyBytes[i])
	}

	output = hex.EncodeToString(outputBytes)
	log.Print("Input:           ", input)
	log.Print("Key:             ", key)
	log.Print("Expected Output: ", expectedOutput)
	log.Print("Actual Output:   ", output)

	if output != expectedOutput {
		return fmt.Errorf("Computed output does not match the expected output")
	}

	return nil
}
