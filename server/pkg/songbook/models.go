package songbook

type Song struct {
	Title    string   `json:"title"`
	Artist   string   `json:"artist"`
	Lyrics   []Stanza `json:"lyrics"`
	RootNote string   `json:"root_note,omitempty"`
}

type Stanza struct {
	Verses []Verse `json:"verses"`
	Type   string  `json:"type,omitempty"`
}

type Verse = []Word

type Word struct {
	Text  string   `json:"text"`
	Chord []string `json:"chords,omitempty"`
}

// GENIUS API MODELS

type SearchResult struct {
	Meta     SearchResultMeta     `json:"meta"`
	Response SearchResultResponse `json:"response"`
}

type SearchResultMeta struct {
	Status int `json:"status"`
}

type SearchResultResponse struct {
	Hits []Hits `json:"hits"`
}

type Hits struct {
	Highlights any    `json:"highlights"`
	Index      string `json:"index"`
	Type       string `json:"type"`
	Result     Hit    `json:"result"`
}

type Hit struct {
	Id                int    `json:"id"`
	Path              string `json:"path"`
	ApiPath           string `json:"api_path"`
	Url               string `json:"url"`
	RelationsIndexUlr string `json:"relationships_index_url"`

	Title         string `json:"title"`
	TitleFeatured string `json:"title_with_featured"`
	FullTitle     string `json:"full_title"`

	HeaderImageThumUrl string `json:"header_image_thumbnail_url"`
	HeaderImageUrl     string `json:"header_image_url"`
	SongImageThumbUrl  string `json:"song_art_image_thumbnail_url"`
	SongArtImageUrl    string `json:"song_art_image_url"`

	ReleaseDateComponents HitDate `json:"release_date_components"`
	ReleaseDateDisplay    string  `json:"release_date_for_display"`
	ReleaseDateAbbr       string  `json:"release_date_with_abbreviated_month_for_display"`

	AnnotationCount int      `json:"annotaion_count"`
	PyongsCount     int      `json:"pyongs_count"`
	Stats           HitStats `json:"stats"`

	LyricsOwnerId int    `json:"lyrics_owner_id"`
	LyricsSatte   string `json:"lyrics_state"`

	ArtistNames        string   `json:"artist_names"`
	PrimaryArtistNames string   `json:"primary_artist_names"`
	PrimaryArtist      Artist   `json:"primary_artist"`
	PrimaryArtists     []Artist `json:"primary_artists"`
	FeaturedArtists    []Artist `json:"featured_artists"`
}

type HitDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type HitStats struct {
	Hot                  bool `json:"hot"`
	PageViews            int  `json:"pageviews"`
	UnreviewdAnnotations int  `json:"unreviewed_annotations"`
}

type Artist struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Url            string `json:"url"`
	ApiPath        string `json:"api_path"`
	ImageUrl       string `json:"image_url"`
	HeaderImageUrl string `json:"header_image_url"`
	IsMemeVerified bool   `json:"is_meme_verified"`
	IsVerified     bool   `json:"is_verified"`
}
