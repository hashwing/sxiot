create table sxito_admin (
	`admin_id` varchar(40) not null primary key,
	`admin_role` int(1) not null,
	`admin_account` varchar(32) not null,
	`admin_password` varchar(32) not null,
	`admin_alias` varchar(32) not null,
	`admin_email` varchar(32),
	`admin_phone` varchar(15)
)ENGINE=InnoDB CHARSET=utf8;

INSERT INTO `sxito_admin` (`admin_id`,`admin_role`, `admin_account`, `admin_password`, `admin_alias`, `admin_email`, `admin_phone`) VALUES ('6ba7b810-9dad-11d1-80b4-00c04fd430c8','1', 'admin', '21232f297a57a5a743894a0e4a801fc3', 'admin', '', '');

create table sxito_user (
	`user_id` varchar(40) not null  primary key,
	`user_account` varchar(32) not null,
	`user_password` varchar(32) not null,
	`user_alias` varchar(32) not null,
	`user_eamil` varchar(32),
	`user_phone` varchar(15)
)ENGINE=InnoDB CHARSET=utf8;

create table sxito_brand (
	`brand_id` varchar(40) not null  primary key,
	`brand_name` varchar(64) not null,
	`brand_metadata` varchar(255)
)ENGINE=InnoDB CHARSET=utf8;

create table sxito_gateway (
	`gateway_id` varchar(40) not null primary key,
	`admin_id` varchar(40) not null,
	`gateway_name` varchar(64) not null
)ENGINE=InnoDB CHARSET=utf8;

create table sxito_device (
	`device_id` varchar(40) not null  primary key,
	`admin_id` varchar(40) not null,
	`brand_id` varchar(40) not null,
	`device_alias` varchar(64) not null,
	foreign key (`admin_id`) references sxito_admin(`admin_id`),
	foreign key (`brand_id`) references sxito_brand(`brand_id`)
)ENGINE=InnoDB CHARSET=utf8;

create table sxito_user_device (
	`id` varchar(40) not null primary key,
	`user_id` varchar(40) not null,
	`device_id` varchar(40) not null,
	`device_alias` varchar(64) not null,
	`device_metadata` varchar(255),
	foreign key (`device_id`) references sxito_device(`device_id`),
	foreign key (`user_id`) references sxito_user(`user_id`)
)ENGINE=InnoDB CHARSET=utf8;

create table sxito_news (
	`news_id` varchar(40) not null primary key,
	`news_title` varchar(255) not null,
	`news_content` text,
	`created` datetime
)ENGINE=InnoDB CHARSET=utf8;

create table sxito_message (
	`message_id` varchar(40) not null primary key,
	`user_id` varchar(40) not null,
	`message_title` varchar(255) not null,
	`message_content` text,
	`created` datetime,
	foreign key (`user_id`) references sxito_user(`user_id`)
)ENGINE=InnoDB CHARSET=utf8;