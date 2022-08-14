package modules

//这里是whosbug路由组所需的结构体

type Project struct {
	Pid string `form:"pid" json:"pid" binding:"required"`
}

type Release struct {
	Version    string `form:"version" json:"version" binding:"required"`
	CommitHash string `form:"last_commit_hash" json:"last_commit_hash" binding:"required"`
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
	Project Project  `json:"project"`
	Release Release  `json:"release"`
	Commit  []Commit `json:"commits"`
}

type T3 struct {
	Project Project  `json:"project"`
	Release Release  `json:"release"`
	Object  []Object `json:"objects"`
}

type ReleaseModules struct {
	Version        string
	LastCommitHash string
	Project        int //or string?
}

type UncountedObject struct {
	Hash             string `form:"hash" json:"hash" binding:"required"`
	ObjectId         string `form:"object_id" json:"pbject_id" binding:"required"`
	OldObjectId      string `form:"old_object_id " json:"old_object_id " binding:"required"`
	Parameters       string `form:"parameters " json:"parameters " binding:"required"`
	StartLine        int    `form:"start_line   " json:"start_line   " binding:"required"`
	Path             string `form:"path   " json:"path   " binding:"required"`
	EndLine          int    `form:"end_line   " json:"end_line   " binding:"required"`
	OldLineCount     int    `form:"old_line_count   " json:"old_line_count   " binding:"required"`
	NewLineCount     int    `form:"new_line_count   " json:"new_line_count   " binding:"required"`
	DeletedLineCount int    `form:"deleted_line_count   " json:"deleted_line_count   " binding:"required"`
	AddedLineCount   int    `form:"added_line_count    " json:"added_line_count    " binding:"required"`
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

type T4 struct {
	Project         Project           `json:"project"`
	Release         Release           `json:"release"`
	UncountedObject []UncountedObject `json:"uncountedObjects"`
}
