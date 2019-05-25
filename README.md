# TeachingAffairsAdministrationSystem


## setup

### development
```bash
# solve frontend dependencies
yarn

# run frontend in dev mode
yarn run dev

# run backend
go run main.go

# then you can visit frontend on http://localhost:3000, but there remains some troubles. To avoid it, you should use webservers like nginx to proxy :3000 and :9090 to the same port to visit.
```

### release
```bash
# run frontend and backend together
docker-compose up --build
# then you should use webservers like nginx to proxy :30003 and :30004 to the same port to visit.
```