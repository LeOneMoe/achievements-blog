create schema achievement_blog;

create table achievement_blog.achievement
(
    id              int auto_increment
        primary key,
    title           text not null,
    achievementText text not null,
    date            text not null,
    authorID        int  not null
);

create table achievement_blog.user
(
    id        int auto_increment
        primary key,
    nickName  text not null,
    firstName text not null,
    lastName  text not null
);

create table achievement_blog.comment
(
    id            int auto_increment
        primary key,
    userID        int  not null,
    achievementID int  not null,
    commentText   text not null
);

create table achievement_blog.like
(
    id            int auto_increment
        primary key,
    userID        int not null,
    achievementID int null
);

CREATE TRIGGER after_achievement_delete AFTER DELETE on achievement_blog.achievement
    FOR EACH ROW
BEGIN
    DELETE FROM `like`
    WHERE `like`.achievementID = old.id;
    DELETE FROM comment
    WHERE comment.achievementID = old.id;
END
