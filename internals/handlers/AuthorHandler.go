package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/sqlcVsGorm_GO/internals/repositories"
)

// AuthorHandler handles HTTP requests related to authors.
type AuthorHandler struct {
	repo repositories.AuthorRepository
}

// NewAuthorHandler creates a new AuthorHandler with the given repository.
func NewAuthorHandler(repo repositories.AuthorRepository) *AuthorHandler {
	return &AuthorHandler{repo: repo}
}

// CreateAuthor handles the creation of a new author.
func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := h.repo.CreateAuthor(c.Request.Context(), req.Name, req.Bio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetAuthor handles retrieving an author by ID.
func (h *AuthorHandler) GetAuthor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	author, err := h.repo.GetAuthor(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get author"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// ListAuthors handles listing all authors.
func (h *AuthorHandler) ListAuthors(c *gin.Context) {
	authors, err := h.repo.ListAuthors(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list authors"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

// DeleteAuthor handles deleting an author by ID.
func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	if err := h.repo.DeleteAuthor(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Author deleted"})
}
