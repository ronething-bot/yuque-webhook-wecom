package param

import "time"

type NewComment struct {
	Data CommentData `json:"data"`
}

type NewCommentReply = NewComment

type CommentData struct {
	Serializer         string           `json:"_serializer"`
	ActionType         string           `json:"action_type"`
	ActorID            int              `json:"actor_id"`
	BodyHTML           string           `json:"body_html"`
	Commentable        Commentable      `json:"commentable"`
	CreatedAt          time.Time        `json:"created_at"`
	ID                 int              `json:"id"`
	LikesT             int              `json:"likes_t"`
	Mood               int              `json:"mood"`
	ParentID           interface{}      `json:"parent_id"`
	Path               string           `json:"path"`
	Status             int              `json:"status"`
	ToUserID           interface{}      `json:"to_user_id"`
	Type               interface{}      `json:"type"`
	UpdatedAt          time.Time        `json:"updated_at"`
	User               CommentTopicUser `json:"user"`
	UserID             int              `json:"user_id"`
	WebhookSubjectType string           `json:"webhook_subject_type"`
}

type Book struct {
	Serializer       string      `json:"_serializer"`
	ContentUpdatedAt time.Time   `json:"content_updated_at"`
	CreatedAt        time.Time   `json:"created_at"`
	CreatorID        int         `json:"creator_id"`
	Description      string      `json:"description"`
	ID               int         `json:"id"`
	ItemsCount       int         `json:"items_count"`
	LikesCount       int         `json:"likes_count"`
	Name             string      `json:"name"`
	Public           int         `json:"public"`
	Slug             string      `json:"slug"`
	Type             string      `json:"type"`
	UpdatedAt        time.Time   `json:"updated_at"`
	User             interface{} `json:"user"`
	UserID           int         `json:"user_id"`
	WatchesCount     int         `json:"watches_count"`
}

type Commentable struct {
	Serializer       string      `json:"_serializer"`
	Body             string      `json:"body"`
	BodyDraft        string      `json:"body_draft"`
	BodyHTML         string      `json:"body_html"`
	Book             Book        `json:"book"`
	BookID           int         `json:"book_id"`
	CommentsCount    int         `json:"comments_count"`
	ContentUpdatedAt time.Time   `json:"content_updated_at"`
	CreatedAt        time.Time   `json:"created_at"`
	DeletedAt        interface{} `json:"deleted_at"`
	FirstPublishedAt time.Time   `json:"first_published_at"`
	Format           string      `json:"format"`
	ID               int         `json:"id"`
	LikesCount       int         `json:"likes_count"`
	Path             string      `json:"path"`
	Public           int         `json:"public"`
	PublishedAt      time.Time   `json:"published_at"`
	ReadStatus       int         `json:"read_status"`
	Slug             string      `json:"slug"`
	Status           int         `json:"status"`
	Title            string      `json:"title"`
	UpdatedAt        time.Time   `json:"updated_at"`
	User             interface{} `json:"user"`
	UserID           int         `json:"user_id"`
	ViewStatus       int         `json:"view_status"`
	WordCount        int         `json:"word_count"`
}

type CommentTopicUser struct {
	Serializer     string    `json:"_serializer"`
	AvatarURL      string    `json:"avatar_url"`
	CreatedAt      time.Time `json:"created_at"`
	Description    string    `json:"description"`
	FollowersCount int       `json:"followers_count"`
	FollowingCount int       `json:"following_count"`
	ID             int       `json:"id"`
	Login          string    `json:"login"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	UpdatedAt      time.Time `json:"updated_at"`
}
