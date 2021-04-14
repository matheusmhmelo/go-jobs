package search

import (
	"github.com/matheusmhmelo/go-jobs/internal/repository"
	"os"
	"strconv"
)

func Jobs(title, page string) (map[string]interface{}, error) {
	var err error
	var jobs []map[string]interface{}
	var total int64

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	if pageInt <= 0 {
		pageInt = 1
	}

	r := repository.NewSearch()

	switch title {
		case "":
			jobs, total, err = r.AllJobs(pageInt)
		default:
			jobs, total, err = r.Jobs(title, pageInt)
	}

	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"results": jobs,
		"pagination": map[string]interface{}{
			"page": pageInt,
			"results_per_page": os.Getenv("RESULTS_PER_PAGE"),
			"total_results": total,
		},
	}

	return response, nil
}
