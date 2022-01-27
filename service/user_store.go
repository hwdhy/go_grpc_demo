/*
@Time : 2022/1/23 19:08
@Author : Hwdhy
@File : user_store
@Software: GoLand
*/
package service

import "sync"

type UserStore interface {
	Save(user *User) error
	Find(username string) (*User, error)
}

type InMempryUserStore struct {
	mutex sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserStore() *InMempryUserStore {
	return &InMempryUserStore{
		users: make(map[string]*User),
	}
}

func (store *InMempryUserStore) Save(user *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return ErrAlreadyExists
	}

	store.users[user.Username] = user.Clone()
	return nil
}

func (store *InMempryUserStore) Find(username string) (*User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[username]
	if user == nil {
		return nil, nil
	}
	return user.Clone(), nil
}
