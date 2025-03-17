FROM swaggerapi/swagger-ui:v5.7.2

ENV PORT=3001 URLS="[ { url: '/api/openapi.yaml', name: 'API' } ]"

EXPOSE 3001

COPY ./openapi/openapi.yaml /usr/share/nginx/html/api/openapi.yaml
