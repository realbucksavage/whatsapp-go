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
	ID       string `json:"id,omitempty"`
	Link     string `json:"link,omitempty"`
	Caption  string `json:"caption,omitempty"`
	Filename string `json:"filename,omitempty"`
}

type Contact struct {
	Addresses []ContactAddress `json:"addresses,omitempty"`
	Birthday  string           `json:"birthday,omitempty"`
	Emails    []ContactEmail   `json:"emails,omitempty"`
	Name      *ContactName     `json:"name,omitempty"`
	Org       *ContactOrg      `json:"org,omitempty"`
	Phones    []ContactPhone   `json:"phones,omitempty"`
	URLs      []URL            `json:"URLs,omitempty"`
}

type ContactAddress struct {
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Type        string `json:"type,omitempty"`
}

type ContactEmail struct {
	Email string `json:"email,omitempty"`
	Type  string `json:"type,omitempty"`
}

type ContactName struct {
	FormattedName string `json:"formatted_name,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
	Suffix        string `json:"suffix,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
}

type ContactOrg struct {
	Company    string `json:"company,omitempty"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`
}

type ContactPhone struct {
	Phone string `json:"phone,omitempty"`
	Type  string `json:"type,omitempty"`
	WAID  string `json:"wa_id,omitempty"`
}

type Context struct {
	MessageID string `json:"message_id,omitempty"`
}

type Interactive struct {
	Action *InteractiveAction `json:"action,omitempty"`
	Body   *TextString        `json:"body,omitempty"`
	Footer *TextString        `json:"footer,omitempty"`
	Header *InteractiveHeader `json:"header,omitempty"`
	Type   InteractiveType    `json:"type,omitempty"`
}

type InteractiveAction struct {
	Button             string               `json:"button,omitempty"`
	Buttons            []Button             `json:"buttons,omitempty"`
	CatalogID          string               `json:"catalog_id,omitempty"`
	ProductRetailerID  string               `json:"product_retailer_id,omitempty"`
	Sections           []InteractiveSection `json:"sections,omitempty"`
	Mode               string               `json:"mode,omitempty"`
	FlowMessageVersion string               `json:"flow_message_version,omitempty"`
	FlowToken          string               `json:"flow_token,omitempty"`
	FlowID             string               `json:"flow_id,omitempty"`
	FlowCTA            string               `json:"flow_cta,omitempty"`
	FlowAction         string               `json:"flow_action,omitempty"`
	FlowActionPayload  string               `json:"flow_action_payload,omitempty"`
}

type Button struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
}

type InteractiveSection struct {
	ProductItems []Product `json:"product_items,omitempty"`
	Rows         []Row     `json:"rows,omitempty"`
	Title        string    `json:"title,omitempty"`
}

type InteractiveHeader struct {
	Document *Media      `json:"document,omitempty"`
	Image    *Media      `json:"image,omitempty"`
	Text     *Media      `json:"text,omitempty"`
	Type     MessageType `json:"type,omitempty"`
	Video    *Media      `json:"video,omitempty"`
}

type Product struct {
	ProductRetailerID string `json:"product_retailer_id,omitempty"`
}

type Row struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type Location struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Name      string `json:"name,omitempty"`
	Address   string `json:"address,omitempty"`
}

type Template struct {
	Name       string      `json:"name,omitempty"`
	Language   *Language   `json:"language,omitempty"`
	Components []Component `json:"components,omitempty"`
}

type Component struct {
	Type       string      `json:"type,omitempty"`
	SubType    string      `json:"sub_type,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	Index      int         `json:"index,omitempty"`
}

type Language struct {
	Policy string `json:"policy,omitempty"`
	Code   string `json:"code,omitempty"`
}

type Parameter struct {
	Type     string    `json:"type,omitempty"`
	Name     string    `json:"name,omitempty"`
	Text     string    `json:"text,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
	DateTime *DateTime `json:"date_time,omitempty"`
	Image    *Media    `json:"image,omitempty"`
	Document *Media    `json:"document,omitempty"`
	Video    *Media    `json:"video,omitempty"`
}

type Currency struct {
	FallbackValue string  `json:"fallback_value,omitempty"`
	Code          string  `json:"code,omitempty"`
	Amount1000    float32 `json:"amount_1000,omitempty"`
}

type DateTime struct {
	FallbackValue string `json:"fallback_value,omitempty"`
}

type Text struct {
	Body       string `json:"body,omitempty"`
	PreviewURL bool   `json:"preview_url,omitempty"`
}

type URL struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

type TextString struct {
	Text string `json:"text,omitempty"`
}
