package controllers

import (
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Empty", func() {
	Context("Success", func() {
		BeforeEach(func() {
			// 初期化処理1
		})

		It("should be true", func() {
			// テストケースその1
		})

		It("should be not nil", func() {
			// テストケースその2
		})
	})

	Context("Failure", func() {
		BeforeEach(func() {
			// 初期化処理2
		})

		It("should be false", func() {
			// テストケースその3
		})

		It("should be nil", func() {
			// テストケースその4
		})
	})
})
