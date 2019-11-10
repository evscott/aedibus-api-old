package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/evscott/z3-e2c-api/models"
	"github.com/google/go-github/github"
)

type Config struct {
	GAL *github.Client
}

func (c *Config) UpdateFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	buffer := bytes.Buffer{}

	/***** Unpack request *****/

	// Unpack file metadata
	repo := r.FormValue("repo")
	branch := r.FormValue("branch")
	fileName := r.FormValue("fileName")

	if repo == "" || branch == "" || fileName == "" {
		log.Fatal("Must include form values for repo, branch, and fileName")
	}

	// Unpack file to upload
	file, header, err := r.FormFile(fileName)
	if err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		name := strings.Split(header.Filename, ".")
		fmt.Printf("Received file: %s\n", name[0])
		defer file.Close()
	}

	if _, err := io.Copy(&buffer, file); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	}
	contents := buffer.Bytes()
	buffer.Reset()

	/***** Get blob sha of file from Github *****/
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", branch)}

	var sha string
	if contents, _, res, err := c.GAL.Repositories.GetContents(ctx, Z3E2C, repo, fileName, &getOptions); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		fmt.Printf("Got sha for file %s %v\n", fileName, res)
		sha = *contents.SHA
	}

	/***** Upload file to Github *****/
	fileOptions := github.RepositoryContentFileOptions{
		Message: String(UpdatingFile),
		Content: contents,
		Branch:  &branch,
		SHA:     &sha,
	}

	if _, res, err := c.GAL.Repositories.UpdateFile(ctx, Z3E2C, repo, fileName, &fileOptions); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		fmt.Printf("File %s uploaded to branch %s %v\n", fileName, branch, res)
	}

	w.WriteHeader(Status(OK))
}

func (c *Config) UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	buffer := bytes.Buffer{}

	/***** Unpack request *****/

	// Unpack file metadata
	repo := r.FormValue("repo")
	branch := r.FormValue("branch")
	fileName := r.FormValue("fileName")

	if repo == "" || branch == "" || fileName == "" {
		log.Fatal("Must include form values for repo, branch, and fileName")
	}

	// Unpack file to upload
	file, header, err := r.FormFile(fileName)
	if err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		name := strings.Split(header.Filename, ".")
		fmt.Printf("Received file: %s\n", name[0])
		defer file.Close()
	}

	if _, err := io.Copy(&buffer, file); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	}
	contents := buffer.Bytes()
	buffer.Reset()

	/***** Upload file to Github *****/
	fileOptions := github.RepositoryContentFileOptions{
		Message: String(UploadingFile),
		Content: contents,
		Branch:  &branch,
	}

	if _, res, err := c.GAL.Repositories.CreateFile(ctx, Z3E2C, repo, fileName, &fileOptions); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		fmt.Printf("File %s uploaded to branch %s %v\n", fileName, branch, res)
	}

	w.WriteHeader(Status(OK))
}

func (c *Config) CreateRepository(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	/***** Unpack create repository request *****/
	req := &models.ReqCreateRepo{}
	ParseReqJsonBody(req, w, r)
	if req.Repo == "" {
		w.WriteHeader(Status(OK))
		log.Fatalf("Invalid request: %v\n", req)
	}

	/***** Create repository *****/
	repo := github.Repository{
		Name:          &req.Repo,
		DefaultBranch: String(MASTER),
	}
	if _, res, err := c.GAL.Repositories.Create(ctx, Z3E2C, &repo); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		fmt.Printf("Repository %s created %v\n", req.Repo, res)
	}

	w.WriteHeader(Status(OK))
}

func (c *Config) CreateReference(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	/***** Unpack create reference request *****/
	req := &models.ReqCreateRef{}
	ParseReqJsonBody(req, w, r)
	if req.Repo == "" || req.Branch == "" {
		w.WriteHeader(Status(InternalServerError))
		log.Fatalf("Invalid request: %v\n", req)
	}

	/***** Get master reference *****/
	masterRef, res, err := c.GAL.Git.GetRef(ctx, Z3E2C, req.Repo, fmt.Sprintf("refs/heads/%s", MASTER))
	if err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		fmt.Printf("Got master reference: %v\n", res)
	}

	/***** Create branch *****/
	reference := github.Reference{
		Ref: String(fmt.Sprintf("refs/heads/%s", req.Branch)),
		URL: String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s", Z3E2C, req.Repo, req.Branch)),
		Object: &github.GitObject{
			Type: String("commit"),
			SHA:  masterRef.Object.SHA,
			URL:  String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s", Z3E2C, req.Repo, masterRef)),
		},
	}

	if _, res, err := c.GAL.Git.CreateRef(ctx, Z3E2C, req.Repo, &reference); err != nil {
		w.WriteHeader(Status(InternalServerError))
		log.Fatal(err)
	} else {
		fmt.Printf("Reference %s/%s created: %v\n", req.Repo, req.Branch, res)
	}

	w.WriteHeader(Status(OK))
}
