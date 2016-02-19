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
	ProjectId            string
	Id                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	ProjectType          string                 `json:"project_type"`
	Public               bool                   `json:"public"`
	Period               string                 `json:"period"`
	Source               string                 `json:"source"`
	DependencyNumber     int                    `json:"dep_number"`
	OutNumber            int                    `json:"out_number"`
	LicensesRed          int                    `json:"licenses_red"`
	LicensesUnknown      int                    `json:"licenses_unknown"`
	DependencyNumberSum  int                    `json:"dep_number_sum"`
	OutNumberSum         int                    `json:"out_number_sum"`
	LicensesRedSum       int                    `json:"licenses_red_sum"`
	LicensesUnknownSum   int                    `json:"licenses_unknown_sum"`
	LicenseWhiteListName string                 `json:"license_whitelist_name"`
	CreatedAt            CustomTime `json:"created_at"`
	UpdatedAt            CustomTime `json:"updated_at"`
	// Format of dates is "20.01.2016-22:32"
}

type CustomTime struct {
	time.Time
}

const ctLayout = "02.01.2006-15:04"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.Parse(ctLayout, string(b))
	return
}

func GetAvailableUpdates(project Project) bool {
	veye := lookupProject(project)
	if veye.OutNumberSum > 0 {
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

func getAllProjectsFromVersionEye() VersionEyeProjects {
	resp, err := http.Get("https://www.versioneye.com/api/v2/projects?api_key=c78c87ec4d8f647d818c")
	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	vprojects, err := parseVEProjects(body)
	return vprojects
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
