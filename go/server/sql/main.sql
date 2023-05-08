CREATE TABLE IF NOT EXISTS reports (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  data TEXT NULL,
  create_time DATETIME DEFAULT (datetime('now')),
  modify_time DATETIME DEFAULT (datetime('now', 'localtime'))
);
CREATE TABLE IF NOT EXISTS cases (
  id INTEGER,
  name TEXT NULL,
  report_id INTEGER NULL,
  status TEXT NULL,
  data TEXT NULL,
  apis TEXT NULL,
  create_time DATETIME DEFAULT (datetime('now')),
  modify_time DATETIME DEFAULT (datetime('now', 'localtime'))
);