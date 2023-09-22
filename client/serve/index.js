const express = require('express');
const path = require('path');
const app = express();

const APP_PORT = Number(process.env.APP_PORT) || 3000

const PROJECT_ROOT_PATH = process.cwd()
const BUILD_DIR_PATH = path.join(PROJECT_ROOT_PATH, 'build')
const INDEX_HTML_FILE_PATH = path.join(BUILD_DIR_PATH, 'index.html')

app.disable('x-powered-by')


// dev
// app.use('/floral', express.static(BUILD_DIR_PATH));
// prod
app.use(express.static(BUILD_DIR_PATH));

app.get('/*', function (_req, res) {
  res.sendFile(INDEX_HTML_FILE_PATH);
});

app.listen(APP_PORT, () => {
  console.info(`Server listening on port ${APP_PORT}`)
});