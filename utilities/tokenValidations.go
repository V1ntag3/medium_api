package utilities

import (
	"medium_api/database"
	"medium_api/models"
)

// Função para remover um token
func UnauthorizedToken(token string) {

	var revokeToken models.RevokedToken

	revokeToken.Token = token

	database.DB.Where("token = ?", token).First(&revokeToken).Delete(&revokeToken)

}

// Função para cadastrar um token valido
func AuthorizedToken(token string) {

	var revokeToken models.RevokedToken

	revokeToken.Token = token

	database.DB.Create(&revokeToken)

}

// Função para verificar se um token está revogado
func IsAuthorizedToken(token string) bool {

	var revokedToken models.RevokedToken

	database.DB.Where("token = ?", token).First(&revokedToken)
	if revokedToken.Token == token {
		// O token está na tabela
		return true
	}
	print(revokedToken.Token)
	// O token não esta na tabela
	return false
}
