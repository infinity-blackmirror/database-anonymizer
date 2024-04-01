package faker

import (
	"fmt"
	base_faker "github.com/jaswdr/faker"
	"strconv"
)

type FakeManager struct {
	Fakers map[string]func() string
}

func NewFakeManager() FakeManager {
	manager := FakeManager{}
	datas := make(map[string]func() string)

	fake := base_faker.New()

	datas["address"] = func() string { return fake.Address().Address() }
	datas["address_buildingnumber"] = func() string { return fake.Address().BuildingNumber() }
	datas["address_city"] = func() string { return fake.Address().City() }
	datas["address_cityprefix"] = func() string { return fake.Address().CityPrefix() }
	datas["address_citysuffix"] = func() string { return fake.Address().CitySuffix() }
	datas["address_country"] = func() string { return fake.Address().Country() }
	datas["address_countryabbr"] = func() string { return fake.Address().CountryAbbr() }
	datas["address_countrycode"] = func() string { return fake.Address().CountryCode() }
	datas["address_latitude"] = func() string { return fmt.Sprintf("%f", fake.Address().Latitude()) }
	datas["address_longitude"] = func() string { return fmt.Sprintf("%f", fake.Address().Longitude()) }
	datas["address_postcode"] = func() string { return fake.Address().PostCode() }
	datas["address_secondaryaddress"] = func() string { return fake.Address().SecondaryAddress() }
	datas["address_state"] = func() string { return fake.Address().State() }
	datas["address_stateabbr"] = func() string { return fake.Address().StateAbbr() }
	datas["address_streetaddress"] = func() string { return fake.Address().StreetAddress() }
	datas["address_streetname"] = func() string { return fake.Address().StreetName() }
	datas["address_streetsuffix"] = func() string { return fake.Address().StreetSuffix() }
	datas["app_name"] = func() string { return fake.App().Name() }
	datas["app_version"] = func() string { return fake.App().Version() }
	datas["beer_alcohol"] = func() string { return fake.Beer().Alcohol() }
	datas["beer_blg"] = func() string { return fake.Beer().Blg() }
	datas["beer_hop"] = func() string { return fake.Beer().Hop() }
	datas["beer_ibu"] = func() string { return fake.Beer().Ibu() }
	datas["beer_malt"] = func() string { return fake.Beer().Malt() }
	datas["beer_name"] = func() string { return fake.Beer().Name() }
	datas["beer_style"] = func() string { return fake.Beer().Style() }
	datas["blood_name"] = func() string { return fake.Blood().Name() }
	datas["boolean_bool"] = func() string {
		if fake.Boolean().Bool() {
			return "1"
		} else {
			return "0"
		}
	}
	datas["car_category"] = func() string { return fake.Car().Category() }
	datas["car_fueltype"] = func() string { return fake.Car().FuelType() }
	datas["car_maker"] = func() string { return fake.Car().Maker() }
	datas["car_model"] = func() string { return fake.Car().Model() }
	datas["car_plate"] = func() string { return fake.Car().Plate() }
	datas["car_transmissiongear"] = func() string { return fake.Car().TransmissionGear() }
	datas["color_css"] = func() string { return fake.Color().CSS() }
	datas["color_colorname"] = func() string { return fake.Color().ColorName() }
	datas["color_hex"] = func() string { return fake.Color().Hex() }
	datas["color_rgb"] = func() string { return fake.Color().RGB() }
	datas["color_safecolorname"] = func() string { return fake.Color().SafeColorName() }
	datas["company_bs"] = func() string { return fake.Company().BS() }
	datas["company_catchphrase"] = func() string { return fake.Company().CatchPhrase() }
	datas["company_ein"] = func() string { return strconv.Itoa(fake.Company().EIN()) }
	datas["company_jobtitle"] = func() string { return fake.Company().JobTitle() }
	datas["company_name"] = func() string { return fake.Company().Name() }
	datas["company_suffix"] = func() string { return fake.Company().Suffix() }
	datas["crypto_bech32address"] = func() string { return fake.Crypto().Bech32Address() }
	datas["crypto_bitcoinaddress"] = func() string { return fake.Crypto().BitcoinAddress() }
	datas["crypto_etheriumaddress"] = func() string { return fake.Crypto().EtheriumAddress() }
	datas["crypto_p2pkhaddress"] = func() string { return fake.Crypto().P2PKHAddress() }
	datas["crypto_p2shaddress"] = func() string { return fake.Crypto().P2SHAddress() }
	datas["currency_code"] = func() string { return fake.Currency().Code() }
	datas["currency_country"] = func() string { return fake.Currency().Country() }
	datas["currency_currency"] = func() string { return fake.Currency().Currency() }
	datas["currency_number"] = func() string { return strconv.Itoa(fake.Currency().Number()) }
	datas["emoji_emoji"] = func() string { return fake.Emoji().Emoji() }
	datas["emoji_emojicode"] = func() string { return fake.Emoji().EmojiCode() }
	datas["file_extension"] = func() string { return fake.File().Extension() }
	datas["file_filenamewithextension"] = func() string { return fake.File().FilenameWithExtension() }
	datas["food_fruit"] = func() string { return fake.Food().Fruit() }
	datas["food_vegetable"] = func() string { return fake.Food().Vegetable() }
	datas["gamer_tag"] = func() string { return fake.Gamer().Tag() }
	datas["gender_abbr"] = func() string { return fake.Gender().Abbr() }
	datas["gender_name"] = func() string { return fake.Gender().Name() }
	datas["genre_name"] = func() string { return fake.Genre().Name() }
	datas["internet_companyemail"] = func() string { return fake.Internet().CompanyEmail() }
	datas["internet_domain"] = func() string { return fake.Internet().Domain() }
	datas["internet_email"] = func() string { return fake.Internet().Email() }
	datas["internet_freeemail"] = func() string { return fake.Internet().FreeEmail() }
	datas["internet_freeemaildomain"] = func() string { return fake.Internet().FreeEmailDomain() }
	datas["internet_httpmethod"] = func() string { return fake.Internet().HTTPMethod() }
	datas["internet_ipv4"] = func() string { return fake.Internet().Ipv4() }
	datas["internet_ipv6"] = func() string { return fake.Internet().Ipv6() }
	datas["internet_localipv4"] = func() string { return fake.Internet().LocalIpv4() }
	datas["internet_macaddress"] = func() string { return fake.Internet().MacAddress() }
	datas["internet_password"] = func() string { return fake.Internet().Password() }
	datas["internet_query"] = func() string { return fake.Internet().Query() }
	datas["internet_safeemail"] = func() string { return fake.Internet().SafeEmail() }
	datas["internet_slug"] = func() string { return fake.Internet().Slug() }
	datas["internet_statuscode"] = func() string { return strconv.Itoa(fake.Internet().StatusCode()) }
	datas["internet_statuscodemessage"] = func() string { return fake.Internet().StatusCodeMessage() }
	datas["internet_statuscodewithmessage"] = func() string { return fake.Internet().StatusCodeWithMessage() }
	datas["internet_tld"] = func() string { return fake.Internet().TLD() }
	datas["internet_url"] = func() string { return fake.Internet().URL() }
	datas["internet_user"] = func() string { return fake.Internet().User() }
	datas["language_language"] = func() string { return fake.Language().Language() }
	datas["language_languageabbr"] = func() string { return fake.Language().LanguageAbbr() }
	datas["language_programminglanguage"] = func() string { return fake.Language().ProgrammingLanguage() }
	datas["mimetype_mimetype"] = func() string { return fake.MimeType().MimeType() }
	datas["music_genre"] = func() string { return fake.Music().Genre() }
	datas["music_name"] = func() string { return fake.Music().Name() }
	datas["payment_creditcardexpirationdatestring"] = func() string { return fake.Payment().CreditCardExpirationDateString() }
	datas["payment_creditcardnumber"] = func() string { return fake.Payment().CreditCardNumber() }
	datas["payment_creditcardtype"] = func() string { return fake.Payment().CreditCardType() }
	datas["person_firstname"] = func() string { return fake.Person().FirstName() }
	datas["person_firstnamefemale"] = func() string { return fake.Person().FirstNameFemale() }
	datas["person_firstnamemale"] = func() string { return fake.Person().FirstNameMale() }
	datas["person_gender"] = func() string { return fake.Person().Gender() }
	datas["person_lastname"] = func() string { return fake.Person().LastName() }
	datas["person_name"] = func() string { return fake.Person().Name() }
	datas["person_namefemale"] = func() string { return fake.Person().NameFemale() }
	datas["person_namemale"] = func() string { return fake.Person().NameMale() }
	datas["person_ssn"] = func() string { return fake.Person().SSN() }
	datas["person_suffix"] = func() string { return fake.Person().Suffix() }
	datas["person_title"] = func() string { return fake.Person().Title() }
	datas["pet_cat"] = func() string { return fake.Pet().Cat() }
	datas["pet_dog"] = func() string { return fake.Pet().Dog() }
	datas["pet_name"] = func() string { return fake.Pet().Name() }
	datas["phone_areacode"] = func() string { return fake.Phone().AreaCode() }
	datas["phone_e164number"] = func() string { return fake.Phone().E164Number() }
	datas["phone_exchangecode"] = func() string { return fake.Phone().ExchangeCode() }
	datas["phone_number"] = func() string { return fake.Phone().Number() }
	datas["phone_tollfreeareacode"] = func() string { return fake.Phone().TollFreeAreaCode() }
	datas["phone_toolfreenumber"] = func() string { return fake.Phone().ToolFreeNumber() }
	datas["time_ampm"] = func() string { return fake.Time().AmPm() }
	datas["time_century"] = func() string { return fake.Time().Century() }
	datas["time_dayofmonth"] = func() string { return strconv.Itoa(fake.Time().DayOfMonth()) }
	datas["time_monthname"] = func() string { return fake.Time().MonthName() }
	datas["time_timezone"] = func() string { return fake.Time().Timezone() }
	datas["time_year"] = func() string { return strconv.Itoa(fake.Time().Year()) }
	datas["useragent_chrome"] = func() string { return fake.UserAgent().Chrome() }
	datas["useragent_firefox"] = func() string { return fake.UserAgent().Firefox() }
	datas["useragent_internetexplorer"] = func() string { return fake.UserAgent().InternetExplorer() }
	datas["useragent_opera"] = func() string { return fake.UserAgent().Opera() }
	datas["useragent_safari"] = func() string { return fake.UserAgent().Safari() }
	datas["useragent_useragent"] = func() string { return fake.UserAgent().UserAgent() }

	manager.Fakers = datas

	return manager
}

func (f *FakeManager) IsValidFaker(name string) bool {
	if name == "" || name == "_" {
		return true
	}

	_, exists := f.Fakers[name]

	return exists
}
