package gonetic

import (
	"math/rand"
	"sort"
)

type GAConfig struct {
	MaxIteration     int
	PopulationSize   int
	CrossoverPercent float64
	MutationPercent  float64
}

type Chromosome interface {
	Fitness() float64
	Crossover(other Chromosome) Chromosome
	Mutate(mutationRate float64) Chromosome
}

type Handler interface {
	InitializePopulation(size int) []Chromosome
}

type GeneticAlgorithm struct {
	Population []Chromosome
	Config     GAConfig
}

// NewGeneticAlgorithm initializes a new GeneticAlgorithm instance.
func NewGeneticAlgorithm(handler Handler, config GAConfig) *GeneticAlgorithm {
	population := handler.InitializePopulation(config.PopulationSize)
	return &GeneticAlgorithm{
		Population: population,
		Config:     config,
	}
}

// InitializePopulation generates an initial population of chromosomes.
func InitializePopulation(size int, prototype Chromosome) []Chromosome {
	population := make([]Chromosome, size)
	for i := range population {
		population[i] = prototype.Mutate(0) // Initialize with mutated copies of the prototype
	}
	return population
}

func (ga *GeneticAlgorithm) Run() Chromosome {
	for generation := 0; generation < ga.Config.MaxIteration; generation++ {
		// Sort population by fitness
		sort.Slice(ga.Population, func(i, j int) bool {
			return ga.Population[i].Fitness() > ga.Population[j].Fitness()
		})

		// Select parents and create new population
		//newPopulation := make([]Chromosome, ga.Config.PopulationSize)
		childPopulation := make([]Chromosome, 0)
		copy(childPopulation, ga.Population)
		for i := 0; i < ga.Config.PopulationSize; i += 2 {
			parent1, parent2 := ga.selectParents()
			child1 := parent1.Crossover(parent2)
			child2 := parent2.Crossover(parent1)
			childPopulation = append(childPopulation, child1)
			childPopulation = append(childPopulation, child2)
			//newPopulation[i] = child1.Mutate(ga.Config.MutationPercent)
			//newPopulation[i+1] = child2.Mutate(ga.Config.MutationPercent)
		}

		for i := 0; i < len(childPopulation); i += 1 {
			if rand.Intn(2) == 0 {
				mutate := childPopulation[i].Mutate(ga.Config.MutationPercent)
				childPopulation = append(childPopulation, mutate)
			}
		}

		sort.Slice(childPopulation, func(i, j int) bool {
			return childPopulation[i].Fitness() > childPopulation[j].Fitness()
		})
		// Replace old population with the new one
		ga.Population = childPopulation[:ga.Config.PopulationSize]
	}

	// Return the best solution
	sort.Slice(ga.Population, func(i, j int) bool {
		return ga.Population[i].Fitness() > ga.Population[j].Fitness()
	})

	return ga.Population[0]
}

// selectParents uses tournament selection to choose two parents.
func (ga *GeneticAlgorithm) selectParents() (Chromosome, Chromosome) {
	// Randomly select two individuals for the tournament
	index1 := rand.Intn(len(ga.Population))
	index2 := rand.Intn(len(ga.Population))

	// Return the better-fit individual as parent1
	if ga.Population[index1].Fitness() > ga.Population[index2].Fitness() {
		return ga.Population[index1], ga.Population[index2]
	}
	return ga.Population[index2], ga.Population[index1]
}
