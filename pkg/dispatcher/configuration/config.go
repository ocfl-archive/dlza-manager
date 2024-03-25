package configuration

import "time"

type Config struct {
	PortHandler        string
	PortStorageHandler string
	CycleLength        time.Duration
}
