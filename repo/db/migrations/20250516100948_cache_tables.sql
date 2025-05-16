-- migrate:up
CREATE    TABLE "queue_number" (
          "date" DATE PRIMARY KEY NOT NULL,
          "number" INT NOT NULL
          );

-- migrate:down
DROP      TABLE if EXISTS "queue_number";