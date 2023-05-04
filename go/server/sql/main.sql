CREATE TABLE IF NOT EXISTS test_results (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  create_time DATETIME DEFAULT (datetime('now')),
  modify_time DATETIME DEFAULT (datetime('now', 'localtime'))
);

CREATE TABLE IF NOT EXISTS test_result_reports (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  test_result_id INTEGER NOT NULL,
  report_date DATE NOT NULL,
  test_result TEXT NOT NULL,
  create_time DATETIME DEFAULT (datetime('now')),
  modify_time DATETIME DEFAULT (datetime('now', 'localtime')),
  FOREIGN KEY (test_result_id) REFERENCES test_results(id)
);