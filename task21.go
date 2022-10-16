package main

import "fmt"

type adapter interface {
	showStats()
}

type iphone struct {
	model  string
	memory int
}

func (i *iphone) stats() {
	fmt.Printf("model: %s, memory: %d gb\n", i.model, i.memory)
}

type samsung struct {
	model  string
	memory int
}

func (s *samsung) stats() {
	fmt.Printf("model: %s, memory: %d gb\n", s.model, s.memory)
}

type iphoneAdapter struct {
	iphone
}

func (phone *iphoneAdapter) showStats() {
	phone.iphone.stats()
}

type samsungAdapter struct {
	samsung
}

func (phone *samsungAdapter) showStats() {
	phone.samsung.stats()
}

func newIphone() *iphoneAdapter {
	return &iphoneAdapter{
		iphone{
			model:  "iphone x",
			memory: 64,
		},
	}
}

func newSamsung() *samsungAdapter {
	return &samsungAdapter{samsung{
		model:  "samsung 10",
		memory: 128,
	}}
}

func main() {
	iphone := newIphone()
	samsung := newSamsung()
	iphone.showStats()
	samsung.showStats()
}
