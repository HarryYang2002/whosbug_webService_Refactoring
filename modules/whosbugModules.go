package modules

//这里是whosbug路由组所需的结构体

type Project struct {
	Pid string `form:"pid" json:"pid" binding:"required"`
}

type Release struct {
	Version    string `form:"version" json:"version" binding:"required"`
	CommitHash string `form:"commit_hash" json:"commit_hash" binding:"required"`
}

type Commit struct {
	Hash   string `form:"hash" json:"hash" binding:"required"`
	Email  string `form:"email" json:"email" binding:"required"`
	Author string `form:"author" json:"author" binding:"required"`
	Time   string `form:"time" json:"time" binding:"required"`
}

type T struct {
	Project Project
	Release Release
	// commit 暂定
}

type T2 struct {
	Project Project
	Release Release
	Commit  []Commit
}

type ReleaseModules struct {
	Version        string
	LastCommitHash string
	Project        int //or string?
}

type Object struct {
	Path          string
	ObjectId      string
	OwnerInfo     string
	OldConfidence float64
	StartLine     int
	EndLine       int
	//commit string ForeignKey 暂定
	parameters string
}
