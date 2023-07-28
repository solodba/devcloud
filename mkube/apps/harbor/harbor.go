package harbor

// Projects结构体初始化函数
func NewProjects(curPage, pageSize int64) *Projects {
	page := NewPage()
	page.CurrentPage = curPage
	page.PageSize = pageSize
	return &Projects{
		Page: page,
		Data: make([]*ProjectData, 0),
	}
}

// Page初始化函数
func NewPage() *Page {
	return &Page{}
}

// ProjectData初始化函数
func NewProjectData() *ProjectData {
	return &ProjectData{
		CurrentUserRoleIds: make([]int64, 0),
		CveAllowlist: &CveAllowlist{
			Items: make([]string, 0),
		},
		Metadata: &Metadata{},
	}
}