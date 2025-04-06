package response

type Response struct {
	ID         int64  `json:"id"`
	AnimeID    int64  `json:"anime_id"`
	AnimeArt   string `json:"anime_image"`
	Acimage    string `json:"character_image"`
	AnimeTitle string `json:"origin"`
	Cgender    string `json:"gender"`
	Cname      string `json:"name"`
	Cdesc      string `json:"desc"`
}
