package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

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
