package v1

import (
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/MaksimYudenko/training/finalTask/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// TraineeServiceServer is implementation of v1.TraineeServiceServer proto interface
type TraineeServiceServer struct {
	db *sql.DB
}

// Create training service
func NewTraineeServiceServer(db *sql.DB) v1.TraineeServiceServer {
	return &TraineeServiceServer{db: db}
}

// CheckAPI checks if the API version requested by client is supported by server
func (s *TraineeServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: "+
					"service implements API version '%s', but asked for '%s'",
				apiVersion, api)
		}
	}
	return nil
}

// Connect returns SQL database connection from the pool
func (s *TraineeServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil,
			status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

// Create new trainee
func (s *TraineeServiceServer) Create(
	ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var id int64
	err = c.QueryRowContext(ctx, `INSERT INTO training(
                     name, surname, email, hasattend, score) 
                     VALUES($1,$2, $3,$4, $5) RETURNING id`,
		req.Trainee.Name, req.Trainee.Surname, req.Trainee.Email,
		req.Trainee.HasAttend, req.Trainee.Score).Scan(&id)
	if err != nil {
		return nil,
			status.Error(codes.Unknown, "failed to insert into training-> "+err.Error())
	}
	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Read trainee
func (s *TraineeServiceServer) Read(
	ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	// query trainee by ID
	rows, err := c.QueryContext(ctx,
		`SELECT id, name, surname, email, hasattend, score 
						FROM training WHERE id=$1`, req.Id)
	if err != nil {
		return nil,
			status.Error(codes.Unknown, "failed to select from trainee-> "+err.Error())
	}
	defer rows.Close()
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil,
				status.Error(codes.Unknown,
					"failed to retrieve data from trainee-> "+err.Error())
		}
		return nil,
			status.Error(codes.NotFound,
				fmt.Sprintf("Trainee with ID='%d' is not found", req.Id))
	}
	// get trainee data
	var td v1.Trainee
	if err := rows.Scan(
		&td.Id, &td.Name, &td.Surname, &td.Email, &td.HasAttend, &td.Score); err != nil {
		return nil,
			status.Error(codes.Unknown,
				"failed to retrieve field values from trainee row-> "+err.Error())
	}
	if rows.Next() {
		return nil, status.Error(codes.Unknown,
			fmt.Sprintf("found multiple trainee rows with ID='%d'", req.Id))
	}
	return &v1.ReadResponse{
		Api:     apiVersion,
		Trainee: &td,
	}, nil
}

// Update trainee
func (s *TraineeServiceServer) Update(
	ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	// update trainee
	res, err := c.ExecContext(ctx, `
					UPDATE training SET name=$2,surname=$3, email=$4, 
                    hasattend=$5, score=$6 WHERE id=$1`,
		req.Trainee.Id, req.Trainee.Name, req.Trainee.Surname,
		req.Trainee.Email, req.Trainee.HasAttend, req.Trainee.Score)
	if err != nil {
		return nil, status.Error(codes.Unknown,
			"failed to update trainee-> "+err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown,
			"failed to retrieve rows affected value-> "+err.Error())
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound,
			fmt.Sprintf("Trainee with ID='%d' is not found", req.Trainee.Id))
	}
	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete trainee
func (s *TraineeServiceServer) Delete(
	ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	// delete trainee
	res, err := c.ExecContext(ctx, `DELETE FROM training 
																	WHERE id=$1`, req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown,
			"failed to delete trainee-> "+err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown,
			"failed to retrieve rows affected value-> "+err.Error())
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound,
			fmt.Sprintf("Trainee with ID='%d' is not found", req.Id))
	}
	return &v1.DeleteResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}

// Read all trainees
func (s *TraineeServiceServer) ReadAll(
	ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	// get trainee list
	rows, err := c.QueryContext(ctx,
		`SELECT id, name, surname, email, hasattend, score 
						FROM training`)
	if err != nil {
		return nil, status.Error(codes.Unknown,
			"failed to select from trainee-> "+err.Error())
	}
	defer rows.Close()
	list := []*v1.Trainee{}
	for rows.Next() {
		td := new(v1.Trainee)
		if err := rows.Scan(&td.Id, &td.Name, &td.Surname, &td.Email,
			&td.HasAttend, &td.Score); err != nil {
			return nil, status.Error(codes.Unknown,
				"failed to retrieve field values from trainee row-> "+err.Error())
		}
		list = append(list, td)
	}
	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown,
			"failed to retrieve data from trainee-> "+err.Error())
	}
	return &v1.ReadAllResponse{
		Api:      apiVersion,
		Trainees: list,
	}, nil
}
