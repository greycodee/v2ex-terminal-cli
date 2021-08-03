package api

import (
	"encoding/json"
	"github.com/greycodee/v2ex-terminal-cli/urls"
	"io"
	"io/ioutil"
	"net/http"
)

type TopicInfo struct {
	Node struct {
		AvatarLarge      string        `json:"avatar_large"`
		Name             string        `json:"name"`
		AvatarNormal     string        `json:"avatar_normal"`
		Title            string        `json:"title"`
		URL              string        `json:"url"`
		Topics           int           `json:"topics"`
		Footer           string        `json:"footer"`
		Header           string        `json:"header"`
		TitleAlternative string        `json:"title_alternative"`
		AvatarMini       string        `json:"avatar_mini"`
		Stars            int           `json:"stars"`
		Aliases          []interface{} `json:"aliases"`
		Root             bool          `json:"root"`
		ID               int           `json:"id"`
		ParentNodeName   string        `json:"parent_node_name"`
	} `json:"node"`
	Member struct {
		Username     string `json:"username"`
		Website      string `json:"website"`
		Github       string `json:"github"`
		Psn          string `json:"psn"`
		AvatarNormal string `json:"avatar_normal"`
		Bio          string `json:"bio"`
		URL          string `json:"url"`
		Tagline      string `json:"tagline"`
		Twitter      string `json:"twitter"`
		Created      int    `json:"created"`
		AvatarLarge  string `json:"avatar_large"`
		AvatarMini   string `json:"avatar_mini"`
		Location     string `json:"location"`
		Btc          string `json:"btc"`
		ID           int    `json:"id"`
	} `json:"member"`
	LastReplyBy     string `json:"last_reply_by"`
	LastTouched     int    `json:"last_touched"`
	Title           string `json:"title"`
	URL             string `json:"url"`
	Created         int    `json:"created"`
	Content         string `json:"content"`
	ContentRendered string `json:"content_rendered"`
	LastModified    int    `json:"last_modified"`
	Replies         int    `json:"replies"`
	ID              int    `json:"id"`
}

func GetTopicInfo(id string) (topicInfo TopicInfo,err error)  {
	url := urls.TopicShow+"?id="+id
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body,err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	list := make([]TopicInfo,0)
	err = json.Unmarshal(body, &list)
	if err!=nil {
		return
	}
	topicInfo = list[0]
	return
}