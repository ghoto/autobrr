package domain

type IndexerRepo interface {
	Store(indexer Indexer) (*Indexer, error)
	Update(indexer Indexer) (*Indexer, error)
	List() ([]Indexer, error)
	Delete(id int) error
	FindByFilterID(id int) ([]Indexer, error)
}

type Indexer struct {
	ID         int               `json:"id"`
	Name       string            `json:"name"`
	Identifier string            `json:"identifier"`
	Enabled    bool              `json:"enabled"`
	Type       string            `json:"type,omitempty"`
	Settings   map[string]string `json:"settings,omitempty"`
}

type IndexerDefinition struct {
	ID          int               `json:"id,omitempty"`
	Name        string            `json:"name"`
	Identifier  string            `json:"identifier"`
	Enabled     bool              `json:"enabled,omitempty"`
	Description string            `json:"description"`
	Language    string            `json:"language"`
	Privacy     string            `json:"privacy"`
	Protocol    string            `json:"protocol"`
	URLS        []string          `json:"urls"`
	Settings    []IndexerSetting  `json:"settings"`
	SettingsMap map[string]string `json:"-"`
	IRC         *IndexerIRC       `json:"irc"`
	Parse       IndexerParse      `json:"parse"`
}

type IndexerSetting struct {
	Name        string `json:"name"`
	Required    bool   `json:"required,omitempty"`
	Type        string `json:"type"`
	Value       string `json:"value,omitempty"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Regex       string `json:"regex,omitempty"`
}

type IndexerIRC struct {
	Network    string
	Server     string
	Channels   []string
	Announcers []string
}

type IndexerParse struct {
	Type  string                `json:"type"`
	Lines []IndexerParseExtract `json:"lines"`
	Match IndexerParseMatch     `json:"match"`
}

type IndexerParseExtract struct {
	Test    []string `json:"test"`
	Pattern string   `json:"pattern"`
	Vars    []string `json:"vars"`
}

type IndexerParseMatch struct {
	TorrentURL string   `json:"torrenturl"`
	Encode     []string `json:"encode"`
}