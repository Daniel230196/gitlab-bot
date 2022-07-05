package gitlab

type GitlabHook struct {
	EventType        string           `json:"event_type"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	Repositoty       Repositoty       `json:"repositoty"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Repositoty struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ObjectAttributes struct {
	Id           int    `json:"id"`
	TargetBranch string `json:"target_branch"`
	SourceBranch string `json:"source_branch"`
	AuthorId     int    `json:"author_id"`
	AssigneeId   int    `json:"assignee_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	State        string `json:"state"`
}
