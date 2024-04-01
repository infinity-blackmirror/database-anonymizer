# Database Anonimize

**Database Anonymizer** is a tool written in GO that allows **anonymizing or deleting data from a MySQL or PostgreSQL database**.

It caters to various use cases such as **providing developers with an anonymized copy of a database** or **fulfilling the need to anonymize or delete data in compliance with GDPR (General Data Protection Regulation) requirements**, based on retention periods defined in the treatment register.

The project includes a vast array of random data generators. It also enables data generation via Twig-written templates. You can specify precise rules for each table or global rules applied to all tables in your configuration.

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
