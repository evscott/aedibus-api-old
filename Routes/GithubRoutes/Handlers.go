package GithubRoutes

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	consts "github.com/evscott/z3-e2c-api/shared"
	"github.com/google/go-github/github"
)

type Config struct {
	GAL *github.Client
}

func (c *Config) GetInfo(w http.ResponseWriter, r *http.Request) {}

func (c *Config) CreateRepository(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Getting hit")
	ctx := context.Background()

	repo := github.Repository{
		Name:          String("test"),
		DefaultBranch: String("master"),
	}

	if _, res, err := c.GAL.Repositories.Create(ctx, consts.Z3E2C, &repo); err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	} else {
		fmt.Printf("Repository created %v\n", res)
	}

	/***** Create master file *****/
	byteFile, err := ioutil.ReadFile("test.md") // b has type []byte
	if err != nil {
		log.Fatal(err)
	}

	message := "Adding README"

	content := github.RepositoryContentFileOptions{
		Message: &message,
		Content: byteFile,
	}

	if _, res, err := c.GAL.Repositories.CreateFile(ctx, consts.Z3E2C, "test", "README.md", &content); err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	} else {
		fmt.Printf("README created %v\n", res)
	}

	w.WriteHeader(200)
}

func (c *Config) CreateRef(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	/***** Get master reference *****/
	masterRef, res, err := c.GAL.Git.GetRef(ctx, consts.Z3E2C, "test", "refs/heads/master")
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	} else {
		fmt.Printf("Got master reference: %v\n", res)
	}

	/***** Create branch *****/
	branchName := "test-branch"

	reference := github.Reference{
		Ref: String(fmt.Sprintf("refs/heads/%s", branchName)),
		URL: String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s", consts.Z3E2C, "test", branchName)),
		Object: &github.GitObject{
			Type: String("commit"),
			SHA:  masterRef.Object.SHA,
			URL:  String(fmt.Sprintf("https://api.github.com/repos/z3-e2c/test/git/commits/%v", masterRef)),
		},
	}

	fmt.Printf("%v\n%v", c.GAL.Git, reference)

	if _, res, err := c.GAL.Git.CreateRef(ctx, consts.Z3E2C, "test", &reference); err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	} else {
		fmt.Printf("Reference created: %v\n", res)
	}
}

func String(s string) *string {
	return &s
}
