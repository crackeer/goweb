CREATE TABLE "object"(
  id    INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT,
  content TEXT,
  tag TEXT,
  create_time text,
  update_time text
)

CREATE TABLE "code"(
  id    INTEGER PRIMARY KEY AUTOINCREMENT,
  comment TEXT,
  content TEXT,
  type TEXT,
  keyword TEXT,
  create_time text,
  update_time text
)