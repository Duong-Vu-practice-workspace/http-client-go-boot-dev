package __mutli_query_param

func fetchTasks(baseURL, availability string) []Issue {
	// ?

	var limit string
	switch availability {
	case "Low":
		limit = "1"
	case "Medium":
		limit = "3"
	case "High":
		limit = "5"
	}
	fullURL := baseURL + "?sort=estimate&limit=" + limit
	return getIssues(fullURL)
}
