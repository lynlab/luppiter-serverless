package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Execution", func() {
	var e *Execution

	JustBeforeEach(func() {
		e = NewExecution(&NewExecutionOption{})
	})

	Describe("NewExecution", func() {
		It("state should be \"created\"", func() {
			Expect(e.GetState()).Should(Equal("created"))
		})
	})

	Describe("Start", func() {
		It("should no error", func() {
			err := e.Start()

			Expect(err).NotTo(HaveOccurred())
		})

		It("finished/canceled execution can't be started", func() {
			e.SetState("finished")
			err := e.Start()

			Expect(err).To(HaveOccurred())
		})

		It("state should be \"started\"", func() {
			e.Start()

			Expect(e.GetState()).Should(Equal("started"))
		})
	})
})
