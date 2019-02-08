package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Job", func() {
	var job *Job

	JustBeforeEach(func() {
		newJob, err := NewJob(&NewJobOption{})
		if err != nil {
			Panic()
		}

		job = newJob
	})

	Describe("Execute", func() {
		It("should success", func() {
			_, err := job.Execute("")
			Expect(err).NotTo(HaveOccurred())
		})

		It("new job should be started", func() {
			exe, _ := job.Execute("")
			Expect(exe.GetState()).Should(Equal("started"))
		})
	})
})
