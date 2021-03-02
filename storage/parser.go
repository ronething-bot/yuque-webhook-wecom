package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"yuque-webhook-wecom/param"

	"github.com/pkg/errors"
	"github.com/wujiyu115/yuqueg"
)

//yuque parser

const (
	Publish            = "publish" //新建
	CommentCreate      = "comment_create"
	CommentReplyCreate = "comment_reply_create"
	YuquePrefix        = "https://www.yuque.com/"
)

type YuqueParser struct {
	inotify INotify
	service *yuqueg.Service
}

func NewYuqueParser(inotify INotify, service *yuqueg.Service) *YuqueParser {
	return &YuqueParser{inotify: inotify, service: service}
}

func (y *YuqueParser) Parser(src map[string]interface{}) (err error) {
	data, ok := src["data"].(map[string]interface{})
	if !ok {
		return errors.New("解析 data 失败")
	}
	webhookType, ok := data["action_type"].(string)
	log.Printf("webhookType is %s\n", webhookType)
	if !ok {
		return errors.New("解析 action type 失败")
	}
	switch webhookType {
	case Publish:
		var newTopic param.NewTopic
		if err = map2struct(src, &newTopic); err != nil {
			return err
		}
		y.Publish(newTopic)
	case CommentCreate:
		var newComment param.NewComment
		if err = map2struct(src, &newComment); err != nil {
			return err
		}
		y.CommentCreate(newComment)
	case CommentReplyCreate:
		var newCommentReply param.NewCommentReply
		if err = map2struct(src, &newCommentReply); err != nil {
			return err
		}
		y.CommentReply(newCommentReply)
	default:
		return errors.New("暂不支持此类型")
	}

	return nil
}

/*
事件：Event

> 作者：atUserName
> 详情：[Title](Url)
> 时间: Time
*/

type messageTemplate struct {
	event      string
	atUserName string
	repo       string
	title      string
	url        string
	time       string
}

func (y *YuqueParser) getNotifyContent(template messageTemplate) string {
	return fmt.Sprintf(`
事件: %s

> 作者: %s
> 仓库: %s
> 详情: [%s](%s)
> 时间: %s
`, template.event, template.atUserName, template.repo, template.title, template.url, template.time)
}

func (y *YuqueParser) Publish(topic param.NewTopic) {
	// 需要对 user id 进行解析 因为此时 user 获取到的是 group
	var atUserName string
	userInfo, err := y.service.User.Get(fmt.Sprintf("%d", topic.Data.UserID))
	if err != nil {
		log.Printf("解析用户出现错误, err: %v\n", err)
		atUserName = "@unknown(unknown)"
	} else {
		// :name(:login)
		atUserName = fmt.Sprintf("@%s(%s)", userInfo.Data.Name, userInfo.Data.Login)
	}
	y.inotify.Notify(y.getNotifyContent(messageTemplate{
		event:      "新增话题",
		atUserName: atUserName,
		repo:       topic.Data.Book.Name,
		title:      topic.Data.Title,
		url:        fmt.Sprintf("%s%s", YuquePrefix, topic.Data.Path),
		time:       topic.Data.CreatedAt.In(time.Local).String(),
	}))
}

func (y *YuqueParser) CommentCreate(comment param.NewComment) {
	atUserName := fmt.Sprintf("@%s(%s)", comment.Data.User.Name, comment.Data.User.Login)
	fmt.Printf("time is %v\n", comment.Data.CreatedAt.In(time.Local).String())
	y.inotify.Notify(y.getNotifyContent(messageTemplate{
		event:      "新增评论",
		atUserName: atUserName,
		repo:       comment.Data.Commentable.Book.Name,
		title:      comment.Data.Commentable.Title,
		url:        fmt.Sprintf("%s%s", YuquePrefix, comment.Data.Path),
		time:       comment.Data.CreatedAt.In(time.Local).String(),
	}))
}

func (y *YuqueParser) CommentReply(reply param.NewCommentReply) {
	atUserName := fmt.Sprintf("@%s(%s)", reply.Data.User.Name, reply.Data.User.Login)
	y.inotify.Notify(y.getNotifyContent(messageTemplate{
		event:      "新增评论回复",
		atUserName: atUserName,
		repo:       reply.Data.Commentable.Book.Name,
		title:      reply.Data.Commentable.Title,
		url:        fmt.Sprintf("%s%s", YuquePrefix, reply.Data.Path),
		time:       reply.Data.CreatedAt.In(time.Local).String(),
	}))
}

func map2struct(src map[string]interface{}, dst interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return errors.Wrap(err, "序列化失败")
	}
	fmt.Printf("b is %v\n", string(b))
	if err = json.Unmarshal(b, dst); err != nil {
		return errors.Wrap(err, "反序列化失败")
	}
	return nil
}
