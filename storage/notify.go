package storage

// wecom notify

import (
	"fmt"
	"log"
	"time"

	"github.com/imroc/req"
)

const wecomBase = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="

type weComSendTextContent struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

type MardkownContent struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}
type Markdown struct {
	Content string `json:"content"`
}

type NewsContent struct {
	Msgtype string `json:"msgtype"`
	News    News   `json:"news"`
}
type Articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
}
type News struct {
	Articles []Articles `json:"articles"`
}

type WeCom struct {
	Token  string
	ReqUrl string
}

//NewWeCom
func NewWeCom(token string) *WeCom {
	return &WeCom{Token: token, ReqUrl: wecomBase + token}
}

//type Message struct {
//	Title  string
//	Desc   string
//	Url    string
//	PicUrl string
//}

type INotify interface {
	Notify(content string)
}

func (w *WeCom) Notify(content string) {
	retry := 3
	for i := 0; i < retry; i++ {
		err := w.SendMarkdown(content)
		if err != nil {
			log.Printf("send markdown msg err: %v, then retry: %d\n", err, i)
			time.Sleep(time.Second)
			continue
		}
		goto End
	}
	log.Printf("发生不成功, 并且重试次数已用完\n")
	return
End:
	fmt.Printf("消息发送成功\n")
}

//func (w *WeCom) Notify(message Message) {
//	retry := 3
//	for i := 0; i < retry; i++ {
//		err := w.SendNews(message.Title, message.Desc, message.Url, message.PicUrl)
//		if err != nil {
//			log.Printf("send news err: %v, then retry: %d\n", err, i)
//			time.Sleep(time.Second)
//			continue
//		}
//		goto End
//	}
//	log.Printf("发生不成功, 并且重试次数已用完\n")
//	return
//End:
//	fmt.Printf("消息发送成功\n")
//}

func (w *WeCom) SendText(content string) error {
	resp, err := req.Post(w.ReqUrl, req.BodyJSON(&weComSendTextContent{
		Msgtype: "text",
		Text: struct {
			Content             string   `json:"content"`
			MentionedList       []string `json:"mentioned_list"`
			MentionedMobileList []string `json:"mentioned_mobile_list"`
		}{
			Content: content,
		},
	}))
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp.String())
	return nil
}

func (w *WeCom) SendMarkdown(content string) error {
	resp, err := req.Post(w.ReqUrl, req.BodyJSON(&MardkownContent{
		Msgtype: "markdown",
		Markdown: Markdown{
			Content: content,
		},
	}))
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp.String())
	return nil
}

func (w *WeCom) SendNews(title, desc, url, picUrl string) error {
	resp, err := req.Post(w.ReqUrl, req.BodyJSON(&NewsContent{
		Msgtype: "news",
		News: News{
			Articles: []Articles{{
				Title:       title,
				Description: desc,
				URL:         url,
				Picurl:      picUrl,
			}},
		},
	}))
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp.String())
	return nil
}
