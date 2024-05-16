package main

import (
	"github.com/gin-gonic/gin"
	"math/rand/v2"
	"net/http"
	"slices"
)

const (
	Low      = 100000000000
	High     = 999999999999
	alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type GenerateInput struct {
	URL string `json:"url"`
}

func GenerateRandomNumber(low, high int64) int64 {
	return low + rand.Int64N(high-low)
}

func IntToBase62(n int64) string {
	var chars []byte
	for n > 0 {
		rem := n % 62
		n /= 62

		chars = append(chars, alphabet[rem])
	}

	slices.Reverse(chars)
	return string(chars)
}

func Generate(c *gin.Context) {
	var input GenerateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	n := GenerateRandomNumber(Low, High)
	shortHash := IntToBase62(n)

	c.JSON(http.StatusOK, gin.H{"hash": shortHash})
}
