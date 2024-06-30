# tkbai-management-dashboard

tkbai-management-dashboard

### Docker build

## FE

```
 docker build . -f dockerfile.fe -t tkbai-fe:latest
```

## BE

```
 docker build . -f dockerfile.be -t tkbai-be:latest
```

### Migrate DB Up

```
migrate -database "mysql://root:03IZmt7eRMukIHdoZahl@tcp(mysql:3306)/sv" -path migration up
```

### Migrate DB Down

```
migrate -database "mysql://root:03IZmt7eRMukIHdoZahl@tcp(mysql:3306)/sv" -path migration down
```

### Generate migration file command

```
migrate create -ext sql -dir {directory_path} -seq {migration_name}
```
