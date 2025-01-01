package main

import (
	"log"
	"time"
)

type Job interface {
	PerformOperation(Worker)
	PerformBiOperation(Worker)
	IsOperationSet() bool
	IsBiOperationSet() bool
}

type GenericJob[R, S comparable] struct {
	recv        <-chan R
	send        chan<- S
	operation   func(R) S
	biOperation func(R) []S
}

func NewGenericJob[R, S comparable](recv <-chan R, send chan<- S) *GenericJob[R, S] {
	return &GenericJob[R, S]{
		recv,
		send,
		nil,
		nil,
	}
}

func (gj *GenericJob[R, S]) IsOperationSet() bool {
	return gj.operation != nil
}

func (gj *GenericJob[R, S]) IsBiOperationSet() bool {
	return gj.biOperation != nil
}

func (gj *GenericJob[R, S]) AddOperation(o func(R) S) {
	if gj.operation != nil || gj.biOperation != nil {
		log.Panic("Cannot update job. Operation is already set on job.")
	}

	gj.operation = o
}

func (gj *GenericJob[R, S]) AddBiOperation(bo func(R) []S) {
	if gj.operation != nil || gj.biOperation != nil {
		log.Panic("Cannot update job. Operation is already set on job.")
	}

	gj.biOperation = bo
}

func (j *GenericJob[R, S]) PerformOperation(w Worker) {
	if j.operation == nil {
		log.Println("Cannot perform operation. Operation not set on job.")
		return
	}

	for i := range j.recv {
		log.Printf("%s: Worker %d performing operation\n", time.Now(), w.id)
		j.send <- j.operation(i)
	}
}

func (j *GenericJob[R, S]) PerformBiOperation(w Worker) {
	if j.biOperation == nil {
		log.Println("Cannot perform bi operation. Bi operation not set on job.")
		return
	}

	for i := range j.recv {
		log.Printf("%s: Worker %d performing bi operation\n", time.Now(), w.id)
		res := j.biOperation(i)
		for _, r := range res {
			j.send <- r
		}
	}
}

type Worker struct {
	id int
}

func NewWorker(id int) *Worker {
	return &Worker{
		id,
	}
}

func (w *Worker) PerformJob(j Job) {
	if !j.IsOperationSet() && !j.IsBiOperationSet() {
		log.Println("Worker cannot perform job. No operation not set on job.")
		return
	}

	if j.IsOperationSet() {
		j.PerformOperation(*w)
	} else if j.IsBiOperationSet() {
		j.PerformBiOperation(*w)
	}
}
