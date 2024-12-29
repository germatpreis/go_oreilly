package main

import "strconv"

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore(names ...string) SimpleDataStore {
	store := SimpleDataStore{userData: map[string]string{}}
	for i, v := range names {
		store.userData[strconv.Itoa(i)] = v
	}
	return store
}
