package whatsapp

type MessageType string

const (
	TypeText        MessageType = "text"
	TypeAudio                   = "audio"
	TypeDocument                = "document"
	TypeImage                   = "image"
	TypeInteractive             = "interactive"
	TypeLocation                = "location"
	TypeSticker                 = "sticker"
	TypeTemplate                = "template"
)

type InteractiveType string

const (
	TypeButton         InteractiveType = "button"
	TypeCatalogMessage                 = "catalog_message"
	TypeList                           = "list"
	TypeProduct                        = "product"
	TypeProductList                    = "product_list"
	TypeFlow                           = "flow"
)

type Message struct {
	Type                  MessageType  `json:"type,omitempty"`
	Audio                 *Media       `json:"audio,omitempty"`
	BizOpaqueCallbackData string       `json:"biz_opaque_callback_data,omitempty"`
	Contacts              []Contact    `json:"contacts,omitempty"`
	Context               *Context     `json:"context,omitempty"`
	Document              *Media       `json:"document,omitempty"`
	Image                 *Media       `json:"image,omitempty"`
	Interactive           *Interactive `json:"interactive,omitempty"`
	Location              *Location    `json:"location,omitempty"`
	MessagingProduct      string       `json:"messaging_product,omitempty"`
	PreviewURL            bool         `json:"preview_url,omitempty"`
	Status                string       `json:"status,omitempty"`
	Sticker               *Media       `json:"sticker,omitempty"`
	Template              *Template    `json:"template,omitempty"`
	Text                  *Text        `json:"text,omitempty"`
	To                    string       `json:"to,omitempty"`
}

type Media struct {
	ID       string
	Link     string
	Caption  string
	Filename string
}

type Contact struct {
	Addresses []ContactAddress
	Birthday  string
	Emails    []ContactEmail
	Name      *ContactName
	Org       *ContactOrg
	Phones    []ContactPhone
	URLs      []URL
}

type ContactAddress struct {
	Street      string
	City        string
	Zip         string
	Country     string
	CountryCode string
	Type        string
}

type ContactEmail struct {
	Email string
	Type  string
}

type ContactName struct {
	FormattedName string
	FirstName     string
	LastName      string
	MiddleName    string
	Suffix        string
	Prefix        string
}

type ContactOrg struct {
	Company    string
	Department string
	Title      string
}

type ContactPhone struct {
	Phone string
	Type  string
	WAID  string
}

type Context struct {
	MessageID string
}

type Interactive struct {
	Action *InteractiveAction
	Body   *TextString
	Footer *TextString
	Header *InteractiveHeader
	Type   InteractiveType
}

type InteractiveAction struct {
	Button             string
	Buttons            []Button
	CatalogID          string
	ProductRetailerID  string
	Sections           []InteractiveSection
	Mode               string
	FlowMessageVersion string
	FlowToken          string
	FlowID             string
	FlowCTA            string
	FlowAction         string
	FlowActionPayload  string
}

type Button struct {
	ID    string
	Type  string
	Title string
}

type InteractiveSection struct {
	ProductItems []Product
	Rows         []Row
	Title        string
}

type InteractiveHeader struct {
	Document *Media
	Image    *Media
	Text     *Media
	Type     MessageType
	Video    *Media
}

type Product struct {
	ProductRetailerID string
}

type Row struct {
	ID          string
	Title       string
	Description string
}

type Location struct {
	Latitude  string
	Longitude string
	Name      string
	Address   string
}

type Template struct {
	Name     string
	Language *Language
}

type Component struct {
	Type       string
	SubType    string
	Parameters []Parameter
	Index      int
}

type Language struct {
	Policy string
	Code   string
}

type Parameter struct {
	Type     string
	Name     string
	Currency *Currency
	DateTime *DateTime
	Image    *Media
	Document *Media
	Video    *Media
}

type Currency struct {
	FallbackValue string
	Code          string
	Amount1000    float32
}

type DateTime struct {
	FallbackValue string
}

type Text struct {
	Body       string
	PreviewURL bool
}

type URL struct {
	URL  string
	Type string
}

type TextString struct {
	Text string
}
