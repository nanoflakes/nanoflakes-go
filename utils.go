package nanoflakes

const (
	TIMESTAMP_BITS     = 14
	GENERATOR_ID_BITS  = 10
	SEQUENCE_BITS      = 12
	MAX_GENERATOR_ID   = 1023
	MAX_SEQUENCE       = 4095
	GENERATOR_ID_SHIFT = SEQUENCE_BITS
	TIMESTAMP_SHIFT    = SEQUENCE_BITS + GENERATOR_ID_SHIFT
)

// Gets the timestamp of a nanoflake.
func TimeStampValue(value int64) int64 {
	return value >> TIMESTAMP_SHIFT
}

// Gets the generator ID of a nanoflake
func GeneratorValue(value int64) int64 {
	return value >> GENERATOR_ID_SHIFT & MAX_GENERATOR_ID
}

// Gets the sequence ID of a nanoflake
func SequenceValue(value int64) int64 {
	return value & MAX_SEQUENCE
}
