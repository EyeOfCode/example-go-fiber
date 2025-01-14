package dto

// ArtworkResponse represents the top-level response from the Art Institute API
type ArtworkResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []Artwork  `json:"data"`
	Info       Info       `json:"info"`
	Config     Config     `json:"config"`
}

// Pagination contains pagination information
type Pagination struct {
	Total       int    `json:"total"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
	TotalPages  int    `json:"total_pages"`
	CurrentPage int    `json:"current_page"`
	NextURL     string `json:"next_url"`
}

// Artwork represents a single artwork item
type Artwork struct {
	ID                          int       `json:"id"`
	APIModel                    string    `json:"api_model"`
	APILink                     string    `json:"api_link"`
	IsBoosted                   bool      `json:"is_boosted"`
	Title                       string    `json:"title"`
	AltTitles                   *string   `json:"alt_titles"`
	Thumbnail                   Thumbnail `json:"thumbnail"`
	MainReferenceNumber         string    `json:"main_reference_number"`
	HasNotBeenViewedMuch        bool      `json:"has_not_been_viewed_much"`
	BoostRank                   *int      `json:"boost_rank"`
	DateStart                   int       `json:"date_start"`
	DateEnd                     int       `json:"date_end"`
	DateDisplay                 string    `json:"date_display"`
	DateQualifierTitle          string    `json:"date_qualifier_title"`
	DateQualifierID             *int      `json:"date_qualifier_id"`
	ArtistDisplay               string    `json:"artist_display"`
	PlaceOfOrigin               string    `json:"place_of_origin"`
	Description                 *string   `json:"description"`
	Dimensions                  string    `json:"dimensions"`
	MediumDisplay               string    `json:"medium_display"`
	Inscriptions                *string   `json:"inscriptions"`
	CreditLine                  string    `json:"credit_line"`
	CatalogueDisplay            *string   `json:"catalogue_display"`
	PublicationHistory          *string   `json:"publication_history"`
	ExhibitionHistory           *string   `json:"exhibition_history"`
	ProvenanceText              *string   `json:"provenance_text"`
	PublishingVerificationLevel string    `json:"publishing_verification_level"`
	InternalDepartmentID        int       `json:"internal_department_id"`
	FiscalYear                  *int      `json:"fiscal_year"`
	FiscalYearDeaccession       *int      `json:"fiscal_year_deaccession"`
	IsPublicDomain              bool      `json:"is_public_domain"`
	IsZoomable                  bool      `json:"is_zoomable"`
	MaxZoomWindowSize           int       `json:"max_zoom_window_size"`
	CopyrightNotice             *string   `json:"copyright_notice"`
	HasMultimediaResources      bool      `json:"has_multimedia_resources"`
	HasEducationalResources     bool      `json:"has_educational_resources"`
	HasAdvancedImaging          bool      `json:"has_advanced_imaging"`
	Colorfulness                float64   `json:"colorfulness"`
	Color                       Color     `json:"color"`
	Latitude                    *float64  `json:"latitude"`
	Longitude                   *float64  `json:"longitude"`
	Latlon                      *string   `json:"latlon"`
	IsOnView                    bool      `json:"is_on_view"`
	OnLoanDisplay               *string   `json:"on_loan_display"`
	GalleryTitle                *string   `json:"gallery_title"`
	GalleryID                   *int      `json:"gallery_id"`
	ArtworkTypeTitle            string    `json:"artwork_type_title"`
	DepartmentTitle             string    `json:"department_title"`
	ArtistID                    int       `json:"artist_id"`
	ArtistTitle                 string    `json:"artist_title"`
	ImageID                     string    `json:"image_id"`
	StyleTitle                  *string   `json:"style_title"`
}

// Thumbnail represents the thumbnail image information
type Thumbnail struct {
	LQIP    string `json:"lqip"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	AltText string `json:"alt_text"`
}

// Color represents color information of the artwork
type Color struct {
	H          int     `json:"h"`
	L          int     `json:"l"`
	S          int     `json:"s"`
	Percentage float64 `json:"percentage"`
	Population int     `json:"population"`
}

// Info represents API information
type Info struct {
	LicenseText  string   `json:"license_text"`
	LicenseLinks []string `json:"license_links"`
	Version      string   `json:"version"`
}

// Config represents API configuration
type Config struct {
	IIIFURL    string `json:"iiif_url"`
	WebsiteURL string `json:"website_url"`
}
