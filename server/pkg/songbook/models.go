package songbook

type Song struct {
	Title    string   `json:"title"`
	Artist   string   `json:"artist"`
	Lyrics   []Stanza `json:"lyrics"`
	RootNote string   `json:"root_note,omitempty"`
}

type Stanza struct {
	Verses []Verse `json:"verse"`
	Type   string  `json:"type,omitempty"`
}

type Verse = []Word

type Word struct {
	Text  string   `json:"text"`
	Chord []string `json:"chords,omitempty"`
}