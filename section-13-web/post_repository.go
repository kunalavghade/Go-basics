package main

import (
	"database/sql"
	"errors"
	"math"
	"strings"
	"time"
)

var (
	ErrDuplicatePostTitle = errors.New("Duplicate Post Title")
)

type Post struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	URL           string    `json:"url"`
	UserID        int       `json:"user_id"`
	UserName      string    `json:"user_name"`
	VotesCount    int       `json:"votes_count"`
	CommentsCount int       `json:"comments_count"`
	TotalRecords  int       `json:"total_records"`
	CreatedAt     time.Time `json:"created_at"`
}

type Comment struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

type Filter struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
	OrderBy     string `json:"order_by"`
	OrderDir    string `json:"order_dir"`
	SearchQuery string `json:"search_query"`
}

func (f *Filter) validate() error {
	if f.Page <= 0 || f.Page >= 10_000_000 {
		return errors.New("Invalid Page Range: 1 to 10000000")
	}
	if f.PageSize <= 0 || f.PageSize >= 100 {
		return errors.New("Invalid Page Size: 1 to 100")
	}
	return nil
}

type MetaData struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	FirstPage    int `json:"first_page"`
	LastPage     int `json:"last_page"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
	TotalRecords int `json:"total_records"`
}

func calculateMataData(TotalRecords, Page, PageSize int) MetaData {
	if TotalRecords == 0 {
		return MetaData{}
	}
	mata := MetaData{
		CurrentPage:  Page,
		PageSize:     PageSize,
		TotalRecords: TotalRecords,
		LastPage:     int(math.Ceil(float64(TotalRecords) / float64(PageSize))),
		FirstPage:    1,
		NextPage:     0,
		PreviousPage: 0,
	}

	if Page > 1 {
		mata.PreviousPage = Page - 1
	}
	if Page < mata.LastPage {
		mata.NextPage = Page + 1
	}
	return mata
}

type PostRepository interface {
	CreatePost(userId int, title, url string) (int, error)
	AddComment(userId, postId int, body string) (int, error)
	AddVote(userId, postId int) (int, error)
	GetAllPosts(filter Filter) ([]Post, MetaData, error)
	GetById(id int) (*Post, error)
	GetComments(postID int) ([]Comment, error)
}

type SQLPostRepository struct {
	db *sql.DB
}

func NewSQLPostRepository(db *sql.DB) *SQLPostRepository {
	return &SQLPostRepository{db: db}
}

func (r *SQLPostRepository) CreatePost(title, url string, userId int) (int, error) {
	stmt := `INSERT INTO posts (title, url, user_id) VALUES (?, ?, ?)`
	res, err := r.db.Exec(stmt, title, url, userId)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: post.title") {
			return 0, ErrDuplicatePostTitle
		}
		return 0, err
	}
	postId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(postId), nil
}

func (r *SQLPostRepository) AddComment(userId, postId int, body string) (int, error) {
	stmt := `INSERT INTO comments (user_id, post_id, body) VALUES (?, ?, ?)`
	res, err := r.db.Exec(stmt, userId, postId, body)
	if err != nil {
		return 0, err
	}
	commentId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(commentId), nil
}

func (r *SQLPostRepository) AddVote(userId, postId int) (int, error) {
	stmt := `INSERT INTO votes (user_id, post_id) VALUES (?, ?)`
	_, err := r.db.Exec(stmt, userId, postId)
	if err != nil {
		if strings.Contains(err.Error(), "Unique constraint failed") || strings.Contains(err.Error(), "PRIMARY KEY constrint failed") {
			return 0, errors.New("User already voted for this post")
		}
		return 0, err
	}
	return 1, nil
}

func (r *SQLPostRepository) GetByID(id int) (*Post, error) {
	query := `
		SELECT 
			p.id,
			p.title,
			p.url, 
			p.user_id,
			u.name,
			COUNT(DISTINCT v.id) AS votes_count,
			COUNT(DISTINCT c.id) AS comments_count,
			p.created_at 
		FROM posts p
		INNER JOIN users u ON p.user_id = u.id 
		LEFT JOIN votes v ON v.post_id = p.id
		LEFT JOIN comments c ON c.post_id = p.id
		WHERE p.id = ?
		GROUP BY 
			p.id,
			p.title,
			p.url,
			p.user_id,
			u.name,
			p.created_at
	`
	post := &Post{}
	err := r.db.QueryRow(query, id).Scan(
		&post.ID,
		&post.Title,
		&post.URL,
		&post.UserID,
		&post.UserName,
		&post.VotesCount,
		&post.CommentsCount,
		&post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *SQLPostRepository) GetAllPosts(filter Filter) ([]Post, MetaData, error) {
	if err := filter.validate(); err != nil {
		return nil, MetaData{}, err
	}

	query := `
		SELECT 
			COUNT(*) OVER() AS total_records,
			p.id,
			p.title,
			p.url, 
			p.user_id,
			u.name,
			COUNT(DISTINCT v.id) AS votes_count,
			COUNT(DISTINCT c.id) AS comments_count,
			p.created_at 
		FROM posts p
		INNER JOIN users u ON p.user_id = u.id 
		LEFT JOIN votes v ON v.post_id = p.id
		LEFT JOIN comments c ON c.post_id = p.id
	`
	var args []any

	if filter.SearchQuery != "" {
		query += "WHERE p.title ILIKE ? OR p.url ILIKE ? OR u.name ILIKE ?"
		search := strings.ToLower(filter.SearchQuery)
		args = append(args, "%"+search+"%")
	}

	query += `
		GROUP BY 
			p.id,
			p.title,
			p.url,
			p.user_id,
			u.name,
			p.created_at
	`

	if filter.OrderBy != "" {
		query += "ORDER BY " + filter.OrderBy
		if filter.OrderDir != "" {
			query += " " + filter.OrderDir
		}
	} else {
		query += "ORDER BY p.created_at DESC"
	}

	query += `
		LIMIT ? OFFSET ?
	`
	args = append(args, filter.PageSize, (filter.Page-1)*filter.PageSize)

	println(query)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, MetaData{}, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		post := Post{}
		err := rows.Scan(
			&post.TotalRecords,
			&post.ID,
			&post.Title,
			&post.URL,
			&post.UserID,
			&post.UserName,
			&post.VotesCount,
			&post.CommentsCount,
			&post.CreatedAt)
		if err != nil {
			return nil, MetaData{}, err
		}
		posts = append(posts, post)
	}
	return posts, calculateMataData(posts[0].TotalRecords, filter.Page, filter.PageSize), nil
}

func (r *SQLPostRepository) GetComments(postID int) ([]Comment, error) {
	query := `
		SELECT c.id,c.body,c.user_id,c.post_id,u.name,c.created_at 
		FROM comments c
		LEFT JOIN users u ON c.user_id = u.id
		WHERE c.post_id = ?
		ORDER BY c.created_at DESC
	`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := []Comment{}
	for rows.Next() {
		comment := Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.Body,
			&comment.UserID,
			&comment.PostID,
			&comment.UserName,
			&comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
