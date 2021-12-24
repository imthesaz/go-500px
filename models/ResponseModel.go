package models

import "time"

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
					Location       string      `json:"location"`
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

func (G *GraphQLResponse) GetHasNextPage() bool {
	return G.Data.PhotoSearch.PageInfo.HasNextPage
}

type PhotoDetail struct {
	Photos struct {
		PhotoInfo struct {
			ID                int       `json:"id"`
			CreatedAt         time.Time `json:"created_at"`
			Privacy           bool      `json:"privacy"`
			Profile           bool      `json:"profile"`
			URL               string    `json:"url"`
			UserID            int       `json:"user_id"`
			Status            int       `json:"status"`
			Width             int       `json:"width"`
			Height            int       `json:"height"`
			Rating            float64   `json:"rating"`
			HighestRating     float64   `json:"highest_rating"`
			HighestRatingDate time.Time `json:"highest_rating_date"`
			ImageFormat       string    `json:"image_format"`
			Images            []struct {
				Format   string `json:"format"`
				Size     int    `json:"size"`
				URL      string `json:"url"`
				HTTPSURL string `json:"https_url"`
			} `json:"images"`
			ImageURL        []string    `json:"image_url"`
			Name            string      `json:"name"`
			Description     string      `json:"description"`
			Category        int         `json:"category"`
			TakenAt         time.Time   `json:"taken_at"`
			ShutterSpeed    string      `json:"shutter_speed"`
			FocalLength     string      `json:"focal_length"`
			Aperture        string      `json:"aperture"`
			Camera          string      `json:"camera"`
			Lens            string      `json:"lens"`
			Iso             string      `json:"iso"`
			Location        string      `json:"location"`
			Latitude        float64     `json:"latitude"`
			Longitude       float64     `json:"longitude"`
			Nsfw            bool        `json:"nsfw"`
			PrivacyLevel    int         `json:"privacy_level"`
			Watermark       bool        `json:"watermark"`
			ShowExifData    bool        `json:"show_exif_data"`
			Tags            []string    `json:"tags"`
			HasNsfwTags     bool        `json:"has_nsfw_tags"`
			Liked           interface{} `json:"liked"`
			Voted           interface{} `json:"voted"`
			LocationDetails struct {
				Attraction []string    `json:"attraction"`
				County     []string    `json:"county"`
				State      []string    `json:"state"`
				City       []string    `json:"city"`
				Country    []string    `json:"country"`
				LocationID interface{} `json:"location_id"`
			} `json:"location_details"`
			LicensingInfo struct {
				ID                      int         `json:"id"`
				UserID                  int         `json:"user_id"`
				StoredAt                string      `json:"stored_at"`
				Width                   int         `json:"width"`
				Height                  int         `json:"height"`
				Status                  string      `json:"status"`
				ModelRelease            interface{} `json:"model_release"`
				Location                interface{} `json:"location"`
				Latitude                interface{} `json:"latitude"`
				Longitude               interface{} `json:"longitude"`
				PropertyRelease         interface{} `json:"property_release"`
				ExclusiveUse            bool        `json:"exclusive_use"`
				SubmissionDate          string      `json:"submission_date"`
				Title                   string      `json:"title"`
				Description             string      `json:"description"`
				Keywords                string      `json:"keywords"`
				AiKeywords              interface{} `json:"ai_keywords"`
				GettyID                 string      `json:"getty_id"`
				VcgID                   string      `json:"vcg_id"`
				TakenAt                 string      `json:"taken_at"`
				CreatedAt               string      `json:"created_at"`
				UpdatedAt               string      `json:"updated_at"`
				Brand                   string      `json:"brand"`
				FileName                interface{} `json:"file_name"`
				Orientation             interface{} `json:"orientation"`
				ShutterSpeed            string      `json:"shutter_speed"`
				Camera                  string      `json:"camera"`
				Aperture                string      `json:"aperture"`
				Iso                     string      `json:"iso"`
				FocalLength             string      `json:"focal_length"`
				Lens                    string      `json:"lens"`
				Category                int         `json:"category"`
				ExtraInfo               string      `json:"extra_info"`
				Nsfw                    interface{} `json:"nsfw"`
				LicensingPhotoExtraInfo struct {
					RemovedAt              interface{} `json:"removed_at"`
					AcceptedAt             string      `json:"accepted_at"`
					RemovedBy              int         `json:"removed_by"`
					Reuploaded             bool        `json:"reuploaded"`
					ModelReleaseChanged    bool        `json:"model_release_changed"`
					PropertyReleaseChanged bool        `json:"property_release_changed"`
				} `json:"licensing_photo_extra_info"`
				StoreOn        bool   `json:"store_on"`
				LicensingUsage string `json:"licensing_usage"`
				ImageURL       struct {
					Num1 string `json:"1"`
					Num4 string `json:"4"`
				} `json:"image_url"`
				ReasonList          []interface{} `json:"reason_list"`
				PropertyReleaseList []interface{} `json:"property_release_list"`
				ModelReleaseList    []interface{} `json:"model_release_list"`
				VendorPhotos        []struct {
					VendorID      int `json:"vendor_id"`
					VendorPhotoID int `json:"vendor_photo_id"`
				} `json:"vendor_photos"`
				FieldName              interface{} `json:"field_name"`
				Privacy                interface{} `json:"privacy"`
				SubmitToCms            bool        `json:"submit_to_cms"`
				RecognizableProperties bool        `json:"recognizable_properties"`
				RecognizablePeople     bool        `json:"recognizable_people"`
				DownloadLink           string      `json:"download_link"`
			} `json:"licensing_info"`
			StoreWidth         int           `json:"store_width"`
			StoreHeight        int           `json:"store_height"`
			LicensingStatus    int           `json:"licensing_status"`
			LicensingType      string        `json:"licensing_type"`
			LicensingUsage     string        `json:"licensing_usage"`
			StoreLicense       bool          `json:"store_license"`
			Comments           []interface{} `json:"comments"`
			CommentsCount      int           `json:"comments_count"`
			VotesCount         int           `json:"votes_count"`
			PositiveVotesCount int           `json:"positive_votes_count"`
			LikedBy            []interface{} `json:"liked_by"`
			TimesViewed        int           `json:"times_viewed"`
			User               struct {
				ID               int       `json:"id"`
				Username         string    `json:"username"`
				Fullname         string    `json:"fullname"`
				AvatarVersion    int       `json:"avatar_version"`
				RegistrationDate time.Time `json:"registration_date"`
				Avatars          struct {
					Tiny struct {
						HTTPS string `json:"https"`
					} `json:"tiny"`
					Small struct {
						HTTPS string `json:"https"`
					} `json:"small"`
					Large struct {
						HTTPS string `json:"https"`
					} `json:"large"`
					Cover struct {
						HTTPS string `json:"https"`
					} `json:"cover"`
					Default struct {
						HTTPS string `json:"https"`
					} `json:"default"`
				} `json:"avatars"`
				UserpicURL      string `json:"userpic_url"`
				UserpicHTTPSURL string `json:"userpic_https_url"`
				Usertype        int    `json:"usertype"`
				Active          int    `json:"active"`
				StoreOn         bool   `json:"store_on"`
				Firstname       string `json:"firstname"`
				Lastname        string `json:"lastname"`
				About           string `json:"about"`
				City            string `json:"city"`
				State           string `json:"state"`
				Country         string `json:"country"`
				CoverURL        string `json:"cover_url"`
				UpgradeStatus   int    `json:"upgrade_status"`
				Affection       int    `json:"affection"`
				FollowersCount  int    `json:"followers_count"`
				Following       bool   `json:"following"`
			} `json:"user"`
			EditorsChoice     bool        `json:"editors_choice"`
			EditorsChoiceDate interface{} `json:"editors_choice_date"`
			EditoredBy        interface{} `json:"editored_by"`
			Feature           string      `json:"feature"`
			FeatureDate       time.Time   `json:"feature_date"`
			CameraInfo        struct {
				ID           int    `json:"id"`
				Name         string `json:"name"`
				FriendlyName string `json:"friendly_name"`
				CameraType   int    `json:"camera_type"`
				Verified     bool   `json:"verified"`
				Slug         string `json:"slug"`
				Features     string `json:"features"`
				Brand        struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Slug string `json:"slug"`
				} `json:"brand"`
			} `json:"camera_info"`
			LensInfo struct {
				ID           int         `json:"id"`
				Name         string      `json:"name"`
				FriendlyName string      `json:"friendly_name"`
				Slug         string      `json:"slug"`
				Features     interface{} `json:"features"`
				Brand        struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Slug string `json:"slug"`
				} `json:"brand"`
			} `json:"lens_info"`
			FillSwitch struct {
				AccessDeleted        bool        `json:"access_deleted"`
				AccessPrivate        bool        `json:"access_private"`
				IncludeDeleted       bool        `json:"include_deleted"`
				ExcludePrivate       bool        `json:"exclude_private"`
				ExcludeNude          bool        `json:"exclude_nude"`
				AlwaysExcludeNude    bool        `json:"always_exclude_nude"`
				ExcludeBlock         bool        `json:"exclude_block"`
				CurrentUserID        interface{} `json:"current_user_id"`
				OnlyUserActive       bool        `json:"only_user_active"`
				IncludeTags          bool        `json:"include_tags"`
				IncludeGeo           bool        `json:"include_geo"`
				IncludeLicensing     bool        `json:"include_licensing"`
				IncludeAdminLocks    bool        `json:"include_admin_locks"`
				IncludeLikeBy        bool        `json:"include_like_by"`
				IncludeComments      bool        `json:"include_comments"`
				IncludeUserInfo      bool        `json:"include_user_info"`
				IncludeFollowInfo    bool        `json:"include_follow_info"`
				IncludeEquipmentInfo bool        `json:"include_equipment_info"`
			} `json:"fill_switch"`
		} `json:"photo_info"`
	} `json:"photos"`
}
