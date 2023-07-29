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

// Repositories初始化函数
func NewRepositories(curPage, pageSize int64) *Repositories {
	page := NewPage()
	page.CurrentPage = curPage
	page.PageSize = pageSize
	return &Repositories{
		Page: page,
		Data: make([]*RepositoryData, 0),
	}
}

// RepositoryData初始化函数
func NewRepositoryData() *RepositoryData {
	return &RepositoryData{}
}

// Artifacts初始化函数
func NewArtifacts(curPage, pageSize int64) *Artifacts {
	page := NewPage()
	page.CurrentPage = curPage
	page.PageSize = pageSize
	return &Artifacts{
		Page: page,
		Data: make([]*ArtifactData, 0),
	}
}

// ArtifactData初始化函数
func NewArtifactData() *ArtifactData {
	return &ArtifactData{
		Accessories:   &Accessories{},
		AdditionLinks: &AdditionLinks{},
		ExtraAttrs: &ExtraAttrs{
			Config: &Config{
				Entrypoint: make([]string, 0),
				Env:        make([]string, 0),
				ExposedPorts: &ExposedPorts{
					Tcp: &Tcp{},
				},
				Labels: &Labels{},
			},
		},
		Labels:     &ArtifactDataLabels{},
		References: &References{},
		Tags:       make([]*Tags, 0),
	}
}

// MatchImages初始化函数
func NewMatchImages() *MatchImages {
	return &MatchImages{
		Images: make([]string, 0),
	}
}

// MatchImages添加方法
func (m *MatchImages) AddItems(items ...string) {
	m.Images = append(m.Images, items...)
}
