package datalayer

import (
    "database/sql"
    sq "github.com/Masterminds/squirrel"
    _ "github.com/lib/pq"
    "time"
)

type PostgresDataAdapter struct {
	db *sql.DB
}

func NewPostgresDataAdapter(dsn string) (*PostgresDataAdapter, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &PostgresDataAdapter{
		db: db,
	}, nil
}

func (p *PostgresDataAdapter) GetInnerSqlDb() *sql.DB {
    return p.db
}

func (p *PostgresDataAdapter) Ping() bool {
	return p.db.Ping() == nil
}

func (p *PostgresDataAdapter) Close() {
	p.Close()
}

// --------------

func (p *PostgresDataAdapter) GetHouseholdMemberById(id int) (*HouseholdMember, error) {
	users := sq.Select("id", "name", "birthday", "email").From("house_members").Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)
	row := users.RunWith(p.db).QueryRow()

	var member HouseholdMember

	if err := row.Scan(&member.id, &member.Name, &member.Birthday, &member.Email); err != nil {
	    return nil, err
    }

    member.Birthday = member.Birthday.In(time.UTC)
    return &member, nil
}