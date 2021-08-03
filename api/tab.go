package api

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/greycodee/v2ex-terminal-cli/urls"
	"io"
	"log"
	"net/http"
	"strings"
)

type Topic struct {
	Id              string 	`json:"id" comment:"主题ID"`
	Member          string `json:"member_name" comment:"会员名"`
	MemberAvatar    string `json:"member_avatar" comment:"会员头像"`
	Title           string `json:"title" comment:"标题"`
	Node            string `json:"node" comment:"节点"`
	NodeAlias       string `json:"node_alias" comment:"节点名称"`
	LastReplyMember string `json:"last_reply_member" comment:"最后回复会员名"`
	LastReplyTime   string `json:"last_reply_time" comment:"最后回复时间"`
	ReplyCount      string `json:"reply_count" comment:"回复总数"`
}

func TabTopicList(url urls.ApiUrl) (topics []Topic, err error) {
	res, err := http.Get(string(url))
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	doc.Find(".item").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(i int, tr *goquery.Selection) {
			topic := Topic{}
			// 主题ID
			topicLink,_ := tr.Find(".topic-link").Attr("href")
			topicId := strings.Split(topicLink,"#")[0][3:]
			topic.Id = topicId

			// 头像
			avatar,_ := tr.Find(".avatar").Attr("src")
			topic.MemberAvatar = avatar

			// 标题
			title := tr.Find(".topic-link ").Text()
			topic.Title = title

			// 节点名
			nodeAlias := tr.Find(".node").Text()
			topic.NodeAlias = nodeAlias

			// 节点
			node, _ := tr.Find(".node").Attr("href")
			topic.Node = node[4:]

			// 会员信息
			memberInfo := tr.Find(".topic_info strong a")
			// 主题发表会员
			topicMember := memberInfo.Get(0).FirstChild.Data
			topic.Member = topicMember

			if memberInfo.Size() >= 2 {
				// 最后回复会员
				lastMember := memberInfo.Get(1).FirstChild.Data
				topic.LastReplyMember = lastMember
			}

			// 最后修改时间
			lastTime := tr.Find(".topic_info span").Text()
			topic.LastReplyTime = lastTime

			replyCount := tr.Find(".count_livid").Text()
			topic.ReplyCount = replyCount

			topics = append(topics,topic)
		})
	})
	return
}

