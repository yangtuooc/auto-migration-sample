FROM liquibase/liquibase:4.21

RUN lpm add mysql --global

ENV TZ=Asia/Shanghai

ADD ./changelog /liquibase/changelog

ENTRYPOINT []

CMD ["sh", "-c", "/liquibase/docker-entrypoint.sh --url=${DB_URL} --username=${DB_USERNAME} --password=${DB_PASSWORD} --classpath=/liquibase/changelog --changeLogFile=changelogs.xml --database-changelog-lock-table-name=liquibase_changelog_lock --database-changelog-table-name=liquibase_changelog update"]

