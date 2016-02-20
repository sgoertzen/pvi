package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func getTestFileContents(filename string) []byte {
	testFile, err := os.Open("./test-data/" + filename)
	checkError(err)
	defer testFile.Close()
	b, err := ioutil.ReadAll(testFile)
	checkError(err)
	return b
}

func TestParseVEProjects(t *testing.T) {
	veProjects, err := parseVEProjects(getTestFileContents("versionEyeProjects.json"))
	assert.Nil(t, err)
	assert.Equal(t, 2, len(veProjects))

	first := veProjects[0]
	assert.Equal(t, "abc123", first.Id)
	assert.Equal(t, "user/name", first.Name)
	assert.Equal(t, true, first.Public)
	assert.Equal(t, "daily", first.Period)
	assert.Equal(t, "sourcesafe", first.Source)
	assert.Equal(t, 118, first.DependencyCount)
	assert.Equal(t, 36, first.OutdatedCount)
	assert.Equal(t, 0, first.LicensesRed)
	assert.Equal(t, 6, first.LicensesUnknown)
	assert.Equal(t, 118, first.DependencyCountSum)
	assert.Equal(t, 36, first.OutdatedCountSum)
	assert.Equal(t, 0, first.LicensesRedSum)
	assert.Equal(t, 6, first.LicensesUnknownSum)
	assert.Empty(t, first.LicenseWhitelist)
	assert.Equal(t, 2016, first.CreatedAt.Year())
	assert.Equal(t, 13, first.CreatedAt.Hour())
}

func TestParseProject(t *testing.T) {
	veProject, err := parseVEProject(getTestFileContents("versionEyeProjectDetails.json"))
	checkError(err)
	assert.Equal(t, "56983e26af789b0027001e5b", veProject.Id)
	assert.True(t, veProject.Public)
	assert.Equal(t, 15, len(veProject.Dependencies))

	veDependency := veProject.Dependencies[0]
	assert.Nil(t, err)
	assert.Equal(t, "some-maven-plugin", veDependency.Name)
	assert.Equal(t, "org.apache.maven.plugins/some-maven-plugin", veDependency.ProdKey)
	assert.Equal(t, "org.apache.maven.plugins", veDependency.GroupId)
	assert.Equal(t, "some-maven-plugin-id", veDependency.ArtifactId)
	assert.Equal(t, "go", veDependency.Language)
	assert.Equal(t, "3.6", veDependency.VersionCurrent)
	assert.Equal(t, "2.7.1", veDependency.VersionRequested)
	assert.Equal(t, "=", veDependency.Comparator)
	assert.Equal(t, false, veDependency.Unknown)
	assert.Equal(t, true, veDependency.Outdated)
	assert.Equal(t, true, veDependency.Stable)

	assert.Equal(t, 1, len(veDependency.Licenses))
	assert.Equal(t, "Apache-2.0", veDependency.Licenses[0].Name)
	assert.Equal(t, "http://www.apache.org/licenses/LICENSE-2.0.txt", veDependency.Licenses[0].URL)
	assert.Nil(t, veDependency.Licenses[0].OnWhitelist)
	assert.Nil(t, veDependency.Licenses[0].OnCwl)

	assert.Equal(t, 1, len(veDependency.SecurityVulnerabilities))
	assert.Equal(t, "Java", veDependency.SecurityVulnerabilities[0].Language)
	assert.Equal(t, "org.bad-library", veDependency.SecurityVulnerabilities[0].ProdKey)
	assert.Equal(t, "2014-1444", veDependency.SecurityVulnerabilities[0].NameId)
	assert.Equal(t, "someone", veDependency.SecurityVulnerabilities[0].Author)
	assert.Equal(t, "This library is bad", veDependency.SecurityVulnerabilities[0].Summary)
	assert.Equal(t, "desc", veDependency.SecurityVulnerabilities[0].Description)
	assert.Equal(t, "platform", veDependency.SecurityVulnerabilities[0].Platform)
	assert.Equal(t, "bn", veDependency.SecurityVulnerabilities[0].OSVDB)
	assert.Equal(t, "2014-1444", veDependency.SecurityVulnerabilities[0].CVE)
	assert.Equal(t, "3.3", veDependency.SecurityVulnerabilities[0].CVSS)
	assert.Equal(t, "", veDependency.SecurityVulnerabilities[0].PublishDate)
	assert.Equal(t, "", veDependency.SecurityVulnerabilities[0].Framework)
	assert.Equal(t, ">=4.1.0.Beta1,4 && <=4.3.1.Final,4 && <=5.1.1.Final,5", veDependency.SecurityVulnerabilities[0].AffectedVersions)
	assert.Equal(t, ">=4.3.2.Final,4 && >=5.2.0.Final,5", veDependency.SecurityVulnerabilities[0].PatchedVersions)
	assert.Empty(t, veDependency.SecurityVulnerabilities[0].UnaffectedVersions)
}
