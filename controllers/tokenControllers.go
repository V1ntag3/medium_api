package controllers

import (
	"medium_api/database"
	"medium_api/models"
)

// Função para revogar um token
func RevokeToken(token string) {

	var revokeToken models.RevokedToken
	revokeToken.Token = token
	database.DB.Create(&revokeToken)

}

// Função para verificar se um token está revogado
func IsTokenRevoked(token string) bool {

	var revokedToken models.RevokedToken

	err := database.DB.Where("token = ?", token).First(&revokedToken)
	if err == nil {
		// O token não está na tabela de tokens revogados
		return false
	}
	// O token está na tabela de tokens revogados
	return true
}
