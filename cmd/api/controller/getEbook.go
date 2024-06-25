package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// GetEbook godoc
// @Summary Retorna um eBook específico
// @Description Exibe um eBook em formato PDF com base no parâmetro fornecido
// @Tags eBooks
// @Accept json
// @Produce application/pdf
// @Param ebook path string true "ID do eBook (1 ou 2)"
// @Success 200 {file} application/pdf "Ebook"
// @Router /viewbook/{ebook} [get]
func GetEbook(ctx *gin.Context) {
	ebook := ctx.Param("ebook")
	fmt.Printf("o valor do ebook é:%s\n", ebook)
	if ebook == "1" {
		ebook = "Ebook - Código de Conduta e Integridade"
	} else {
		ebook = "Ebook - Política  de Porta-Vozes"
	}
	filePath := filepath.Join("./pkg/ebooks/", ebook+" (1).pdf")
	fileBytes, err := ioutil.ReadFile(filePath)

	fmt.Printf("%s\n", filePath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Arquivo não encontrado"})
		return
	}

	ctx.Header("Content-Type", "application/pdf")

	ctx.Data(http.StatusOK, "application/pdf", fileBytes)
}
