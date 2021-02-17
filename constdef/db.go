package constdef

const (
	// DSNTemplate .
	DSNTemplate string = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local"
)

// Table names
const (
	BlogArticleTable = "blog_article"
	BlogUserTable    = "blog_user"
)
