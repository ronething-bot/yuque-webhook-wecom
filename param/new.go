package param

import "time"

type NewTopic struct {
	Data NewTopicData `json:"data"`
}

type NewTopicData struct {
	Serializer         string       `json:"_serializer"`
	ActionType         string       `json:"action_type"`
	ActorID            int          `json:"actor_id"`
	Body               string       `json:"body"`
	BodyDraft          string       `json:"body_draft"`
	BodyHTML           string       `json:"body_html"`
	Book               Book         `json:"book"`
	BookID             int          `json:"book_id"`
	CommentsCount      int          `json:"comments_count"`
	ContentUpdatedAt   time.Time    `json:"content_updated_at"`
	CreatedAt          time.Time    `json:"created_at"`
	DeletedAt          interface{}  `json:"deleted_at"`
	FirstPublishedAt   time.Time    `json:"first_published_at"`
	Format             string       `json:"format"`
	ID                 int          `json:"id"`
	LikesCount         int          `json:"likes_count"`
	Path               string       `json:"path"`
	Public             int          `json:"public"`
	Publish            bool         `json:"publish"`
	PublishedAt        time.Time    `json:"published_at"`
	ReadStatus         int          `json:"read_status"`
	Slug               string       `json:"slug"`
	Status             int          `json:"status"`
	Title              string       `json:"title"`
	UpdatedAt          time.Time    `json:"updated_at"`
	User               NewTopicUser `json:"user"`
	UserID             int          `json:"user_id"`
	ViewStatus         int          `json:"view_status"`
	WebhookSubjectType string       `json:"webhook_subject_type"`
	WordCount          int          `json:"word_count"`
}

type NewTopicUser struct {
	Serializer       string    `json:"_serializer"`
	AvatarURL        string    `json:"avatar_url"`
	BooksCount       int       `json:"books_count"`
	CreatedAt        time.Time `json:"created_at"`
	Description      string    `json:"description"`
	FollowersCount   int       `json:"followers_count"`
	FollowingCount   int       `json:"following_count"`
	ID               int       `json:"id"`
	Login            string    `json:"login"`
	Name             string    `json:"name"`
	PublicBooksCount int       `json:"public_books_count"`
	Type             string    `json:"type"`
	UpdatedAt        time.Time `json:"updated_at"`
}
