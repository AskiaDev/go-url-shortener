package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



var testStoreService *StorageService

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}


func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.google.com"
	userID := "tcYd/i9istfk6IU0dNXwzw=="
	shortURL := "j5ljOWdOF+E="


	// Persist the mapping
	SaveUrlMapping(shortURL, initialLink, userID)

	// Retrieve the mapping
	retrievedUrl := RetrieveInitialUrl(shortURL)
	assert.Equal(t, initialLink, retrievedUrl)



	assert.Equal(t, initialLink, retrievedUrl)
}
