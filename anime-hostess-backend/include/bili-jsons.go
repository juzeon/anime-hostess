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
