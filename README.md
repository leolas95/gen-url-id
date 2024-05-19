# Generate short hash

This service generates a short hash string used by the (mini-url)[https://github.com/leolas95/mini-url]  service
to generate short URLs

The short hash is a base 62 (i.e: alphanumerical) string of length 7. Only one change is needed to increase/decrease
the hash length

## Endpoints

### Shorten URL
```http request
POST /urls

url={long_url}
```

Returns the hash for the given URL