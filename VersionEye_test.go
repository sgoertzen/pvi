package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testJson = `[{
		"id": "abc123",
		"name": "user/name",
		"project_type": "Maven2",
		"public": true,
		"period": "daily",
		"source": "sourcesafe",
		"dep_number": 118,
		"out_number": 36,
		"licenses_red": 0,
		"licenses_unknown": 6,
		"dep_number_sum": 118,
		"out_number_sum": 36,
		"licenses_red_sum": 0,
		"licenses_unknown_sum": 6,
		"license_whitelist_name": null,
		"created_at": "15.01.2016-13:32",
		"updated_at": "15.01.2016-00:32"
	},
	{
		"id": "xyz456",
	    "name": "user/name2",
	    "project_type": "Maven2",
	    "public": true,
	    "period": "weekly",
	    "source": "flatfile",
	    "dep_number": 13,
	    "out_number": 0,
	    "licenses_red": 0,
	    "licenses_unknown": 0,
	    "dep_number_sum": 13,
	    "out_number_sum": 0,
	    "licenses_red_sum": 0,
	    "licenses_unknown_sum": 0,
	    "license_whitelist_name": null,
		"created_at": "15.01.2016-00:32",
		"updated_at": "15.01.2016-00:32"
	}]`

func TestParseVEProjects(t *testing.T) {
	veProjects, err := parseVEProjects([]byte(testJson))
	assert.Nil(t, err)
	assert.Equal(t, 2, len(veProjects))

	first := veProjects[0]
	assert.Equal(t, "abc123", first.Id)
	assert.Equal(t, "user/name", first.Name)
	assert.Equal(t, true, first.Public)
	assert.Equal(t, "daily", first.Period)
	assert.Equal(t, "sourcesafe", first.Source)
	assert.Equal(t, 118, first.DependencyNumber)
	assert.Equal(t, 36, first.OutNumber)
	assert.Equal(t, 0, first.LicensesRed)
	assert.Equal(t, 6, first.LicensesUnknown)
	assert.Equal(t, 118, first.DependencyNumberSum)
	assert.Equal(t, 36, first.OutNumberSum)
	assert.Equal(t, 0, first.LicensesRedSum)
	assert.Equal(t, 6, first.LicensesUnknownSum)
	assert.Empty(t, first.LicenseWhiteListName)
	assert.Equal(t, 2016, first.CreatedAt.Year())
	assert.Equal(t, 13, first.CreatedAt.Hour())
}
