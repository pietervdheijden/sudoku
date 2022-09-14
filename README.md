# Sudoku

## Quick start

```bash
$ docker-compose up
```

## Environments

|Environment|Frontend URL|Backend URL|
|---|---|---|
|Local Development|http://127.0.0.1:5173|http://127.0.0.1:8080|
|Docker Compose|http://127.0.0.1:5000|http://127.0.0.1:5001|
|Production|https://productionURL.com|https://productionURL.com/api|


## Sudoku info
Links:
- Sudoku source: https://sudoku.com/easy/

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