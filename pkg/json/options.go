package json

import (
	"fmt"
)

// Option defines a function that can act on the event
type Option func(e *Event)

// Add adds a key value pair to the object
func Add(key string, value interface{}) Option {
	return func(e *Event) {
		if _, ok := (*e)[key]; !ok {
			(*e)[key] = value
		}
	}
}

// Reject rejects an object base on a key value pair
func Reject(key string, value interface{}) Option {
	return func(e *Event) {
		if (*e)[key] == value {
			*e = nil
		}
	}
}

// Remove removes all occurences of key from objects
func Remove(key string) Option {
	return func(e *Event) {
		if _, ok := (*e)[key]; ok {
			delete(*e, key)
		}
	}
}

// Prefix adds a prefix to the key
func Prefix(key, prefix string) Option {
	return func(e *Event) {
		if _, ok := (*e)[key]; ok {
			prefixedKey := fmt.Sprintf("%s_%s", prefix, key)
			(*e)[prefixedKey] = (*e)[key]
			delete(*e, key)
		}
	}
}
