package nanoflakes

import (
	"fmt"
	"time"
)

type nanoflakeLocalGenerator struct {
	epoch         int64
	generatorId   int
	sequence      int
	lastTimeStamp int64
}

// returns a pointer to a newly created local generator
func LocalGenerator(epoch int64, generatorId int) (*nanoflakeLocalGenerator, error) {

	if generatorId > MAX_GENERATOR_ID || generatorId < 0 {
		err := fmt.Errorf("Invalid generator ID")
		return nil, err
	}

	return &nanoflakeLocalGenerator{epoch: epoch, generatorId: generatorId, sequence: 0, lastTimeStamp: -1}, nil
}

// returns a pointer to the next nanoflake
func (generator nanoflakeLocalGenerator) Next() (*Nanoflake, error) {
	var timestamp = time.Now().UnixMilli()

	if timestamp <= generator.lastTimeStamp {
		err := fmt.Errorf("Clock moved backwards. Refusing to generate nanoflakes for %d milliseconds", (generator.lastTimeStamp - timestamp))
		return nil, err
	}

	if !(generator.lastTimeStamp != timestamp) {
		generator.sequence = 0
	} else {
		generator.sequence = (generator.sequence + 1) & MAX_SEQUENCE
		if generator.sequence == 0 {
			timestamp = generator.tillNextMillis(timestamp)
		}
	}

	generator.lastTimeStamp = timestamp

	var value = timestamp - generator.epoch<<TIMESTAMP_SHIFT | int64(generator.generatorId)<<GENERATOR_ID_SHIFT | int64(generator.sequence)

	return &Nanoflake{epoch: generator.epoch, value: value}, nil
}

// Gets this nanoflake generator's epoch in milliseconds
func (generator nanoflakeLocalGenerator) EpochMillis() int64 {
	return generator.epoch
}

func (generator nanoflakeLocalGenerator) tillNextMillis(lastTimeStamp int64) int64 {
	var timestamp = time.Now().UnixMilli()

	for timestamp <= lastTimeStamp {
		timestamp = time.Now().UnixMilli()
	}

	return timestamp
}
