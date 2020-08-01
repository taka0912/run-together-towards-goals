# migrationの実行方法

- gooseをインストール

    `go get -u github.com/pressly/goose/cmd/goose`

- 下記のコマンド操作は、プロジェクトパスにて行う

## local

- 実行状況を確認

    `goose -dir ./db/migrations/ mysql "user:pass@tcp(127.0.0.1:3306)/my_goal?parseTime=true&charset=utf8&loc=Asia%2FTokyo" status`

- goose upで実行

    `goose -dir ./db/migrations/ mysql "user:pass@tcp(127.0.0.1:3306)/my_goal?parseTime=true&charset=utf8&loc=Asia%2FTokyo" up`

- 既にGormのマイグレーションによってDBが補正されておりGooseが正常に作動しない場合は、下記の様なSQLをDBに対して実行し、マイグレーションを実行させたことにする

    `INSERT INTO my_goal.goose_db_version (version_id, is_applied, tstamp) VALUES (20200727181157, 1, now());`

    version_idはマイグレーションファイルの個々の名前を入れる。実行させたことにしたいマイグレーションのversion_idを変更したSQLを作成し、実行する。

## 本番（Heroku）環境

- 実行状況を確認

    `goose -dir ./db/migrations/ mysql "ba4b88c4d6ec1b:ad2382f7@tcp(us-cdbr-east-02.cleardb.com:3306)/heroku_84029d8861a40e9?parseTime=true&charset=utf8&loc=Asia%2FTokyo" status`

- goose upで実行

    `goose -dir ./db/migrations/ mysql "ba4b88c4d6ec1b:ad2382f7@tcp(us-cdbr-east-02.cleardb.com:3306)/heroku_84029d8861a40e9?parseTime=true&charset=utf8&loc=Asia%2FTokyo" up`

- 既にGormのマイグレーションによってDBが補正されておりGooseが正常に作動しない場合は、下記の様なSQLをDBに対して実行し、マイグレーションを実行させたことにする

    `INSERT INTO my_goal.goose_db_version (version_id, is_applied, tstamp) VALUES (20200727181157, 1, now());`

    version_idはマイグレーションファイルの個々のファイル名を上記の参考に代入する。  
    実行させたことにしたいマイグレーションのversion_idを変更したSQLを作成し、実行する。  

