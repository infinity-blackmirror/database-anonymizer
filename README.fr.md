# Database Anonimizer

**Database Anonimizer** est un outil écrit en GO et qui permet **d'anonymiser ou supprimer des données** d'une base de données **MySQL** ou **PostgreSQL**.

Il répond à plusieurs cas d'usage. Il **permet de transférer une copie de base de données anonymisée à des développeurs et des développeuses** ou répondre à la nécessité d'**anonymiser ou supprimer des données dans le cadre du RGPD** (Règlement général sur la protection des données) selon les durées de conservation définies dans le registre des traitements.

Le projet inclue une grande quantité de générateurs de données aléatoires. Il permet également de générer des données via des modèles écrits en Twig. Vous pouvez spécifier des règles précises pour chaque tables ou bien des règles globales appliquées sur chacunes des tables de votre configuration.

## Usage

### Configuration

La configuration est écrite en YAML. Voici un exemple complet :

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

Pour afficher l'aide, utiliser `-h` :

```
database-anonymizer -h
```

Voici des exemples pour MySQL et PostgreSQL :

```
database-anonymizer --dsn "mysql://username:password@tcp(db_host)/db_name" --schema ./schema.yaml
database-anonymizer --dsn "postgres://postgres:postgres@localhost:5432/test" --schema ./schema.yaml
```

### Liste des générateurs

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
