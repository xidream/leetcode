package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test0003LongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "0003LongestSubstringWithoutRepeatingCharacters Suite")
}