# buat skema
migrate create -ext sql -dir db/migrations -seq initial_schema

#  jalanin migration
migrate -path db/migrations -database "postgresql://root:secret@localhost:5000/simple_bank?sslmode=disable" -verbose up