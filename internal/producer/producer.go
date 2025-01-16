package producer

import (
	"math/rand"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type Food struct {
	Name string
}

type Producer interface {
	Produce(chan<- Food)
}

type Consumer interface {
	Consume(<-chan Food)
}

type NakulProducer struct {
}

type ChayConsumer struct {
}

func (p *NakulProducer) Produce(foodChan chan<- Food) {
	foodItems := []string{"pizza", "burger", "dumplings"}

	for {
		food := Food{
			Name: foodItems[rand.Intn(len(foodItems))],
		}

		foodChan <- food
		log.Info().Str("food", food.Name).Msg("nakul gave some food")
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	}

}

func (p *ChayConsumer) Consume(foodChan <-chan Food) {
	for food := range foodChan {
		log.Info().Str("food", food.Name).Msg("Chay ate some food")
		time.Sleep(time.Second)
	}
}

func RunExamples() error {
	foodChan := make(chan Food, 10)

	producers := []Producer{
		&NakulProducer{},
		&NakulProducer{},
	}

	consumers := []Consumer{
		&ChayConsumer{},
		&ChayConsumer{},
		&ChayConsumer{},
	}

	var wg sync.WaitGroup

	for _, p := range producers {
		wg.Add(1)
		go func(prod Producer) {
			defer wg.Done()
			prod.Produce(foodChan)

		}(p)
	}

	for _, c := range consumers {
		wg.Add(1)
		go func(cons Consumer) {
			defer wg.Done()
			cons.Consume(foodChan)

		}(c)
	}

	wg.Wait()
	return nil
}
