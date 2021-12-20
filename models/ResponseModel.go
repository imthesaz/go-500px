package models

type GraphQLResponse struct {
	Data struct {
		PhotoSearch struct {
			Edges []struct {
				Node struct {
					ID             string      `json:"id"`
					LegacyID       string      `json:"legacyId"`
					CanonicalPath  string      `json:"canonicalPath"`
					Name           string      `json:"name"`
					Description    string      `json:"description"`
					Category       string      `json:"category"`
					UploadedAt     string      `json:"uploadedAt"`
					Location       interface{} `json:"location"`
					Width          int         `json:"width"`
					Height         int         `json:"height"`
					IsLikedByMe    interface{} `json:"isLikedByMe"`
					NotSafeForWork bool        `json:"notSafeForWork"`
					Tags           []string    `json:"tags"`
					Photographer   struct {
						ID            string `json:"id"`
						LegacyID      string `json:"legacyId"`
						Username      string `json:"username"`
						DisplayName   string `json:"displayName"`
						CanonicalPath string `json:"canonicalPath"`
						Avatar        struct {
							Images []struct {
								URL string `json:"url"`
								ID  string `json:"id"`
							} `json:"images"`
							ID string `json:"id"`
						} `json:"avatar"`
						FollowedByUsers struct {
							TotalCount     int         `json:"totalCount"`
							IsFollowedByMe interface{} `json:"isFollowedByMe"`
						} `json:"followedByUsers"`
					} `json:"photographer"`
					Images []struct {
						Size    int    `json:"size"`
						URL     string `json:"url"`
						JpegURL string `json:"jpegUrl"`
						WebpURL string `json:"webpUrl"`
						ID      string `json:"id"`
					} `json:"images"`
					Typename string `json:"__typename"`
				} `json:"node"`
				Cursor string `json:"cursor"`
			} `json:"edges"`
			TotalCount int `json:"totalCount"`
			PageInfo   struct {
				EndCursor   string `json:"endCursor"`
				HasNextPage bool   `json:"hasNextPage"`
			} `json:"pageInfo"`
		} `json:"photoSearch"`
	} `json:"data"`
}
