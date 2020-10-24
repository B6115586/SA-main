// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/thanawat/app/ent/migrate"

	"github.com/thanawat/app/ent/problem"
	"github.com/thanawat/app/ent/problemstatus"
	"github.com/thanawat/app/ent/problemtitle"
	"github.com/thanawat/app/ent/room"
	"github.com/thanawat/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Problem is the client for interacting with the Problem builders.
	Problem *ProblemClient
	// ProblemStatus is the client for interacting with the ProblemStatus builders.
	ProblemStatus *ProblemStatusClient
	// ProblemTitle is the client for interacting with the ProblemTitle builders.
	ProblemTitle *ProblemTitleClient
	// Room is the client for interacting with the Room builders.
	Room *RoomClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Problem = NewProblemClient(c.config)
	c.ProblemStatus = NewProblemStatusClient(c.config)
	c.ProblemTitle = NewProblemTitleClient(c.config)
	c.Room = NewRoomClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Problem:       NewProblemClient(cfg),
		ProblemStatus: NewProblemStatusClient(cfg),
		ProblemTitle:  NewProblemTitleClient(cfg),
		Room:          NewRoomClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:        cfg,
		Problem:       NewProblemClient(cfg),
		ProblemStatus: NewProblemStatusClient(cfg),
		ProblemTitle:  NewProblemTitleClient(cfg),
		Room:          NewRoomClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Problem.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Problem.Use(hooks...)
	c.ProblemStatus.Use(hooks...)
	c.ProblemTitle.Use(hooks...)
	c.Room.Use(hooks...)
	c.User.Use(hooks...)
}

// ProblemClient is a client for the Problem schema.
type ProblemClient struct {
	config
}

// NewProblemClient returns a client for the Problem from the given config.
func NewProblemClient(c config) *ProblemClient {
	return &ProblemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `problem.Hooks(f(g(h())))`.
func (c *ProblemClient) Use(hooks ...Hook) {
	c.hooks.Problem = append(c.hooks.Problem, hooks...)
}

// Create returns a create builder for Problem.
func (c *ProblemClient) Create() *ProblemCreate {
	mutation := newProblemMutation(c.config, OpCreate)
	return &ProblemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Problem.
func (c *ProblemClient) Update() *ProblemUpdate {
	mutation := newProblemMutation(c.config, OpUpdate)
	return &ProblemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProblemClient) UpdateOne(pr *Problem) *ProblemUpdateOne {
	mutation := newProblemMutation(c.config, OpUpdateOne, withProblem(pr))
	return &ProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProblemClient) UpdateOneID(id int) *ProblemUpdateOne {
	mutation := newProblemMutation(c.config, OpUpdateOne, withProblemID(id))
	return &ProblemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Problem.
func (c *ProblemClient) Delete() *ProblemDelete {
	mutation := newProblemMutation(c.config, OpDelete)
	return &ProblemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProblemClient) DeleteOne(pr *Problem) *ProblemDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProblemClient) DeleteOneID(id int) *ProblemDeleteOne {
	builder := c.Delete().Where(problem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProblemDeleteOne{builder}
}

// Create returns a query builder for Problem.
func (c *ProblemClient) Query() *ProblemQuery {
	return &ProblemQuery{config: c.config}
}

// Get returns a Problem entity by its id.
func (c *ProblemClient) Get(ctx context.Context, id int) (*Problem, error) {
	return c.Query().Where(problem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProblemClient) GetX(ctx context.Context, id int) *Problem {
	pr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pr
}

// QueryUser queries the user edge of a Problem.
func (c *ProblemClient) QueryUser(pr *Problem) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problem.UserTable, problem.UserColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRoom queries the room edge of a Problem.
func (c *ProblemClient) QueryRoom(pr *Problem) *RoomQuery {
	query := &RoomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, id),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problem.RoomTable, problem.RoomColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProblemtitle queries the problemtitle edge of a Problem.
func (c *ProblemClient) QueryProblemtitle(pr *Problem) *ProblemTitleQuery {
	query := &ProblemTitleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, id),
			sqlgraph.To(problemtitle.Table, problemtitle.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problem.ProblemtitleTable, problem.ProblemtitleColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProblemstatus queries the problemstatus edge of a Problem.
func (c *ProblemClient) QueryProblemstatus(pr *Problem) *ProblemStatusQuery {
	query := &ProblemStatusQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(problem.Table, problem.FieldID, id),
			sqlgraph.To(problemstatus.Table, problemstatus.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, problem.ProblemstatusTable, problem.ProblemstatusColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProblemClient) Hooks() []Hook {
	return c.hooks.Problem
}

// ProblemStatusClient is a client for the ProblemStatus schema.
type ProblemStatusClient struct {
	config
}

// NewProblemStatusClient returns a client for the ProblemStatus from the given config.
func NewProblemStatusClient(c config) *ProblemStatusClient {
	return &ProblemStatusClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `problemstatus.Hooks(f(g(h())))`.
func (c *ProblemStatusClient) Use(hooks ...Hook) {
	c.hooks.ProblemStatus = append(c.hooks.ProblemStatus, hooks...)
}

// Create returns a create builder for ProblemStatus.
func (c *ProblemStatusClient) Create() *ProblemStatusCreate {
	mutation := newProblemStatusMutation(c.config, OpCreate)
	return &ProblemStatusCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for ProblemStatus.
func (c *ProblemStatusClient) Update() *ProblemStatusUpdate {
	mutation := newProblemStatusMutation(c.config, OpUpdate)
	return &ProblemStatusUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProblemStatusClient) UpdateOne(ps *ProblemStatus) *ProblemStatusUpdateOne {
	mutation := newProblemStatusMutation(c.config, OpUpdateOne, withProblemStatus(ps))
	return &ProblemStatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProblemStatusClient) UpdateOneID(id int) *ProblemStatusUpdateOne {
	mutation := newProblemStatusMutation(c.config, OpUpdateOne, withProblemStatusID(id))
	return &ProblemStatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ProblemStatus.
func (c *ProblemStatusClient) Delete() *ProblemStatusDelete {
	mutation := newProblemStatusMutation(c.config, OpDelete)
	return &ProblemStatusDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProblemStatusClient) DeleteOne(ps *ProblemStatus) *ProblemStatusDeleteOne {
	return c.DeleteOneID(ps.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProblemStatusClient) DeleteOneID(id int) *ProblemStatusDeleteOne {
	builder := c.Delete().Where(problemstatus.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProblemStatusDeleteOne{builder}
}

// Create returns a query builder for ProblemStatus.
func (c *ProblemStatusClient) Query() *ProblemStatusQuery {
	return &ProblemStatusQuery{config: c.config}
}

// Get returns a ProblemStatus entity by its id.
func (c *ProblemStatusClient) Get(ctx context.Context, id int) (*ProblemStatus, error) {
	return c.Query().Where(problemstatus.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProblemStatusClient) GetX(ctx context.Context, id int) *ProblemStatus {
	ps, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ps
}

// QueryProblemstatus2problem queries the problemstatus2problem edge of a ProblemStatus.
func (c *ProblemStatusClient) QueryProblemstatus2problem(ps *ProblemStatus) *ProblemQuery {
	query := &ProblemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ps.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(problemstatus.Table, problemstatus.FieldID, id),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, problemstatus.Problemstatus2problemTable, problemstatus.Problemstatus2problemColumn),
		)
		fromV = sqlgraph.Neighbors(ps.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProblemStatusClient) Hooks() []Hook {
	return c.hooks.ProblemStatus
}

// ProblemTitleClient is a client for the ProblemTitle schema.
type ProblemTitleClient struct {
	config
}

// NewProblemTitleClient returns a client for the ProblemTitle from the given config.
func NewProblemTitleClient(c config) *ProblemTitleClient {
	return &ProblemTitleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `problemtitle.Hooks(f(g(h())))`.
func (c *ProblemTitleClient) Use(hooks ...Hook) {
	c.hooks.ProblemTitle = append(c.hooks.ProblemTitle, hooks...)
}

// Create returns a create builder for ProblemTitle.
func (c *ProblemTitleClient) Create() *ProblemTitleCreate {
	mutation := newProblemTitleMutation(c.config, OpCreate)
	return &ProblemTitleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for ProblemTitle.
func (c *ProblemTitleClient) Update() *ProblemTitleUpdate {
	mutation := newProblemTitleMutation(c.config, OpUpdate)
	return &ProblemTitleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProblemTitleClient) UpdateOne(pt *ProblemTitle) *ProblemTitleUpdateOne {
	mutation := newProblemTitleMutation(c.config, OpUpdateOne, withProblemTitle(pt))
	return &ProblemTitleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProblemTitleClient) UpdateOneID(id int) *ProblemTitleUpdateOne {
	mutation := newProblemTitleMutation(c.config, OpUpdateOne, withProblemTitleID(id))
	return &ProblemTitleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ProblemTitle.
func (c *ProblemTitleClient) Delete() *ProblemTitleDelete {
	mutation := newProblemTitleMutation(c.config, OpDelete)
	return &ProblemTitleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProblemTitleClient) DeleteOne(pt *ProblemTitle) *ProblemTitleDeleteOne {
	return c.DeleteOneID(pt.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProblemTitleClient) DeleteOneID(id int) *ProblemTitleDeleteOne {
	builder := c.Delete().Where(problemtitle.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProblemTitleDeleteOne{builder}
}

// Create returns a query builder for ProblemTitle.
func (c *ProblemTitleClient) Query() *ProblemTitleQuery {
	return &ProblemTitleQuery{config: c.config}
}

// Get returns a ProblemTitle entity by its id.
func (c *ProblemTitleClient) Get(ctx context.Context, id int) (*ProblemTitle, error) {
	return c.Query().Where(problemtitle.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProblemTitleClient) GetX(ctx context.Context, id int) *ProblemTitle {
	pt, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pt
}

// QueryProblemtitle2problem queries the problemtitle2problem edge of a ProblemTitle.
func (c *ProblemTitleClient) QueryProblemtitle2problem(pt *ProblemTitle) *ProblemQuery {
	query := &ProblemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(problemtitle.Table, problemtitle.FieldID, id),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, problemtitle.Problemtitle2problemTable, problemtitle.Problemtitle2problemColumn),
		)
		fromV = sqlgraph.Neighbors(pt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProblemTitleClient) Hooks() []Hook {
	return c.hooks.ProblemTitle
}

// RoomClient is a client for the Room schema.
type RoomClient struct {
	config
}

// NewRoomClient returns a client for the Room from the given config.
func NewRoomClient(c config) *RoomClient {
	return &RoomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `room.Hooks(f(g(h())))`.
func (c *RoomClient) Use(hooks ...Hook) {
	c.hooks.Room = append(c.hooks.Room, hooks...)
}

// Create returns a create builder for Room.
func (c *RoomClient) Create() *RoomCreate {
	mutation := newRoomMutation(c.config, OpCreate)
	return &RoomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Room.
func (c *RoomClient) Update() *RoomUpdate {
	mutation := newRoomMutation(c.config, OpUpdate)
	return &RoomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoomClient) UpdateOne(r *Room) *RoomUpdateOne {
	mutation := newRoomMutation(c.config, OpUpdateOne, withRoom(r))
	return &RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoomClient) UpdateOneID(id int) *RoomUpdateOne {
	mutation := newRoomMutation(c.config, OpUpdateOne, withRoomID(id))
	return &RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Room.
func (c *RoomClient) Delete() *RoomDelete {
	mutation := newRoomMutation(c.config, OpDelete)
	return &RoomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RoomClient) DeleteOne(r *Room) *RoomDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RoomClient) DeleteOneID(id int) *RoomDeleteOne {
	builder := c.Delete().Where(room.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoomDeleteOne{builder}
}

// Create returns a query builder for Room.
func (c *RoomClient) Query() *RoomQuery {
	return &RoomQuery{config: c.config}
}

// Get returns a Room entity by its id.
func (c *RoomClient) Get(ctx context.Context, id int) (*Room, error) {
	return c.Query().Where(room.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoomClient) GetX(ctx context.Context, id int) *Room {
	r, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return r
}

// QueryUser queries the user edge of a Room.
func (c *RoomClient) QueryUser(r *Room) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(room.Table, room.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, room.UserTable, room.UserColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRoom2problem queries the room2problem edge of a Room.
func (c *RoomClient) QueryRoom2problem(r *Room) *ProblemQuery {
	query := &ProblemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(room.Table, room.FieldID, id),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, room.Room2problemTable, room.Room2problemColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoomClient) Hooks() []Hook {
	return c.hooks.Room
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryUser2room queries the user2room edge of a User.
func (c *UserClient) QueryUser2room(u *User) *RoomQuery {
	query := &RoomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.User2roomTable, user.User2roomColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser2problem queries the user2problem edge of a User.
func (c *UserClient) QueryUser2problem(u *User) *ProblemQuery {
	query := &ProblemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(problem.Table, problem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.User2problemTable, user.User2problemColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}