package models

import "fmt"

const baseURL = "https://api.500px.com/v1/photos?ids="
const queryParameters = "&image_size%5B%5D=1&image_size%5B%5D=2&image_size%5B%5D=32&image_size%5B%5D=31&image_size%5B%5D=33&image_size%5B%5D=34&image_size%5B%5D=35&image_size%5B%5D=36&image_size%5B%5D=2048&image_size%5B%5D=4&image_size%5B%5D=14&include_states=1&expanded_user_info=true&include_tags=true&include_geo=true&is_following=true&include_equipment_info=true&include_licensing=true&include_releases=true&liked_by=1&include_vendor_photos=true"

type PhotoSearchPaginationContainerQuery struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Cursor string `json:"cursor"`
		Search string `json:"search"`
		Sort   string `json:"sort"`
	} `json:"variables"`
	Query string `json:"query"`
}

type PhotoSearchQueryRendererQuery struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Sort   string `json:"sort"`
		Search string `json:"search"`
	} `json:"variables"`
	Query string `json:"query"`
}

func CreatePhotoInfoQuery(id string) string {
	return fmt.Sprintf("%s%s%s", baseURL, id, queryParameters)
}

func (P *PhotoSearchQueryRendererQuery) InitPhotoSearchQueryRendererQueryBody(searchStr string, sortStr string) {
	P.OperationName = "PhotoSearchQueryRendererQuery"
	P.Query = "query PhotoSearchQueryRendererQuery($sort: PhotoSort, $search: String!) {\n  ...PhotoSearchPaginationContainer_query_67nah\n}\n\nfragment PhotoSearchPaginationContainer_query_67nah on Query {\n  photoSearch(sort: $sort, first: 20, search: $search) {\n    edges {\n      node {\n        id\n        legacyId\n        canonicalPath\n        name\n        description\n        category\n        uploadedAt\n        location\n        width\n        height\n        isLikedByMe\n        notSafeForWork\n        tags\n        photographer: uploader {\n          id\n          legacyId\n          username\n          displayName\n          canonicalPath\n          avatar {\n            images {\n              url\n              id\n            }\n            id\n          }\n          followedByUsers {\n            totalCount\n            isFollowedByMe\n          }\n        }\n        images(sizes: [33, 35]) {\n          size\n          url\n          jpegUrl\n          webpUrl\n          id\n        }\n        __typename\n      }\n      cursor\n    }\n    totalCount\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n  }\n}\n"
	P.Variables.Search = searchStr
	P.Variables.Sort = sortStr
}
func (P *PhotoSearchPaginationContainerQuery) InitPhotoSearchPaginationContainerQueryBody(cursorStr string, searchStr string, sortStr string) {
	P.OperationName = "PhotoSearchPaginationContainerQuery"
	P.Query = "query PhotoSearchPaginationContainerQuery($cursor: String, $search: String!, $sort: PhotoSort) {\n  ...PhotoSearchPaginationContainer_query_2SsbxT\n}\n\nfragment PhotoSearchPaginationContainer_query_2SsbxT on Query {\n  photoSearch(sort: $sort, first: 20, after: $cursor, search: $search) {\n    edges {\n      node {\n        id\n        legacyId\n        canonicalPath\n        name\n        description\n        category\n        uploadedAt\n        location\n        width\n        height\n        isLikedByMe\n        notSafeForWork\n        tags\n        photographer: uploader {\n          id\n          legacyId\n          username\n          displayName\n          canonicalPath\n          avatar {\n            images {\n              url\n              id\n            }\n            id\n          }\n          followedByUsers {\n            totalCount\n            isFollowedByMe\n          }\n        }\n        images(sizes: [33, 35]) {\n          size\n          url\n          jpegUrl\n          webpUrl\n          id\n        }\n        __typename\n      }\n      cursor\n    }\n    totalCount\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n  }\n}\n"
	P.Variables.Cursor = cursorStr
	P.Variables.Search = searchStr
	P.Variables.Sort = sortStr
}
