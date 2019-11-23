package gh

import (
	"context"
	"fmt"

	"github.com/evscott/z3-e2c-api/models"
	consts "github.com/evscott/z3-e2c-api/shared/constants"
	"github.com/evscott/z3-e2c-api/shared/utils"
	"github.com/google/go-github/github"
)

type Config struct {
	gal *github.Client
}

func Init(gal *github.Client) *Config {
	return &Config{
		gal: gal,
	}
}

// TODO
//
func (c *Config) CreateComment(ctx context.Context, fileName, assignmentName, commitID, body string, position int) (*github.PullRequestComment, error) {
	comment := github.PullRequestComment{
		Path:     &fileName,
		CommitID: &commitID,
		Body:     &body,
		Position: &position,
	}
	res, _, err := c.gal.PullRequests.CreateComment(ctx, consts.Z3E2C, assignmentName, 1, &comment)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TODO
//
func (c *Config) CreatePullRequest(ctx context.Context, dropboxName, assignmentName, title, body string) (*github.PullRequest, error) {
	pullRequest := github.NewPullRequest{
		Title:               &title,
		Head:                &dropboxName,
		Base:                utils.String(consts.MASTER),
		Body:                &body,
		Issue:               nil,
		MaintainerCanModify: utils.Bool(true),
	}
	res, _, err := c.gal.PullRequests.Create(ctx, consts.Z3E2C, assignmentName, &pullRequest)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// TODO
//
func (c *Config) CreateRepository(ctx context.Context, assignmentName string) error {
	repo := github.Repository{
		Name:          &assignmentName,
		DefaultBranch: utils.String(consts.MASTER),
	}
	if _, _, err := c.gal.Repositories.Create(ctx, consts.Z3E2C, &repo); err != nil {
		return err
	}
	return nil
}

// TODO
//
func (c *Config) CreateFile(ctx context.Context, assignmentName, dropboxName, fileName string, contents []byte) error {
	fileOptions := github.RepositoryContentFileOptions{
		Message: utils.String("Uploading file"),
		Content: contents,
		Branch:  &dropboxName,
	}
	if _, _, err := c.gal.Repositories.CreateFile(ctx, consts.Z3E2C, assignmentName, fileName, &fileOptions); err != nil {
		return err
	}
	return nil
}

// TODO
//
//
func (c *Config) GetReadme(ctx context.Context, assignmentName, dropboxName string) (*models.ResGetFile, error) {
	// Get blob sha of file from GithubHelpers to be used as target of update
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", dropboxName)}
	fileContent, _, err := c.gal.Repositories.GetReadme(ctx, consts.Z3E2C, assignmentName, &getOptions)
	if err != nil {
		return nil, err
	}
	content, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}

	res := &models.ResGetFile{
		FileName:    assignmentName,
		DropboxName: dropboxName,
		Content:     content,
	}

	return res, nil
}

// TODO
//
//
func (c *Config) GetFileContents(ctx context.Context, assignmentName, dropboxName string) (*models.ResGetFile, error) {
	// Get blob sha of file from GithubHelpers to be used as target of update
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", dropboxName)}
	fileContent, _, _, err := c.gal.Repositories.GetContents(ctx, consts.Z3E2C, assignmentName, dropboxName, &getOptions)
	if err != nil {
		return nil, err
	}
	content, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}

	res := &models.ResGetFile{
		FileName:    assignmentName,
		DropboxName: dropboxName,
		Content:     content,
	}

	return res, nil
}

// TODO
//
func (c *Config) UpdateFile(ctx context.Context, assignmentName, dropboxName, fileName string, newContents []byte) error {
	// Get blob sha of file from GithubHelpers to be used as target of update
	var sha string
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", dropboxName)}
	oldContents, _, _, err := c.gal.Repositories.GetContents(ctx, consts.Z3E2C, assignmentName, fileName, &getOptions)
	if err != nil {
		return err
	}
	sha = *oldContents.SHA

	// Upload file to GithubHelpers
	fileOptions := github.RepositoryContentFileOptions{
		Message: utils.String("Updating file"), // TODO
		Content: newContents,
		Branch:  &dropboxName,
		SHA:     &sha,
	}
	if _, _, err := c.gal.Repositories.UpdateFile(ctx, consts.Z3E2C, assignmentName, fileName, &fileOptions); err != nil {
		return err
	}
	return nil
}

// TODO
//
func (c *Config) CreateDropbox(ctx context.Context, dropboxName, assignmentName string) error {
	masterBranch, _, err := c.gal.Git.GetRef(ctx, consts.Z3E2C, assignmentName, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		return err
	}

	reference := github.Reference{
		Ref: utils.String(fmt.Sprintf("refs/heads/%s", dropboxName)),
		URL: utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s", consts.Z3E2C, assignmentName, dropboxName)),
		Object: &github.GitObject{
			Type: utils.String("commit"),
			SHA:  masterBranch.Object.SHA,
			URL:  utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s", consts.Z3E2C, assignmentName, consts.MASTER)),
		},
	}

	if _, _, err := c.gal.Git.CreateRef(ctx, consts.Z3E2C, assignmentName, &reference); err != nil {
		return err
	}
	return nil
}

// TODO
//
func (c *Config) GetMasterBlobSha(ctx context.Context, assignmentName string) (*string, error) {
	masterBranch, _, err := c.gal.Git.GetRef(ctx, consts.Z3E2C, assignmentName, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		return nil, err
	}
	return masterBranch.Object.SHA, nil
}

// TODO
//
func (c *Config) GetPrComments(ctx context.Context, assignmentName, dropboxName string) error {
	return nil
}
