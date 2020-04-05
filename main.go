package main

import (
    "errors"
    "flag"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/kouhin/envflag"
    _ "github.com/lib/pq"
    "github.com/nickkeers/housemate/datalayer"
    "github.com/nickkeers/housemate/handlers"
    log "github.com/sirupsen/logrus"
)

var (
    listenPort = flag.Int("listen-port", 8080, "The port to listen on")
    database = flag.String("database", "postgres", "Database to use, for now, postgres")
    dsn = flag.String("dsn", "", "connection string for postgres")
    resetDb = flag.Bool("reset", false, "reset the database - run all down migrations")
)

func main() {
    if err := run(); err != nil {
        log.Error(err)
    }
}

func run() error {
    if err := envflag.Parse(); err != nil {
        log.Fatal(err)
    }

    if *database != "postgres" {
        return errors.New("unsupported database - only postgres is supported for now")
    }

    db, err := datalayer.NewPostgresDataAdapter(*dsn)

    if err != nil {
        return err
    }

    driver, err := postgres.WithInstance(db.GetInnerSqlDb(), &postgres.Config{})
    migrator, err := migrate.NewWithDatabaseInstance("file://./migrations/", "postgres", driver)

    if err != nil {
        return err
    }

    if *resetDb {
        _ = migrator.Down()
        return nil
    }

    _ = migrator.Up()

    if *listenPort < 443 || *listenPort > 65535 {
        return errors.New("invalid port number")
    }

    return setupAndRunGin(db)

}

func setupAndRunGin(db *datalayer.PostgresDataAdapter) error {
    handlerContext := handlers.NewHandlerContext(db)

    r := gin.New()

    healthGroup := r.Group("health")
    healthGroup.GET("/ping", func(ctx *gin.Context) {
        ctx.Status(200)
    })
    healthGroup.GET("/healthcheck", handlerContext.HealthcheckHandler)

    users := r.Group("/users")
    {
        users.GET("/:id", handlerContext.GetUserByIdHandler)
    }


    return r.Run(fmt.Sprintf(":%d", *listenPort))
}