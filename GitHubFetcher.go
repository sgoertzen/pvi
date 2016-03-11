package pvi

import (
	"bufio"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
    "os"
	"os/exec"
)

// CloneAllRepos clones all repos for an orgnaization
func CloneAllRepos(orgname string) error {
	client := github.NewClient(nil)
	org, resp, err := client.Organizations.Get(orgname)
	check(err)
	check(github.CheckResponse(resp.Response))
	log.Printf("total private repos: %d", org.TotalPrivateRepos)

	repos, resp, err := client.Repositories.ListByOrg(*org.Name, nil)
	for _, repo := range repos {
		// clone!!
		log.Println("would clone " + *repo.Name + " (" + *repo.CloneURL + ")")
		//clone(*repo.CloneURL)
	}
	return nil
}

func refreshAllRepos(org string) error {
	return nil
}

func getClient() *github.Client {
	token := oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")}
	ts := oauth2.StaticTokenSource(&token)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	return client
}

func clone(cloneURL string) (int, error) {
	app, err := exec.LookPath("git")
	cmd := exec.Command(app, "clone", cloneURL)

	path := "./"
	cmd.Dir = path
	stdout, err := cmd.StdoutPipe()
	check(err)

	err = cmd.Start()
	check(err)

	in := bufio.NewScanner(stdout)

	for in.Scan() {
		// Uncomment if we want to include maven output in the logs
		//log.Printf(in.Text())
	}

	err = cmd.Wait()
	if err != nil {
		log.Println(err)
		return 1, err
	}
	return 0, nil
}
