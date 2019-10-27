package models

import "sync"

type ChopStick struct {
	sync.Mutex
}
