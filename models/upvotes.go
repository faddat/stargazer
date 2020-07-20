// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Upvote is an object representing the database table.
type Upvote struct {
	ID              int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	VendorID        int       `boil:"vendor_id" json:"vendorID" toml:"vendorID" yaml:"vendorID"`
	PostID          string    `boil:"post_id" json:"postID" toml:"postID" yaml:"postID"`
	Creator         string    `boil:"creator" json:"creator" toml:"creator" yaml:"creator"`
	RewardAddress   string    `boil:"reward_address" json:"rewardAddress" toml:"rewardAddress" yaml:"rewardAddress"`
	VoteNumber      int       `boil:"vote_number" json:"voteNumber" toml:"voteNumber" yaml:"voteNumber"`
	VoteAmount      int64     `boil:"vote_amount" json:"voteAmount" toml:"voteAmount" yaml:"voteAmount"`
	VoteDenom       string    `boil:"vote_denom" json:"voteDenom" toml:"voteDenom" yaml:"voteDenom"`
	DepositAmount   int64     `boil:"deposit_amount" json:"depositAmount" toml:"depositAmount" yaml:"depositAmount"`
	DepositDenom    string    `boil:"deposit_denom" json:"depositDenom" toml:"depositDenom" yaml:"depositDenom"`
	Timestamp       time.Time `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	CurationEndTime time.Time `boil:"curation_end_time" json:"curationEndTime" toml:"curationEndTime" yaml:"curationEndTime"`
	Body            string    `boil:"body" json:"body" toml:"body" yaml:"body"`
	CreatedAt       time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt       time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	DeletedAt       null.Time `boil:"deleted_at" json:"deletedAt,omitempty" toml:"deletedAt" yaml:"deletedAt,omitempty"`

	R *upvoteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L upvoteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UpvoteColumns = struct {
	ID              string
	VendorID        string
	PostID          string
	Creator         string
	RewardAddress   string
	VoteNumber      string
	VoteAmount      string
	VoteDenom       string
	DepositAmount   string
	DepositDenom    string
	Timestamp       string
	CurationEndTime string
	Body            string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
}{
	ID:              "id",
	VendorID:        "vendor_id",
	PostID:          "post_id",
	Creator:         "creator",
	RewardAddress:   "reward_address",
	VoteNumber:      "vote_number",
	VoteAmount:      "vote_amount",
	VoteDenom:       "vote_denom",
	DepositAmount:   "deposit_amount",
	DepositDenom:    "deposit_denom",
	Timestamp:       "timestamp",
	CurationEndTime: "curation_end_time",
	Body:            "body",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// Generated where

var UpvoteWhere = struct {
	ID              whereHelperint64
	VendorID        whereHelperint
	PostID          whereHelperstring
	Creator         whereHelperstring
	RewardAddress   whereHelperstring
	VoteNumber      whereHelperint
	VoteAmount      whereHelperint64
	VoteDenom       whereHelperstring
	DepositAmount   whereHelperint64
	DepositDenom    whereHelperstring
	Timestamp       whereHelpertime_Time
	CurationEndTime whereHelpertime_Time
	Body            whereHelperstring
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpertime_Time
	DeletedAt       whereHelpernull_Time
}{
	ID:              whereHelperint64{field: "\"upvotes\".\"id\""},
	VendorID:        whereHelperint{field: "\"upvotes\".\"vendor_id\""},
	PostID:          whereHelperstring{field: "\"upvotes\".\"post_id\""},
	Creator:         whereHelperstring{field: "\"upvotes\".\"creator\""},
	RewardAddress:   whereHelperstring{field: "\"upvotes\".\"reward_address\""},
	VoteNumber:      whereHelperint{field: "\"upvotes\".\"vote_number\""},
	VoteAmount:      whereHelperint64{field: "\"upvotes\".\"vote_amount\""},
	VoteDenom:       whereHelperstring{field: "\"upvotes\".\"vote_denom\""},
	DepositAmount:   whereHelperint64{field: "\"upvotes\".\"deposit_amount\""},
	DepositDenom:    whereHelperstring{field: "\"upvotes\".\"deposit_denom\""},
	Timestamp:       whereHelpertime_Time{field: "\"upvotes\".\"timestamp\""},
	CurationEndTime: whereHelpertime_Time{field: "\"upvotes\".\"curation_end_time\""},
	Body:            whereHelperstring{field: "\"upvotes\".\"body\""},
	CreatedAt:       whereHelpertime_Time{field: "\"upvotes\".\"created_at\""},
	UpdatedAt:       whereHelpertime_Time{field: "\"upvotes\".\"updated_at\""},
	DeletedAt:       whereHelpernull_Time{field: "\"upvotes\".\"deleted_at\""},
}

// UpvoteRels is where relationship names are stored.
var UpvoteRels = struct {
}{}

// upvoteR is where relationships are stored.
type upvoteR struct {
}

// NewStruct creates a new relationship struct
func (*upvoteR) NewStruct() *upvoteR {
	return &upvoteR{}
}

// upvoteL is where Load methods for each relationship are stored.
type upvoteL struct{}

var (
	upvoteAllColumns            = []string{"id", "vendor_id", "post_id", "creator", "reward_address", "vote_number", "vote_amount", "vote_denom", "deposit_amount", "deposit_denom", "timestamp", "curation_end_time", "body", "created_at", "updated_at", "deleted_at"}
	upvoteColumnsWithoutDefault = []string{"vendor_id", "post_id", "creator", "reward_address", "vote_number", "vote_amount", "vote_denom", "deposit_amount", "deposit_denom", "timestamp", "curation_end_time", "body", "deleted_at"}
	upvoteColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	upvotePrimaryKeyColumns     = []string{"id"}
)

type (
	// UpvoteSlice is an alias for a slice of pointers to Upvote.
	// This should generally be used opposed to []Upvote.
	UpvoteSlice []*Upvote
	// UpvoteHook is the signature for custom Upvote hook methods
	UpvoteHook func(context.Context, boil.ContextExecutor, *Upvote) error

	upvoteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	upvoteType                 = reflect.TypeOf(&Upvote{})
	upvoteMapping              = queries.MakeStructMapping(upvoteType)
	upvotePrimaryKeyMapping, _ = queries.BindMapping(upvoteType, upvoteMapping, upvotePrimaryKeyColumns)
	upvoteInsertCacheMut       sync.RWMutex
	upvoteInsertCache          = make(map[string]insertCache)
	upvoteUpdateCacheMut       sync.RWMutex
	upvoteUpdateCache          = make(map[string]updateCache)
	upvoteUpsertCacheMut       sync.RWMutex
	upvoteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var upvoteBeforeInsertHooks []UpvoteHook
var upvoteBeforeUpdateHooks []UpvoteHook
var upvoteBeforeDeleteHooks []UpvoteHook
var upvoteBeforeUpsertHooks []UpvoteHook

var upvoteAfterInsertHooks []UpvoteHook
var upvoteAfterSelectHooks []UpvoteHook
var upvoteAfterUpdateHooks []UpvoteHook
var upvoteAfterDeleteHooks []UpvoteHook
var upvoteAfterUpsertHooks []UpvoteHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Upvote) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range upvoteBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Upvote) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range upvoteBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Upvote) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range upvoteBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Upvote) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range upvoteBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Upvote) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range upvoteAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Upvote) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range upvoteAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Upvote) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range upvoteAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Upvote) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range upvoteAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Upvote) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range upvoteAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUpvoteHook registers your hook function for all future operations.
func AddUpvoteHook(hookPoint boil.HookPoint, upvoteHook UpvoteHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		upvoteBeforeInsertHooks = append(upvoteBeforeInsertHooks, upvoteHook)
	case boil.BeforeUpdateHook:
		upvoteBeforeUpdateHooks = append(upvoteBeforeUpdateHooks, upvoteHook)
	case boil.BeforeDeleteHook:
		upvoteBeforeDeleteHooks = append(upvoteBeforeDeleteHooks, upvoteHook)
	case boil.BeforeUpsertHook:
		upvoteBeforeUpsertHooks = append(upvoteBeforeUpsertHooks, upvoteHook)
	case boil.AfterInsertHook:
		upvoteAfterInsertHooks = append(upvoteAfterInsertHooks, upvoteHook)
	case boil.AfterSelectHook:
		upvoteAfterSelectHooks = append(upvoteAfterSelectHooks, upvoteHook)
	case boil.AfterUpdateHook:
		upvoteAfterUpdateHooks = append(upvoteAfterUpdateHooks, upvoteHook)
	case boil.AfterDeleteHook:
		upvoteAfterDeleteHooks = append(upvoteAfterDeleteHooks, upvoteHook)
	case boil.AfterUpsertHook:
		upvoteAfterUpsertHooks = append(upvoteAfterUpsertHooks, upvoteHook)
	}
}

// One returns a single upvote record from the query.
func (q upvoteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Upvote, error) {
	o := &Upvote{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for upvotes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Upvote records from the query.
func (q upvoteQuery) All(ctx context.Context, exec boil.ContextExecutor) (UpvoteSlice, error) {
	var o []*Upvote

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Upvote slice")
	}

	if len(upvoteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Upvote records in the query.
func (q upvoteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count upvotes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q upvoteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if upvotes exists")
	}

	return count > 0, nil
}

// Upvotes retrieves all the records using an executor.
func Upvotes(mods ...qm.QueryMod) upvoteQuery {
	mods = append(mods, qm.From("\"upvotes\""))
	return upvoteQuery{NewQuery(mods...)}
}

// FindUpvote retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUpvote(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Upvote, error) {
	upvoteObj := &Upvote{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"upvotes\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, upvoteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from upvotes")
	}

	return upvoteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Upvote) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no upvotes provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(upvoteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	upvoteInsertCacheMut.RLock()
	cache, cached := upvoteInsertCache[key]
	upvoteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			upvoteAllColumns,
			upvoteColumnsWithDefault,
			upvoteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(upvoteType, upvoteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(upvoteType, upvoteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"upvotes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"upvotes\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into upvotes")
	}

	if !cached {
		upvoteInsertCacheMut.Lock()
		upvoteInsertCache[key] = cache
		upvoteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Upvote.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Upvote) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	upvoteUpdateCacheMut.RLock()
	cache, cached := upvoteUpdateCache[key]
	upvoteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			upvoteAllColumns,
			upvotePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update upvotes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"upvotes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, upvotePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(upvoteType, upvoteMapping, append(wl, upvotePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update upvotes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for upvotes")
	}

	if !cached {
		upvoteUpdateCacheMut.Lock()
		upvoteUpdateCache[key] = cache
		upvoteUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q upvoteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for upvotes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for upvotes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UpvoteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), upvotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"upvotes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, upvotePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in upvote slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all upvote")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Upvote) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no upvotes provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(upvoteColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	upvoteUpsertCacheMut.RLock()
	cache, cached := upvoteUpsertCache[key]
	upvoteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			upvoteAllColumns,
			upvoteColumnsWithDefault,
			upvoteColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			upvoteAllColumns,
			upvotePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert upvotes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(upvotePrimaryKeyColumns))
			copy(conflict, upvotePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"upvotes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(upvoteType, upvoteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(upvoteType, upvoteMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert upvotes")
	}

	if !cached {
		upvoteUpsertCacheMut.Lock()
		upvoteUpsertCache[key] = cache
		upvoteUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Upvote record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Upvote) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Upvote provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), upvotePrimaryKeyMapping)
	sql := "DELETE FROM \"upvotes\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from upvotes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for upvotes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q upvoteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no upvoteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from upvotes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for upvotes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UpvoteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(upvoteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), upvotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"upvotes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, upvotePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from upvote slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for upvotes")
	}

	if len(upvoteAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Upvote) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUpvote(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UpvoteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UpvoteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), upvotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"upvotes\".* FROM \"upvotes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, upvotePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UpvoteSlice")
	}

	*o = slice

	return nil
}

// UpvoteExists checks if the Upvote row exists.
func UpvoteExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"upvotes\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if upvotes exists")
	}

	return exists, nil
}
