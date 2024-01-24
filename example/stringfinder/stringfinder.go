package main

import (
	"fmt"
	"github.com/mj73sudo/gonetic"
	"math/rand"
	"time"
)

type StringChromosome struct {
	Chr string
}

func (s StringChromosome) Fitness() float64 {
	result := "Hello My name is Mohammad Jalili"
	fitness := 0.0
	for i := 0; i < len(s.Chr); i++ {
		if s.Chr[i] == result[i] {
			fitness++
		}
	}
	return fitness
}

func (s StringChromosome) Crossover(other gonetic.Chromosome) gonetic.Chromosome {
	otherSC := other.(StringChromosome)
	crossoverPoint := rand.Intn(len(s.Chr) - 1)
	child := s.Chr[:crossoverPoint] + otherSC.Chr[crossoverPoint:]
	return StringChromosome{Chr: child}
}
func (s StringChromosome) Mutate(mutationRate float64) gonetic.Chromosome {
	mute := s.Chr
	muteRunes := []rune(mute)
	for i := range mute {
		if rand.Float64() < mutationRate {
			chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ,.!?1234567890"
			muteRunes[i] = []rune(chars)[rand.Intn(len(chars))]
		}
	}
	mute = string(muteRunes)

	return StringChromosome{Chr: mute}
}

// NewStringChromosome creates a new StringChromosome with random characters.
func NewStringChromosome(length int) StringChromosome {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ,.!?1234567890"
	chromosome := make([]byte, length)
	for i := range chromosome {
		chromosome[i] = chars[rand.Intn(len(chars))]
	}
	return StringChromosome{Chr: string(chromosome)}
}

type StringHandler struct {
	Prototype gonetic.Chromosome
}

func (sh StringHandler) InitializePopulation(size int) []gonetic.Chromosome {
	chrs := make([]gonetic.Chromosome, size)
	for i := 0; i < size; i++ {
		chrs[i] = sh.Prototype
	}
	return chrs
}
func main() {
	rand.Seed(time.Now().UnixNano())
	config := gonetic.GAConfig{
		PopulationSize:  1000,
		MaxIteration:    1000,
		MutationPercent: 0.3,
	}

	prototype := NewStringChromosome(len("Hello My name is Mohammad Jalili"))
	stringHandler := StringHandler{
		Prototype: prototype,
	}
	ga := gonetic.NewGeneticAlgorithm(stringHandler, config)
	bestSolution := ga.Run()
	fmt.Println("Best Solution:", bestSolution)
}
