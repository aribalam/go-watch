package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) getGenres(rows *sql.Rows) (map[int]string, error) {
	genres := make(map[int]string)
	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.ID] = mg.Genre.GenreName
	}

	return genres, nil
}

func (m *DBModel) Get(id int) (*Movie, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `select * from movies where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	query = "select mg.id, mg.movie_id, mg.genre_id, g.genre_name from movies_genres mg left join genres g on (g.id = mg.genre_id) where mg.movie_id = $1"
	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	genres, err := m.getGenres(rows)
	if err != nil {
		return nil, err
	}

	movie.MovieGenre = genres

	return &movie, nil
}

func (m *DBModel) All() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	query := "select * from movies order by title"
	rows, _ := m.DB.QueryContext(ctx, query)

	var movies []*Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		query = "select mg.id, mg.movie_id, mg.genre_id, g.genre_name from movies_genres mg left join genres g on (g.id = mg.genre_id) where mg.movie_id = $1"
		rows, _ := m.DB.QueryContext(ctx, query, movie.ID)
		defer rows.Close()

		genres, err := m.getGenres(rows)
		if err != nil {
			return nil, err
		}
		movie.MovieGenre = genres

		movies = append(movies, &movie)
	}

	return movies, nil
}
