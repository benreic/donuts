drop table if exists `participant`;

create table `participant` (
	`participant_id` int(11) unsigned not null auto_increment,
	`email` varchar(200) not null,
	`salt`  varchar(100) not null,
	`pwd`   varchar(1000) not null,
	PRIMARY KEY ( `participant_id` )
) ENGINE=InnoDB default charset=utf8;

