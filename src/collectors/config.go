package collectors

import (
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func FormatTitle(s string) string {
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, "/", "_")
	return s
}

var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"

var limitRules = &colly.LimitRule{
	RandomDelay: 2 * time.Second,
	Parallelism: 4,
}

type DecathlonPostData struct {
	Components []struct {
		ID    string `json:"id"`
		Input struct {
			AsyncRequest bool     `json:"asyncRequest"`
			Count        int      `json:"count"`
			Ids          []string `json:"ids"`
			Page         int      `json:"page"`
		} `json:"input"`
		Type string `json:"type"`
	} `json:"components"`
}
type DecathlonResponseBody struct {
	Num0 struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Input struct {
			Ids          []string `json:"ids"`
			Page         int      `json:"page"`
			Count        int      `json:"count"`
			AsyncRequest bool     `json:"asyncRequest"`
		} `json:"input"`
		Stats struct {
			Execution float64 `json:"execution"`
		} `json:"stats"`
		Params struct {
		} `json:"params"`
		Data struct {
			Stats struct {
				AverageRating           float64 `json:"averageRating"`
				ReviewsCount            int     `json:"reviewsCount"`
				SatisfiedReviewsCount   int     `json:"satisfiedReviewsCount"`
				RecommendedReviewsCount int     `json:"recommendedReviewsCount"`
				Ratings                 []struct {
					Note                  int `json:"note"`
					Reviews               int `json:"reviews"`
					SatisfiedReviewsCount int `json:"satisfiedReviewsCount"`
				} `json:"ratings"`
				SatisfiedReviewsPercentage float64 `json:"satisfiedReviewsPercentage"`
			} `json:"stats"`
			Pager struct {
				TotalPages     int `json:"totalPages"`
				ReviewsPerPage int `json:"reviewsPerPage"`
				CurrentPage    int `json:"currentPage"`
				NextPage       int `json:"nextPage"`
			} `json:"pager"`
			Reviews []struct {
				ReviewID int `json:"reviewId"`
				Author   struct {
					Country      string `json:"country"`
					AuthorID     int    `json:"authorId"`
					Firstname    string `json:"firstname"`
					Age          string `json:"age"`
					CountryLabel string `json:"countryLabel"`
				} `json:"author"`
				Review struct {
					Locale         string `json:"locale"`
					ModelID        string `json:"modelId"`
					Title          string `json:"title"`
					Body           string `json:"body"`
					PublishedAt    string `json:"publishedAt"`
					CountUpVotes   int    `json:"countUpVotes"`
					CountVotes     int    `json:"countVotes"`
					Useful         bool   `json:"useful"`
					Rating         int    `json:"rating"`
					IsTester       bool   `json:"isTester"`
					IsCustomer     bool   `json:"isCustomer"`
					IsCollaborator bool   `json:"isCollaborator"`
					VoteURL        string `json:"voteUrl"`
				} `json:"review"`
				UsageContext struct {
					RangeUsage string `json:"rangeUsage"`
					ModelID    string `json:"modelId"`
				} `json:"usageContext"`
				VerifiedPurchase bool `json:"verifiedPurchase"`
				Criterias        []struct {
					IsCriteriaEnabled bool          `json:"is_criteria_enabled"`
					ID                int           `json:"id"`
					Name              string        `json:"name"`
					Note              int           `json:"note"`
					Reasons           []interface{} `json:"reasons"`
				} `json:"criterias,omitempty"`
				SizeFit struct {
					Label   string `json:"label"`
					ID      int    `json:"id"`
					Message string `json:"message"`
				} `json:"sizeFit,omitempty"`
				WidthFit struct {
					Label   string `json:"label"`
					ID      int    `json:"id"`
					Message string `json:"message"`
				} `json:"widthFit,omitempty"`
			} `json:"reviews"`
			Sizometer struct {
				Message string `json:"message"`
				Popup   struct {
					LabelRange         string `json:"label_range"`
					ValuePercent       string `json:"value_percent"`
					Message            string `json:"message"`
					MessagePlaceholder string `json:"message_placeholder"`
				} `json:"popup"`
				SizeFit []struct {
					Value int    `json:"value"`
					Label string `json:"label"`
					Total int    `json:"total"`
				} `json:"sizeFit"`
				WidthFit []struct {
					Value int    `json:"value"`
					Label string `json:"label"`
					Total int    `json:"total"`
				} `json:"widthFit"`
			} `json:"sizometer"`
			Infos struct {
			} `json:"infos"`
		} `json:"data"`
	} `json:"0"`
}
