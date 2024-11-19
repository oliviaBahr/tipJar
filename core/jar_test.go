package core

import (
	"testing"
	"tipJar/utils"

	"github.com/stretchr/testify/assert"
)

func TestAddTip(t *testing.T) {
	assert := assert.New(t)

	jar, err := LoadTestJar()
	assert.NoError(err, "Failed to load jar")

	err = jar.AddTip("Test Tip", "This is a test tip", "test,tip", "https://example.com")
	assert.NoError(err, "jar.AddTip failed", "e", err)

	allTips := jar.GetAllTips()

	assert.Equal(1, len(allTips), "Expected 1 tip")
}

func TestRemoveTip(t *testing.T) {
	assert := assert.New(t)

	jar, err := LoadTestJar()
	assert.NoError(err, "Failed to load jar")

	jar.AddTip("Test Tip", "This is a test tip", "test,tip", "https://example.com")

	allTips := jar.GetAllTips()
	assert.Equal(1, len(allTips), "Expected 1 tip")

	err = jar.RemoveTip(allTips[0])
	assert.NoError(err, "RemoveTip failed")

	allTips = jar.GetAllTips()
	assert.Equal(0, len(allTips), "Expected 0 tips")
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)

	jar, err := LoadTestJar()
	assert.NoError(err, "Failed to load jar", "e", err)

	jar.AddTip("Test Tip", "1", "test,tip", "https://example.com")
	jar.AddTip("Test Tip 2", "2", "test", "https://example.com")
	jar.AddTip("Test Tip 3", "3", "test", "https://example.com")
	jar.AddTip("ABCD", "4", "test,tip", "https://example.com")

	allTips := jar.GetAllTips()
	assert.Equal(4, len(allTips), "Expected 4 tips")

	tips := jar.SearchTips("test", []string{})
	assert.Equal(3, len(tips), "Expected 3 tips")

	tips = jar.SearchTips("", []string{"tip"})
	assert.Equal(2, len(tips), "Expected 2 tips")
}

func TestSearchByTags(t *testing.T) {
	assert := assert.New(t)

	jar, err := LoadTestJar()
	assert.NoError(err, "Failed to load jar")

	jar.AddTip("Test Tip", "This is a test tip", "test,tip", "https://example.com")
	jar.AddTip("Test Tip 2", "2", "test", "https://example.com")
	jar.AddTip("Test Tip 3", "3", "test", "https://example.com")

	tips := jar.SearchByTags([]string{"test"})
	assert.Equal(3, len(tips), "Expected 3 tips")
}

func TestTemp(t *testing.T) {
	assert := assert.New(t)

	_, err := utils.GetRepoDir()
	assert.NoError(err, "Failed to get repo dir")

}
