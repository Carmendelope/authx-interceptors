/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package apikey

import "github.com/nalej/derrors"

// APIKeyAccess interface to delegate checking whether an API key is valid.
type APIKeyAccess interface {
	// Connect to the appropriate backend.
	Connect() derrors.Error
	// Check if the API Key is valid
	IsValid(apiKey string) derrors.Error
}