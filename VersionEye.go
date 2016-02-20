package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Update struct {
	CurrentVersion   string
	AvailableUpgrade string
	Project          *Project
}

type VersionEyeProjects []*VersionEyeProject

type VersionEyeProject struct {
	Id                 string      `json:"id"`
	Name               string      `json:"name"`
	ProjectType        string      `json:"project_type"`
	Public             bool        `json:"public"`
	PrivateScm         bool        `json:"private_scm"`
	Period             string      `json:"period"`
	Source             string      `json:"source"`
	DependencyCount    int         `json:"dep_number"`
	OutdatedCount      int         `json:"out_number"`
	DependencyCountSum int         `json:"dep_number_sum"`
	OutdatedCountSum   int         `json:"out_number_sum"`
	LicensesRed        int         `json:"licenses_red"`
	LicensesUnknown    int         `json:"licenses_unknown"`
	LicensesRedSum     int         `json:"licenses_red_sum"`
	LicensesUnknownSum int         `json:"licenses_unknown_sum"`
	SvCount            int         `json:"sv_count"`
	CreatedAt          CustomTime  `json:"created_at"`
	UpdatedAt          CustomTime  `json:"updated_at"`
	LicenseWhitelist   interface{} `json:"license_whitelist"`
	Dependencies       []struct {
		Name             string `json:"name"`
		ProdKey          string `json:"prod_key"`
		GroupId          string `json:"group_id"`
		ArtifactId       string `json:"artifact_id"`
		Language         string `json:"language"`
		VersionCurrent   string `json:"version_current"`
		VersionRequested string `json:"version_requested"`
		Comparator       string `json:"comparator"`
		Unknown          bool   `json:"unknown"`
		Outdated         bool   `json:"outdated"`
		Stable           bool   `json:"stable"`
		Licenses         []struct {
			Name        string      `json:"name"`
			URL         string      `json:"url"`
			OnWhitelist interface{} `json:"on_whitelist"`
			OnCwl       interface{} `json:"on_cwl"`
		} `json:"licenses"`
		SecurityVulnerabilities []struct {
			Language           string `json:"language"`
			ProdKey            string `json:"prod_key"`
			NameId             string `json:"name_id"`
			Author             string `json:"author"`
			Summary            string `json:"summary"`
			Description        string `json:"description"`
			Platform           string `json:"platform"`
			OSVDB              string `json:"osvdb"`
			CVE                string `json:"cve"`
			CVSS               string `json:"cvss_v2"`
			PublishDate        string `json:"publish_date"`
			Framework          string `json:"framework"`
			AffectedVersions   string `json:"affected_versions_string"`
			PatchedVersions    string `json:"patched_versions_string"`
			UnaffectedVersions string `json:"unaffected_versions_string"`
		} `json:"security_vulnerabilities"`
	} `json:"dependencies"`
}

type CustomTime struct {
	time.Time
}

const shortVEFormat = "02.01.2006-15:04"
const longVEFormat = "2006-01-02T15:04:05.000Z"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	// Version eye will sometimes send dates in shortened format and other times in long format.  Try both.
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.Parse(shortVEFormat, string(b))
	if err != nil {
		ct.Time, err = time.Parse(longVEFormat, string(b))
		check(err)
	}
	return
}

var apikey string

func SetKey(key string) {
	apikey = key
}

func GetAvailableUpdates(project Project) bool {
	veye := lookupProject(project)
	if veye.OutdatedCountSum > 0 {
		return true
	}
	return false
}

func lookupProject(project Project) *VersionEyeProject {
	vProjects := getAllProjectsFromVersionEye()
	for _, eye := range vProjects {
		if eye.Id == project.ArtifactId {
			return eye
		}
	}
	// Parse the response into a VersionEyeProject
	return new(VersionEyeProject)
}

type SecurityVulnerability struct {
	Language           string `json:"language"`
	ProdKey            string `json:"prod_key"`
	NameId             string `json:"name_id"`
	Author             string `json:"author"`
	Summary            string `json:"summary"`
	Description        string `json:"description"`
	Platform           string `json:"platform"`
	OSVDB              string `json:"osvdb"`
	CVE                string `json:"cve"`
	CVSS               string `json:"cvss_v2"`
	PublishDate        string `json:"publish_date"`
	Framework          string `json:"framework"`
	AffectedVersions   string `json:"affected_versions_string"`
	PatchedVersions    string `json:"patched_versions_string"`
	UnaffectedVersions string `json:"unaffected_versions_string"`
	//Links string `json:"links"`
}

func getAllProjectsFromVersionEye() VersionEyeProjects {
	checkApiKey()
	resp, err := http.Get("https://www.versioneye.com/api/v2/projects?api_key=" + apikey)
	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	vprojects, err := parseVEProjects(body)
	return vprojects
}

func getProjectDetailsFromVersionEye(versionEyeProjectId string) VersionEyeProject {
	checkApiKey()
	resp, err := http.Get("https://www.versioneye.com/api/v2/projects/" + versionEyeProjectId + "?api_key=" + apikey)
	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	vprojects, err := parseVEProject(body)
	return vprojects
}

func parseVEProject(body []byte) (VersionEyeProject, error) {
	var s = new(VersionEyeProject)
	err := json.Unmarshal(body, &s)
	checkError(err)
	return *s, err
}

func parseVEProjects(body []byte) (VersionEyeProjects, error) {
	var s = new(VersionEyeProjects)
	err := json.Unmarshal(body, &s)
	checkError(err)
	return *s, err
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkApiKey() {
	if len(apikey) == 0 {
		panic("No Version Eye API Key set")
	}
}
