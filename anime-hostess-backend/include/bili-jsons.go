package include

type BiliSearchAnime struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Seid           string `json:"seid"`
		Page           int    `json:"page"`
		Pagesize       int    `json:"pagesize"`
		NumResults     int    `json:"numResults"`
		NumPages       int    `json:"numPages"`
		SuggestKeyword string `json:"suggest_keyword"`
		RqtType        string `json:"rqt_type"`
		CostTime       struct {
			ParamsCheck         string `json:"params_check"`
			IllegalHandler      string `json:"illegal_handler"`
			AsResponseFormat    string `json:"as_response_format"`
			AsRequest           string `json:"as_request"`
			DeserializeResponse string `json:"deserialize_response"`
			AsRequestFormat     string `json:"as_request_format"`
			Total               string `json:"total"`
			MainHandler         string `json:"main_handler"`
		} `json:"cost_time"`
		ExpList struct {
			Field1 bool `json:"6608"`
			Field2 bool `json:"5505"`
		} `json:"exp_list"`
		EggHit int `json:"egg_hit"`
		Result []struct {
			Type           string   `json:"type"`
			MediaId        int      `json:"media_id"`
			Title          string   `json:"title"`
			OrgTitle       string   `json:"org_title"`
			MediaType      int      `json:"media_type"`
			Cv             string   `json:"cv"`
			Staff          string   `json:"staff"`
			SeasonId       int      `json:"season_id"`
			IsAvid         bool     `json:"is_avid"`
			HitColumns     []string `json:"hit_columns"`
			HitEpids       string   `json:"hit_epids"`
			SeasonType     int      `json:"season_type"`
			SeasonTypeName string   `json:"season_type_name"`
			SelectionStyle string   `json:"selection_style"`
			EpSize         int      `json:"ep_size"`
			Url            string   `json:"url"`
			ButtonText     string   `json:"button_text"`
			IsFollow       int      `json:"is_follow"`
			IsSelection    int      `json:"is_selection"`
			Eps            []struct {
				Id          int    `json:"id"`
				Cover       string `json:"cover"`
				Title       string `json:"title"`
				Url         string `json:"url"`
				ReleaseDate string `json:"release_date"`
				Badges      []struct {
					Text             string `json:"text"`
					TextColor        string `json:"text_color"`
					TextColorNight   string `json:"text_color_night"`
					BgColor          string `json:"bg_color"`
					BgColorNight     string `json:"bg_color_night"`
					BorderColor      string `json:"border_color"`
					BorderColorNight string `json:"border_color_night"`
					BgStyle          int    `json:"bg_style"`
				} `json:"badges"`
				IndexTitle string `json:"index_title"`
				LongTitle  string `json:"long_title"`
			} `json:"eps"`
			Badges []struct {
				Text             string `json:"text"`
				TextColor        string `json:"text_color"`
				TextColorNight   string `json:"text_color_night"`
				BgColor          string `json:"bg_color"`
				BgColorNight     string `json:"bg_color_night"`
				BorderColor      string `json:"border_color"`
				BorderColorNight string `json:"border_color_night"`
				BgStyle          int    `json:"bg_style"`
			} `json:"badges"`
			Cover         string `json:"cover"`
			Areas         string `json:"areas"`
			Styles        string `json:"styles"`
			GotoUrl       string `json:"goto_url"`
			Desc          string `json:"desc"`
			Pubtime       int    `json:"pubtime"`
			MediaMode     int    `json:"media_mode"`
			FixPubtimeStr string `json:"fix_pubtime_str"`
			MediaScore    struct {
				Score     float64 `json:"score"`
				UserCount int     `json:"user_count"`
			} `json:"media_score"`
			DisplayInfo []struct {
				Text             string `json:"text"`
				TextColor        string `json:"text_color"`
				TextColorNight   string `json:"text_color_night"`
				BgColor          string `json:"bg_color"`
				BgColorNight     string `json:"bg_color_night"`
				BorderColor      string `json:"border_color"`
				BorderColorNight string `json:"border_color_night"`
				BgStyle          int    `json:"bg_style"`
			} `json:"display_info"`
			PgcSeasonId int `json:"pgc_season_id"`
			Corner      int `json:"corner"`
		} `json:"result"`
		ShowColumn int `json:"show_column"`
	} `json:"data"`
}
type BiliAnimeDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		Activity struct {
			HeadBgUrl string `json:"head_bg_url"`
			Id        int    `json:"id"`
			Title     string `json:"title"`
		} `json:"activity"`
		Alias string `json:"alias"`
		Areas []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"areas"`
		BkgCover string `json:"bkg_cover"`
		Cover    string `json:"cover"`
		Episodes []struct {
			Aid       int    `json:"aid"`
			Badge     string `json:"badge"`
			BadgeInfo struct {
				BgColor      string `json:"bg_color"`
				BgColorNight string `json:"bg_color_night"`
				Text         string `json:"text"`
			} `json:"badge_info"`
			BadgeType int    `json:"badge_type"`
			Bvid      string `json:"bvid"`
			Cid       int    `json:"cid"`
			Cover     string `json:"cover"`
			Dimension struct {
				Height int `json:"height"`
				Rotate int `json:"rotate"`
				Width  int `json:"width"`
			} `json:"dimension"`
			Duration    int    `json:"duration"`
			From        string `json:"from"`
			Id          int    `json:"id"`
			IsViewHide  bool   `json:"is_view_hide"`
			Link        string `json:"link"`
			LongTitle   string `json:"long_title"`
			PubTime     int    `json:"pub_time"`
			Pv          int    `json:"pv"`
			ReleaseDate string `json:"release_date"`
			Rights      struct {
				AllowDemand   int `json:"allow_demand"`
				AllowDm       int `json:"allow_dm"`
				AllowDownload int `json:"allow_download"`
				AreaLimit     int `json:"area_limit"`
			} `json:"rights"`
			ShareCopy string `json:"share_copy"`
			ShareUrl  string `json:"share_url"`
			ShortLink string `json:"short_link"`
			Status    int    `json:"status"`
			Subtitle  string `json:"subtitle"`
			Title     string `json:"title"`
			Vid       string `json:"vid"`
		} `json:"episodes"`
		Evaluate string `json:"evaluate"`
		Freya    struct {
			BubbleDesc    string `json:"bubble_desc"`
			BubbleShowCnt int    `json:"bubble_show_cnt"`
			IconShow      int    `json:"icon_show"`
		} `json:"freya"`
		JpTitle string `json:"jp_title"`
		Link    string `json:"link"`
		MediaId int    `json:"media_id"`
		Mode    int    `json:"mode"`
		NewEp   struct {
			Desc  string `json:"desc"`
			Id    int    `json:"id"`
			IsNew int    `json:"is_new"`
			Title string `json:"title"`
		} `json:"new_ep"`
		Payment struct {
			Discount int `json:"discount"`
			PayType  struct {
				AllowDiscount    int `json:"allow_discount"`
				AllowPack        int `json:"allow_pack"`
				AllowTicket      int `json:"allow_ticket"`
				AllowTimeLimit   int `json:"allow_time_limit"`
				AllowVipDiscount int `json:"allow_vip_discount"`
				ForbidBb         int `json:"forbid_bb"`
			} `json:"pay_type"`
			Price             string `json:"price"`
			Promotion         string `json:"promotion"`
			Tip               string `json:"tip"`
			ViewStartTime     int    `json:"view_start_time"`
			VipDiscount       int    `json:"vip_discount"`
			VipFirstPromotion string `json:"vip_first_promotion"`
			VipPromotion      string `json:"vip_promotion"`
		} `json:"payment"`
		Positive struct {
			Id    int    `json:"id"`
			Title string `json:"title"`
		} `json:"positive"`
		Publish struct {
			IsFinish      int    `json:"is_finish"`
			IsStarted     int    `json:"is_started"`
			PubTime       string `json:"pub_time"`
			PubTimeShow   string `json:"pub_time_show"`
			UnknowPubDate int    `json:"unknow_pub_date"`
			Weekday       int    `json:"weekday"`
		} `json:"publish"`
		Rating struct {
			Count int     `json:"count"`
			Score float64 `json:"score"`
		} `json:"rating"`
		Record string `json:"record"`
		Rights struct {
			AllowBp         int    `json:"allow_bp"`
			AllowBpRank     int    `json:"allow_bp_rank"`
			AllowDownload   int    `json:"allow_download"`
			AllowReview     int    `json:"allow_review"`
			AreaLimit       int    `json:"area_limit"`
			BanAreaShow     int    `json:"ban_area_show"`
			CanWatch        int    `json:"can_watch"`
			Copyright       string `json:"copyright"`
			ForbidPre       int    `json:"forbid_pre"`
			FreyaWhite      int    `json:"freya_white"`
			IsCoverShow     int    `json:"is_cover_show"`
			IsPreview       int    `json:"is_preview"`
			OnlyVipDownload int    `json:"only_vip_download"`
			Resource        string `json:"resource"`
			WatchPlatform   int    `json:"watch_platform"`
		} `json:"rights"`
		SeasonId    int    `json:"season_id"`
		SeasonTitle string `json:"season_title"`
		Seasons     []struct {
			Badge     string `json:"badge"`
			BadgeInfo struct {
				BgColor      string `json:"bg_color"`
				BgColorNight string `json:"bg_color_night"`
				Text         string `json:"text"`
			} `json:"badge_info"`
			BadgeType           int    `json:"badge_type"`
			Cover               string `json:"cover"`
			HorizontalCover1610 string `json:"horizontal_cover_1610"`
			HorizontalCover169  string `json:"horizontal_cover_169"`
			MediaId             int    `json:"media_id"`
			NewEp               struct {
				Cover     string `json:"cover"`
				Id        int    `json:"id"`
				IndexShow string `json:"index_show"`
			} `json:"new_ep"`
			SeasonId    int    `json:"season_id"`
			SeasonTitle string `json:"season_title"`
			SeasonType  int    `json:"season_type"`
			Stat        struct {
				Favorites    int `json:"favorites"`
				SeriesFollow int `json:"series_follow"`
				Views        int `json:"views"`
			} `json:"stat"`
		} `json:"seasons"`
		Section []struct {
			EpisodeId int `json:"episode_id"`
			Episodes  []struct {
				Aid       int    `json:"aid"`
				Badge     string `json:"badge"`
				BadgeInfo struct {
					BgColor      string `json:"bg_color"`
					BgColorNight string `json:"bg_color_night"`
					Text         string `json:"text"`
				} `json:"badge_info"`
				BadgeType int    `json:"badge_type"`
				Bvid      string `json:"bvid"`
				Cid       int    `json:"cid"`
				Cover     string `json:"cover"`
				Dimension struct {
					Height int `json:"height"`
					Rotate int `json:"rotate"`
					Width  int `json:"width"`
				} `json:"dimension"`
				Duration    int    `json:"duration"`
				From        string `json:"from"`
				Id          int    `json:"id"`
				IsViewHide  bool   `json:"is_view_hide"`
				Link        string `json:"link"`
				LongTitle   string `json:"long_title"`
				PubTime     int    `json:"pub_time"`
				Pv          int    `json:"pv"`
				ReleaseDate string `json:"release_date"`
				Rights      struct {
					AllowDemand   int `json:"allow_demand"`
					AllowDm       int `json:"allow_dm"`
					AllowDownload int `json:"allow_download"`
					AreaLimit     int `json:"area_limit"`
				} `json:"rights"`
				ShareCopy string `json:"share_copy"`
				ShareUrl  string `json:"share_url"`
				ShortLink string `json:"short_link"`
				Stat      struct {
					Coin     int `json:"coin"`
					Danmakus int `json:"danmakus"`
					Likes    int `json:"likes"`
					Play     int `json:"play"`
					Reply    int `json:"reply"`
				} `json:"stat"`
				Status   int    `json:"status"`
				Subtitle string `json:"subtitle"`
				Title    string `json:"title"`
				Vid      string `json:"vid"`
			} `json:"episodes"`
			Id    int    `json:"id"`
			Title string `json:"title"`
			Type  int    `json:"type"`
		} `json:"section"`
		Series struct {
			SeriesId    int    `json:"series_id"`
			SeriesTitle string `json:"series_title"`
		} `json:"series"`
		ShareCopy     string `json:"share_copy"`
		ShareSubTitle string `json:"share_sub_title"`
		ShareUrl      string `json:"share_url"`
		Show          struct {
			WideScreen int `json:"wide_screen"`
		} `json:"show"`
		SquareCover string `json:"square_cover"`
		Stat        struct {
			Coins     int `json:"coins"`
			Danmakus  int `json:"danmakus"`
			Favorite  int `json:"favorite"`
			Favorites int `json:"favorites"`
			Likes     int `json:"likes"`
			Reply     int `json:"reply"`
			Share     int `json:"share"`
			Views     int `json:"views"`
		} `json:"stat"`
		Status   int    `json:"status"`
		Subtitle string `json:"subtitle"`
		Title    string `json:"title"`
		Total    int    `json:"total"`
		Type     int    `json:"type"`
		UpInfo   struct {
			Avatar   string `json:"avatar"`
			Follower int    `json:"follower"`
			IsFollow int    `json:"is_follow"`
			Mid      int    `json:"mid"`
			Pendant  struct {
				Image string `json:"image"`
				Name  string `json:"name"`
				Pid   int    `json:"pid"`
			} `json:"pendant"`
			ThemeType  int    `json:"theme_type"`
			Uname      string `json:"uname"`
			VerifyType int    `json:"verify_type"`
			VipLabel   struct {
				LabelTheme string `json:"label_theme"`
				Path       string `json:"path"`
				Text       string `json:"text"`
			} `json:"vip_label"`
			VipStatus int `json:"vip_status"`
			VipType   int `json:"vip_type"`
		} `json:"up_info"`
		UserStatus struct {
			AreaLimit    int `json:"area_limit"`
			BanAreaShow  int `json:"ban_area_show"`
			Follow       int `json:"follow"`
			FollowStatus int `json:"follow_status"`
			Login        int `json:"login"`
			Pay          int `json:"pay"`
			PayPackPaid  int `json:"pay_pack_paid"`
			Sponsor      int `json:"sponsor"`
		} `json:"user_status"`
	} `json:"result"`
}
