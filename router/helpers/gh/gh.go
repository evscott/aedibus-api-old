package gh

import (
	"context"
	"fmt"
	"github.com/evscott/aedibus-api/models"
	consts "github.com/evscott/aedibus-api/shared/constants"
	"github.com/evscott/aedibus-api/shared/utils"
	"github.com/google/go-github/github"
)

// TODO
//
func (c *Config) CreateComment(ctx context.Context, fileName, assignmentName, commitID, body string, pullRequestNumber, position int) (*github.PullRequestComment, error) {
	comment := github.PullRequestComment{
		Path:     &fileName,
		CommitID: &commitID,
		Body:     &body,
		Position: &position,
	}
	res, _, err := c.gal.PullRequests.CreateComment(ctx, consts.AEDIBUS, assignmentName, pullRequestNumber, &comment)
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
	res, _, err := c.gal.PullRequests.Create(ctx, consts.AEDIBUS, assignmentName, &pullRequest)
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
	if _, _, err := c.gal.Repositories.Create(ctx, consts.AEDIBUS, &repo); err != nil {
		return err
	}
	return nil
}

// TODO
//
func (c *Config) CreateFile(ctx context.Context, assignmentName, dropboxName, fileName string, contents []byte) (*github.RepositoryContentResponse, error) {
	fileOptions := github.RepositoryContentFileOptions{
		Message: utils.String("Uploading file"),
		Content: contents,
		Branch:  &dropboxName,
	}
	res, _, err := c.gal.Repositories.CreateFile(ctx, consts.AEDIBUS, assignmentName, fileName, &fileOptions)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TODO
//
func (c *Config) UpdateFile(ctx context.Context, assignmentName, dropboxName, fileName string, newContents []byte) (*github.RepositoryContentResponse, error) {
	// Get blob sha of file from GithubHelpers to be used as target of update
	var sha string
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", dropboxName)}
	oldContents, _, _, err := c.gal.Repositories.GetContents(ctx, consts.AEDIBUS, assignmentName, fileName, &getOptions)
	if err != nil {
		return nil, err
	}
	sha = *oldContents.SHA

	// Upload file to GithubHelpers
	fileOptions := github.RepositoryContentFileOptions{
		Message: utils.String("Updating file"), // TODO
		Content: newContents,
		Branch:  &dropboxName,
		SHA:     &sha,
	}
	res, _, err := c.gal.Repositories.UpdateFile(ctx, consts.AEDIBUS, assignmentName, fileName, &fileOptions)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// TODO
//
//
func (c *Config) GetReadme(ctx context.Context, assignmentName string) (*models.ResGetFile, error) {
	// Get blob sha of file from GithubHelpers to be used as target of update
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", consts.MASTER)}
	fileContent, _, err := c.gal.Repositories.GetReadme(ctx, consts.AEDIBUS, assignmentName, &getOptions)
	if err != nil {
		return nil, err
	}
	content, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}

	res := &models.ResGetFile{
		AssignmentName: assignmentName,
		FileName:       consts.README,
		DropboxName:    consts.MASTER,
		Content:        content,
	}

	return res, nil
}

// TODO
//
//
func (c *Config) GetFileContents(ctx context.Context, fileName, assignmentName string) (*models.ResGetFile, error) {
	// Get blob sha of file from GithubHelpers to be used as target of update
	getOptions := github.RepositoryContentGetOptions{Ref: fmt.Sprintf("heads/%s", consts.MASTER)}
	fileContent, _, _, err := c.gal.Repositories.GetContents(ctx, consts.AEDIBUS, assignmentName, consts.MASTER, &getOptions)
	if err != nil {
		return nil, err
	}
	content, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}

	res := &models.ResGetFile{
		FileName:    fileName,
		DropboxName: consts.MASTER,
		Content:     content,
	}

	return res, nil
}

// TODO
//
func (c *Config) CreateDropbox(ctx context.Context, dropboxName, assignmentName string) error {
	masterBranch, _, err := c.gal.Git.GetRef(ctx, consts.AEDIBUS, assignmentName, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		return err
	}

	reference := github.Reference{
		Ref: utils.String(fmt.Sprintf("refs/heads/%s", dropboxName)),
		URL: utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/%s", consts.AEDIBUS, assignmentName, dropboxName)),
		Object: &github.GitObject{
			Type: utils.String("commit"),
			SHA:  masterBranch.Object.SHA,
			URL:  utils.String(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s", consts.AEDIBUS, assignmentName, consts.MASTER)),
		},
	}

	if _, _, err := c.gal.Git.CreateRef(ctx, consts.AEDIBUS, assignmentName, &reference); err != nil {
		return err
	}
	return nil
}

// TODO
//
func (c *Config) GetMasterBlobSha(ctx context.Context, assignmentName string) (*string, error) {
	masterBranch, _, err := c.gal.Git.GetRef(ctx, consts.AEDIBUS, assignmentName, fmt.Sprintf("refs/heads/%s", consts.MASTER))
	if err != nil {
		return nil, err
	}
	return masterBranch.Object.SHA, nil
}

// TODO
//
func (c *Config) GetPrComments(ctx context.Context, assignmentName string, prNumber int) ([]*github.PullRequestComment, error) {
	options := &github.PullRequestListCommentsOptions{}
	comments, _, err := c.gal.PullRequests.ListComments(ctx, consts.AEDIBUS, assignmentName, prNumber, options)
	return comments, err
}
