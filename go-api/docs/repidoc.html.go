package docs

import _ "embed"

//go:embed swagger.json
var SwaggerJSON []byte

const RapiDocHTML = `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>API Documentation - RapiDoc</title>
  <script type="module" src="https://unpkg.com/rapidoc/dist/rapidoc-min.js"></script>
  <style>
    body { margin: 0; padding: 0; }
  </style>
</head>
<body>
  <rapi-doc
    spec-url="/docs/swagger.json"
    theme="dark"
    render-style="read"
    show-header="true"
    allow-authentication="true"
    primary-color="#6366f1"
    bg-color="#0f172a"
    text-color="#f8fafc"
    nav-bg-color="#1e293b"
    font-size="default"
  >
  </rapi-doc>
</body>
</html>`

const SwaggerUIHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>API Documentation - Swagger UI</title>
  <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css" />
  <style>
    html { box-sizing: border-box; overflow: -moz-scrollbars-vertical; overflow-y: scroll; }
    *, *:before, *:after { box-sizing: inherit; }
    body { margin: 0; background: #fafafa; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
  <script src="https://unpkg.com/swagger-ui-dist/swagger-ui-standalone-preset.js"></script>
  <script>
    window.onload = function() {
      window.ui = SwaggerUIBundle({
        url: "/docs/swagger.json",
        dom_id: '#swagger-ui',
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout"
      });
    };
  </script>
</body>
</html>`
