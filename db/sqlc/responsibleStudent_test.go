package db

import (
	"context"
	"database/sql"
	"myclass/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createResponsibleStudentForTest(t *testing.T) ResponsibleStudent {
	responsibleStudentparams := CreateResponsibleStudentParams{
		Name:   sql.NullString{String: "Monica", Valid: true},
		Gender: sql.NullString{String: "Feminino", Valid: true},
		Email:  sql.NullString{String: "mshcneider@gamil.com", Valid: true},
		DateOfBirth: sql.NullTime{Time: time.Date(
			1992, 01, 13, 20, 34, 58, 651387237, time.UTC), Valid: true},
		Cpf:     sql.NullInt32{Int32: 84737262, Valid: true},
		PhoneID: util.CreateRandomIds(0, 10000), //esse Id tem que ser único em cada teste
		UpdatedAt: sql.NullTime{Time: time.Date(
			2023, 11, 17, 20, 34, 58, 651387237, time.UTC), Valid: true},
	}

	responsibleStudent, err := testQueries.CreateResponsibleStudent(context.Background(), responsibleStudentparams)
	require.NoError(t, err)
	require.NotEmpty(t, responsibleStudent)
	require.Equal(t, responsibleStudentparams.Name, responsibleStudent.Name)
	require.Equal(t, responsibleStudentparams.Gender, responsibleStudent.Gender)
	require.Equal(t, responsibleStudentparams.Email, responsibleStudent.Email)
	require.Equal(t, responsibleStudentparams.Cpf, responsibleStudent.Cpf)
	require.Equal(t, responsibleStudentparams.PhoneID, responsibleStudent.PhoneID)

	return responsibleStudent

}
func TestCreateResponsibleStudent(t *testing.T) {
	createResponsibleStudentForTest(t)
}

func TestGetResponsibleStudent(t *testing.T) {
	responsibleStudent1 := createResponsibleStudentForTest(t)
	getResponsibleStudent1, err := testQueries.GetResponsibleStudent(context.Background(), responsibleStudent1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getResponsibleStudent1)
	require.Equal(t, responsibleStudent1.ID, getResponsibleStudent1.ID)
	require.Equal(t, responsibleStudent1.Name, getResponsibleStudent1.Name)
	require.Equal(t, responsibleStudent1.Gender, getResponsibleStudent1.Gender)
	require.Equal(t, responsibleStudent1.Email, getResponsibleStudent1.Email)
	require.Equal(t, responsibleStudent1.Cpf, getResponsibleStudent1.Cpf)
	require.Equal(t, responsibleStudent1.PhoneID, getResponsibleStudent1.PhoneID)
}

func TestUpdateResponsibleStudent(t *testing.T) {
	newResponsibleStudent := createResponsibleStudentForTest(t)

	updateArgs := UpdateResponsibleStudentParams{
		ID:    newResponsibleStudent.ID,
		Name:  sql.NullString{String: "Monica", Valid: true},
		Email: sql.NullString{String: "monicaschnei@gmail.com", Valid: true},
	}
	updatedResponsibleStudent, err := testQueries.UpdateResponsibleStudent(context.Background(), updateArgs)
	require.NoError(t, err)
	require.NotEmpty(t, updatedResponsibleStudent)
	require.Equal(t, newResponsibleStudent.ID, updatedResponsibleStudent.ID)
	require.Equal(t, newResponsibleStudent.Gender, updatedResponsibleStudent.Gender)
	require.Equal(t, updateArgs.Email, updatedResponsibleStudent.Email)
	require.Equal(t, newResponsibleStudent.Cpf, updatedResponsibleStudent.Cpf)
	require.Equal(t, newResponsibleStudent.PhoneID, updatedResponsibleStudent.PhoneID)

}

func TestListResponsibleStudent(t *testing.T) {
	for i := 0; i < 4; i++ {
		createResponsibleStudentForTest(t)
	}

	argList := ListResponsibleStudentParams{
		Limit:  2,
		Offset: 2,
	}
	responsibleStudents, err := testQueries.ListResponsibleStudent(context.Background(), argList)
	require.NoError(t, err)
	require.Len(t, responsibleStudents, 2)
	for _, responsible := range responsibleStudents {
		require.NotEmpty(t, responsible)
	}
}

func TestDeleteResponsibleStudent(t *testing.T) {
	responsibleStudent := createResponsibleStudentForTest(t)
	_, err := testQueries.DeleteResponsibleStudent(context.Background(), responsibleStudent.ID)
	require.NoError(t, err)
	getDeletedResponsibleStudent, err := testQueries.GetResponsibleStudent(context.Background(), responsibleStudent.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getDeletedResponsibleStudent)
}
