package urls

type (
	urls struct {
		STATIC_PATH          string
		HOME_PATH            string
		PROJECT_VERSION_PATH string
	}
)

func ReturnURLS() urls {
	var url_patterns urls
	url_patterns.STATIC_PATH = "/static/"
	url_patterns.PROJECT_VERSION_PATH = "api/v1/"
	url_patterns.HOME_PATH = url_patterns.PROJECT_VERSION_PATH + "/"
	return url_patterns
}
