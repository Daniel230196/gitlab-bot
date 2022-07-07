package gitlab

import (
	"fmt"
	"strconv"
	"strings"
)

type MrEvent struct {
	authorId    int
	description string
	source      string
	target      string
	title       string
	state       string
	id          int
	url         string
}

type MergeRequestData struct {
	AuthorId int
	Source   string
	Target   string
}

func createEventByType(eType string, attributes EventAttributes) Event {
	switch eType {
	case "merge_request":
		return MrEvent{
			attributes.MrAuthor,
			attributes.Description,
			attributes.Source,
			attributes.Target,
			attributes.Title,
			attributes.State,
			attributes.Id,
			attributes.Url,
		}
	case "note":
		return CommentEvent{
			attributes.Id,
			attributes.Type,
			attributes.Description,
			attributes.Url,
			MergeRequestData{
				attributes.MrAuthor,
				attributes.Source,
				attributes.Target,
			},
		}
	default:
		return nil
	}
}

func (e MrEvent) prepareMessage() string {
	message := fmt.Sprintf("Merge request *%s* *_%s_* FROM *%s* TO *%s*\n", e.title, e.state, e.source, e.target)
	url := strings.Replace(e.url, "_", "\\_", -1)
	url = strings.Replace(url, "*", "\\*", -1)
	message = message + "*URL*:\t" + url
	return message
}

func (e MrEvent) getAuthorId() string {
	return strconv.Itoa(e.authorId)
}

type CommentEvent struct {
	authorId     int
	noteableType string
	note         string
	url          string
	mergeRequest MergeRequestData
}

func (e CommentEvent) getAuthorId() string {
	return strconv.Itoa(e.mergeRequest.AuthorId)
}

func (e CommentEvent) prepareMessage() string {
	return fmt.Sprintf("You have new comment in MR from %d\n*Comment*:%s\nURL:%s", e.authorId, e.note, e.url)
}
