package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ReygaFitra/inc-final-project.git/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyCard = []model.Card{
	{
		CardID:         1,
		UserID:         1,
		CardType:       "BRI",
		CardNumber:     "123456789",
		ExpirationDate: "08/26",
		CVV:            "123",
	},
	{
		CardID:         2,
		UserID:         2,
		CardType:       "Mandiri",
		CardNumber:     "987654321",
		ExpirationDate: "04/25",
		CVV:            "321",
	},
}

var dummyCardResponse = []model.CardResponse{
	{
		UserID:         1,
		CardType:       "BRI",
		CardNumber:     "123456789",
		ExpirationDate: "08/26",
		CVV:            "123",
	},
	{
		UserID:         2,
		CardType:       "Mandiri",
		CardNumber:     "987654321",
		ExpirationDate: "04/25",
		CVV:            "321",
	},
}

type CardRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *CardRepositoryTestSuite) TestGetAll_Success() {
	var users = dummyCardResponse[0]
	suite.mockSql.ExpectQuery("SELECT user_id, card_type, card_number, expiration_date, cvv FROM mst_card").WillReturnRows(sqlmock.NewRows([]string{"card_type", "card_number", "expiration_date", "CVV"}).AddRow(users.CardType, users.CardNumber, users.ExpirationDate, users.CVV))
	cardRepository := NewCardRepository(suite.mockDB)
	result := cardRepository.GetAll()
	assert.NotNil(suite.T(), result)
}

func (suite *CardRepositoryTestSuite) TestGetAll_Failed() {
	suite.mockSql.ExpectQuery("SELECT user_id, card_type, card_number, expiration_date, cvv FROM mst_card").WillReturnError(fmt.Errorf("no data"))
	cardRepository := NewCardRepository(suite.mockDB)
	result := cardRepository.GetAll()
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "no data", result)
}

// func (suite *CardRepositoryTestSuite) TestGetByUserID_Success() {
// 	card := dummyCardResponse[0]
// 	suite.mockSql.ExpectQuery("SELECT user_id, card_type, card_number, expiration_date, cvv FROM mst_card WHERE user_id = \\$1").WithArgs(card.UserID).WillReturnRows(sqlmock.NewRows([]string{"user_id, card_type, card_number, expiration_date, cvv"}).AddRow(card.UserID, card.CardType, card.CardNumber, card.ExpirationDate, card.CVV))
// 	cardRepository := NewCardRepository(suite.mockDB)
// 	result, err := cardRepository.GetByUserID(card.UserID)
// 	assert.NotNil(suite.T(), result)
// 	assert.Nil(suite.T(), err)
// }

func (suite *CardRepositoryTestSuite) TestGetByUserID_Failed() {
	suite.mockSql.ExpectQuery("SELECT user_id, card_type, card_number, expiration_date, cvv FROM mst_card WHERE user_id = \\$1").WithArgs(1).WillReturnError(errors.New("some error"))
	cardRepository := NewCardRepository(suite.mockDB)
	result, err := cardRepository.GetByUserID(1)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *CardRepositoryTestSuite) TestGetByCardID_Success() {
	card := dummyCard[0]
	suite.mockSql.ExpectQuery("SELECT card_id, card_type, card_number, expiration_date, cvv, user_id FROM mst_card WHERE card_id = \\$1").WithArgs(card.CardID).WillReturnRows(sqlmock.NewRows([]string{
		"card_id", "card_type", "card_number", "expiration_date", "cvv", "user_id"}).AddRow(card.CardID, card.CardType, card.CardNumber, card.ExpirationDate, card.CVV, card.UserID))
	cardRepository := NewCardRepository(suite.mockDB)
	result, err := cardRepository.GetByCardID(card.CardID)
	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), card, *result)
}

func (suite *CardRepositoryTestSuite) TestGetByCardID_Failed() {
	card := dummyCard[0]
	suite.mockSql.ExpectQuery("SELECT card_id, card_type, card_number, expiration_date, cvv, user_id FROM mst_card WHERE card_id = \\$1").WithArgs(card.CardID).WillReturnError(errors.New("card not found"))
	cardRepository := NewCardRepository(suite.mockDB)
	result, err := cardRepository.GetByCardID(card.CardID)
	assert.Nil(suite.T(), result)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "card ID not found", err.Error())
}

func (suite *CardRepositoryTestSuite) TestCreateCard_Success() {
	newCard := dummyCardResponse[0]
	userID := uint(1)
	suite.mockSql.ExpectExec("INSERT INTO mst_card \\(user_id, card_type, card_number, expiration_date, cvv\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)").WithArgs(userID, newCard.CardType, newCard.CardNumber, newCard.ExpirationDate, newCard.CVV).WillReturnResult(sqlmock.NewResult(1, 1))
	cardRepository := NewCardRepository(suite.mockDB)
	result, err := cardRepository.Create(userID, &newCard)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.EqualValues(suite.T(), &newCard, result)
}

// Setup test
func (suite *CardRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		log.Println("Error database", err)
	}
	suite.mockDB = mockDb
	suite.mockSql = mockSql
}

func (suite *CardRepositoryTestSuite) TearDownTest() {
	suite.mockDB.Close()
}

func TestCardRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CardRepositoryTestSuite))
}
