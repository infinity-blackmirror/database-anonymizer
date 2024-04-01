# Database Anonimizer

**Database Anonymizer** is a tool written in GO that allows **anonymizing or deleting data from a MySQL or PostgreSQL database**.

It addresses various use cases such as **providing developers with an anonymized copy of a database** or **satisfying the need to anonymize or delete data in accordance with GDPR (General Data Protection Regulation) requirements. data protection)**, depending on the retention periods. defined in the processing register.

The project includes a vast array of fakers. It also enables data generation via Twig-written templates. You can specify precise rules for each table or global rules applied to all tables in your configuration.

## Usage

### Configuration

The configuration is written in YAML. Here's a complete example:

```
rules:
  columns:
    phone: phone_e164number
  generators:
    person_name: [display_name]
  actions:
    - table: user
      virtual_columns:
        domain: internet_domain
      columns:
        firstname: person_firstname
        lastname: person_lastname
        email: "{{ (firstname ~ '.' ~ lastname ~ '@' ~ domain)|lower }}"
    - table: company
      columns:
        name: company_name
    - table: access_log
      query: 'select * from access_log where date < (NOW() - INTERVAL 6 MONTH)'
      delete: true
    - table: user_ip
      primary_key: [user_id, ip_id]
      delete: true
```

### Exécution

To display help, use `-h`:

```
database-anonymizer -h
```

Here are examples for MySQL and PostgreSQL:

```
database-anonymizer --dsn "mysql://username:password@tcp(db_host)/db_name" --schema ./schema.yaml
database-anonymizer --dsn "postgres://postgres:postgres@localhost:5432/test" --schema ./schema.yaml
```

### List of fakers

- address
- address_buildingnumber
- address_city
- address_cityprefix
- address_citysuffix
- address_country
- address_countryabbr
- address_countrycode
- address_latitude
- address_longitude
- address_postcode
- address_secondaryaddress
- address_state
- address_stateabbr
- address_streetaddress
- address_streetname
- address_streetsuffix
- app_name
- app_version
- beer_alcohol
- beer_blg
- beer_hop
- beer_ibu
- beer_malt
- beer_name
- beer_style
- blood_name
- boolean_bool
- car_category
- car_fueltype
- car_maker
- car_model
- car_plate
- car_transmissiongear
- color_css
- color_colorname
- color_hex
- color_rgb
- color_safecolorname
- company_bs
- company_catchphrase
- company_ein
- company_jobtitle
- company_name
- company_suffix
- crypto_bech32address
- crypto_bitcoinaddress
- crypto_etheriumaddress
- crypto_p2pkhaddress
- crypto_p2shaddress
- currency_code
- currency_country
- currency_currency
- currency_number
- emoji_emoji
- emoji_emojicode
- file_extension
- file_filenamewithextension
- food_fruit
- food_vegetable
- gamer_tag
- gender_abbr
- gender_name
- genre_name
- internet_companyemail
- internet_domain
- internet_email
- internet_freeemail
- internet_freeemaildomain
- internet_httpmethod
- internet_ipv4
- internet_ipv6
- internet_localipv4
- internet_macaddress
- internet_password
- internet_query
- internet_safeemail
- internet_slug
- internet_statuscode
- internet_statuscodemessage
- internet_statuscodewithmessage
- internet_tld
- internet_url
- internet_user
- language_language
- language_languageabbr
- language_programminglanguage
- mimetype_mimetype
- music_genre
- music_name
- payment_creditcardexpirationdatestring
- payment_creditcardnumber
- payment_creditcardtype
- person_firstname
- person_firstnamefemale
- person_firstnamemale
- person_gender
- person_lastname
- person_name
- person_namefemale
- person_namemale
- person_ssn
- person_suffix
- person_title
- pet_cat
- pet_dog
- pet_name
- phone_areacode
- phone_e164number
- phone_exchangecode
- phone_number
- phone_tollfreeareacode
- phone_toolfreenumber
- time_ampm
- time_century
- time_dayofmonth
- time_monthname
- time_timezone
- time_year
- useragent_chrome
- useragent_firefox
- useragent_internetexplorer
- useragent_opera
- useragent_safari
- useragent_useragent
