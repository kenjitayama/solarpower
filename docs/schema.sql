
create table solar_power_logs (
	id INT NOT NULL AUTO_INCREMENT,
	log_date DATETIME,
	generated FLOAT,
	sold FLOAT,
	bought FLOAT,
	used FLOAT,
	p1_generated FLOAT,
	external_generated FLOAT,
	income INT,
	outgo INT,
	created_at DATETIME,
	updated_at DATETIME,
	deleted_at DATETIME,
	PRIMARY KEY (`id`)
);

create index idx_solar_power_logs_log_date on solar_power_logs(log_date);