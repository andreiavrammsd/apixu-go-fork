package apixu

// Config is used to set up the Apixu service instance with the following:
// 	ApiKey: API key generated on Apixu website
// 	Language: Language code for condition text
type Config struct {
	APIKey   string
	Language string
}
