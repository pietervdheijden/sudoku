# Sudoku

## Backend
Backend is written in Go.

Notes:
- HTTP framework for Go: gin: gin-gonic/gin (https://github.com/gin-gonic/gin)

Links:
- Sudoku source: https://sudoku.com/easy/

Golang links:
- https://go.dev/tour/concurrency/11
- https://go.dev/doc/
- https://go.dev/doc/code
- https://go.dev/ref/spec
- https://vimeo.com/53221558
- https://go.dev/doc/codewalk/functions/


Sudoku tactics:
- XYZ wing: https://www.sudokuwiki.org/XYZ_Wing
- A Sudoku should have a unique solution.
- Single cell = naked singles

Helper: https://www.sudoku-solutions.com/

Technique per rating / difficulty:
- Simple: Naked Single, Hidden Single
- Easy: Naked Pair, Hidden Pair, Pointing Pairs
- Medium: Naked Triple, Naked Quad, Pointing Triples, Hidden Triple, Hidden Quad
- Hard: XWing, Swordfish, Jellyfish, XYWing, XYZWing

Advanced sudoku strategies:
- https://www.sudokuonline.io/tips/advanced-sudoku-strategies
- https://www.kristanix.com/sudokuepic/sudoku-solving-techniques.php
- Swordfish: https://www.sudokuwiki.org/Sword_Fish_Strategy
- Jellyfish: https://www.sudokuwiki.org/Jelly_Fish_Strategy 