# Sudoku Frontend
The frontend is written in Vue.js.

## Quick start
Run the app with npm:

```bash
$ npm run dev
```

Run the app with Docker:

```bash
$ docker build . -t sudoku-app && docker run -p 5173:8080 sudoku-app
```

The app is now available on http://localhost:5173.

## dotenv configuration

The environments (see root README.md) are configured with dotenv files:

- Local Development: `.env`
- Docker Compose: `.env.dockercompose`
- Production: `.env.production`

For more details about Modes and Environment variables in Vue.js, see: https://cli.vuejs.org/guide/mode-and-env.html.

# Credits
Favicon source: https://www.vectorstock.com/royalty-free-vector/symmetrical-sudoku-icon-vector-9412445
