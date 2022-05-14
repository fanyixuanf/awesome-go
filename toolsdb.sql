-- toolsdb.sys_users definition

CREATE TABLE toolsdb.`sys_users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `uuid` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nick_name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `header_img` text COLLATE utf8mb4_unicode_ci,
  `authority_id` int NOT NULL DEFAULT '0',
  `user_profile` text COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  KEY `sys_user_uuid_IDX` (`uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';


INSERT INTO toolsdb.sys_users
(created_at, updated_at, uuid, username, password, nick_name, header_img, authority_id, user_profile)
VALUES(CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '00000000-0000-0000-0000-000000000000', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2Ffc%2F64%2Fab%2Ffc64ab13576aeb6522356f19c4426aee.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1647743817&t=1c9ace914b246cc1f17bbd043a54f5bf', 888, 'test1');

CREATE TABLE `jwt_blacklists` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `created_at` timestamp NULL DEFAULT NULL,
                            `updated_at` timestamp NULL DEFAULT NULL,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            `jwt` text COLLATE utf8mb4_unicode_ci,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='jwt_blacklists';



CREATE TABLE toolsdb.`sys_operation_records` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `created_at` timestamp NULL DEFAULT NULL,
                                  `updated_at` timestamp NULL DEFAULT NULL,
                                  `deleted_at` timestamp NULL DEFAULT NULL,
                                  ip varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                  method text COLLATE utf8mb4_unicode_ci,
                                  path varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                  status int(1) DEFAULT 0,
                                  latency int(10) DEFAULT 0,
                                  agent text COLLATE utf8mb4_unicode_ci,
                                  error_message text COLLATE utf8mb4_unicode_ci,
                                  body text COLLATE utf8mb4_unicode_ci,
                                  resp text COLLATE utf8mb4_unicode_ci,
                                  user_id int(10) DEFAULT 0,
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='sys_operation_records';
