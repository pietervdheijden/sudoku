package routers

import (
	"github.com/gin-gonic/gin"

	sudokuV1 "github.com/pietervdheijden/sudoku/routers/api/v1"
)

func Run()  {
	router := gin.Default()
	
	router.Use(CORSMiddleware())
	
	router.POST("/api/v1/solve", sudokuV1.Solve)
	router.POST("/api/v1/hint", sudokuV1.Hint)
	router.POST("/api/v1/check", sudokuV1.Check)
	router.POST("/api/v1/options", sudokuV1.Options) // TODO: consider renaming to candidates

	// TODO: add endpoint /api/v1/difficulty or return difficulty in endpoint /api/v1/check
    // Difficulty rating from https://www.sudoku-solutions.com/:
    // - Simple: Naked Single, Hidden Single
    // - Easy: Naked Pair, Hidden Pair, Pointing Pairs
    // - Medium: Naked Triple, Naked Quad, Pointing Triples, Hidden Triple, Hidden Quad
    // - Hard: XWing, Swordfish, Jellyfish, XYWing, XYZWing
	
	router.Run(":8080")
}

// TODO: replace CORS middleware with BFF, or host backend and frontend on same URL.
// Source: https://stackoverflow.com/a/29439630/3737152
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}