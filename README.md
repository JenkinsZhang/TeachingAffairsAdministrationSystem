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

# then you can visit http://localhost:3000
# but we suggest that you should configure your webserver e.g. nginx to prevent cross-origin problems due to our idleness ☺️.
```

### release
```bash
# run frontend and backend together
docker-compose up --build
# then you can visit http://localhost:30003
# by the way, don't forget configure your webserver to make the service work.
```