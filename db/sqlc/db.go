// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addAvailabilityStmt, err = db.PrepareContext(ctx, addAvailability); err != nil {
		return nil, fmt.Errorf("error preparing query AddAvailability: %w", err)
	}
	if q.createPhoneStmt, err = db.PrepareContext(ctx, createPhone); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePhone: %w", err)
	}
	if q.createProfessionalInformationStmt, err = db.PrepareContext(ctx, createProfessionalInformation); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProfessionalInformation: %w", err)
	}
	if q.createProfessionalUserStmt, err = db.PrepareContext(ctx, createProfessionalUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProfessionalUser: %w", err)
	}
	if q.createResponsibleStudentStmt, err = db.PrepareContext(ctx, createResponsibleStudent); err != nil {
		return nil, fmt.Errorf("error preparing query CreateResponsibleStudent: %w", err)
	}
	if q.createStudentUserStmt, err = db.PrepareContext(ctx, createStudentUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateStudentUser: %w", err)
	}
	if q.createSubjectMatterStmt, err = db.PrepareContext(ctx, createSubjectMatter); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSubjectMatter: %w", err)
	}
	if q.createSubjectMatterClassStmt, err = db.PrepareContext(ctx, createSubjectMatterClass); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSubjectMatterClass: %w", err)
	}
	if q.deleteAvailabilityStmt, err = db.PrepareContext(ctx, deleteAvailability); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAvailability: %w", err)
	}
	if q.deletePhoneStmt, err = db.PrepareContext(ctx, deletePhone); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePhone: %w", err)
	}
	if q.deleteProfessionalInformationStmt, err = db.PrepareContext(ctx, deleteProfessionalInformation); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProfessionalInformation: %w", err)
	}
	if q.deleteProfessionalUserStmt, err = db.PrepareContext(ctx, deleteProfessionalUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProfessionalUser: %w", err)
	}
	if q.deleteResponsibleStudentStmt, err = db.PrepareContext(ctx, deleteResponsibleStudent); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteResponsibleStudent: %w", err)
	}
	if q.deleteStudentUserStmt, err = db.PrepareContext(ctx, deleteStudentUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteStudentUser: %w", err)
	}
	if q.deleteSubjectMatterStmt, err = db.PrepareContext(ctx, deleteSubjectMatter); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSubjectMatter: %w", err)
	}
	if q.deleteSubjectMatterClassStmt, err = db.PrepareContext(ctx, deleteSubjectMatterClass); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSubjectMatterClass: %w", err)
	}
	if q.getAvailabilityStmt, err = db.PrepareContext(ctx, getAvailability); err != nil {
		return nil, fmt.Errorf("error preparing query GetAvailability: %w", err)
	}
	if q.getPhoneStmt, err = db.PrepareContext(ctx, getPhone); err != nil {
		return nil, fmt.Errorf("error preparing query GetPhone: %w", err)
	}
	if q.getProfessionalInformationStmt, err = db.PrepareContext(ctx, getProfessionalInformation); err != nil {
		return nil, fmt.Errorf("error preparing query GetProfessionalInformation: %w", err)
	}
	if q.getProfessionalUserStmt, err = db.PrepareContext(ctx, getProfessionalUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetProfessionalUser: %w", err)
	}
	if q.getResponsibleStudentStmt, err = db.PrepareContext(ctx, getResponsibleStudent); err != nil {
		return nil, fmt.Errorf("error preparing query GetResponsibleStudent: %w", err)
	}
	if q.getStudentUserStmt, err = db.PrepareContext(ctx, getStudentUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetStudentUser: %w", err)
	}
	if q.getSubjectMatterStmt, err = db.PrepareContext(ctx, getSubjectMatter); err != nil {
		return nil, fmt.Errorf("error preparing query GetSubjectMatter: %w", err)
	}
	if q.getSubjectMatterClassStmt, err = db.PrepareContext(ctx, getSubjectMatterClass); err != nil {
		return nil, fmt.Errorf("error preparing query GetSubjectMatterClass: %w", err)
	}
	if q.listAvailabilityStmt, err = db.PrepareContext(ctx, listAvailability); err != nil {
		return nil, fmt.Errorf("error preparing query ListAvailability: %w", err)
	}
	if q.listPhoneStmt, err = db.PrepareContext(ctx, listPhone); err != nil {
		return nil, fmt.Errorf("error preparing query ListPhone: %w", err)
	}
	if q.listProfessionalInformationStmt, err = db.PrepareContext(ctx, listProfessionalInformation); err != nil {
		return nil, fmt.Errorf("error preparing query ListProfessionalInformation: %w", err)
	}
	if q.listProfessionalInformationByUserStmt, err = db.PrepareContext(ctx, listProfessionalInformationByUser); err != nil {
		return nil, fmt.Errorf("error preparing query ListProfessionalInformationByUser: %w", err)
	}
	if q.listProfessionalUserStmt, err = db.PrepareContext(ctx, listProfessionalUser); err != nil {
		return nil, fmt.Errorf("error preparing query ListProfessionalUser: %w", err)
	}
	if q.listResponsibleStudentStmt, err = db.PrepareContext(ctx, listResponsibleStudent); err != nil {
		return nil, fmt.Errorf("error preparing query ListResponsibleStudent: %w", err)
	}
	if q.listStudentUserStmt, err = db.PrepareContext(ctx, listStudentUser); err != nil {
		return nil, fmt.Errorf("error preparing query ListStudentUser: %w", err)
	}
	if q.listSubjectMatterStmt, err = db.PrepareContext(ctx, listSubjectMatter); err != nil {
		return nil, fmt.Errorf("error preparing query ListSubjectMatter: %w", err)
	}
	if q.listSubjectMatterClassStmt, err = db.PrepareContext(ctx, listSubjectMatterClass); err != nil {
		return nil, fmt.Errorf("error preparing query ListSubjectMatterClass: %w", err)
	}
	if q.updatePhoneStmt, err = db.PrepareContext(ctx, updatePhone); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePhone: %w", err)
	}
	if q.updateProfessionalInformationStmt, err = db.PrepareContext(ctx, updateProfessionalInformation); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProfessionalInformation: %w", err)
	}
	if q.updateProfessionalUserStmt, err = db.PrepareContext(ctx, updateProfessionalUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProfessionalUser: %w", err)
	}
	if q.updateResponsibleStudentStmt, err = db.PrepareContext(ctx, updateResponsibleStudent); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateResponsibleStudent: %w", err)
	}
	if q.updateStudentUserStmt, err = db.PrepareContext(ctx, updateStudentUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateStudentUser: %w", err)
	}
	if q.updateSubjectMatterStmt, err = db.PrepareContext(ctx, updateSubjectMatter); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSubjectMatter: %w", err)
	}
	if q.updateSubjectMatterClassStmt, err = db.PrepareContext(ctx, updateSubjectMatterClass); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSubjectMatterClass: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addAvailabilityStmt != nil {
		if cerr := q.addAvailabilityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addAvailabilityStmt: %w", cerr)
		}
	}
	if q.createPhoneStmt != nil {
		if cerr := q.createPhoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPhoneStmt: %w", cerr)
		}
	}
	if q.createProfessionalInformationStmt != nil {
		if cerr := q.createProfessionalInformationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProfessionalInformationStmt: %w", cerr)
		}
	}
	if q.createProfessionalUserStmt != nil {
		if cerr := q.createProfessionalUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProfessionalUserStmt: %w", cerr)
		}
	}
	if q.createResponsibleStudentStmt != nil {
		if cerr := q.createResponsibleStudentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createResponsibleStudentStmt: %w", cerr)
		}
	}
	if q.createStudentUserStmt != nil {
		if cerr := q.createStudentUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStudentUserStmt: %w", cerr)
		}
	}
	if q.createSubjectMatterStmt != nil {
		if cerr := q.createSubjectMatterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSubjectMatterStmt: %w", cerr)
		}
	}
	if q.createSubjectMatterClassStmt != nil {
		if cerr := q.createSubjectMatterClassStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSubjectMatterClassStmt: %w", cerr)
		}
	}
	if q.deleteAvailabilityStmt != nil {
		if cerr := q.deleteAvailabilityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAvailabilityStmt: %w", cerr)
		}
	}
	if q.deletePhoneStmt != nil {
		if cerr := q.deletePhoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePhoneStmt: %w", cerr)
		}
	}
	if q.deleteProfessionalInformationStmt != nil {
		if cerr := q.deleteProfessionalInformationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProfessionalInformationStmt: %w", cerr)
		}
	}
	if q.deleteProfessionalUserStmt != nil {
		if cerr := q.deleteProfessionalUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProfessionalUserStmt: %w", cerr)
		}
	}
	if q.deleteResponsibleStudentStmt != nil {
		if cerr := q.deleteResponsibleStudentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteResponsibleStudentStmt: %w", cerr)
		}
	}
	if q.deleteStudentUserStmt != nil {
		if cerr := q.deleteStudentUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteStudentUserStmt: %w", cerr)
		}
	}
	if q.deleteSubjectMatterStmt != nil {
		if cerr := q.deleteSubjectMatterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSubjectMatterStmt: %w", cerr)
		}
	}
	if q.deleteSubjectMatterClassStmt != nil {
		if cerr := q.deleteSubjectMatterClassStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSubjectMatterClassStmt: %w", cerr)
		}
	}
	if q.getAvailabilityStmt != nil {
		if cerr := q.getAvailabilityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAvailabilityStmt: %w", cerr)
		}
	}
	if q.getPhoneStmt != nil {
		if cerr := q.getPhoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPhoneStmt: %w", cerr)
		}
	}
	if q.getProfessionalInformationStmt != nil {
		if cerr := q.getProfessionalInformationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProfessionalInformationStmt: %w", cerr)
		}
	}
	if q.getProfessionalUserStmt != nil {
		if cerr := q.getProfessionalUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProfessionalUserStmt: %w", cerr)
		}
	}
	if q.getResponsibleStudentStmt != nil {
		if cerr := q.getResponsibleStudentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getResponsibleStudentStmt: %w", cerr)
		}
	}
	if q.getStudentUserStmt != nil {
		if cerr := q.getStudentUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStudentUserStmt: %w", cerr)
		}
	}
	if q.getSubjectMatterStmt != nil {
		if cerr := q.getSubjectMatterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSubjectMatterStmt: %w", cerr)
		}
	}
	if q.getSubjectMatterClassStmt != nil {
		if cerr := q.getSubjectMatterClassStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSubjectMatterClassStmt: %w", cerr)
		}
	}
	if q.listAvailabilityStmt != nil {
		if cerr := q.listAvailabilityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAvailabilityStmt: %w", cerr)
		}
	}
	if q.listPhoneStmt != nil {
		if cerr := q.listPhoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPhoneStmt: %w", cerr)
		}
	}
	if q.listProfessionalInformationStmt != nil {
		if cerr := q.listProfessionalInformationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProfessionalInformationStmt: %w", cerr)
		}
	}
	if q.listProfessionalInformationByUserStmt != nil {
		if cerr := q.listProfessionalInformationByUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProfessionalInformationByUserStmt: %w", cerr)
		}
	}
	if q.listProfessionalUserStmt != nil {
		if cerr := q.listProfessionalUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProfessionalUserStmt: %w", cerr)
		}
	}
	if q.listResponsibleStudentStmt != nil {
		if cerr := q.listResponsibleStudentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listResponsibleStudentStmt: %w", cerr)
		}
	}
	if q.listStudentUserStmt != nil {
		if cerr := q.listStudentUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listStudentUserStmt: %w", cerr)
		}
	}
	if q.listSubjectMatterStmt != nil {
		if cerr := q.listSubjectMatterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listSubjectMatterStmt: %w", cerr)
		}
	}
	if q.listSubjectMatterClassStmt != nil {
		if cerr := q.listSubjectMatterClassStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listSubjectMatterClassStmt: %w", cerr)
		}
	}
	if q.updatePhoneStmt != nil {
		if cerr := q.updatePhoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePhoneStmt: %w", cerr)
		}
	}
	if q.updateProfessionalInformationStmt != nil {
		if cerr := q.updateProfessionalInformationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProfessionalInformationStmt: %w", cerr)
		}
	}
	if q.updateProfessionalUserStmt != nil {
		if cerr := q.updateProfessionalUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProfessionalUserStmt: %w", cerr)
		}
	}
	if q.updateResponsibleStudentStmt != nil {
		if cerr := q.updateResponsibleStudentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateResponsibleStudentStmt: %w", cerr)
		}
	}
	if q.updateStudentUserStmt != nil {
		if cerr := q.updateStudentUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateStudentUserStmt: %w", cerr)
		}
	}
	if q.updateSubjectMatterStmt != nil {
		if cerr := q.updateSubjectMatterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSubjectMatterStmt: %w", cerr)
		}
	}
	if q.updateSubjectMatterClassStmt != nil {
		if cerr := q.updateSubjectMatterClassStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSubjectMatterClassStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                    DBTX
	tx                                    *sql.Tx
	addAvailabilityStmt                   *sql.Stmt
	createPhoneStmt                       *sql.Stmt
	createProfessionalInformationStmt     *sql.Stmt
	createProfessionalUserStmt            *sql.Stmt
	createResponsibleStudentStmt          *sql.Stmt
	createStudentUserStmt                 *sql.Stmt
	createSubjectMatterStmt               *sql.Stmt
	createSubjectMatterClassStmt          *sql.Stmt
	deleteAvailabilityStmt                *sql.Stmt
	deletePhoneStmt                       *sql.Stmt
	deleteProfessionalInformationStmt     *sql.Stmt
	deleteProfessionalUserStmt            *sql.Stmt
	deleteResponsibleStudentStmt          *sql.Stmt
	deleteStudentUserStmt                 *sql.Stmt
	deleteSubjectMatterStmt               *sql.Stmt
	deleteSubjectMatterClassStmt          *sql.Stmt
	getAvailabilityStmt                   *sql.Stmt
	getPhoneStmt                          *sql.Stmt
	getProfessionalInformationStmt        *sql.Stmt
	getProfessionalUserStmt               *sql.Stmt
	getResponsibleStudentStmt             *sql.Stmt
	getStudentUserStmt                    *sql.Stmt
	getSubjectMatterStmt                  *sql.Stmt
	getSubjectMatterClassStmt             *sql.Stmt
	listAvailabilityStmt                  *sql.Stmt
	listPhoneStmt                         *sql.Stmt
	listProfessionalInformationStmt       *sql.Stmt
	listProfessionalInformationByUserStmt *sql.Stmt
	listProfessionalUserStmt              *sql.Stmt
	listResponsibleStudentStmt            *sql.Stmt
	listStudentUserStmt                   *sql.Stmt
	listSubjectMatterStmt                 *sql.Stmt
	listSubjectMatterClassStmt            *sql.Stmt
	updatePhoneStmt                       *sql.Stmt
	updateProfessionalInformationStmt     *sql.Stmt
	updateProfessionalUserStmt            *sql.Stmt
	updateResponsibleStudentStmt          *sql.Stmt
	updateStudentUserStmt                 *sql.Stmt
	updateSubjectMatterStmt               *sql.Stmt
	updateSubjectMatterClassStmt          *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                    tx,
		tx:                                    tx,
		addAvailabilityStmt:                   q.addAvailabilityStmt,
		createPhoneStmt:                       q.createPhoneStmt,
		createProfessionalInformationStmt:     q.createProfessionalInformationStmt,
		createProfessionalUserStmt:            q.createProfessionalUserStmt,
		createResponsibleStudentStmt:          q.createResponsibleStudentStmt,
		createStudentUserStmt:                 q.createStudentUserStmt,
		createSubjectMatterStmt:               q.createSubjectMatterStmt,
		createSubjectMatterClassStmt:          q.createSubjectMatterClassStmt,
		deleteAvailabilityStmt:                q.deleteAvailabilityStmt,
		deletePhoneStmt:                       q.deletePhoneStmt,
		deleteProfessionalInformationStmt:     q.deleteProfessionalInformationStmt,
		deleteProfessionalUserStmt:            q.deleteProfessionalUserStmt,
		deleteResponsibleStudentStmt:          q.deleteResponsibleStudentStmt,
		deleteStudentUserStmt:                 q.deleteStudentUserStmt,
		deleteSubjectMatterStmt:               q.deleteSubjectMatterStmt,
		deleteSubjectMatterClassStmt:          q.deleteSubjectMatterClassStmt,
		getAvailabilityStmt:                   q.getAvailabilityStmt,
		getPhoneStmt:                          q.getPhoneStmt,
		getProfessionalInformationStmt:        q.getProfessionalInformationStmt,
		getProfessionalUserStmt:               q.getProfessionalUserStmt,
		getResponsibleStudentStmt:             q.getResponsibleStudentStmt,
		getStudentUserStmt:                    q.getStudentUserStmt,
		getSubjectMatterStmt:                  q.getSubjectMatterStmt,
		getSubjectMatterClassStmt:             q.getSubjectMatterClassStmt,
		listAvailabilityStmt:                  q.listAvailabilityStmt,
		listPhoneStmt:                         q.listPhoneStmt,
		listProfessionalInformationStmt:       q.listProfessionalInformationStmt,
		listProfessionalInformationByUserStmt: q.listProfessionalInformationByUserStmt,
		listProfessionalUserStmt:              q.listProfessionalUserStmt,
		listResponsibleStudentStmt:            q.listResponsibleStudentStmt,
		listStudentUserStmt:                   q.listStudentUserStmt,
		listSubjectMatterStmt:                 q.listSubjectMatterStmt,
		listSubjectMatterClassStmt:            q.listSubjectMatterClassStmt,
		updatePhoneStmt:                       q.updatePhoneStmt,
		updateProfessionalInformationStmt:     q.updateProfessionalInformationStmt,
		updateProfessionalUserStmt:            q.updateProfessionalUserStmt,
		updateResponsibleStudentStmt:          q.updateResponsibleStudentStmt,
		updateStudentUserStmt:                 q.updateStudentUserStmt,
		updateSubjectMatterStmt:               q.updateSubjectMatterStmt,
		updateSubjectMatterClassStmt:          q.updateSubjectMatterClassStmt,
	}
}
