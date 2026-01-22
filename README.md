flyway.url=jdbc:oracle:thin:@dev-db-host:1521/DEVDB
flyway.user=flyway_user
flyway.password=********

flyway.schemas=APP_CORE,APP_TXN,APP_REPORT,APP_AUDIT
flyway.defaultSchema=FLYWAY_META

flyway.locations=filesystem:/opt/flyway/sql/common,filesystem:/opt/flyway/sql/dev

flyway.baselineOnMigrate=true
flyway.baselineVersion=1
flyway.baselineDescription=Existing database baseline

flyway.validateOnMigrate=true
flyway.outOfOrder=false
