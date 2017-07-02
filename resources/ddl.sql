CREATE TABLE IF NOT EXISTS users (
		`id` INTEGER NOT NULL PRIMARY KEY,
		`login` TEXT NOT NULL,
		`name` TEXT NOT NULL,
		`status` VARCHAR(128) NOT NULL,
		`selected_game_id` INTEGER NOT NULL,
		`cur_game_id` INTEGER NOT NULL,
		`cur_question_id` INTEGER NOT NULL,
		`cur_attempts` INTEGER NOT NULL,
		`cur_hint_attempts` INTEGER NOT NULL,
		`cur_points` INTEGER NOT NULL
)