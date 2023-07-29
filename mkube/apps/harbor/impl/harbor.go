package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/solodba/devcloud/mkube/apps/harbor"
)

// 获取Harbor Projects
func (i *impl) QueryProjects(ctx context.Context, in *harbor.QueryProjectsRequest) (*harbor.Projects, error) {
	projects := harbor.NewProjects(in.CurPage, in.PageSize)
	url := fmt.Sprintf("%s://%s%s?page=%d&page_size=%d", i.harbor.Scheme, i.harbor.Host, GET_PROJECTS, in.CurPage, in.PageSize)
	if in.Keyword != "" {
		url = fmt.Sprintf("%s&q=name=~%s", url, in.Keyword)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(i.harbor.Username, i.harbor.Password)
	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	xTotalCount := resp.Header.Get("x-total-count")
	newxTotalCount, err := strconv.Atoi(xTotalCount)
	if err != nil {
		return nil, err
	}
	projectDatas := make([]*harbor.ProjectData, newxTotalCount)
	err = json.Unmarshal(respBody, &projectDatas)
	if err != nil {
		return nil, err
	}
	projects.Data = projectDatas
	projects.Page.TotalPage = int64(math.Ceil(float64(newxTotalCount) / float64(in.PageSize)))
	projects.Page.TotalCount = int64(newxTotalCount)
	return projects, nil
}

// 获取Harbor Respositories
func (i *impl) QueryRepositories(ctx context.Context, in *harbor.QueryRepositoriesRequest) (*harbor.Repositories, error) {
	repositories := harbor.NewRepositories(in.CurPage, in.PageSize)
	path := fmt.Sprintf(GET_REPOSITORIES, in.ProjectName)
	url := fmt.Sprintf("%s://%s%s?page=%d&page_size=%d", i.harbor.Scheme, i.harbor.Host, path, in.CurPage, in.PageSize)
	if in.Keyword != "" {
		url = fmt.Sprintf("%s&q=name=~%s", url, in.Keyword)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(i.harbor.Username, i.harbor.Password)
	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	xTotalCount := resp.Header.Get("x-total-count")
	newXTotalCount, err := strconv.Atoi(xTotalCount)
	if err != nil {
		return nil, err
	}
	repositoryDatas := make([]*harbor.RepositoryData, newXTotalCount)
	err = json.Unmarshal(respBody, &repositoryDatas)
	if err != nil {
		return nil, err
	}
	repositories.Data = repositoryDatas
	repositories.Page.TotalPage = int64(math.Ceil(float64(newXTotalCount) / float64(in.PageSize)))
	repositories.Page.TotalCount = int64(newXTotalCount)
	return repositories, nil
}

// 获取Harbor Artifacts
func (i *impl) QueryArtifacts(ctx context.Context, in *harbor.QueryArtifactsRequest) (*harbor.Artifacts, error) {
	artifacts := harbor.NewArtifacts(in.CurPage, in.PageSize)
	path := fmt.Sprintf(GET_ARTIFACTS, in.ProjectName, in.RepositoryName)
	url := fmt.Sprintf("%s://%s%s?page=%d&page_size=%d", i.harbor.Scheme, i.harbor.Host, path, in.CurPage, in.PageSize)
	if in.Keyword != "" {
		url = fmt.Sprintf("%s&q=tags=~%s", url, in.Keyword)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(i.harbor.Username, i.harbor.Password)
	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	xTotalCount := resp.Header.Get("x-total-count")
	newXTotalCount, err := strconv.Atoi(xTotalCount)
	if err != nil {
		return nil, fmt.Errorf("解析x-total-count失败! 原因: %s", err.Error())
	}
	artifactDatas := make([]*harbor.ArtifactData, newXTotalCount)
	err = json.Unmarshal(respBody, &artifactDatas)
	if err != nil {
		return nil, err
	}
	for index := range artifactDatas {
		artifactDatas[index].Host = i.harbor.Host
	}
	artifacts.Data = artifactDatas
	artifacts.Page.TotalPage = int64(math.Ceil(float64(newXTotalCount) / float64(in.PageSize)))
	artifacts.Page.TotalCount = int64(newXTotalCount)
	return artifacts, nil
}

// 匹配镜像仓库
func (i *impl) QueryMatchImages(ctx context.Context, in *harbor.QueryMatchImagesRequest) (*harbor.MatchImages, error) {
	keywordArr := strings.Split(in.Keyword, ":")
	image := ""
	tag := ""
	if len(keywordArr) == 1 {
		image = keywordArr[0]
	} else {
		image = keywordArr[0]
		tag = keywordArr[1]
	}
	matchImages := harbor.NewMatchImages()
	req := harbor.NewQueryProjectsRequest()
	req.CurPage = 1
	req.PageSize = 20
	projects, err := i.QueryProjects(ctx, req)
	if err != nil {
		return nil, err
	}
	if projects.Page.TotalCount == 0 {
		return matchImages, nil
	}
	var countFlag int64
	for _, projectData := range projects.Data {
		req := harbor.NewQueryRepositoriesRequest()
		req.ProjectName = projectData.Name
		req.CurPage = 1
		req.PageSize = 10
		req.Keyword = image
		repositories, err := i.QueryRepositories(ctx, req)
		if err != nil {
			return nil, err
		}
		for _, repositoryData := range repositories.Data {
			repositoryName := filepath.Base(repositoryData.Name)
			req := harbor.NewQueryArtifactsRequest()
			req.ProjectName = projectData.Name
			req.RepositoryName = repositoryName
			req.CurPage = 1
			req.PageSize = 10
			req.Keyword = tag
			artifacts, err := i.QueryArtifacts(ctx, req)
			if err != nil {
				return nil, err
			}
			if artifacts.Page.TotalCount == 0 {
				continue
			}
			for _, artifact := range artifacts.Data {
				for _, tag := range artifact.Tags {
					imageInfo := fmt.Sprintf("%s/%s/%s:%s", artifact.Host, projectData.Name, repositoryName, tag.Name)
					matchImages.AddItems(imageInfo)
					countFlag++
					if countFlag == 10 {
						return matchImages, nil
					}
				}
			}
		}
	}
	return matchImages, nil
}
