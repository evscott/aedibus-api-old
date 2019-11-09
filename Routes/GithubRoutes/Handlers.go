package GithubRoutes

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

type Config struct {
	GAL *github.Client
}

func (c *Config) Test(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	fmt.Printf("Getting branch")

	//name := "test-project"
	//
	//repo := github.Repository{
	//	Name: &name,
	//}
	//
	//if repo, res, err := c.GAL.Repositories.Create(ctx, "sakjfh", &repo); err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Printf("Repo created: %v\n", repo)
	//	fmt.Printf("Res: %v\n", res)
	//}
	//
	b, err := ioutil.ReadFile("README.md") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	message := "Testing uploading a file through go api"
	branch := "test2"

	options := github.RepositoryContentFileOptions{
		Message:   &message,
		Content:   b,
		SHA:       nil,
		Branch:    &branch,
		Author:    nil,
		Committer: nil,
	}

	contents, res, err := c.GAL.Repositories.CreateFile(ctx, "sakjfh", "test-project", "READMEeeee.md", &options)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Contents: %v\n", contents)
		fmt.Printf("Response: %v\n", res)
	}

	fmt.Printf("Got branch")
}
