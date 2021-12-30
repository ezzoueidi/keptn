package common

import (
	"errors"
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	common_mock "github.com/keptn/keptn/resource-service/common/fake"
	"github.com/keptn/keptn/resource-service/common_models"
	config2 "github.com/keptn/keptn/resource-service/config"
	. "gopkg.in/check.v1"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type BaseSuite struct {
	//Suite      fixtures.Suite
	Repository *git.Repository
	url        string
}

var _ = Suite(&BaseSuite{})

func (s *BaseSuite) SetUpSuite(c *C) {
	// init fixture repo
	//s.Suite.SetUpSuite(c)
	s.buildBasicRepository(c)
}

func (s *BaseSuite) TearDownSuite(c *C) {
	//s.Suite.TearDownSuite(c)
	err := os.RemoveAll("./debug")
	c.Assert(err, IsNil)
}

func (s *BaseSuite) SetUpTest(c *C) {
	s.SetUpSuite(c)
}

func (s *BaseSuite) buildBasicRepository(c *C) {
	err := os.RemoveAll("./debug")
	c.Assert(err, IsNil)
	//url := fixtures.ByURL("https://github.com/git-fixtures/basic.git").One().DotGit().Root()
	s.url = "./debug/remote"
	// make a local remote
	_, err = git.PlainClone(s.url, true, &git.CloneOptions{URL: "https://github.com/git-fixtures/basic.git"})
	c.Assert(err, IsNil)

	// make local git repo
	fs, err := memfs.New().Chroot(config2.ConfigDir + "/sockshop")
	s.Repository, err = git.Clone(memory.NewStorage(), fs, &git.CloneOptions{URL: s.url})
	c.Assert(err, IsNil)
}

func (s *BaseSuite) NewGitContext() common_models.GitContext {
	return common_models.GitContext{
		Project: "sockshop",
		Credentials: &common_models.GitCredentials{
			User:      "Me",
			Token:     "blabla",
			RemoteURI: s.url},
	}
}

func (s *BaseSuite) TestGit_CreateBranch(c *C) {

	tests := []struct {
		name         string
		gitContext   common_models.GitContext
		branch       string
		sourceBranch string
		wantErr      bool
		error        error
	}{
		{
			name:         "simple branch from master",
			gitContext:   s.NewGitContext(),
			branch:       "dev",
			sourceBranch: "master",
			wantErr:      false,
			error:        nil,
		},
		{
			name:         "add existing",
			gitContext:   s.NewGitContext(),
			branch:       "dev",
			sourceBranch: "master",
			wantErr:      true,
			error:        errors.New("branch already exists"),
		},
		{
			name:         "illegal add to non existing branch",
			gitContext:   s.NewGitContext(),
			branch:       "dev",
			sourceBranch: "refs/heads/branch",
			wantErr:      true,
			error:        errors.New("reference not found"),
		},
	}
	r := s.Repository
	g := Git{
		s.NewTestGit(),
	}

	expected := []byte("[core]\n\tbare = false\n[remote \"origin\"]\n\turl = " +
		"./debug/remote\n\tfetch = +refs/heads/*:refs/remotes/origin/*\n[branch \"dev\"]\n" +
		"\tremote = origin\n\tmerge = refs/heads/dev\n[branch \"master\"]\n" +
		"\tremote = origin\n\tmerge = refs/heads/master\n")

	for _, tt := range tests {
		c.Logf("Test: %s", tt.name)

		err := g.CreateBranch(tt.gitContext, tt.branch, tt.sourceBranch)

		if (err != nil) && tt.wantErr {
			c.Assert(err.Error(), Equals, tt.error.Error())
			continue
		}
		if err != nil {
			c.Errorf("CreateBranch() error = %v, wantErr %v", err, tt.wantErr)
		}

		// check git config files
		cfg, err := r.Config()
		c.Assert(err, IsNil)
		marshaled, err := cfg.Marshal()
		c.Assert(err, IsNil)
		c.Assert(string(expected), Equals, string(marshaled))
	}
}

func (s *BaseSuite) TestGit_CheckoutBranch(c *C) {

	tests := []struct {
		name       string
		gitContext common_models.GitContext
		branch     string
		wantErr    bool
	}{
		{
			name:       "checkout master branch full ref",
			gitContext: s.NewGitContext(),
			branch:     "refs/heads/master",
		},
		{
			name:       "checkout master branch",
			gitContext: s.NewGitContext(),
			branch:     "master",
		},
		{
			name:       "checkout not existing branch",
			gitContext: s.NewGitContext(),
			branch:     "refs/heads/dev",
			wantErr:    true,
		},
	}
	g := Git{s.NewTestGit()}
	for _, tt := range tests {
		c.Log("Test: ", tt.name)
		if err := g.CheckoutBranch(tt.gitContext, tt.branch); (err != nil) != tt.wantErr {
			c.Errorf("CheckoutBranch() error = %v, wantErr %v", err, tt.wantErr)
		}

	}
}

func (s *BaseSuite) TestGit_GetFileRevision(c *C) {

	tests := []struct {
		name       string
		gitContext common_models.GitContext
		file       string
		content    string
		wantErr    bool
	}{
		{
			name:       "get from commitID",
			gitContext: s.NewGitContext(),
			file:       "foo/example.go",
			content:    "ciao",
			wantErr:    false,
		},
	}
	for _, tt := range tests {

		g := Git{
			s.NewTestGit(),
		}
		id := s.commitAndPush(tt.file, tt.content, c)
		got, err := g.GetFileRevision(tt.gitContext, id.String(), tt.file)
		if (err != nil) != tt.wantErr {
			c.Errorf("GetFileRevision() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		b := []byte(fmt.Sprintf("%s", tt.content))
		if !reflect.DeepEqual(got, b) {
			c.Errorf("GetFileRevision() got = %v, want %v", got, b)
		}

	}
}

func (s *BaseSuite) NewTestGit() *common_mock.GogitMock {

	return &common_mock.GogitMock{
		PlainCloneFunc: func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
			return s.Repository, nil
		},
		PlainInitFunc: nil,
		PlainOpenFunc: func(path string) (*git.Repository, error) {
			return s.Repository, nil //git.PlainOpen(path)
		},
	}
}

func (s *BaseSuite) commitAndPush(file string, content string, c *C) plumbing.Hash {
	r := s.Repository
	w, err := r.Worktree()
	f, err := w.Filesystem.Create(file)
	c.Assert(err, IsNil)
	f.Write([]byte(fmt.Sprintf("%s", content)))
	f.Close()

	_, err = w.Add(file)
	c.Assert(err, IsNil)

	id, err := w.Commit("added a file",
		&git.CommitOptions{
			All: true,
			Author: &object.Signature{
				Name:  "Test Create Branch",
				Email: "createBranch@gogit-test.com",
				When:  time.Now(),
			},
		})

	c.Assert(err, IsNil)
	//push to repo
	err = r.Push(&git.PushOptions{
		//Force: true,
		Auth: &http.BasicAuth{
			Username: "whatever",
			Password: "whatever",
		}})
	c.Assert(err, IsNil)

	return id
}
