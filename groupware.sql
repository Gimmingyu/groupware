CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL COMMENT '기본키',
    `email`      longtext  NOT NULL COMMENT '이메일',
    `password`   longtext  NOT NULL COMMENT '패스워드',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '생성일자',
    `updated_at` timestamp NOT NULL COMMENT '수정일자'
);

CREATE TABLE `group`
(
    `id`          bigint unsigned NOT NULL COMMENT '기본키',
    `owner_id`    bigint unsigned NOT NULL COMMENT '그룹 소유자',
    `name`        longtext NOT NULL COMMENT '그룹 이름',
    `description` longtext NULL COMMENT '긴 설명',
    `summary`     longtext NULL COMMENT '짧은 설명'
);

CREATE TABLE `document`
(
    `id`         bigint unsigned NOT NULL COMMENT '기본키',
    `group_id`   bigint unsigned NOT NULL COMMENT '그룹 id',
    `author_id`  bigint unsigned NOT NULL COMMENT '작성자 id',
    `email`      longtext  NOT NULL COMMENT '제목',
    `password`   longtext  NOT NULL COMMENT '패스워드',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '생성일자',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '수정일자'
);

CREATE TABLE `role`
(
    `id`       bigint unsigned NOT NULL COMMENT '기본키',
    `group_id` bigint unsigned NOT NULL COMMENT '그룹에서는 여러가지 역할을 정의할 수 있음',
    `name`     varchar(255) NOT NULL COMMENT '역할 이름'
);

CREATE TABLE `membership`
(
    `id`         bigint unsigned NOT NULL COMMENT '기본키',
    `user_id`    bigint unsigned NOT NULL COMMENT '유저 id',
    `group_id`   bigint unsigned NOT NULL COMMENT '그룹 id',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '멤버가 된 날짜',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '수정일자',
    `deleted_at` timestamp NULL COMMENT '삭제일자'
);

CREATE TABLE `permission`
(
    `id`      bigint unsigned NOT NULL COMMENT '기본키',
    `role_id` bigint unsigned NOT NULL COMMENT '역할은 여러가지 권한을 가진다',
    `name`    enum('READ_DOC', 'WRITE_DOC', 'UPDATE_DOC', 'DELETE_DOC', 'ALLOW_INVITATION', 'ALLOW_KICK') NOT NULL COMMENT '권한 이름'
);

CREATE TABLE `invitation`
(
    `id`         bigint unsigned NOT NULL COMMENT '기본키',
    `user_id`    bigint unsigned NOT NULL COMMENT '링크 생성자',
    `group_id`   bigint unsigned NOT NULL COMMENT '그룹 id',
    `link`       longtext  NOT NULL COMMENT '링크',
    `expired_at` timestamp NOT NULL COMMENT '만료일자',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '생성일자',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '수정일자'
);

CREATE TABLE `message`
(
    `id`          bigint unsigned NOT NULL COMMENT '기본키',
    `sender_id`   bigint unsigned NOT NULL COMMENT '보내는 사람',
    `receiver_id` bigint unsigned NOT NULL COMMENT '받는 사람',
    `content`     longtext  NOT NULL COMMENT '내용',
    `created_at`  timestamp NOT NULL DEFAULT current_timestamp() COMMENT '생성일자',
    `updated_at`  timestamp NOT NULL DEFAULT current_timestamp() COMMENT '수정일자'
);

ALTER TABLE `user`
    ADD CONSTRAINT `PK_USER` PRIMARY KEY (
                                          `id`
        );

ALTER TABLE `group`
    ADD CONSTRAINT `PK_GROUP` PRIMARY KEY (
                                           `id`
        );

ALTER TABLE `document`
    ADD CONSTRAINT `PK_DOCUMENT` PRIMARY KEY (
                                              `id`
        );

ALTER TABLE `role`
    ADD CONSTRAINT `PK_ROLE` PRIMARY KEY (
                                          `id`
        );

ALTER TABLE `membership`
    ADD CONSTRAINT `PK_MEMBERSHIP` PRIMARY KEY (
                                                `id`
        );

ALTER TABLE `permission`
    ADD CONSTRAINT `PK_PERMISSION` PRIMARY KEY (
                                                `id`
        );

ALTER TABLE `invitation`
    ADD CONSTRAINT `PK_INVITATION` PRIMARY KEY (
                                                `id`
        );

ALTER TABLE `message`
    ADD CONSTRAINT `PK_MESSAGE` PRIMARY KEY (
                                             `id`
        );

ALTER TABLE `group`
    ADD CONSTRAINT `FK_user_TO_group_1` FOREIGN KEY (
                                                     `owner_id`
        )
        REFERENCES `user` (
                           `id`
            );

ALTER TABLE `document`
    ADD CONSTRAINT `FK_group_TO_document_1` FOREIGN KEY (
                                                         `group_id`
        )
        REFERENCES `group` (
                            `id`
            );

ALTER TABLE `document`
    ADD CONSTRAINT `FK_user_TO_document_1` FOREIGN KEY (
                                                        `author_id`
        )
        REFERENCES `user` (
                           `id`
            );

ALTER TABLE `role`
    ADD CONSTRAINT `FK_group_TO_role_1` FOREIGN KEY (
                                                     `group_id`
        )
        REFERENCES `group` (
                            `id`
            );

ALTER TABLE `membership`
    ADD CONSTRAINT `FK_user_TO_membership_1` FOREIGN KEY (
                                                          `user_id`
        )
        REFERENCES `user` (
                           `id`
            );

ALTER TABLE `membership`
    ADD CONSTRAINT `FK_group_TO_membership_1` FOREIGN KEY (
                                                           `group_id`
        )
        REFERENCES `group` (
                            `id`
            );

ALTER TABLE `permission`
    ADD CONSTRAINT `FK_role_TO_permission_1` FOREIGN KEY (
                                                          `role_id`
        )
        REFERENCES `role` (
                           `id`
            );

ALTER TABLE `invitation`
    ADD CONSTRAINT `FK_user_TO_invitation_1` FOREIGN KEY (
                                                          `user_id`
        )
        REFERENCES `user` (
                           `id`
            );

ALTER TABLE `invitation`
    ADD CONSTRAINT `FK_group_TO_invitation_1` FOREIGN KEY (
                                                           `group_id`
        )
        REFERENCES `group` (
                            `id`
            );

ALTER TABLE `message`
    ADD CONSTRAINT `FK_user_TO_message_1` FOREIGN KEY (
                                                       `sender_id`
        )
        REFERENCES `user` (
                           `id`
            );

ALTER TABLE `message`
    ADD CONSTRAINT `FK_user_TO_message_2` FOREIGN KEY (
                                                       `receiver_id`
        )
        REFERENCES `user` (
                           `id`
            );

