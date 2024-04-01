# Database Anonimize

**Database Anonimizer** est un outil écrit en GO et qui permet **d'anonymiser ou supprimer des données** d'une base de données **MySQL** ou **PostgreSQL**.

Il répond à plusieurs cas d'usags comme le **permettre de transférer une copie de base de données anonymisée à des développeurs et des développeuses** ou répondre à la nécessité d'**anonymiser ou supprimer des données dans le cadre du RGPD** (Règlement général sur la protection des données) selon des durées de conservation définies dans un registre de traitement.

Le projet inclue une grande quantité de générateurs de données aléatoires. Il permet également de générer des données via des modèles écrits en Twig. Vous pouvez spécifier des règles précises pour chaque tables ou bien des règles globales appliquées sur chacunes des tables de votre configration.

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
