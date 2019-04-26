/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package apikey

import (
	"github.com/nalej/derrors"
	"sync"
	"time"
)


// DefaultTokenTTL with a default TTL
const DefaultTokenTTL = time.Hour

type InMemoryAPIKeyAccess struct{
	// Mutex for managing mockup access.
	sync.Mutex
	// token map with the join tokens and their expiration date.
	token map[string]int64
	// ttl with the token ttl.
	ttl time.Duration
}

// NewInMemoryAPIKeyAccess with the default TTL.
func NewInMemoryAPIKeyAccess() * InMemoryAPIKeyAccess{
	return &InMemoryAPIKeyAccess{
		token: make(map[string]int64, 0),
		ttl: DefaultTokenTTL,
	}
}

// Add a new token
func (ima * InMemoryAPIKeyAccess) Add(token string){
	ima.Lock()
	ima.token[token] = time.Now().Add(ima.ttl).Unix()
	ima.Unlock()
}

// Connect to the appropriate backend.
func (ima * InMemoryAPIKeyAccess)Connect() derrors.Error{
	return nil
}

// Check if the API Key is valid
func (ima * InMemoryAPIKeyAccess) IsValid(apiKey string) derrors.Error{
	ima.Lock()
	defer ima.Unlock()
	expire, exists := ima.token[apiKey]
	if exists{
		if expire >= time.Now().Unix(){
			return nil
		}else{
			// Expire the token
			delete(ima.token, apiKey)
		}
	}
	return derrors.NewUnauthenticatedError("invalid token")
}

func (ima * InMemoryAPIKeyAccess) Clear() {
	ima.Lock()
	ima.token = make(map[string]int64, 0)
	ima.Unlock()
}