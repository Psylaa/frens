package models_test

import (
	"testing"

	"github.com/bwoff11/frens/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestPrivacyToString(t *testing.T) {
	testCases := []struct {
		name     string
		privacy  models.Privacy
		expected string
	}{
		{
			name:     "public",
			privacy:  models.PrivacyPublic,
			expected: "public",
		},
		{
			name:     "protected",
			privacy:  models.PrivacyProtected,
			expected: "protected",
		},
		{
			name:     "private",
			privacy:  models.PrivacyPrivate,
			expected: "private",
		},
		{
			name:     "invalid",
			privacy:  models.Privacy("invalid"),
			expected: "private",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.privacy.ToString())
		})
	}
}
